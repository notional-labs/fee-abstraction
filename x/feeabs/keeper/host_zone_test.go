package keeper_test

// import (
// 	"strconv"

// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/notional-labs/feeabstraction/v1/x/feeabs/keeper"
// 	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
// )

// func createNHostZone(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HostChainFeeAbsConfig {
// 	items := make([]types.HostChainFeeAbsConfig, n)
// 	for i := range items {
// 		items[i].IbcDenom = strconv.Itoa(i)
// 		items[i].HostChainNativeDenomIbcedOnOsmosis = strconv.Itoa(i)
// 		items[i].MiddlewareAddress = "cosmos123"
// 		items[i].IbcTransferChannel = "channel-1"
// 		items[i].HostZoneIbcTransferChannel = "channel-2"
// 		items[i].CrosschainSwapAddress = "osmo123456"
// 		items[i].PoolId = 1
// 		items[i].IsOsmosis = false
// 		items[i].Frozen = false
// 		keeper.SetHostZoneConfig(ctx, "ibc/123", items[i])
// 	}
// 	return items
// }
