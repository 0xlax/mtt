package cmd

import (
	"github.com/spf13/cobra"
)

func addTxCmdCommonFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&flagSequence, "sequence", "s", 0, "sequence number for the tx")
	cmd.Flags().IntVarP(&flagAccount, "account", "a", 0, "account number for the tx")
	cmd.Flags().StringVarP(&flagNode, "node", "n", "", "tendermint rpc node to get sequence and account number from")
	cmd.Flags().BoolVarP(&flagForce, "force", "f", false, "overwrite files already there")
	cmd.Flags().BoolVarP(&flagAdditional, "additional", "x", false, "add additional txs with higher sequence number")
}

func addDenomFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&flagDenom, "denom", "d", "", "fee denom, for offline creation")
}

func addSignCmdFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&flagTxIndex, "index", "i", 0, "index of the tx to sign")
	cmd.Flags().StringVarP(&flagFrom, "from", "f", "", "name of your local key to sign with")
	cmd.MarkFlagRequired("from")
}

func addListCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&flagAll, "all", "a", false, "list files for all chains and keys")
}

func addBroadcastCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&flagNode, "node", "n", "", "node address to broadcast too. flag overrides config")
	cmd.Flags().IntVarP(&flagTxIndex, "index", "i", 0, "index of the tx to broadcast")
}

func addDeleteCmdFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&flagTxIndex, "index", "i", 0, "index of the tx to delete")
}
