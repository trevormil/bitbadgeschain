package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

var _ = strconv.Itoa(0)

func CmdNewBadge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-badge [uri] [permissions] [freeze-addresses-digest] [subasset-uris]",
		Short: "Broadcast message newBadge",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUri := args[0]
			argFreezeAddressesDigest := args[2]
			argSubassetUris := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			permissions, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgNewBadge(
				clientCtx.GetFromAddress().String(),
				argUri,
				permissions,
				argFreezeAddressesDigest,
				argSubassetUris,
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
