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

	if err := m.ValidateAuthority(sdkCtx, msg.Authority); err != nil {
		return nil, err
	}

	if msg.Params == nil {
		return nil, types.ErrInvalidRequest.Wrap("params cannot be nil")
	}

	if err := m.SetParams(sdkCtx, *msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
