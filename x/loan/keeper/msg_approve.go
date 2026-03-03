package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) ApproveLoan(
	ctx context.Context,
	msg *types.MsgApproveLoan,
) (*types.MsgApproveLoanResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if err := m.ValidateAuthority(
		sdkCtx,
		msg.Authority,
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
	if err := types.CanApprove(loan); err != nil {
		return nil, err
	}

	// Update state
	now := sdkCtx.BlockTime()

	loan.Status = types.LoanStatus_LOAN_STATUS_APPROVED
	loan.LazPolicy = msg.Authority
	loan.ApprovedAt = &now

	// Persist
	m.SetLoan(sdkCtx, loan)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanApproved,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority),
		),
	)

	return &types.MsgApproveLoanResponse{}, nil
}
