package keeper

import (
	"context"
	"fmt"

	loanv1 "cosmossdk.io/api/overloan/loan/v1"
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
	loan, found := m.GetLoan(sdkCtx, msg.LoanId)
	if !found {
		return nil, types.ErrLoanNotFound
	}

	// Validasi state machine
	if err := types.CanApprove(&loan); err != nil {
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
