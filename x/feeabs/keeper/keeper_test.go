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

	// ctx          sdk.Context
	// feeAbsApp    *app.FeeAbs
	// feeAbsKeeper keeper.Keeper
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

// func (suite *KeeperTestSuite) SetupTest() {
// 	suite.feeAbsApp = apphelpers.Setup(suite.T(), false, 1)
// 	suite.ctx = suite.feeAbsApp.BaseApp.NewContext(false, tmproto.Header{
// 		ChainID: fmt.Sprintf("test-chain-%s", tmrand.Str(4)),
// 		Height:  1,
// 	})
// 	suite.feeAbsKeeper = suite.feeAbsApp.FeeabsKeeper

// 	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.feeAbsApp.InterfaceRegistry())
// 	types.RegisterQueryServer(queryHelper, keeper.NewQuerier(suite.feeAbsKeeper))
// 	suite.queryClient = types.NewQueryClient(queryHelper)

// 	suite.msgServer = keeper.NewMsgServerImpl(suite.feeAbsKeeper)
// }

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
