package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter Store Keys
var (
	KeySettlementDenom    = []byte("SettlementDenom")
	KeyMinLoanAmount      = []byte("MinLoanAmount")
	KeyMaxLoanAmount      = []byte("MaxLoanAmount")
	KeyMaxTenorMonths     = []byte("MaxTenorMonths")
	KeyLazAuthorities     = []byte("LazAuthorities")
	KeyOmnibusAuthorities = []byte("OmnibusAuthorities")
)

// Params wrapper lokal agar bisa implement ParamSet

// NewParams creates a new Params instance.
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters.
//
//	func DefaultParams() Params {
//		return NewParams()
//	}
func DefaultParams() Params {
	return Params{
		SettlementDenom: "stake",
		MaxLoanAmount:   1_000_000,
		MinLoanAmount:   100,
		MaxTenorMonths:  12,
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs Implement ParamSet Interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeySettlementDenom,
			&p.SettlementDenom,
			validateDenom,
		),
		paramtypes.NewParamSetPair(
			KeyMinLoanAmount,
			&p.MinLoanAmount,
			validateUint64,
		),
		paramtypes.NewParamSetPair(
			KeyMaxLoanAmount,
			&p.MaxLoanAmount,
			validateUint64,
		),
		paramtypes.NewParamSetPair(
			KeyMaxTenorMonths,
			&p.MaxTenorMonths,
			validateUint64,
		),
		paramtypes.NewParamSetPair(
			KeyLazAuthorities,
			&p.LazAuthorities,
			validateAddress,
		),
		paramtypes.NewParamSetPair(
			KeyOmnibusAuthorities,
			&p.OmnibusAuthorities,
			validateAddress,
		),
	}
}

func (p Params) Validate() error {

	// Denom wajib diisi
	// if p.SettlementDenom == "" {
	// 	return fmt.Errorf("settlement denom required")
	// }

	// // Min tidak boleh lebih besar dari Max
	// if p.MinLoanAmount > p.MaxLoanAmount {
	// 	return fmt.Errorf("min loan amount cannot exceed max loan amount")
	// }

	// // Max tenor harus > 0
	// if p.MaxTenorMonths == 0 {
	// 	return fmt.Errorf("max tenor months must be positive")
	// }

	return nil
}

// // ValidateParams Stateless Validation (bisnis rule)
// func ValidateParams(p *types.Params) error {

// 	// // SettlementDenom wajib diisi sebagai denom token utama
// 	// if p.SettlementDenom == "" {
// 	// 	return ErrInvalidRequest.Wrap("settlement denom required")
// 	// }

// 	// // MinLoanAmount harus lebih besar dari nol
// 	// if p.MinLoanAmount == 0 {
// 	// 	return ErrInvalidRequest.Wrap("min loan amount must be > 0")
// 	// }

// 	// // MaxLoanAmount harus lebih besar dari MinLoanAmount
// 	// if p.MaxLoanAmount <= p.MinLoanAmount {
// 	// 	return ErrInvalidRequest.Wrap("max must be > min")
// 	// }

// 	// // MaxTenorMonths tidak boleh nol
// 	// if p.MaxTenorMonths == 0 {
// 	// 	return ErrInvalidRequest.Wrap("max tenor must be > 0")
// 	// }

// 	// // Validasi format alamat LazAuthorities
// 	// if _, err := sdk.AccAddressFromBech32(p.LazAuthorities); err != nil {
// 	// 	return ErrInvalidAddress.Wrap("invalid laz group policy")
// 	// }

// 	// // Validasi format alamat OmnibusAuthorities
// 	// if _, err := sdk.AccAddressFromBech32(p.OmnibusAuthorities); err != nil {
// 	// 	return ErrInvalidAddress.Wrap("invalid omnibus group policy")
// 	// }

// 	return nil
// }

// Validators untuk Param Store
func validateDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type")
	}
	if v == "" {
		return fmt.Errorf("denom cannot be empty")
	}
	return nil
}

func validateUint64(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type")
	}
	return nil
}

func validateAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type")
	}
	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid address format")
	}
	return nil
}

// func (p Params) Validate() error {
// 	// if p.SettlementDenom == "" {
// 	// 	return fmt.Errorf("settlement denom required")
// 	// }

// 	// if p.MinLoanAmount == 0 {
// 	// 	return fmt.Errorf("min loan amount must be > 0")
// 	// }

// 	// if p.MaxLoanAmount <= p.MinLoanAmount {
// 	// 	return fmt.Errorf("max loan amount must be greater than min loan amount")
// 	// }

// 	// if p.MaxTenorMonths == 0 {
// 	// 	return fmt.Errorf("max tenor must be > 0")
// 	// }

// 	// if _, err := sdk.AccAddressFromBech32(p.LazAuthorities); err != nil {
// 	// 	return fmt.Errorf("invalid laz group policy address")
// 	// }

// 	// if _, err := sdk.AccAddressFromBech32(p.OmnibusAuthorities); err != nil {
// 	// 	return fmt.Errorf("invalid omnibus group policy address")
// 	// }

// 	return nil
// }
