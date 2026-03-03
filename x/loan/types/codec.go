package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(reg codectypes.InterfaceRegistry) {
	reg.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLoan{},
		&MsgApproveLoan{},
		&MsgRejectLoan{},
		&MsgRepayLoan{},
		&MsgConfirmDisbursement{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(reg, &_Msg_serviceDesc)
}
