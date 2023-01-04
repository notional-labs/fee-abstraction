package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

// SetIbcTokenExchangeRate set native-token/ibc-token pair exchange rate
func (k Keeper) SetIbcTokenFeeExchangeRate(ctx sdk.Context, denom string, exchangeRate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: exchangeRate})
	store.Set(types.GetIbcTokenFeeExchangeRateKey(denom), bz)
}

// GetIbcTokenExchangeRate get native-token/ibc-token pair exchange rate
func (k Keeper) GetIbcTokenFeeExchangeRate(ctx sdk.Context, denom string) (sdk.Dec, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetIbcTokenFeeExchangeRateKey(denom))
	if bz == nil {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidExchangeRate, "Token %s not have exchange rate data", denom)
	}
	var decProto sdk.DecProto
	k.cdc.MustUnmarshal(bz, &decProto)
	return decProto.Dec, nil
}
