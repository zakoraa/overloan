package keeper

import (
	"context"
	"fmt"

	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) RepayLoan(
	ctx context.Context,
	msg *loanv1.MsgRepayLoan,
) (*loanv1.MsgRepayLoanResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params := m.GetParams(sdkCtx)

	//  Validasi authority
	if err := m.ValidateAuthority(sdkCtx, msg.Authority); err != nil {
		return nil, err
	}

	//  Ambil loan
	loan, found := m.GetLoan(sdkCtx, msg.LoanId)
	if !found {
		return nil, types.ErrLoanNotFound
	}

	//  Harus sudah disbursed
	if loan.Status != loanv1.LoanStatus_LOAN_STATUS_DISBURSED {
		return nil, types.ErrInvalidStateTransition.
			Wrap("loan must be disbursed")
	}

	//  Validasi denom
	if msg.Amount.Denom != params.SettlementDenom {
		return nil, types.ErrInvalidPrincipal.
			Wrap("invalid settlement denom")
	}

	//  Parse repayment amount
	repayInt, ok := sdkmath.NewIntFromString(msg.Amount.Amount)
	if !ok || !repayInt.IsPositive() {
		return nil, types.ErrInvalidPrincipal.
			Wrap("repayment must be positive")
	}

	//  Parse outstanding
	if loan.Outstanding == nil {
		return nil, types.ErrInvalidStateTransition.
			Wrap("outstanding not set")
	}

	outstandingInt, ok := sdkmath.NewIntFromString(loan.Outstanding.Amount)
	if !ok {
		panic("invalid outstanding amount in state")
	}

	// Cek repayment tidak melebihi outstanding
	if repayInt.GT(outstandingInt) {
		return nil, types.ErrInvalidPrincipal.
			Wrap("repayment exceeds outstanding")
	}

	// Convert ke sdk.Coin untuk transfer
	repayCoin := sdk.NewCoin(msg.Amount.Denom, repayInt)
	coins := sdk.NewCoins(repayCoin)

	omnibusAddr, err := sdk.AccAddressFromBech32(params.OmnibusGroupPolicy)
	if err != nil {
		return nil, err
	}

	moduleAddr := m.GetModuleAddress()

	// Transfer dari Omnibus → Module
	if err := m.bankKeeper.SendCoins(
		sdkCtx,
		omnibusAddr,
		moduleAddr,
		coins,
	); err != nil {
		return nil, err
	}

	// Burn dari module account
	if err := m.bankKeeper.BurnCoins(
		sdkCtx,
		types.ModuleName,
		coins,
	); err != nil {
		return nil, err
	}

	// Update outstanding
	newOutstanding := outstandingInt.Sub(repayInt)

	loan.Outstanding.Amount = newOutstanding.String()

	//  Jika lunas
	if newOutstanding.IsZero() {
		loan.Status = loanv1.LoanStatus_LOAN_STATUS_REPAID
	}

	m.SetLoan(sdkCtx, loan)

	//  Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanRepaid,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.Amount+msg.Amount.Denom),
		),
	)

	return &loanv1.MsgRepayLoanResponse{}, nil
}
