package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	apphelpers "github.com/notional-labs/feeabstraction/v1/app/helpers"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/keeper"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

func createNHostZone(t *testing.T, keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HostChainFeeAbsConfig {
	items := make([]types.HostChainFeeAbsConfig, n)
	for i := range items {
		items[i].IbcDenom = "ibc/" + strconv.Itoa(i)
		items[i].OsmosisPoolTokenDenomIn = "ibc/" + strconv.Itoa(i)
		items[i].MiddlewareAddress = "cosmos123"
		items[i].IbcTransferChannel = "channel-1"
		items[i].HostZoneIbcTransferChannel = "channel-2"
		items[i].CrosschainSwapAddress = "osmo123456"
		items[i].PoolId = 1
		items[i].IsOsmosis = false
		items[i].Frozen = false
		items[i].OsmosisQueryChannel = "channel-3"
		err := keeper.SetHostZoneConfig(ctx, items[i].IbcDenom, items[i])
		require.NoError(t, err)
	}
	return items
}
func TestHostZoneGet(t *testing.T) {
	app := apphelpers.Setup(t, false, 1)
	ctx := apphelpers.NewContextForApp(*app)
	items := createNHostZone(t, &app.FeeabsKeeper, ctx, 10)
	for _, item := range items {
		got, err := app.FeeabsKeeper.GetHostZoneConfig(ctx, item.IbcDenom)
		require.NoError(t, err)
		require.Equal(t, item, got)
	}
}

func TestHostZoneRemove(t *testing.T) {
	app := apphelpers.Setup(t, false, 1)
	ctx := apphelpers.NewContextForApp(*app)
	items := createNHostZone(t, &app.FeeabsKeeper, ctx, 10)
	for _, item := range items {
		err := app.FeeabsKeeper.DeleteHostZoneConfig(ctx, item.IbcDenom)
		require.NoError(t, err)
		got, _ := app.FeeabsKeeper.GetHostZoneConfig(ctx, item.IbcDenom)
		require.NotEqual(t, item, got)
	}
}

func TestHostZoneGetAll(t *testing.T) {
	app := apphelpers.Setup(t, false, 1)
	ctx := apphelpers.NewContextForApp(*app)
	items := createNHostZone(t, &app.FeeabsKeeper, ctx, 10)
	got, _ := app.FeeabsKeeper.GetAllHostZoneConfig(ctx)
	require.ElementsMatch(t, items, got)
}
