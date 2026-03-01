package keeper

import (
	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (k Keeper) GetParams(ctx sdk.Context) loanv1.Params {
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)
	return params.Params
}

func (k Keeper) SetParams(ctx sdk.Context, params loanv1.Params) {
	k.paramSpace.SetParamSet(ctx, &types.Params{
		Params: params,
	})
}
