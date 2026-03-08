package types

import (
	errorsmod "cosmossdk.io/errors" // package error resmi SDK terbaru
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pastikan MsgApproveLoan dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*MsgApproveLoan)(nil)

// ValidateMsgApproveLoan melakukan validasi stateless (tanpa akses store)
// Fungsi ini dipanggil sebelum tx diproses keeper
func (m *MsgApproveLoan) ValidateMsgApproveLoan() error {

	// Validasi alamat laz dalam format Bech32
	if _, err := sdk.AccAddressFromBech32(m.Laz); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid laz address: %v",
			err,
		)
	}

	// loanId wajib diisi
	if m.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	return nil
}
