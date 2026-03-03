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

	//  Validasi authority
	if err := m.ValidateAuthority(sdkCtx, msg.Authority); err != nil {
		return nil, err
	}

	//  Ambil loan
	loan, err := m.GetLoan(sdkCtx, msg.LoanId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrLoanNotFound
		}
		return nil, err
	}

	//  Validasi state machine
	if loan.Status != types.LoanStatus_LOAN_STATUS_PENDING {
		return nil, types.ErrInvalidStateTransition.
			Wrap("only pending loan can be rejected")
	}

	//  Update status
	loan.Status = types.LoanStatus_LOAN_STATUS_REJECTED

	m.SetLoan(sdkCtx, loan)

	//  Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanRejected,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority),
		),
	)

	return &types.MsgRejectLoanResponse{}, nil
}
