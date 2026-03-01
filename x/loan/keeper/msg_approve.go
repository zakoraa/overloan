package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	loanv1 "github.com/cosmos/cosmos-sdk/api/overloan/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m msgServer) ApproveLoan(
	ctx context.Context,
	msg *loanv1.MsgApproveLoan,
) (*loanv1.MsgApproveLoanResponse, error) {

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

	loan.Status = loanv1.LoanStatus_LOAN_STATUS_APPROVED
	loan.LazPolicy = msg.Authority
	loan.ApprovedAt = timestamppb.New(now)

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

	return &loanv1.MsgApproveLoanResponse{}, nil
}
