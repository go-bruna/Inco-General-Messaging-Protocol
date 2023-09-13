package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evmos/evmos/v9/x/callback/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCallEvmAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call-evm-add [contract-address] [func-name] [arg]",
		Short: "Broadcast message call-evm-add",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContractAddress := args[0]
			argFuncName := args[1]
			argArg := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCallEvmAdd(
				clientCtx.GetFromAddress().String(),
				argContractAddress,
				argFuncName,
				argArg,
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
