package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) RejectLoan(
	ctx context.Context,
	msg *types.MsgRejectLoan,
) (*types.MsgRejectLoanResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Validasi laz authority
	if err := m.ValidateLazAuthority(
		sdkCtx,
		msg.Laz,
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

	// Pastikan loan memang milik laz ini
	if loan.Laz != msg.Laz {
		return nil, types.ErrUnauthorized
	}

	// Validasi state machine (loan harus PENDING)
	if err := types.CanReject(loan); err != nil {
		return nil, err
	}

	// Update status
	loan.Status = types.LoanStatus_LOAN_STATUS_REJECTED

	m.SetLoan(sdkCtx, loan)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanRejected,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyLaz, msg.Laz),
		),
	)

	return &types.MsgRejectLoanResponse{}, nil
}