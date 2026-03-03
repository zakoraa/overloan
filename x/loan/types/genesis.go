package types

import (
	"fmt"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *loanv1.GenesisState {
	return &loanv1.GenesisState{
		Params: &loanv1.Params{},
		NextId: 0,
		Loans:  []*loanv1.Loan{},
	}
}

func ValidateGenesis(gs *loanv1.GenesisState) error {
	if gs == nil {
		return fmt.Errorf("genesis state cannot be nil")
	}

	// if gs.Params != nil {
	// 	params := Params(*gs.Params)
	// 	if err := params.Validate(); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
