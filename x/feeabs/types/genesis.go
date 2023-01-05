package types

// DefaultGenesisState returns the default genesesis state for
// module.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: &Params{},
	}
}

// ValidateGenesis validates the genesis state
func ValidateGenesis(genState *GenesisState) error {
	return nil
}
