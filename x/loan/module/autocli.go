package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: loanv1.Query_ServiceDesc.ServiceName,
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: loanv1.Msg_ServiceDesc.ServiceName,
		},
	}
}