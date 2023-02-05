package keeper_test

import (
	"encoding/json"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/keeper"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
	"github.com/stretchr/testify/require"
)

// TODO: need to refactor this test, use driven table
func TestParseMsgToMemo(t *testing.T) {
	twapRouter := types.TwapRouter{
		SlippagePercentage: "20",
		WindowSeconds:      10,
	}

	swap := types.Swap{
		InputCoin:   sdk.NewCoin("khanhyeungan", sdk.NewInt(123)),
		OutPutDenom: "khanhyeuchau",
		Slippage:    types.Twap{Twap: twapRouter},
		Receiver:    "123456",
	}

	msgSwap := types.OsmosisSwapMsg{
		OsmosisSwap: swap,
	}

	mockAddress := "cosmos123456789"
	mockReceiver := "cosmos123456789"

	//TODO: need to check assert msg
	memo, err := keeper.ParseMsgToMemo(msgSwap, mockAddress, mockReceiver)
	require.NoError(t, err)

	var memoOsmosis types.OsmosisSpecialMemo
	_ = json.Unmarshal([]byte(memo), &memoOsmosis)
	fmt.Printf("%v\n", memoOsmosis.Wasm)
	fmt.Printf("%v\n", memoOsmosis.Wasm["msg"])

	require.True(t, false)
}
