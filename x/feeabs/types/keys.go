package types

const (
	// Module name store the name of the module
	ModuleName = "feeabs"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the feeabs module
	RouterKey = ModuleName

	// Contract: Coin denoms cannot contain this character
	KeySeparator = "|"
)

var (
	OsmosisFeeRate = []byte{0x01} // Key for fee rate of ibc osmosis token
)

// GetNativeTokenExchangeRateKey return the native-token exchange rate key prefix
func GetOsmosisFeeRateKey() (key []byte) {
	return OsmosisFeeRate
}
