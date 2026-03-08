package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) ConfirmDisbursement(
	ctx context.Context,
	msg *types.MsgConfirmDisbursement,
) (*types.MsgConfirmDisbursementResponse, error) {

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

	loan, err := m.GetLoan(sdkCtx, msg.LoanId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrLoanNotFound
		}
		return nil, err
	}

	// Validasi state machine
	if err := types.CanDisburse(loan); err != nil {
		return nil, err
	}

	// Validasi denom settlement
	if loan.Principal.Denom != params.SettlementDenom {
		return nil, types.ErrInvalidPrincipal
	}

	now := sdkCtx.BlockTime()

	loan.Status = types.LoanStatus_LOAN_STATUS_DISBURSED
	loan.DisbursedAt = &now
	loan.Outstanding = loan.Principal

	m.SetLoan(sdkCtx, loan)

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanDisbursed,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loan.Id)),
			sdk.NewAttribute(types.AttributeKeyOmnibus, msg.Omnibus),
			sdk.NewAttribute(types.AttributeKeyAmount, loan.Principal.String()),
		),
	)

	return &types.MsgConfirmDisbursementResponse{}, nil
}
