package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/x/loan/types"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLoanByBorrower(ctx sdk.Context, borrower sdk.AccAddress, loanID uint64) {
	store := prefix.NewStore(
		runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)),
		types.LoanByBorrowerPrefix,
	)

	key := append(borrower.Bytes(), sdk.Uint64ToBigEndian(loanID)...)
	store.Set(key, []byte{1})
}

func (k Keeper) GetLoansByBorrower(
	ctx sdk.Context,
	borrower sdk.AccAddress,
) []*types.Loan {

	store := prefix.NewStore(
		runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)),
		types.LoanByBorrowerPrefix,
	)

	iterator := store.Iterator(
		borrower.Bytes(),
		storetypes.PrefixEndBytes(borrower.Bytes()),
	)
	defer iterator.Close()

	var loans []*types.Loan

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		loanID := sdk.BigEndianToUint64(key[len(borrower.Bytes()):])

		loan, err := k.GetLoan(ctx, loanID)
		if err != nil {
			if errors.Is(err, collections.ErrNotFound) {
				continue
			}
			return nil
		}

		loans = append(loans, loan)
	}

	return loans
}

func (k Keeper) HasActiveLoan(ctx sdk.Context, borrower sdk.AccAddress) bool {
	loans := k.GetLoansByBorrower(ctx, borrower)

	for _, loan := range loans {
		if types.IsActiveStatus(loan.Status) {
			return true
		}
	}
	return false
}
