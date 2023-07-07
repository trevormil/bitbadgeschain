package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgTransferBadges{}, "badges/TransferBadges", nil)
	cdc.RegisterConcrete(&MsgDeleteCollection{}, "badges/DeleteCollection", nil)
	cdc.RegisterConcrete(&MsgUpdateUserApprovedTransfers{}, "badges/UpdateUserApprovedTransfers", nil)
	cdc.RegisterConcrete(&MsgUpdateCollection{}, "badges/UpdateCollection", nil)
	cdc.RegisterConcrete(&MsgCreateAddressMappings{}, "badges/CreateAddressMappings", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferBadges{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateUserApprovedTransfers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAddressMappings{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
	// AminoCdc is a amino codec created to support amino JSON compatible msgs.
	AminoCdc = codec.NewAminoCodec(Amino)
)
