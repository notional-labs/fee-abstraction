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
	NativeTokenExchangeRateKeyPrefix = []byte{0x01} // Key prefix for native token exchange rate
	IbcTokenExchangeRateKeyPrefix    = []byte{0x02} // Key prefix for ibc token exchange rate
)

// GetNativeTokenExchangeRateKey return the native-token exchange rate key prefix
func GetNativeTokenExchangeRateKey() (key []byte) {
	key = append(key, NativeTokenExchangeRateKeyPrefix...)
	key = append(key, []byte(KeySeparator)...)
	return key
}

// GetIbcTokenFeeExchangeRateKey return the ibc-token exchange rate key prefix
func GetIbcTokenFeeExchangeRateKey(denom string) (key []byte) {
	key = append(key, IbcTokenExchangeRateKeyPrefix...)
	key = append(key, []byte(denom)...)
	key = append(key, []byte(KeySeparator)...)
	return key
}
