package keeper

import (
	"context"
	"fmt"

	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m msgServer) ConfirmDisbursement(
	ctx context.Context,
	msg *loanv1.MsgConfirmDisbursement,
) (*loanv1.MsgConfirmDisbursementResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Ambil params
	params := m.GetParams(sdkCtx)

	// Validasi authority (harus omnibus policy)
	if err := m.ValidateAuthority(
		sdkCtx,
		msg.Authority,
	); err != nil {
		return nil, err
	}

	// Ambil loan
	loan, found := m.GetLoan(sdkCtx, msg.LoanId)
	if !found {
		return nil, types.ErrLoanNotFound
	}

	// Validasi state machine
	if err := types.CanDisburse(loan); err != nil {
		return nil, err
	}

	moduleAddr := m.GetModuleAddress()

	// Mint settlement token ke module account
	amountInt, ok := sdkmath.NewIntFromString(loan.Principal.Amount)
	if !ok {
		return nil, types.ErrInvalidCoin.Wrap("invalid principal amount")
	}

	sdkCoin := sdk.NewCoin(
		loan.Principal.Denom,
		amountInt,
	)

	coins := sdk.NewCoins(sdkCoin)
	// Transfer ke Omnibus policy address
	omnibusAddr, err := sdk.AccAddressFromBech32(params.OmnibusGroupPolicy)
	if err != nil {
		return nil, err
	}

	if err := m.bankKeeper.SendCoins(
		sdkCtx,
		moduleAddr,
		omnibusAddr,
		coins,
	); err != nil {
		return nil, err
	}

	// Update loan state
	now := sdkCtx.BlockTime()

	loan.Status = loanv1.LoanStatus_LOAN_STATUS_DISBURSED
	loan.DisbursedAt = timestamppb.New(now)
	loan.Outstanding = loan.Principal

	m.SetLoan(sdkCtx, loan)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanDisbursed,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyAmount, loan.Principal.String()),
		),
	)

	return &loanv1.MsgConfirmDisbursementResponse{}, nil
}
