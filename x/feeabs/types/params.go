package types

import (
	"fmt"
	"strings"
	time "time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Feeabs params default values .
const (
	DefaultSwapPeriod time.Duration = time.Minute * 100

	DefaultQueryPeriod time.Duration = time.Minute * 1

	DefaultContractAddress string = ""
)

// Parameter keys store keys.
var (
	KeyOsmosisIbcDenom          = []byte("osmosis_ibc_denom")
	KeyOsmosisIbcConnectionId   = []byte("osmosis_ibc_connection_id")
	KeyOsmosisQueryContract     = []byte("osmosis_query_contract")
	KeyFeeRateUpdatePeriod      = []byte("fee_rate_update_period")
	KeyAccumulatedFeeSwapPeriod = []byte("accumulated_fee_swap_period")

	_ paramtypes.ParamSet = &Params{}
)

// ParamTable for lockup module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOsmosisIbcDenom, &p.OsmosisIbcDenom, validateOsmosisIbcDenom),
		paramtypes.NewParamSetPair(KeyOsmosisIbcConnectionId, &p.OsmosisIbcConnectionId, validateIbcConnectionId),
		paramtypes.NewParamSetPair(KeyOsmosisQueryContract, &p.OsmosisQueryContract, validateOsmosisQueryContract),
		paramtypes.NewParamSetPair(KeyFeeRateUpdatePeriod, &p.FeeRateUpdatePeriod, noOp),
		paramtypes.NewParamSetPair(KeyAccumulatedFeeSwapPeriod, &p.AccumulatedFeeSwapPeriod, noOp),
	}
}

func noOp(i interface{}) error {
	return nil
}

func validateOsmosisIbcDenom(i interface{}) error {
	denom, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.HasPrefix(denom, "ibc/") {
		return fmt.Errorf("osmosis ibc denom doesn't have ibc prefix")
	}

	return nil
}

func validateIbcConnectionId(i interface{}) error {
	connectionId, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !strings.HasPrefix(connectionId, "connection-") {
		return fmt.Errorf("wrong connection id format")
	}

	return nil
}

func validateOsmosisQueryContract(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
