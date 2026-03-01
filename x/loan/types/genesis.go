package types

type GenesisState struct{}

func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

