package keeper

import (
	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLoanByBorrower(
	ctx sdk.Context,
	borrower sdk.AccAddress,
	loanID uint64,
) error {

	return k.LoansByBorrower.Set(
		ctx,
		collections.Join(borrower, loanID),
		loanID,
	)
}

func (k Keeper) GetLoansByBorrower(
	ctx sdk.Context,
	borrower sdk.AccAddress,
) ([]*types.Loan, error) {

	rng := collections.NewPrefixedPairRange[sdk.AccAddress, uint64](borrower)

	iter, err := k.LoansByBorrower.Iterate(ctx, rng)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var loans []*types.Loan

	for ; iter.Valid(); iter.Next() {

		kv, err := iter.KeyValue()
		if err != nil {
			continue
		}

		loanID := kv.Key.K2()

		loan, err := k.Loans.Get(ctx, loanID)
		if err != nil {
			continue
		}

		loans = append(loans, &loan)
	}

	return loans, nil
}

func (k Keeper) HasActiveLoan(ctx sdk.Context, borrower sdk.AccAddress) bool {

	loans, err := k.GetLoansByBorrower(ctx, borrower)
	if err != nil {
		return false
	}

	for _, loan := range loans {
		if types.IsActiveStatus(loan.Status) {
			return true
		}
	}

	return false
}
