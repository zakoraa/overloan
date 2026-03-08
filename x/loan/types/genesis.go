package types

import (
	"fmt"
)

// DefaultGenesis mengembalikan state awal modul loan
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(), 
		NextId: 1,               
		Loans:  []Loan{},        
	}
}

// Validate melakukan validasi awal sebelum chain start
func (gs *GenesisState) Validate() error {
	if gs == nil {
		return fmt.Errorf("genesis state cannot be nil")
	}

	// Validasi parameter global
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validasi NextId tidak boleh 0
	if gs.NextId == 0 {
		return fmt.Errorf("next_id must start from 1")
	}

	// Validasi tidak ada duplicate loan ID
	seen := make(map[uint64]bool)
	for _, loan := range gs.Loans {
		if seen[loan.Id] {
			return fmt.Errorf("duplicate loan id: %d", loan.Id)
		}
		seen[loan.Id] = true
	}

	return nil
}
