package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	loanv1 "github.com/cosmos/cosmos-sdk/api/overloan/loan/v1"
)

// Parameter Store Keys
var (
	KeySettlementDenom    = []byte("SettlementDenom")
	KeyMinLoanAmount      = []byte("MinLoanAmount")
	KeyMaxLoanAmount      = []byte("MaxLoanAmount")
	KeyMaxTenorMonths     = []byte("MaxTenorMonths")
	KeyLazGroupPolicy     = []byte("LazGroupPolicy")
	KeyOmnibusGroupPolicy = []byte("OmnibusGroupPolicy")
)

// Params wrapper lokal agar bisa implement ParamSet
type Params struct {
	*loanv1.Params
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
			KeyLazGroupPolicy,
			&p.LazGroupPolicy,
			validateAddress,
		),
		paramtypes.NewParamSetPair(
			KeyOmnibusGroupPolicy,
			&p.OmnibusGroupPolicy,
			validateAddress,
		),
	}
}

// ValidateParams Stateless Validation (bisnis rule)
func ValidateParams(p *loanv1.Params) error {

	// SettlementDenom wajib diisi sebagai denom token utama
	if p.SettlementDenom == "" {
		return ErrInvalidRequest.Wrap("settlement denom required")
	}

	// MinLoanAmount harus lebih besar dari nol
	if p.MinLoanAmount == 0 {
		return ErrInvalidRequest.Wrap("min loan amount must be > 0")
	}

	// MaxLoanAmount harus lebih besar dari MinLoanAmount
	if p.MaxLoanAmount <= p.MinLoanAmount {
		return ErrInvalidRequest.Wrap("max must be > min")
	}

	// MaxTenorMonths tidak boleh nol
	if p.MaxTenorMonths == 0 {
		return ErrInvalidRequest.Wrap("max tenor must be > 0")
	}

	// Validasi format alamat LazGroupPolicy
	if _, err := sdk.AccAddressFromBech32(p.LazGroupPolicy); err != nil {
		return ErrInvalidAddress.Wrap("invalid laz group policy")
	}

	// Validasi format alamat OmnibusGroupPolicy
	if _, err := sdk.AccAddressFromBech32(p.OmnibusGroupPolicy); err != nil {
		return ErrInvalidAddress.Wrap("invalid omnibus group policy")
	}

	return nil
}

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
