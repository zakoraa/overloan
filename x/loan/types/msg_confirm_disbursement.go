package types

import (
	errorsmod "cosmossdk.io/errors" // package error resmi SDK terbaru
	sdk "github.com/cosmos/cosmos-sdk/types"

	loanv1 "github.com/zakoraa/overloan/api/overloan/loan/v1"
)

// Pastikan MsgConfirmDisbursement dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*loanv1.MsgConfirmDisbursement)(nil)

// ValidateBasic melakukan validasi stateless (tanpa akses store)
// Fungsi ini dipanggil sebelum tx diproses keeper
func (msg *loanv1.MsgConfirmDisbursement) ValidateBasic() error {

	// Validasi alamat authority dalam format Bech32
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid authority address: %v",
			err,
		)
	}

	// loanId wajib diisi
	if msg.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	return nil
}
