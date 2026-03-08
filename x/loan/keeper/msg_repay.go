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

	repayInt := msg.Amount.Amount

	// repayment harus positif
	if !repayInt.IsPositive() {
		return nil, types.ErrInvalidPrincipal.
			Wrap("repayment must be positive")
	}

	if loan.Outstanding == nil {
		return nil, types.ErrInvalidStateTransition.
			Wrap("outstanding not set")
	}

	outstandingInt := loan.Outstanding.Amount

	// tidak boleh lebih besar dari outstanding
	if repayInt.GT(outstandingInt) {
		return nil, types.ErrInvalidPrincipal.
			Wrap("repayment exceeds outstanding")
	}

	// update outstanding
	newOutstanding := outstandingInt.Sub(repayInt)

	loan.Outstanding = &sdk.Coin{
		Denom:  msg.Amount.Denom,
		Amount: newOutstanding,
	}

	// jika lunas
	if newOutstanding.IsZero() {
		loan.Status = types.LoanStatus_LOAN_STATUS_REPAID
	}

	m.SetLoan(sdkCtx, loan)

	// emit event
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