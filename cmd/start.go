package cmd

import (
	"github.com/lajosdeme/transaction-relayer/core"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the transaction relayer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		core.RunRouter()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
