package helpers

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/cosmos/ibc-go/v3/testing/simapp"
	feeapp "github.com/notional-labs/feeabstraction/v1/app"
	"github.com/stretchr/testify/suite"
)

// SimAppChainID hardcoded chainID for simulation
const (
	SimAppChainID = "fee-app"
)

type AppTestHelper struct {
	suite.Suite

	App     *feeapp.FeeAbs
	HostApp *simapp.SimApp

	IbcEnabled   bool
	Coordinator  *ibctesting.Coordinator
	FeeAbtractionChain  *ibctesting.TestChain
	HostChain    *ibctesting.TestChain
	TransferPath *ibctesting.Path

	QueryHelper  *baseapp.QueryServiceTestHelper
	TestAccs     []sdk.AccAddress
	IcaAddresses map[string]string
	Ctx          sdk.Context
}

// DefaultConsensusParams defines the default Tendermint consensus params used
// in feeapp testing.
var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

type EmptyAppOptions struct{}

func (EmptyAppOptions) Get(o string) interface{} { return nil }

func Setup(t *testing.T, isCheckTx bool, invCheckPeriod uint) *feeapp.FeeAbs {
	t.Helper()

	app, genesisState := setup(!isCheckTx, invCheckPeriod)
	if !isCheckTx {
		// InitChain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		require.NoError(t, err)

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func setup(withGenesis bool, invCheckPeriod uint) (*feeapp.FeeAbs, feeapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := feeapp.MakeEncodingConfig()
	app := feeapp.NewFeeAbs(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		feeapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		EmptyAppOptions{},
	)
	if withGenesis {
		return app, feeapp.NewDefaultGenesisState()
	}

	return app, feeapp.GenesisState{}
}

// Initializes a ibctesting coordinator to keep track of Stride and a host chain's state
func (s *AppTestHelper) SetupIBCChains(hostChainID string) {
	s.Coordinator = ibctesting.NewCoordinator(s.T(), 0)

	// Initialize a stride testing app by casting a StrideApp -> TestingApp
	ibctesting.DefaultTestingAppInit = InitFeeAbtractionIBCTestingApp
	fmt.Println("testing.T", s.T())
	s.FeeAbtractionChain = ibctesting.NewTestChain(s.T(), s.Coordinator, SimAppChainID)

	// Initialize a host testing app using SimApp -> TestingApp
	ibctesting.DefaultTestingAppInit = ibctesting.SetupTestingApp
	s.HostChain = ibctesting.NewTestChain(s.T(), s.Coordinator, hostChainID)

	// Update coordinator
	s.Coordinator.Chains = map[string]*ibctesting.TestChain{
		SimAppChainID: s.FeeAbtractionChain,
		hostChainID:   s.HostChain,
	}
	s.IbcEnabled = true
}

// Initializes a new Stride App casted as a TestingApp for IBC support
func InitFeeAbtractionIBCTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	app := InitFeeAbtractionTestApp(false)
	return app, feeapp.NewDefaultGenesisState()
}

// Initializes a new StrideApp without IBC functionality
func InitFeeAbtractionTestApp(initChain bool) *feeapp.FeeAbs {
	db := dbm.NewMemDB()
	app := feeapp.NewFeeAbs(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		feeapp.DefaultNodeHome,
		5,
		feeapp.MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
	)
	if initChain {
		genesisState := feeapp.NewDefaultGenesisState()
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

// Creates clients, connections, and a channel between fee-abs and a host chain
func (s *AppTestHelper) CreateChannel(hostChainID string) string {
	// If we have yet to create the host chain, do that here
	s.SetupIBCChains(hostChainID)
	s.Require().Equal(s.HostChain.ChainID, hostChainID,
		"The testing app has already been initialized with a different chainID (%s)", s.HostChain.ChainID)

	// Create clients, connections, and a transfer channel
	s.TransferPath = NewTransferPath(s.FeeAbtractionChain, s.HostChain)
	s.Coordinator.Setup(s.TransferPath)

	// Replace stride and host apps with those from TestingApp
	s.App = s.FeeAbtractionChain.App.(*feeapp.FeeAbs)
	s.HostApp = s.HostChain.GetSimApp()
	s.Ctx = s.FeeAbtractionChain.GetContext()
	// Finally confirm the channel was setup properly
	s.Require().Equal(ibctesting.FirstClientID, s.TransferPath.EndpointA.ClientID, "fee-abs clientID")
	s.Require().Equal(ibctesting.FirstConnectionID, s.TransferPath.EndpointA.ConnectionID, "fee-abs connectionID")
	s.Require().Equal(ibctesting.FirstChannelID, s.TransferPath.EndpointA.ChannelID, "fee-abs transfer channelID")

	s.Require().Equal(ibctesting.FirstClientID, s.TransferPath.EndpointB.ClientID, "host clientID")
	s.Require().Equal(ibctesting.FirstConnectionID, s.TransferPath.EndpointB.ConnectionID, "host connectionID")
	s.Require().Equal(ibctesting.FirstChannelID, s.TransferPath.EndpointB.ChannelID, "host transfer channelID")

	return s.TransferPath.EndpointA.ChannelID
}

func NewTransferPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = transfertypes.Version
	path.EndpointB.ChannelConfig.Version = transfertypes.Version

	return path
}
