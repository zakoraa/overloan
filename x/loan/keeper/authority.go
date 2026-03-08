package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (k Keeper) ValidateAuthority(authority string) error {

	// validasi format address
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		return types.ErrInvalidAddress.Wrap(err.Error())
	}

	if authority != k.authority {
		return types.ErrUnauthorized
	}

	return nil
}

func (k Keeper) ValidateLazAuthority(ctx sdk.Context, authority string) error {

	addr, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		return types.ErrInvalidAddress.Wrap(err.Error())
	}

	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}

	if !isAuthorityInList(addr, params.LazAuthorities) {
		return types.ErrUnauthorized
	}

	return nil
}

func (k Keeper) ValidateOmnibusAuthority(ctx sdk.Context, authority string) error {

	addr, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		return types.ErrInvalidAddress.Wrap(err.Error())
	}

	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}

	if !isAuthorityInList(addr, params.OmnibusAuthorities) {
		return types.ErrUnauthorized
	}

	return nil
}

func isAuthorityInList(authority sdk.AccAddress, list []string) bool {
	for _, auth := range list {

		addr, err := sdk.AccAddressFromBech32(auth)
		if err != nil {
			continue
		}

		if authority.Equals(addr) {
			return true
		}
	}

	return false
}

func (k Keeper) GetModuleAddress() sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.ModuleName)
}