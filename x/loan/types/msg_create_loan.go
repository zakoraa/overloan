package types

import (
	errorsmod "cosmossdk.io/errors" // package error resmi SDK terbaru
	sdk "github.com/cosmos/cosmos-sdk/types"

	loanv1 "github.com/zakoraa/overloan/api/overloan/loan/v1"
)

// Pastikan MsgCreateLoan dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*loanv1.MsgCreateLoan)(nil)

// ValidateBasic melakukan validasi stateless (tanpa akses store)
// Fungsi ini dipanggil sebelum tx diproses keeper
func (msg *loanv1.MsgCreateLoan) ValidateBasic() error {

	// Validasi alamat borrower dalam format Bech32
	if _, err := sdk.AccAddressFromBech32(msg.Borrower); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid borrower address: %v",
			err,
		)
	}

	// Validasi struktur coin (denom & amount valid)
	if !msg.Principal.IsValid() {
		return ErrInvalidPrincipal.Wrap("invalid coin format")
	}

	// Principal tidak boleh nol atau negatif
	if msg.Principal.IsZero() || msg.Principal.Amount.IsNegative() {
		return ErrInvalidPrincipal.Wrap("principal must be positive")
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
