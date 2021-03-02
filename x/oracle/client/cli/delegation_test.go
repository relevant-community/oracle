package cli_test

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/oracle/x/oracle/client/cli"
	"github.com/relevant-community/oracle/x/oracle/types"
)

func (s *IntegrationTestSuite) TestDelegationCmd() {
	val := s.network.Validators[0]

	del, err := sdk.AccAddressFromBech32("cosmos1cxlt8kznps92fwu3j6npahx4mjfutydyene2qw")
	s.Require().NoError(err)

	clientCtx := val.ClientCtx.WithNodeURI(val.RPCAddress)
	clientCtx.OutputFormat = "json"

	args := []string{
		del.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.TxDelegate(), args)
	s.Require().NoError(err)

	txRes := &sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), txRes), out.String())
	s.Require().Equal(uint32(0), txRes.Code)

	args = []string{
		fmt.Sprintf("%s", val.Address.String()),
	}

	out, err = clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDelegeateAddress(), args)
	delRes := &types.QueryDelegeateAddressResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), delRes), out.String())
	s.Require().Equal(del.String(), delRes.Delegate)

	args = []string{
		fmt.Sprintf("%s", delRes.Delegate),
	}

	out, err = clitestutil.ExecTestCLICmd(clientCtx, cli.CmdValidatorAddress(), args)
	valRes := &types.QueryValidatorAddressResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), valRes), out.String())

	s.Require().Equal(val.Address.String(), valRes.Validator)

	// undo delegation
	args = []string{
		val.Address.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	out, err = clitestutil.ExecTestCLICmd(clientCtx, cli.TxDelegate(), args)
	s.Require().NoError(err)

	txRes = &sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), txRes), out.String())
	s.Require().Equal(uint32(0), txRes.Code)

	// test undo delegation
	args = []string{
		fmt.Sprintf("%s", val.Address.String()),
	}
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDelegeateAddress(), args)
	s.Require().Contains(out.String(), "NotFound")
}
