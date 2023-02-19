package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	_ govtypes.Content = &AddHostZoneProposal{}
	_ govtypes.Content = &DeleteHostZoneProposal{}
)

const (
	// ProposalTypeAddHostZone defines the type for a AddHostZoneProposal
	ProposalTypeAddHostZone    = "AddHostZone"
	ProposalTypeDeleteHostZone = "DeleteHostZone"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddHostZone)
	govtypes.RegisterProposalType(ProposalTypeDeleteHostZone)
}

// NewClientUpdateProposal creates a new client update proposal.
func NewAddHostZoneProposal(title, description string, config HostChainFeeAbsConfig) govtypes.Content {
	return &AddHostZoneProposal{
		Title:           title,
		Description:     description,
		HostChainConfig: &config,
	}
}

func NewDeleteHostZoneProposal(title, description string, config HostChainFeeAbsConfig) govtypes.Content {
	return &DeleteHostZoneProposal{
		Title:           title,
		Description:     description,
		HostChainConfig: &config,
	}
}

// GetTitle returns the title of a client update proposal.
func (ahzp *AddHostZoneProposal) GetTitle() string { return ahzp.Title }

// GetDescription returns the description of a client update proposal.
func (ahzp *AddHostZoneProposal) GetDescription() string { return ahzp.Description }

// ProposalRoute returns the routing key of a client update proposal.
func (ahzp *AddHostZoneProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a client update proposal.
func (ahzp *AddHostZoneProposal) ProposalType() string { return ProposalTypeAddHostZone }

// ValidateBasic runs basic stateless validity checks
func (ahzp *AddHostZoneProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(ahzp)
	if err != nil {
		return err
	}

	// TODO: add validate here

	return nil
}

// GetTitle returns the title of a client update proposal.
func (dhzp *DeleteHostZoneProposal) GetTitle() string { return dhzp.Title }

// GetDescription returns the description of a client update proposal.
func (dhzp *DeleteHostZoneProposal) GetDescription() string { return dhzp.Description }

// ProposalRoute returns the routing key of a client update proposal.
func (dhzp *DeleteHostZoneProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a client update proposal.
func (dhzp *DeleteHostZoneProposal) ProposalType() string { return ProposalTypeAddHostZone }

// ValidateBasic runs basic stateless validity checks
func (dhzp *DeleteHostZoneProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(dhzp)
	if err != nil {
		return err
	}

	// TODO: add validate here

	return nil
}
