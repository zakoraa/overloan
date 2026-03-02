package keeper

import (
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs loanv1.GenesisState) {

	if gs.Params != nil {
		if err := k.SetParams(ctx, gs.Params); err != nil {
			panic(err)
		}
	}

	if err := k.NextID.Set(ctx, gs.NextId); err != nil {
		panic(err)
	}

	for _, loan := range gs.Loans {
		if err := k.SetLoan(ctx, loan); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) (*loanv1.GenesisState, error) {

	genesis := &loanv1.GenesisState{}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	genesis.Params = &params

	nextID, err := k.NextID.Get(ctx)
	if err != nil {
		return nil, err
	}
	genesis.NextId = nextID

	var loans []*loanv1.Loan

	err = k.Loans.Walk(ctx, nil, func(_ uint64, loan loanv1.Loan) (bool, error) {
		l := loan
		loans = append(loans, &l)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.Loans = loans

	return genesis, nil
}
