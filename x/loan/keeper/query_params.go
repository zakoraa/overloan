package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (q queryServer) Params(
	ctx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {

	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params, err := q.k.GetParams(sdkCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{
		Params: &params,
	}, nil
}
