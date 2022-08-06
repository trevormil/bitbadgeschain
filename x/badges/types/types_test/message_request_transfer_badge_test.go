package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/trevormil/bitbadgeschain/testutil/sample"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

func TestMsgRequestTransferBadge_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRequestTransferBadge
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRequestTransferBadge{
				Creator: "invalid_address",
				SubbadgeRanges: []*types.NumberRange{
					{
						Start: 0,
						End: 0,
					},
				},
				Amount: 10,

			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgRequestTransferBadge{
				Creator: sample.AccAddress(),
				SubbadgeRanges: []*types.NumberRange{
					{
						Start: 0,
						End: 0,
					},
				},
				Amount: 10,
			},
		}, {
			name: "invalid subbadge range",
			msg: types.MsgRequestTransferBadge{
				Creator: sample.AccAddress(),
				SubbadgeRanges: []*types.NumberRange{
					{
						Start: 10,
						End: 0,
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
