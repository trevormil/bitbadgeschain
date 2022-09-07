package types_test

import (
	"testing"

	"github.com/bitbadges/bitbadgeschain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
)

func TestMsgRevokeBadge_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRevokeBadge
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRevokeBadge{
				Creator: "invalid_address",
				SubbadgeRanges: []*types.IdRange{
					{
						Start: 0,
						End:   0,
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgRevokeBadge{
				Creator: sample.AccAddress(),
				SubbadgeRanges: []*types.IdRange{
					{
						Start: 0,
						End:   0,
					},
				},
			},
		},
		{
			name: "invalid subbadge range",
			msg: types.MsgRevokeBadge{
				Creator: sample.AccAddress(),
				SubbadgeRanges: []*types.IdRange{
					{
						Start: 10,
						End:   1,
					},
				},
			},
			err: types.ErrStartGreaterThanEnd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
