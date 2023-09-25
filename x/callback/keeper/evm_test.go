package keeper_test

// func (suite *KeeperTestSuite) TestDeployIncoContract() {
// 	suite.SetupTest() // reset

// 	_, err := suite.app.CallbackKeeper.DeployIncoContract(suite.ctx)
// 	suite.Require().NoError(err)
// }

func (suite *KeeperTestSuite) TestCallIncoContract() {
	suite.SetupTest() // reset

	_, err := suite.app.CallbackKeeper.CallEvmAdd(suite.ctx, "", "0x0aF24F6e261F91836A57Eca0583DA9D3361253fB", "add", "5")
	suite.Require().NoError(err)
}
