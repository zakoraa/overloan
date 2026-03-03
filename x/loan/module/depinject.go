package loan

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	loanmodulev1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/module/v1"
	"github.com/cosmos/cosmos-sdk/x/loan/keeper"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

var _ appmodule.AppModule = AppModule{}

func (AppModule) IsOnePerModuleType() {}

func init() {
	appmodule.Register(
		&loanmodulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	Config       *loanmodulev1.Module

	BankKeeper    types.BankKeeper
	AccountKeeper types.AccountKeeper
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	authority := authtypes.NewModuleAddress("gov")

	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.BankKeeper,
		in.AccountKeeper,
		authority.String(),
	)

	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{
		Module: m,
		Keeper: k,
	}
}
