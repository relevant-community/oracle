package keeper_test

import (
	"github.com/relevant-community/oracle/x/oracle/types"
)

func (suite *KeeperTestSuite) TestCreateClaim() {
	ctx := suite.ctx.WithIsCheckTx(false)

	claim := types.NewTestClaim(1, "test", "test")

	suite.k.CreateClaim(ctx, claim)

	res := suite.k.GetClaim(ctx, claim.Hash())
	suite.NotNil(res)
	suite.Equal(claim, res)

	suite.k.DeleteClaim(ctx, claim.Hash())

	deleted := suite.k.GetClaim(ctx, claim.Hash())
	suite.Nil(deleted)
}

func (suite *KeeperTestSuite) TestIterateClaims() {
	ctx := suite.ctx.WithIsCheckTx(false)
	numClaims := 20
	suite.populateClaims(ctx, numClaims)

	claims := suite.k.GetAllClaims(ctx)
	suite.Len(claims, numClaims)
}
