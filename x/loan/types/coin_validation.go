package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// ValidateSettlementDenom memastikan denom coin sesuai dengan settlement denom modul
func ValidateSettlementDenom(coin sdk.Coin, denom string) error {

	// Denom coin harus sama dengan settlement denom yang dikonfigurasi
	if coin.Denom != denom {
		return ErrInvalidPrincipal.Wrap("invalid settlement denom")
	}

	return nil
}