package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetLoan(ctx sdk.Context, id uint64) (*loanv1.Loan, error) {
	loan, err := k.Loans.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (k Keeper) SetLoan(ctx sdk.Context, loan *loanv1.Loan) error {
	err := k.Loans.Set(ctx, loan.Id, *loan)

	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) DeleteLoan(ctx sdk.Context, id uint64) error {
	err := k.Loans.Remove(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetNextLoanID(ctx sdk.Context) (uint64, error) {
	id, err := k.NextID.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			id = 0
		} else {
			return 0, err
		}
	}

	id++

	if err := k.NextID.Set(ctx, id); err != nil {
		return 0, err
	}

	return id, nil
}
