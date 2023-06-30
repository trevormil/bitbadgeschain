package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateUserPermissions = "update_user_permissions"

var _ sdk.Msg = &MsgUpdateUserPermissions{}

func NewMsgUpdateUserPermissions(creator string, collectionId sdkmath.Uint, permissions *UserPermissions) *MsgUpdateUserPermissions {
	return &MsgUpdateUserPermissions{
		Creator: creator,
		CollectionId: collectionId,
		Permissions: permissions,
	}
}

func (msg *MsgUpdateUserPermissions) Route() string {
  return RouterKey
}

func (msg *MsgUpdateUserPermissions) Type() string {
  return TypeMsgUpdateUserPermissions
}

func (msg *MsgUpdateUserPermissions) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserPermissions) GetSignBytes() []byte {
  bz := AminoCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserPermissions) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.CollectionId.IsNil() || msg.CollectionId.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidRequest, "invalid collection id")
	}

	if msg.Permissions == nil {
		return sdkerrors.Wrapf(ErrInvalidPermissions, "invalid permissions (%s)", msg.Permissions)
	}

	err = ValidateUserPermissions(msg.Permissions, true)
	if err != nil {
		return err
	}

	for _, mapping := range msg.AddressMappings {
		if err := ValidateAddressMapping(mapping); err != nil {
			return err
		}
	}

  return nil
}

