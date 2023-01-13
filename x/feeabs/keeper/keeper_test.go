package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/notional-labs/feeabstraction/v1/app/helpers"
)

type KeeperTestSuite struct {
	helpers.AppTestHelper
	suite.Suite
	channelId string
}

// Test helpers
func (suite *KeeperTestSuite) SetupTest() {
	suite.channelId = suite.CreateChannel("fee-abs")
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
