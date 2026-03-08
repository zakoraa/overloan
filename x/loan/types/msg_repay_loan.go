package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pastikan MsgRepayLoan dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*MsgRepayLoan)(nil)

// ValidateBasic melakukan validasi stateless
// Fungsi ini dipanggil sebelum tx diproses keeper
func (m *MsgRepayLoan) ValidateBasic() error {

	// Validasi omnibus address (format Bech32)
	if _, err := sdk.AccAddressFromBech32(m.Omnibus); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid omnibus address: %v",
			err,
		)
	}

	// loan_id wajib ada
	if m.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	// Amount tidak boleh nil
	if m.Amount == nil {
		return ErrInvalidRequest.Wrap("amount required")
	}

	// Validasi coin positif
	if err := ValidatePositiveCoin(
		m.Amount.Denom,
		m.Amount.Amount,
	); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidCoin,
			"invalid repayment amount: %v",
			err,
		)
	}

	return nil
}
