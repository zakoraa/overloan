package loan

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	"github.com/cosmos/cosmos-sdk/x/loan/keeper"
	loantypes "github.com/cosmos/cosmos-sdk/x/loan/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.AppModule      = AppModule{}
)

//
// BASIC
//

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return loantypes.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (AppModuleBasic) RegisterInterfaces(reg codectypes.InterfaceRegistry) {}

func (AppModuleBasic) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	return json.RawMessage("{}")
}

func (AppModuleBasic) ValidateGenesis(codec.JSONCodec, json.RawMessage) error {
	return nil
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(
	clientCtx client.Context,
	mux *runtime.ServeMux,
) {}

//
// FULL MODULE
//

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{keeper: k}
}

func (AppModule) IsAppModule() {}

func (AppModule) IsOnePerModuleType() {}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	loanv1.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
}

func (am AppModule) InitGenesis(ctx context.Context, _ json.RawMessage) error {
	return nil
}

func (am AppModule) ExportGenesis(ctx context.Context) (json.RawMessage, error) {
	return json.RawMessage("{}"), nil
}