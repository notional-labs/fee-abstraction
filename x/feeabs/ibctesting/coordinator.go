package ibctesting

import (
	"testing"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const ChainIDPrefix = "testchain"

var (
	globalStartTime = time.Date(2020, 12, 4, 10, 30, 0, 0, time.UTC)
	TimeIncrement   = time.Second * 5
)

// Coordinator is a testing struct which contains N TestChain's. It handles keeping all chains
// in sync with regards to time.
type Coordinator struct {
	t *testing.T

	CurrentTime time.Time
	Chains      map[string]*TestChain
}

// NewCoordinator initializes Coordinator with N TestChain's
func NewCoordinator(t *testing.T, n int, opts ...[]wasmkeeper.Option) *Coordinator {
	chains := make(map[string]*TestChain)
	coord := &Coordinator{
		t:           t,
		CurrentTime: globalStartTime,
	}

	for i := 0; i < n; i++ {
		chainID := GetChainID(i)
		var x []wasmkeeper.Option
		if len(opts) > i {
			x = opts[i]
		}
		chains[chainID] = NewTestChain(t, coord, chainID, x...)
	}
	coord.Chains = chains

	return coord
}
