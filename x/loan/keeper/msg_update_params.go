package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (m msgServer) UpdateParams(
	ctx context.Context,
	msg *types.MsgUpdateParams,
) (*types.MsgUpdateParamsResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Validasi authority
	if err := m.ValidateAuthority(msg.Authority); err != nil {
		return nil, err
	}

	// Params tidak boleh nil
	if msg.Params == nil {
		return nil, types.ErrInvalidRequest.Wrap("params must not be nil")
	}

	// Validasi params
	if err := msg.Params.Validate(); err != nil {
		return nil, err
	}

	// Update params
	if err := m.SetParams(sdkCtx, *msg.Params); err != nil {
		return nil, err
	}

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateParams,
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority),
		),
	)

	return &types.MsgUpdateParamsResponse{}, nil
}