package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	loanv1 "github.com/zakoraa/cosmos-sdk/api/overloan/loan/v1"
)

// Validate melakukan validasi konfigurasi parameter modul secara stateless
func (p loanv1.Params) Validate() error {

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