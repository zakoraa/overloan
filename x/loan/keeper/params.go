package keeper

import (
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (*loanv1.Params, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &params, nil
}

func (k Keeper) SetParams(ctx sdk.Context, params *loanv1.Params) error {
	err := k.Params.Set(ctx, *params)

	if err != nil {
		return err
	}
	return nil
}
