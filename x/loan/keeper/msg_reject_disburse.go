package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) RejectDisbursement(
	ctx context.Context,
	msg *types.MsgRejectDisbursement,
) (*types.MsgRejectDisbursementResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

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
	if err := types.CanRejectDisburse(loan); err != nil {
		return nil, err
	}

	// Validasi reason (opsional tapi direkomendasikan)
	if msg.Reason == "" {
		return nil, types.ErrInvalidRequest.Wrap("reason required")
	}

	now := sdkCtx.BlockTime()

	// Update state
	loan.Status = types.LoanStatus_LOAN_STATUS_CANCELLED
	loan.DisbursedAt = &now

	// Persist state
	m.SetLoan(sdkCtx, loan)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanDisburseRejected,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyOmnibus, msg.Omnibus),
			sdk.NewAttribute(types.AttributeKeyReason, msg.Reason),
		),
	)

	return &types.MsgRejectDisbursementResponse{}, nil
}
