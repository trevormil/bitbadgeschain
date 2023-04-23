package cli

import (
	"encoding/json"
	"strconv"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMintBadge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-badge [collection-id] [supplys] [transfers] [claims] [collection-uri] [badge-uris]",
		Short: "Broadcast message mintBadge",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			
			var badgeSupplyAndAmount []*types.BadgeSupplyAndAmount
			err = json.Unmarshal([]byte(args[1]), &badgeSupplyAndAmount)
			if err != nil {
				return err
			}

			var transfers []*types.Transfers
			err = json.Unmarshal([]byte(args[2]), &transfers)
			if err != nil {
				return err
			}

			var claims []*types.Claim
			err = json.Unmarshal([]byte(args[3]), &claims)
			if err != nil {
				return err
			}

			argCollectionUri, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			var badgeUris []*types.BadgeUri
			err = json.Unmarshal([]byte(args[5]), &badgeUris)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}


			msg := types.NewMsgMintBadge(
				clientCtx.GetFromAddress().String(),
				argId,
				badgeSupplyAndAmount,
				transfers,
				claims,
				argCollectionUri,
				badgeUris,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
