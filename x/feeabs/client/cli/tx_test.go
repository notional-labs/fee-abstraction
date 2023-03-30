package cli

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/notional-labs/feeabstraction/v2/x/feeabs/types"
	"github.com/stretchr/testify/require"
)

func TestParseProposal(t *testing.T) {
	expectedConfig := types.HostChainFeeAbsConfig{
		IbcDenom:                   "ibc/123",
		OsmosisPoolTokenDenomIn:    "ibc/456",
		MiddlewareAddress:          "cosmos123",
		IbcTransferChannel:         "channel-1",
		HostZoneIbcTransferChannel: "channel-2",
		CrosschainSwapAddress:      "osmo123456",
		PoolId:                     1,
		IsOsmosis:                  false,
		Frozen:                     false,
		OsmosisQueryChannel:        "channel-3",
	}
	cdc := codec.NewLegacyAmino()
	okJSON := testutil.WriteToNewTempFile(t, `
{
	"title": "Add Fee Abbtraction Host Zone Proposal",
	"description": "Add Fee Abbtraction Host Zone",
	"host_chain_fee_abs_config": 
		{
			"ibc_denom": "ibc/123",
			"osmosis_pool_token_denom_in": "ibc/456",
			"middleware_address": "cosmos123",
			"ibc_transfer_channel":"channel-1",
			"host_zone_ibc_transfer_channel":"channel-2",
			"crosschain_swap_address":"osmo123456",
			"pool_id": "1",
			"is_osmosis": false,
			"frozen": false,
			"osmosis_query_channel": "channel-3"

		},
	"deposit": "1000stake"
}
	  `)

	proposal, err := ParseAddHostZoneProposalJSON(cdc, okJSON.Name())
	require.NoError(t, err)
	require.Equal(t, "Add Fee Abbtraction Host Zone Proposal", proposal.Title)
	require.Equal(t, "Add Fee Abbtraction Host Zone", proposal.Description)
	require.Equal(t, "1000stake", proposal.Deposit)
	require.Equal(t, expectedConfig, proposal.HostChainFeeAbsConfig)
}
