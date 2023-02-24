package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RegisterAddresses(goCtx context.Context, msg *types.MsgRegisterAddresses) (*types.MsgRegisterAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	CreatorAccountNum := k.Keeper.MustGetAccountNumberForBech32AddressString(ctx, msg.Creator)

	start := uint64(0)
	end := uint64(0)

	addressNums := make([]uint64, len(msg.AddressesToRegister))
	for i, address := range msg.AddressesToRegister {
		convertedAddress, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
		}

		newNum := k.Keeper.GetOrCreateAccountNumberForAccAddressBech32(ctx, convertedAddress) //This panics but is saved
		if i == 0 {
			start = newNum
		}
		end = newNum

		addressNums[i] = newNum
	}

	addressesJson, err := json.Marshal(addressNums)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "badges"),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute("addressNums", string(addressesJson)),
			sdk.NewAttribute("creator", fmt.Sprint(CreatorAccountNum)),
		),
	)

	return &types.MsgRegisterAddressesResponse{
		RegisteredAddressNumbers: CreateIdRange(start, end),
	}, nil
}
