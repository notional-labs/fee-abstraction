package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	wasmibctesting "github.com/notional-labs/feeabstraction/v1/x/feeabs/ibctesting"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/keeper"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *wasmibctesting.Coordinator
	chainA      *wasmibctesting.TestChain
	chainB      *wasmibctesting.TestChain
	chainC      *wasmibctesting.TestChain

	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = wasmibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(wasmibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(wasmibctesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(wasmibctesting.GetChainID(2))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), suite.chainA.GetTestSupport().InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerier(suite.chainA.GetTestSupport().FeeAbsKeeper()))
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
