package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/relevant-community/oracle/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func getValidatorAddr(ctx sdk.Context, k msgServer, signer sdk.AccAddress) sdk.ValAddress {
	// get delegator's validator
	valAddr := sdk.ValAddress(k.GetValidatorAddressFromDelegate(ctx, signer))

	// if there is no delegation it must be the validator
	if valAddr == nil {
		valAddr = sdk.ValAddress(signer)
	}

	return valAddr
}

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claim := msg.GetClaim()
	claimType := claim.Type()
	signer := msg.MustGetSigner()

	valAddr := getValidatorAddr(ctx, k, signer)

	// make sure this message is submitted by a validator
	val := k.StakingKeeper.Validator(ctx, valAddr)

	if val == nil {
		return nil, sdkerrors.Wrap(staking.ErrNoValidatorFound, valAddr.String())
	}

	claimParams := k.ClaimParamsForType(ctx, claimType)
	var claimTypeExists bool
	if claimParams.ClaimType == claimType {
		claimTypeExists = true
	}

	if claimTypeExists != true {
		return nil, sdkerrors.Wrap(types.ErrNoClaimTypeExists, claim.Type())
	}

	var prevoteHash []byte
	if claimParams.Prevote == true {
		// when using prevote claims must be submited within the correct round
		// claim.RoundID == currentRound + roundLength (at least not earlier)
		claimRoundID := claim.GetRoundID()
		currentRound := k.GetCurrentRound(ctx, claimType)

		if claimRoundID+claimParams.VotePeriod != currentRound {
			return nil, sdkerrors.Wrap(types.ErrIncorrectClaimRound, fmt.Sprintf("expected %d, got %d", currentRound-claimParams.VotePeriod, claimRoundID))
		}

		prevoteHash = types.VoteHash(msg.Salt, claim.Hash().String(), signer)
		hasPrevote := k.HasPrevote(ctx, prevoteHash)
		if hasPrevote == false {
			return nil, sdkerrors.Wrap(types.ErrNoPrevote, claim.Hash().String())
		}
	}

	// store the validator vote
	k.CreateVote(ctx, claim, valAddr)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeVote),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyClaimHash, claim.Hash().String()),
		),
	)

	if claimParams.Prevote == true {
		k.DeletePrevote(ctx, prevoteHash)
	}

	return &types.MsgVoteResponse{
		Hash: claim.Hash(),
	}, nil
}

// Delegate implements types.MsgServer
func (k msgServer) Delegate(c context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	val, del := msg.MustGetValidator(), msg.MustGetDelegate()

	if k.Keeper.StakingKeeper.Validator(ctx, sdk.ValAddress(val)) == nil {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	}

	k.SetValidatorDelegateAddress(ctx, val, del)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeDelegate),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyDelegate, msg.Delegate),
		),
	)

	return &types.MsgDelegateResponse{}, nil
}

func (k msgServer) Prevote(goCtx context.Context, msg *types.MsgPrevote) (*types.MsgPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer := msg.MustGetSigner()

	valAddr := getValidatorAddr(ctx, k, signer)

	// make sure this message is submitted by a validator
	val := k.StakingKeeper.Validator(ctx, valAddr)
	if val == nil {
		return nil, sdkerrors.Wrap(staking.ErrNoValidatorFound, valAddr.String())
	}

	k.CreatePrevote(ctx, msg.Hash)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypePrevote),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyPrevoteHash, fmt.Sprintf("%x", msg.Hash)),
		),
	)

	return &types.MsgPrevoteResponse{}, nil
}
