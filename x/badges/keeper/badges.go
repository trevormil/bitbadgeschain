package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

//Gets badge and throws error if it does not exist. Alternative to GetBadgeFromStore which returns a found bool, not an error.
func (k Keeper) GetBadgeE(ctx sdk.Context, badgeId uint64) (types.BitBadge, error) {
	badge, found := k.GetBadgeFromStore(ctx, badgeId)
	if !found {
		return types.BitBadge{}, ErrBadgeNotExists
	}

	return badge, nil
}


// Gets the badge details from the store if it exists. Throws error if subbadge ranges are invalid or the subbadge does not yet exist.
func (k Keeper) GetBadgeAndAssertSubbadgeRangesAreValid(ctx sdk.Context, badgeId uint64, subbadgeRanges []*types.IdRange) (types.BitBadge, error) {
	badge, err := k.GetBadgeE(ctx, badgeId)
	if err != nil {
		return badge, err
	}

	for _, subbadgeRange := range subbadgeRanges {
		// Subbadge ranges can set end == 0 to save storage space. By convention, this means end == start
		if subbadgeRange.End == 0 {
			subbadgeRange.End = subbadgeRange.Start
		} 

		if subbadgeRange.Start > subbadgeRange.End {
			return types.BitBadge{}, ErrInvalidSubbadgeRange
		}

		if subbadgeRange.End >= badge.NextSubassetId {
			return types.BitBadge{}, ErrSubBadgeNotExists
		}
	}

	return badge, nil
}
