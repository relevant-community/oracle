package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/oracle/x/oracle/types"
)

// CreateRound creates a Round (used in genesis file)
func (k Keeper) CreateRound(ctx sdk.Context, round types.Round) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RoundKey)
	store.Set(types.RoundPrefix(round.ClaimType, round.RoundId), k.cdc.MustMarshalBinaryBare(&round))
}

// GetRound retrieves a Round that contains all Votes for a claimType and roundID
func (k Keeper) GetRound(ctx sdk.Context, claimType string, roundID uint64) *types.Round {
	roundKey := types.GetRoundKey(claimType, roundID)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RoundKey)
	var round types.Round
	bz := store.Get(types.KeyPrefix(roundKey))
	if len(bz) == 0 {
		return nil
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &round)
	return &round
}

// GetAllRounds retrieves all the Rounds (used in genesis)
func (k Keeper) GetAllRounds(ctx sdk.Context) (rounds []types.Round) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RoundKey)
	defer iterator.Close()

	rounds = []types.Round{}
	for ; iterator.Valid(); iterator.Next() {
		var round types.Round
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &round)
		rounds = append(rounds, round)
	}
	return
}

// PENDING

// AddPendingRound adds the roundId to the pending que
func (k Keeper) AddPendingRound(ctx sdk.Context, claimType string, roundID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PendingRoundKey)
	bz := []byte(strconv.FormatUint(roundID, 10))
	store.Set(types.RoundPrefix(claimType, roundID), bz)
}

// GetPendingRounds returns an array of pending rounds for a given claimType
func (k Keeper) GetPendingRounds(ctx sdk.Context, claimType string) (rounds []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PendingRoundKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(claimType))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		round, err := strconv.ParseUint(string(iterator.Value()), 10, 64)
		if err != nil {
			// Panic because the count should be always formattable to uint64
			panic("cannot decode count")
		}
		rounds = append(rounds, round)
	}
	return
}

// GetAllPendingRounds returns all pending rounds
func (k Keeper) GetAllPendingRounds(ctx sdk.Context) (allPendingRounds map[string]([]uint64)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PendingRoundKey)

	defer iterator.Close()
	params := k.GetParams(ctx)

	allPendingRounds = map[string][]uint64{}
	for _, param := range params.ClaimParams {
		pending := k.GetPendingRounds(ctx, param.ClaimType)
		allPendingRounds[param.ClaimType] = pending
	}

	return allPendingRounds
}

// DeletePendingRound deletes the roundKey from the store
func (k Keeper) DeletePendingRound(ctx sdk.Context, claimType string, roundID uint64) {
	roundKey := types.GetRoundKey(claimType, roundID)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PendingRoundKey)
	store.Delete(types.KeyPrefix(roundKey))
}

// GetCurrentRound returns the current vote round
func (k Keeper) GetCurrentRound(ctx sdk.Context, claimType string) uint64 {
	claimParams := k.ClaimParamsForType(ctx, claimType)
	block := ctx.BlockHeight()
	return uint64(block) - uint64(block)%claimParams.VotePeriod
}
