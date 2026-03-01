/*
Package keeper berisi implementasi logika bisnis dan
pengelolaan state untuk modul loan.

Package ini menangani lifecycle pinjaman,
validasi otoritas, akuntansi token settlement,
*/
package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	loanv1 "github.com/cosmos/cosmos-sdk/api/overloan/loan/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService

	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	// groupKeeper   types.GroupKeeper

	authority string

	Schema collections.Schema

	Loans  collections.Map[uint64, loanv1.Loan]
	NextID collections.Item[uint64]
	Params collections.Item[loanv1.Params]
}

// NewKeeper membuat instance keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	// groupKeeper types.GroupKeeper,
	authority string,
) Keeper {

	sb := collections.NewSchemaBuilder(storeService)

	loans := collections.NewMap(
		sb,
		collections.NewPrefix("loans"),
		"loans",
		collections.Uint64Key,
		codec.CollValue[loanv1.Loan](cdc),
	)

	nextID := collections.NewItem(
		sb,
		collections.NewPrefix("next_id"),
		"next_id",
		collections.Uint64Value,
	)

	params := collections.NewItem(
		sb,
		collections.NewPrefix("params"),
		"params",
		codec.CollValue[loanv1.Params](cdc),
	)

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	return Keeper{
		cdc:           cdc,
		storeService:  storeService,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		// groupKeeper:   groupKeeper,
		authority:     authority,
		Schema:        schema,
		Loans:         loans,
		NextID:        nextID,
		Params:        params,
	}
}
