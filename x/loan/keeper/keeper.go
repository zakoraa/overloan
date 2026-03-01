/*
Package keeper berisi implementasi logika bisnis dan
pengelolaan state untuk modul loan.

Package ini menangani lifecycle pinjaman,
validasi otoritas, akuntansi token settlement,
*/
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	storetypes "cosmossdk.io/store/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

// Keeper mengelola akses state dan dependency modul loan
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	paramSpace paramtypes.Subspace

	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	groupKeeper   types.GroupKeeper

	authority string
}

// NewKeeper membuat instance keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	groupKeeper types.GroupKeeper,
	authority string,
) Keeper {

	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramSpace:    ps,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		groupKeeper:   groupKeeper,
		authority:     authority,
	}
}
