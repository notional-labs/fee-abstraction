package types

import (
	"strings"
)

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

func GetNativeTokenExchangeRateKey() (key []byte) {
	key = append(key, NativeTokenExchangeRateKeyPrefix...)
	key = append(key, []byte(KeySeparator)...)
	return key
}

func GetIbcTokenExchangeRateKey(denom string) (key []byte) {
	upperDenom := strings.ToUpper(denom)
	key = append(key, IbcTokenExchangeRateKeyPrefix...)
	key = append(key, []byte(upperDenom)...)
	key = append(key, []byte(KeySeparator)...)
	return key
}
