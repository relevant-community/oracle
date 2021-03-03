package cli_test

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/oracle/x/oracle/client/cli"
	"github.com/relevant-community/oracle/x/oracle/exported"
	"github.com/relevant-community/oracle/x/oracle/types"
	"github.com/spf13/cobra"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

var testClaim = types.NewTestClaim(1, "test", types.TestClaimType)
var testPrevoteClaim = types.NewTestClaim(1, "test", types.TestPrevoteClaimType)

func txHandler(cmd *cobra.Command, txEvent ctypes.ResultEvent) error {
	return nil
}

func initBlockHandler(claim exported.Claim) cli.BlockHandler {
	return func(cmd *cobra.Command, blockEvent ctypes.ResultEvent) error {
		helper, err := cli.NewWorkerHelper(cmd, blockEvent)
		if err != nil {
			return err
		}
		if helper.IsRoundStart(claim.Type()) == false {
			return nil
		}
		helper.SubmitWorkerTx(claim)
		return nil
	}
}

func (s *IntegrationTestSuite) TestWorkerNoPrevoteCmd() {
	val := s.network.Validators[0]
	expectedCode := uint32(0)
	cli.InitializeWorker(initBlockHandler(testClaim), txHandler)
	claimType := types.TestClaimType

	args := []string{
		"1",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	clientCtx := val.ClientCtx.WithNodeURI(val.RPCAddress)
	clientCtx.OutputFormat = "json"

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.StartWorkerCmd(), args)
	s.Require().NoError(err)

	txResp := &sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), txResp), out.String())
	s.Require().Equal(expectedCode, txResp.Code)

	// wait for the worker tx to execute and confirm state transition
	s.network.WaitForHeight(2)

	//////////////////
	//// INTEGRARTION - Test some queries
	//////////////////

	// Claim was created
	res, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdClaim(), []string{testClaim.Hash().String()})

	resType := &types.QueryClaimResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(res.Bytes(), resType))
	resClaim := &types.TestClaim{}
	resClaim.Unmarshal(resType.Claim.Value)
	s.Require().Equal(resClaim.String(), testClaim.String())

	// Pending round was created
	res, err = clitestutil.ExecTestCLICmd(clientCtx, cli.CmdPendingRounds(), []string{claimType})
	resPending := &types.QueryPendingRoundsResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(res.Bytes(), resPending))
	s.Require().Equal(len(resPending.PendingRounds), 1)
	s.Require().Equal(resPending.PendingRounds[0], uint64(1))

	// Query round
	res, err = clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRound(), []string{claimType, "1"})
	resRound := &types.QueryRoundResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(res.Bytes(), resRound))
	s.Require().Equal(len(resRound.Round.Votes), 1)
	s.Require().Equal(resRound.Round.RoundId, uint64(1))

	// Query all rounds
	res, err = clitestutil.ExecTestCLICmd(clientCtx, cli.CmdAllRounds(), []string{})
	resAllRounds := &types.QueryAllRoundsResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(res.Bytes(), resAllRounds))
	s.Require().Equal(len(resAllRounds.Rounds), 1)
}

func (s *IntegrationTestSuite) TestWorkerPrevoteCmd() {
	val := s.network.Validators[0]
	expectedCode := uint32(0)

	// we need to be careful about the claims round and how long we wait for
	// the actual claim to go through
	currentHeight, err := s.network.LatestHeight()

	nextRound := (uint64(currentHeight) + types.TestVotePeriod) - (uint64(currentHeight)+types.TestVotePeriod)%types.TestVotePeriod
	testPrevoteClaim.BlockHeight = int64(nextRound)
	cli.InitializeWorker(initBlockHandler(testPrevoteClaim), txHandler)

	// we need to wait 1 extra round to submit claim
	var endHeight = nextRound + types.TestVotePeriod

	args := []string{
		strconv.FormatInt(int64(endHeight), 10),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	clientCtx := val.ClientCtx.WithNodeURI(val.RPCAddress)
	clientCtx.OutputFormat = "json"

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.StartWorkerCmd(), args)
	s.Require().NoError(err)

	txResp := &sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), txResp), out.String())
	s.Require().Equal(expectedCode, txResp.Code)

	// wait for the worker tx to execute and confirm state transition
	s.network.WaitForHeight(int64(endHeight) + 1)

	// Claim was created
	res, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdClaim(), []string{testPrevoteClaim.Hash().String()})
	resType := &types.QueryClaimResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(res.Bytes(), resType), res.String())
	resClaim := &types.TestClaim{}
	resClaim.Unmarshal(resType.Claim.Value)
	s.Require().Equal(resClaim.String(), testPrevoteClaim.String())

}

//  TODO test edge cases
// func (s *IntegrationTestSuite) TestWorkerEdgeCasesCmd() {
// 	val := s.network.Validators[0]
// 	cli.InitializeWorker(initBlockHandler(testClaim), txHandler)

// 	currentHeight, _ := s.network.LatestHeight()

// 	testCases := map[string]struct {
// 		pre          func()
// 		args         []string
// 		expectErr    bool
// 		respType     proto.Message
// 		expectedCode uint32
// 	}{
// 		"submit claim for prev round": {
// 			func() {
// 				testClaim.BlockHeight = currentHeight - 1
// 			},
// 			[]string{
// 				"1",
// 				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
// 				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
// 				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
// 				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
// 			},
// 			true,
// 			&sdk.TxResponse{},
// 			0,
// 		},
// 	}

// 	for name, tc := range testCases {
// 		tc := tc

// 		s.Run(name, func() {
// 			clientCtx := val.ClientCtx.WithNodeURI(val.RPCAddress)

// 			out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.StartWorkerCmd(), tc.args)

// 			if tc.expectErr {
// 				s.Require().Error(err)
// 			} else {
// 				s.Require().NoError(err)
// 				s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
// 				txResp := tc.respType.(*sdk.TxResponse)
// 				s.Require().Equal(tc.expectedCode, txResp.Code)
// 			}
// 		})
// 	}
// }
