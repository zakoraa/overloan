package keeper

import (
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (loanv1.Params, error) {
	return k.Params.Get(ctx)
}

func (k Keeper) SetParams(ctx sdk.Context, params *loanv1.Params) error {
	return k.Params.Set(ctx, *params)
}
