package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeBadge = "revoke_badge"

var _ sdk.Msg = &MsgRevokeBadge{}

func NewMsgRevokeBadge(creator string, addresses []uint64, amounts []uint64, badgeId uint64, subbadgeId uint64) *MsgRevokeBadge {
	return &MsgRevokeBadge{
		Creator:    creator,
		Addresses:    addresses,
		Amounts:     amounts,
		BadgeId:    badgeId,
		SubbadgeId: subbadgeId,
	}
}

func (msg *MsgRevokeBadge) Route() string {
	return RouterKey
}

func (msg *MsgRevokeBadge) Type() string {
	return TypeMsgRevokeBadge
}

func (msg *MsgRevokeBadge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeBadge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeBadge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Amounts) != len(msg.Addresses) {
		return ErrInvalidAmountsAndAddressesLength
	}
	
	for _, amount := range msg.Amounts {
		if amount == 0 {
			return ErrAmountEqualsZero
		}
	}
	return nil
}
