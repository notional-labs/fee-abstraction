package types

// DefaultGenesisState returns the default genesesis state for
// module.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: &Params{},
	}
}

// ValidateGenesis validates the genesis state
func (gs GenesisState) ValidateGenesis() error {
	return nil
}
