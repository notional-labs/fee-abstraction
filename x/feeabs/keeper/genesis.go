package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	return &types.GenesisState{
		Params: params,
	}
}
