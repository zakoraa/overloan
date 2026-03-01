package keeper

import (
	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (k Keeper) GetLoan(ctx sdk.Context, id uint64) (*loanv1.Loan, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LoanKeyPrefix)

	bz := store.Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, false
	}

	var loan loanv1.Loan
	k.cdc.MustUnmarshal(bz, &loan)

	return &loan, true
}

func (k Keeper) SetLoan(ctx sdk.Context, loan *loanv1.Loan) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LoanKeyPrefix)

	bz := k.cdc.MustMarshal(loan)
	store.Set(sdk.Uint64ToBigEndian(loan.Id), bz)
}

func (k Keeper) DeleteLoan(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LoanKeyPrefix)
	store.Delete(sdk.Uint64ToBigEndian(id))
}

func (k Keeper) GetNextLoanID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.LoanIDKey)

	var id uint64
	if bz == nil {
		id = 1
	} else {
		id = sdk.BigEndianToUint64(bz) + 1
	}

	store.Set(types.LoanIDKey, sdk.Uint64ToBigEndian(id))
	return id
}
