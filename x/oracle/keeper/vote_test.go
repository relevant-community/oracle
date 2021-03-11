package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/oracle/x/oracle/types"
)

const claimType = "test"

func (suite *KeeperTestSuite) TestCastVote() {
	ctx := suite.ctx.WithIsCheckTx(false)
	claim := types.NewTestClaim(99, "test", claimType)
	val0 := suite.validators[0]
	val1 := suite.validators[1]

	suite.k.CreateVote(ctx, claim, val0)

	savedClaim := suite.k.GetClaim(ctx, claim.Hash())
	suite.NotNil(savedClaim)

	pending := suite.k.GetPendingRounds(ctx, claimType)
	suite.NotNil(pending)

	roundID := pending[len(pending)-1]
	suite.NotNil(roundID)
	suite.Equal(roundID, claim.GetRoundID())

	round := suite.k.GetRound(ctx, claimType, roundID)
	suite.NotNil(round)

	vote := round.Votes[len(round.Votes)-1]
	suite.NotNil(vote)

	suite.Equal(vote.Validator, val0)
	suite.Equal(vote.ClaimHash, claim.Hash())
	suite.Equal(vote.ClaimType, claim.ClaimType)
	suite.Equal(vote.ConsensusId, claim.GetConcensusKey())
	suite.Equal(vote.RoundId, roundID)

	// Add second vote
	suite.k.CreateVote(ctx, claim, val1)
	round = suite.k.GetRound(ctx, claimType, roundID)
	suite.NotNil(round)
	suite.Equal(len(round.Votes), 2)

	// Test cleanup

	// Remove pending round
	suite.k.FinalizeRound(ctx, claimType, roundID)

	pending = suite.k.GetPendingRounds(ctx, claimType)
	suite.Equal(len(pending), 0)

	// Remove votes + claims
	suite.k.DeleteVotesForRound(ctx, claimType, roundID)
	round = suite.k.GetRound(ctx, claimType, roundID)
	suite.Nil(round)

	savedClaim = suite.k.GetClaim(ctx, claim.Hash())
	suite.Nil(savedClaim)
}

func (suite *KeeperTestSuite) TestVoteTally() {
	ctx := suite.ctx.WithIsCheckTx(false)
	claim := types.NewTestClaim(99, "test", claimType)
	roundID := claim.GetRoundID()

	val0 := suite.validators[0]
	val1 := suite.validators[1]

	suite.k.CreateVote(ctx, claim, val0)

	// Haven't reached threshold
	roundResult := suite.k.TallyVotes(ctx, claimType, roundID)
	suite.Nil(roundResult)

	// Haven't reached threshold (50%)
	suite.k.CreateVote(ctx, claim, val1)
	roundResult = suite.k.TallyVotes(ctx, claimType, roundID)
	suite.NotNil(roundResult)

	totalBondedPower := sdk.TokensToConsensusPower(suite.k.StakingKeeper.TotalBondedTokens(ctx))
	suite.Equal(roundResult.TotalPower, totalBondedPower)

	suite.Equal(roundResult.VotePower, suite.pow[0]+suite.pow[1])

	suite.Equal(len(roundResult.Claims), 1)
	suite.Equal(roundResult.Claims[0].ClaimHash, claim.Hash())
}

// TEST Getters & Setters used in Genesis

func (suite *KeeperTestSuite) TestIterateVotes() {
	ctx := suite.ctx.WithIsCheckTx(false)
	num := 20
	suite.populateVotes(ctx, num)

	votes := suite.k.GetAllRounds(ctx)
	suite.Len(votes, num)
}

func (suite *KeeperTestSuite) populateVotes(ctx sdk.Context, num int) []types.Round {
	rounds := make([]types.Round, num)

	for i := 0; i < num; i++ {
		roundID := uint64(i)

		claim := types.NewTestClaim(int64(roundID), "test", claimType)
		vote := types.NewVote(roundID, claim, suite.validators[0], claimType)
		round := types.Round{
			Votes:     []types.Vote{*vote},
			RoundId:   roundID,
			ClaimType: claimType,
		}
		rounds[i] = round
		suite.k.CreateRound(ctx, round)
	}
	return rounds
}
