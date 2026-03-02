package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

func RegisterInterfaces(reg codectypes.InterfaceRegistry) {
	reg.RegisterImplementations((*sdk.Msg)(nil),
		&loanv1.MsgCreateLoan{},
	)

	reg.RegisterImplementations((*sdk.Msg)(nil),
		&loanv1.MsgApproveLoan{},
	)

	reg.RegisterImplementations((*sdk.Msg)(nil),
		&loanv1.MsgConfirmDisbursement{},
	)

	reg.RegisterImplementations((*sdk.Msg)(nil),
		&loanv1.MsgRejectLoan{},
	)

	reg.RegisterImplementations((*sdk.Msg)(nil),
		&loanv1.MsgRepayLoan{},
	)

}
