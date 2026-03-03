package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {

	if err := k.SetParams(ctx, gs.Params); err != nil {
		panic(err)
	}

	if err := k.NextID.Set(ctx, gs.NextId); err != nil {
		panic(err)
	}

	for i := range gs.Loans {
		loan := gs.Loans[i]
		if err := k.SetLoan(ctx, &loan); err != nil {
			panic(err)
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (*types.GenesisState, error) {

	var genesis types.GenesisState

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	genesis.Params = params

	nextID, err := k.NextID.Get(ctx)
	if err != nil {
		return nil, err
	}
	genesis.NextId = nextID

	var loans []types.Loan

	err = k.Loans.Walk(ctx, nil, func(_ uint64, loan types.Loan) (bool, error) {
		loans = append(loans, loan)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.Loans = loans

	return &genesis, nil
}
