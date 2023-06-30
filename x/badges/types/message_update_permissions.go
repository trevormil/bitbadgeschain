package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateCollectionPermissions = "update_permissions"

var _ sdk.Msg = &MsgUpdateCollectionPermissions{}

func NewMsgUpdateCollectionPermissions(creator string, collectionId sdkmath.Uint, permissions *CollectionPermissions) *MsgUpdateCollectionPermissions {
	return &MsgUpdateCollectionPermissions{
		Creator:      creator,
		CollectionId: collectionId,
		Permissions:  permissions,
	}
}

func (msg *MsgUpdateCollectionPermissions) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCollectionPermissions) Type() string {
	return TypeMsgUpdateCollectionPermissions
}

func (msg *MsgUpdateCollectionPermissions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCollectionPermissions) GetSignBytes() []byte {
	bz := AminoCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCollectionPermissions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = ValidatePermissions(msg.Permissions)
	if err != nil {
		return err
	}

	if msg.CollectionId.IsNil() || msg.CollectionId.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidRequest, "invalid collection id")
	}

	for _, mapping := range msg.AddressMappings {
		if err := ValidateAddressMapping(mapping); err != nil {
			return err
		}
	}

	return nil
}
