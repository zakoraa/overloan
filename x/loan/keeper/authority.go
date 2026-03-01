package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (k Keeper) ValidateAuthority(ctx sdk.Context, authority string) error {

	if authority != k.authority {
		return types.ErrUnauthorized
	}

	params := k.GetParams(ctx)

	if authority != params.LazGroupPolicy {
		return types.ErrUnauthorized
	}

	ok, err := k.groupKeeper.HasGroupPolicy(ctx, authority)
	if err != nil {
		return err
	}
	if !ok {
		return types.ErrInvalidAuthority
	}

	return nil
}

func (k Keeper) GetModuleAddress() sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.ModuleName)
}
