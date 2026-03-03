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
	Loans  collections.Map[uint64, types.Loan]
	NextID collections.Item[uint64]
	Params collections.Item[types.Params]

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
		types.LoanKeyPrefix,
		"loans",
		collections.Uint64Key,
		codec.CollValue[types.Loan](cdc),
	)

	// Secondary index: borrower -> loanID
	loansByBorrower := collections.NewMap(
		sb,
		types.LoanByBorrowerPrefix,
		"loans_by_borrower",
		collections.PairKeyCodec(collections.StringKey, collections.Uint64Key),
		collections.Uint64Value,
	)

	nextID := collections.NewItem(
		sb,
		types.LoanIDKey,
		"next_id",
		collections.Uint64Value,
	)

	params := collections.NewItem(
		sb,
		types.ParamsKey,
		"params",
		codec.CollValue[types.Params](cdc),
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
