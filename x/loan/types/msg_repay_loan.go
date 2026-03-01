package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

// Pastikan MsgRepayLoan dari proto memenuhi interface sdk.Msg
var _ sdk.Msg = (*loanv1.MsgRepayLoan)(nil)

// ValidateMsgRepayLoan melakukan validasi stateless
// Fungsi ini dipanggil sebelum tx diproses keeper
func ValidateMsgRepayLoan(msg *loanv1.MsgRepayLoan) error {

	// Validasi authority address (format Bech32)
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid authority address: %v",
			err,
		)
	}

	// loan_id wajib ada
	if msg.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	// Amount tidak boleh nil
	if msg.Amount == nil {
		return ErrInvalidRequest.Wrap("amount required")
	}

	// Validasi coin positif
	if err := ValidatePositiveCoin(
		msg.Amount.Denom,
		msg.Amount.Amount,
	); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidCoin,
			"invalid repayment amount: %v",
			err,
		)
	}

	return nil
}
