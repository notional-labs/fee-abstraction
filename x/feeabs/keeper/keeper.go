package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       sdk.StoreKey
	paramstore     paramtypes.Subspace
	transferKeeper ibctransferkeeper.Keeper
	// capability
	memKey         storetypes.StoreKey
	capMap         map[uint64]*types.Capability
	// ibc keeper
	portKeeper    types.PortKeeper
	channelKeeper types.ChannelKeeper
	scopedKeeper  types.ScopedKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	transferKeeper ibctransferkeeper.Keeper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper types.ScopedKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		paramstore:     ps,
		transferKeeper: transferKeeper,
		channelKeeper:  channelKeeper,
		scopedKeeper:   scopedKeeper,
		portKeeper:     portKeeper,
	}
}

// need to implement
func (k Keeper) GetModuleAddress() sdk.AccAddress {
	return sdk.AccAddress{}
}

// need to implement
func (k Keeper) CalculateNativeFromIBCCoin(ibcCoin sdk.Coins) (coins sdk.Coins, err error) {
	err = k.verifyIBCCoin(ibcCoin)
	if err != nil {
		return sdk.Coins{}, nil
	}
	return coins, nil
}

// TODO : need to implement
// return err if IBC token isn't in allowed_list
func (k Keeper) verifyIBCCoin(ibcCoin sdk.Coins) error {
	return nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// InitMemStore will assure that the module store is a memory store (it will panic if it's not)
// and willl initialize it. The function is safe to be called multiple times.
// InitMemStore must be called every time the app starts before the keeper is used (so
// `BeginBlock` or `InitChain` - whichever is first). We need access to the store so we
// can't initialize it in a constructor.
func (k *Keeper) InitMemStore(ctx sdk.Context) {
	memStore := ctx.KVStore(k.memKey)
	memStoreType := memStore.GetStoreType()

	if memStoreType != storetypes.StoreTypeMemory {
		panic(fmt.Sprintf("invalid memory store type; got %s, expected: %s", memStoreType, storetypes.StoreTypeMemory))
	}

	// create context with no block gas meter to ensure we do not consume gas during local initialization logic.
	noGasCtx := ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())

	// check if memory store has not been initialized yet by checking if initialized flag is nil.
	if !k.IsInitialized(noGasCtx) {
		prefixStore := prefix.NewStore(noGasCtx.KVStore(k.storeKey), types.KeyPrefixIndexCapability)
		iterator := sdk.KVStorePrefixIterator(prefixStore, nil)

		// initialize the in-memory store for all persisted capabilities
		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			index := types.IndexFromKey(iterator.Key())

			var capOwners types.CapabilityOwners

			k.cdc.MustUnmarshal(iterator.Value(), &capOwners)
			k.InitializeCapability(noGasCtx, index, capOwners)
		}

		// set the initialized flag so we don't rerun initialization logic
		memStore := noGasCtx.KVStore(k.memKey)
		memStore.Set(types.KeyMemInitialized, []byte{1})
	}
}

// IsInitialized returns true if the keeper is properly initialized, and false otherwise.
func (k *Keeper) IsInitialized(ctx sdk.Context) bool {
	memStore := ctx.KVStore(k.memKey)
	return memStore.Get(types.KeyMemInitialized) != nil
}

// InitializeCapability takes in an index and an owners array. It creates the capability in memory
// and sets the fwd and reverse keys for each owner in the memstore.
// It is used during initialization from genesis.
func (k Keeper) InitializeCapability(ctx sdk.Context, index uint64, owners types.CapabilityOwners) {
	memStore := ctx.KVStore(k.memKey)

	cap := types.NewCapability(index)
	for _, owner := range owners.Owners {
		// Set the forward mapping between the module and capability tuple and the
		// capability name in the memKVStore
		memStore.Set(types.FwdCapabilityKey(owner.Module, cap), []byte(owner.Name))

		// Set the reverse mapping between the module and capability name and the
		// index in the in-memory store. Since marshalling and unmarshalling into a store
		// will change memory address of capability, we simply store index as value here
		// and retrieve the in-memory pointer to the capability from our map
		memStore.Set(types.RevCapabilityKey(owner.Module, owner.Name), sdk.Uint64ToBigEndian(index))

		// Set the mapping from index from index to in-memory capability in the go map
		k.capMap[index] = cap
	}
}

