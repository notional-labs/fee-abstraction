package keeper

import (
	"errors"

	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"

	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       sdk.StoreKey
	memKey         sdk.StoreKey
	paramstore     paramtypes.Subspace
	stakingKeeper  types.StakingKeeper
	transferKeeper ibctransferkeeper.Keeper

	// ibc keeper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	scopedKeeper  types.ScopedKeeper
}

// need to implement
func (k Keeper) GetModuleAddress() sdk.AccAddress {
	return sdk.AccAddress{}
}

// need to implement
func (k Keeper) CalculateNativeFromOsmosis(ctx sdk.Context, osmosisAmount sdk.Int) (sdk.Coin, error) {
	// Calculate native token amount to swap for osmosisAmount
	osmosisExchangeRate, err := k.GetOsmosisExchangeRate(ctx)
	if err != nil {
		return sdk.Coin{}, err
	}
	nativeCoinAmount := osmosisExchangeRate.MulInt(osmosisAmount)

	nativeCoin := sdk.Coin{
		Denom:  k.stakingKeeper.BondDenom(ctx),
		Amount: nativeCoinAmount.RoundInt(),
	}

	return nativeCoin, nil
}

// TODO : need to implement
// return err if IBC token isn't in allowed_list
func (k Keeper) verifyIBCCoin(ctx sdk.Context, ibcCoin sdk.Coin) error {
	// Get param
	params := k.GetParams(ctx)
	// Check if ibcCoin not valid
	if ibcCoin.Denom != params.AllowedIbcToken {
		return errors.New("Ibc token not allowed")
	}
	return nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
