package types

import (
	errorsmod "cosmossdk.io/errors" // package error resmi SDK terbaru
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pastikan MsgConfirmDisbursement dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*MsgConfirmDisbursement)(nil)

// ValidateBasic melakukan validasi stateless (tanpa akses store)
// Fungsi ini dipanggil sebelum tx diproses keeper
func (m *MsgConfirmDisbursement) ValidateBasic() error {

	// Validasi alamat Omnibus dalam format Bech32
	if _, err := sdk.AccAddressFromBech32(m.Omnibus); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid Omnibus address: %v",
			err,
		)
	}

	// loanId wajib diisi
	if m.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	return nil
}
