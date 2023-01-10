package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Module name store the name of the module
	ModuleName = "feeabs"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the feeabs module
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_feeabs"

	// Contract: Coin denoms cannot contain this character
	KeySeparator = "|"
)

var (
	OsmosisExchangeRate = []byte{0x01} // Key for the exchange rate of osmosis (to native token)
	KeyChannelID        = []byte{0x02} // Key for IBC channel to osmosis
	// KeyPrefixIndexCapability defines a key prefix that stores index to capability
	// owners mappings.
	KeyPrefixIndexCapability = []byte("capability_index")

	// KeyMemInitialized defines the key that stores the initialized flag in the memory store
	KeyMemInitialized = []byte("mem_initialized")
)

// GetOsmosisExchangeRateKey return the key for set/getting the exchange rate of osmosis (to native token)
func GetOsmosisExchangeRateKey() (key []byte) {
	return OsmosisExchangeRate
}

// IndexFromKey returns an index from a call to IndexToKey for a given capability
// index.
func IndexFromKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key)
}

// FwdCapabilityKey returns a forward lookup key for a given module and capability
// reference.
func FwdCapabilityKey(module string, cap *Capability) []byte {
	return []byte(fmt.Sprintf("%s/fwd/%#016p", module, cap))
}

// RevCapabilityKey returns a reverse lookup key for a given module and capability
// name.
func RevCapabilityKey(module, name string) []byte {
	return []byte(fmt.Sprintf("%s/rev/%s", module, name))
}
