package types

import (
	sdkmath "cosmossdk.io/math"
)

// ValidateSettlementDenom memastikan denom sesuai settlement denom modul
func ValidateSettlementDenom(denom, settlementDenom string) error {

	// Denom tidak boleh berbeda dari settlement denom yang dikonfigurasi
	if denom != settlementDenom {
		return ErrInvalidPrincipal.Wrap("invalid settlement denom")
	}

	return nil
}

// ValidatePositiveCoin memastikan denom ada dan amount positif
func ValidatePositiveCoin(denom, amountStr string) error {

	// Denom wajib diisi
	if denom == "" {
		return ErrInvalidPrincipal.Wrap("denom required")
	}

	// Parse string amount menjadi sdkmath.Int
	amount, ok := sdkmath.NewIntFromString(amountStr)
	if !ok {
		return ErrInvalidPrincipal.Wrap("invalid amount format")
	}

	// Amount harus lebih besar dari nol
	if !amount.IsPositive() {
		return ErrInvalidPrincipal.Wrap("amount must be positive")
	}

	return nil
}
