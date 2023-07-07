package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateUserApprovedTransfers = "update_user_approved_transfers"

var _ sdk.Msg = &MsgUpdateUserApprovedTransfers{}

func NewMsgUpdateUserApprovedTransfers(creator string) *MsgUpdateUserApprovedTransfers {
	return &MsgUpdateUserApprovedTransfers{
		Creator: creator,
	}
}

func (msg *MsgUpdateUserApprovedTransfers) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserApprovedTransfers) Type() string {
	return TypeMsgUpdateUserApprovedTransfers
}

func (msg *MsgUpdateUserApprovedTransfers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserApprovedTransfers) GetSignBytes() []byte {
	bz := AminoCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserApprovedTransfers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := ValidateUserApprovedIncomingTransferTimeline(msg.ApprovedIncomingTransfersTimeline, msg.Creator); err != nil {
		return err
	}

	if err := ValidateUserApprovedOutgoingTransferTimeline(msg.ApprovedOutgoingTransfersTimeline, msg.Creator); err != nil {
		return err
	}

	if msg.Permissions == nil {
		msg.Permissions = &UserPermissions{}
	}

	err = ValidateUserPermissions(msg.Permissions, true)
	if err != nil {
		return err
	}

	return nil
}
