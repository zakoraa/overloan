package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pastikan MsgRejectLoan memenuhi sdk.Msg
var _ sdk.Msg = (*MsgRejectLoan)(nil)

// ValidateBasic melakukan validasi stateless
func (m *MsgRejectLoan) ValidateBasic() error {

	// laz wajib valid Bech32
	if _, err := sdk.AccAddressFromBech32(m.Laz); err != nil {
		return errorsmod.Wrapf(
			ErrInvalidAddress,
			"invalid laz address: %v",
			err,
		)
	}

	// LoanId wajib ada
	if m.LoanId == 0 {
		return ErrInvalidRequest.Wrap("loan_id required")
	}

	return nil
}
