package loan

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	loanv1 "github.com/cosmos/cosmos-sdk/api/overloan/loan/v1"
	"github.com/cosmos/cosmos-sdk/x/loan/keeper"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

type AppModule struct {
	module.AppModule
	keeper keeper.Keeper
}

func NewAppModule(
	cdc codec.Codec,
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
) AppModule {
	return AppModule{
		keeper: k,
	}
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	loanv1.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	// loanv1.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(am.keeper))
}
