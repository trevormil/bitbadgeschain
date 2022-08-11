package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestTransferBadge = "request_transfer_badge"

var _ sdk.Msg = &MsgRequestTransferBadge{}

func NewMsgRequestTransferBadge(creator string, from uint64, amount uint64, badgeId uint64, subbadgeRanges []*IdRange, expirationTime uint64) *MsgRequestTransferBadge {
	return &MsgRequestTransferBadge{
		Creator:        creator,
		From:           from,
		Amount:         amount,
		BadgeId:        badgeId,
		SubbadgeRanges: subbadgeRanges,
		ExpirationTime: expirationTime,
	}
}

func (msg *MsgRequestTransferBadge) Route() string {
	return RouterKey
}

func (msg *MsgRequestTransferBadge) Type() string {
	return TypeMsgRequestTransferBadge
}

func (msg *MsgRequestTransferBadge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestTransferBadge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestTransferBadge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if msg.SubbadgeRanges == nil {
		return ErrRangesIsNil
	}

	for _, subbadgeRange := range msg.SubbadgeRanges {
		if subbadgeRange == nil || subbadgeRange.Start > subbadgeRange.End {
			return ErrStartGreaterThanEnd
		}
	}

	if msg.Amount == 0 {
		return ErrAmountEqualsZero
	}

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
