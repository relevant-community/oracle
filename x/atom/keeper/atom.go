package keeper

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/oracle/x/atom/types"
)

func (k Keeper) SetAtomUsd(ctx sdk.Context, atomUsd types.AtomUsd) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&atomUsd)
	store.Set(types.AtomUsdKey, b)
}

func (k Keeper) GetAtomUsd(ctx sdk.Context) *types.AtomUsd {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AtomUsdKey)

	var atomUsd types.AtomUsd
	if len(bz) == 0 {
		return nil
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &atomUsd)
	return &atomUsd
}

func (k Keeper) UpdateAtomUsd(ctx sdk.Context) {
	claimType := types.AtomClaim
	pending := k.oracleKeeper.GetPendingRounds(ctx, claimType)

	// sort pending rounds in ascending order
	sort.SliceStable(pending, func(i, j int) bool { return pending[i] < pending[j] })

	for _, roundID := range pending {
		result := k.oracleKeeper.TallyVotes(ctx, claimType, roundID)

		if result == nil || result.Claims[0] == nil {
			continue
		}

		// take an average of all claims and commit to chain
		var avgAtomUsd sdk.Dec
		var blockNumber int64
		for i, claimResult := range result.Claims {
			claimHash := claimResult.ClaimHash
			atomClaim, ok := k.oracleKeeper.GetClaim(ctx, claimHash).(*types.AtomUsd)
			if ok == false {
				fmt.Printf("Error retrieving claim")
				continue
			}
			avgAtomUsd = avgAtomUsd.Mul(sdk.NewDec(int64(i - 1))).Add(atomClaim.Price).Quo(sdk.NewDec(int64(i)))
			blockNumber = atomClaim.BlockHeight
		}

		atomUsd := &types.AtomUsd{
			Price:       avgAtomUsd,
			BlockHeight: blockNumber,
		}

		k.SetAtomUsd(ctx, *atomUsd)

		// TODO delete the earlier rounds also
		k.oracleKeeper.FinalizeRound(ctx, claimType, roundID)

		return
	}
}
