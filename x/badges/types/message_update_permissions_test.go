package types_test

import (
	"testing"

	"github.com/bitbadges/bitbadgeschain/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
)

func TestMsgUpdatePermissions_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdatePermissions
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdatePermissions{
				Creator: "invalid_address",
				CollectionId: sdk.NewUint(1),
				Permissions: sdk.NewUint(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgUpdatePermissions{
				Creator: sample.AccAddress(),
				CollectionId: sdk.NewUint(1),
				Permissions: sdk.NewUint(0),
			},
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
