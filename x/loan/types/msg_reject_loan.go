package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	loanv1 "cosmossdk.io/api/overloan/loan/v1"
)

// Pastikan MsgRejectLoan memenuhi sdk.Msg
var _ sdk.Msg = (*loanv1.MsgRejectLoan)(nil)

// ValidateMsgRejectLoan melakukan validasi stateless
func ValidateMsgRejectLoan(msg *loanv1.MsgRejectLoan) error {

	// Authority wajib valid Bech32
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid authority address: %v",
			err,
		)
	}

	// LoanId wajib ada
	if msg.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	return nil
}
