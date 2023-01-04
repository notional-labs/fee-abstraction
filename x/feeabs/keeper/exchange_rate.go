package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

func (k Keeper) SetNativeTokenExchangeRate(ctx sdk.Context, exchangeRate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: exchangeRate})
	store.Set(types.GetNativeTokenExchangeRateKey(), bz)
}

func (k Keeper) GetNativeTokenExchangeRate(ctx sdk.Context) (sdk.Dec, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetNativeTokenExchangeRateKey())
	if bz == nil {
		return sdk.ZeroDec(), errors.New("No native token exchange rate data")
	}
	var decProto sdk.DecProto
	k.cdc.MustUnmarshal(bz, &decProto)
	return decProto.Dec, nil
}

func (k Keeper) SetIbcTokenExchangeRate(ctx sdk.Context, denom string, exchangeRate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: exchangeRate})
	store.Set(types.GetIbcTokenExchangeRateKey(denom), bz)
}

func (k Keeper) GetIbcTokenExchangeRate(ctx sdk.Context, denom string) (sdk.Dec, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetIbcTokenExchangeRateKey(denom))
	if bz == nil {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidExchangeRate, "Token %s not have exchange rate data", denom)
	}
	var decProto sdk.DecProto
	k.cdc.MustUnmarshal(bz, &decProto)
	return decProto.Dec, nil
}
