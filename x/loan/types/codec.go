package types

import (
    codectypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"

    loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

func RegisterInterfaces(reg codectypes.InterfaceRegistry) {
    reg.RegisterImplementations((*sdk.Msg)(nil),
        &loanv1.MsgCreateLoan{},
        &loanv1.MsgRepayLoan{},
        &loanv1.MsgApproveLoan{},
        &loanv1.MsgConfirmDisbursement{},
        &loanv1.MsgRejectLoan{},
    )
}