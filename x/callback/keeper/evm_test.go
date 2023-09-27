package keeper_test

// func (suite *KeeperTestSuite) TestDeployIncoContract() {
// 	suite.SetupTest() // reset

// 	_, err := suite.app.CallbackKeeper.DeployIncoContract(suite.ctx)
// 	suite.Require().NoError(err)
// }

func (suite *KeeperTestSuite) TestCallIncoContract() {
	suite.SetupTest() // reset

	_, err := suite.app.CallbackKeeper.CallEvmAdd(suite.ctx, "", "0xaFe2AC71f883661477ab3c7D9dB2e36A5348eAf1", "add", "5")
	suite.Require().NoError(err)
}
