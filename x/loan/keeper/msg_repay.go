package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) RepayLoan(
	ctx context.Context,
	msg *types.MsgRepayLoan,
) (*types.MsgRepayLoanResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params, err := m.GetParams(sdkCtx)
	if err != nil {
		return nil, err
	}

	// Validasi omnibus authority
	if err := m.ValidateOmnibusAuthority(
		sdkCtx,
		msg.Omnibus,
	); err != nil {
		return nil, err
	}

	// Ambil loan
	loan, err := m.GetLoan(sdkCtx, msg.LoanId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrLoanNotFound
		}
		return nil, err
	}

	// Validasi state machine
	if err := types.CanRepay(loan); err != nil {
		return nil, err
	}

	// Validasi denom
	if msg.Amount.Denom != params.SettlementDenom {
		return nil, types.ErrInvalidPrincipal.
			Wrap("invalid settlement denom")
	}

	// Validasi repayment positif
	repayInt := msg.Amount.Amount
	if !repayInt.IsPositive() {
		return nil, types.ErrInvalidPrincipal.
			Wrap("repayment must be positive")
	}

	if loan.Outstanding == nil {
		return nil, types.ErrInvalidStateTransition.
			Wrap("outstanding not set")
	}

	outstandingInt := loan.Outstanding.Amount

	// Cek repayment tidak melebihi outstanding
	if repayInt.GT(outstandingInt) {
		return nil, types.ErrInvalidPrincipal.
			Wrap("repayment exceeds outstanding")
	}

	// Parse omnibus address
	omnibusAddr, err := sdk.AccAddressFromBech32(msg.Omnibus)
	if err != nil {
		return nil, types.ErrInvalidAddress.Wrap(err.Error())
	}

	moduleAddr := m.GetModuleAddress()

	// Transfer repayment dari omnibus → module
	repayCoin := sdk.NewCoin(msg.Amount.Denom, repayInt)
	coins := sdk.NewCoins(repayCoin)

	if err := m.bankKeeper.SendCoins(
		sdkCtx,
		omnibusAddr,
		moduleAddr,
		coins,
	); err != nil {
		return nil, err
	}

	// Burn settlement token
	if err := m.bankKeeper.BurnCoins(
		sdkCtx,
		types.ModuleName,
		coins,
	); err != nil {
		return nil, err
	}

	// Update outstanding
	newOutstanding := outstandingInt.Sub(repayInt)

	loan.Outstanding = &sdk.Coin{
		Denom:  msg.Amount.Denom,
		Amount: newOutstanding,
	}

	// Jika lunas
	if newOutstanding.IsZero() {
		loan.Status = types.LoanStatus_LOAN_STATUS_REPAID
	}

	m.SetLoan(sdkCtx, loan)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanRepaid,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyOmnibus, msg.Omnibus),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
	)

	return &types.MsgRepayLoanResponse{}, nil
}