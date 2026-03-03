package cli

import (
	"strconv"

	basev1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "loan",
		Short:                      "Loan transactions",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateLoan(),
	)

	return cmd
}

func CmdCreateLoan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-loan [amount] [tenor-months] [metadata]",
		Short: "Create a new loan",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			tenor, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := &loanv1.MsgCreateLoan{
				Borrower: clientCtx.GetFromAddress().String(),
				Principal: &basev1beta1.Coin{
					Denom:  amount.Denom,
					Amount: amount.Amount.String(),
				},
				TenorMonths:  tenor,
				MetadataHash: args[2],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
