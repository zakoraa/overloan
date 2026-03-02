package keeper

import (
	"context"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (q queryServer) Params(
	ctx context.Context,
	req *loanv1.QueryParamsRequest,
) (*loanv1.QueryParamsResponse, error) {

	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params, err := q.k.GetParams(sdkCtx)
	if err != nil {
		return nil, err
	}

	return &loanv1.QueryParamsResponse{
		Params: params,
	}, nil
}
