package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/x/loan/types"

	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLoanByBorrower(ctx sdk.Context, borrower sdk.AccAddress, loanID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LoanByBorrowerPrefix)

	key := append(borrower.Bytes(), sdk.Uint64ToBigEndian(loanID)...)
	store.Set(key, []byte{1})
}

func (k Keeper) GetLoansByBorrower(
	ctx sdk.Context,
	borrower sdk.AccAddress,
) []*loanv1.Loan {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LoanByBorrowerPrefix)

	iterator := store.Iterator(
		borrower.Bytes(),
		storetypes.PrefixEndBytes(borrower.Bytes()),
	)
	defer iterator.Close()

	var loans []*loanv1.Loan

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		loanID := sdk.BigEndianToUint64(key[len(borrower.Bytes()):])

		loan, found := k.GetLoan(ctx, loanID)
		if found {
			loans = append(loans, loan)
		}
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
