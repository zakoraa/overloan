package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	return k.Params.Get(ctx)
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	err := k.Params.Set(ctx, params)

	if err != nil {
		return err
	}
	return nil
}
