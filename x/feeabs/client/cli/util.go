package cli

import (
	"io/ioutil" //nolint:staticcheck

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

// ParseParamChangeProposalJSON reads and parses a ParamChangeProposalJSON from
// file.
func ParseAddHostZoneProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (AddHostZoneProposalJSON, error) {
	proposal := AddHostZoneProposalJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

type (
	AddHostZoneProposalJSON struct {
		Title                 string                      `json:"title" yaml:"title"`
		Description           string                      `json:"description" yaml:"description"`
		HostChainFeeAbsConfig types.HostChainFeeAbsConfig `json:"host_chain_fee_abs_config" yaml:"host_chain_fee_abs_config"`
		Deposit               string                      `json:"deposit" yaml:"deposit"`
	}
)
