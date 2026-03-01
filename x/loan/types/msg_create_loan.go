package types

import (
	errorsmod "cosmossdk.io/errors" // package error resmi SDK terbaru
	sdk "github.com/cosmos/cosmos-sdk/types"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

// Pastikan MsgCreateLoan dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*loanv1.MsgCreateLoan)(nil)

// ValidateCreateLoan melakukan validasi stateless (tanpa akses store)
// Fungsi ini dipanggil sebelum tx diproses keeper
func ValidateCreateLoan(msg *loanv1.MsgCreateLoan) error {

	// Validasi alamat borrower dalam format Bech32
	if _, err := sdk.AccAddressFromBech32(msg.Borrower); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid borrower address: %v",
			err,
		)
	}

	// Validasi bahwa principal memiliki denom valid dan amount positif
	if err := ValidatePositiveCoin(msg.Principal.Denom, msg.Principal.Amount); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidCoin,
			"invalid pricipal: %v",
			err,
		)
	}

	// Tenor wajib lebih dari 0 bulan
	if msg.TenorMonths == 0 {
		return ErrInvalidTenor.Wrap("tenor must be greater than zero")
	}

	// Metadata hash wajib ada
	if len(msg.MetadataHash) == 0 {
		return errorsmod.Wrap(ErrInvalidRequest, "metadata hash required")
	}

	return nil
}
