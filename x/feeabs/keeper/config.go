package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

func (keeper Keeper) GetHostZoneConfig(ctx sdk.Context, chainId string) (chainConfig types.HostChainFeeAbsConfig, err error) {
	store := ctx.KVStore(keeper.storeKey)
	key := types.GetKeyHostZoneConfig(chainId)
	bz := store.Get(key)

	err = keeper.cdc.Unmarshal(bz, &chainConfig)
	if err != nil {
		return types.HostChainFeeAbsConfig{}, err
	}

	return
}

func (keeper Keeper) SetHostZoneConfig(ctx sdk.Context, chainId string, chainConfig types.HostChainFeeAbsConfig) error {
	store := ctx.KVStore(keeper.storeKey)
	key := types.GetKeyHostZoneConfig(chainId)

	bz, err := keeper.cdc.Marshal(&chainConfig)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

func (keeper Keeper) GetAllHostZoneConfig(ctx sdk.Context, chainId string) (allChainConfigs []types.HostChainFeeAbsConfig, err error) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetKeyHostZoneConfig(chainId))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		var chainConfig types.HostChainFeeAbsConfig
		err := keeper.cdc.Unmarshal(bz, &chainConfig)
		if err != nil {
			panic(err)
		}
		allChainConfigs = append(allChainConfigs, chainConfig)
	}

	return allChainConfigs, nil
}