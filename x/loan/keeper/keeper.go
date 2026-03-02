/*
Package keeper berisi implementasi logika bisnis dan
pengelolaan state untuk modul loan.

Package ini menangani lifecycle pinjaman,
validasi otoritas, akuntansi token settlement.
*/
package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService

	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper

	authority string

	Schema collections.Schema

	// Primary storage
	Loans  collections.Map[uint64, loanv1.Loan]
	NextID collections.Item[uint64]
	Params collections.Item[loanv1.Params]

	// Secondary index
	LoansByBorrower collections.Map[
		collections.Pair[string, uint64],
		uint64,
	]
}

// NewKeeper membuat instance keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	authority string,
) Keeper {

	sb := collections.NewSchemaBuilder(storeService)

	// Primary map: loanID -> Loan
	loans := collections.NewMap(
		sb,
		collections.NewPrefix("loans"),
		"loans",
		collections.Uint64Key,
		codec.CollValue[loanv1.Loan](cdc),
	)

	// Secondary index: borrower -> loanID
	loansByBorrower := collections.NewMap(
		sb,
		collections.NewPrefix("loans_by_borrower"),
		"loans_by_borrower",
		collections.PairKeyCodec(collections.StringKey, collections.Uint64Key),
		collections.Uint64Value,
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
		cdc:             cdc,
		storeService:    storeService,
		bankKeeper:      bankKeeper,
		accountKeeper:   accountKeeper,
		authority:       authority,
		Schema:          schema,
		Loans:           loans,
		LoansByBorrower: loansByBorrower,
		NextID:          nextID,
		Params:          params,
	}
}
