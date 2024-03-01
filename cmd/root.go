package cmd

import (
	"fmt"
	"os"

	"github.com/lajosdeme/transaction-relayer/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tx-relay",
	Short: "Entrypoint for the transaction relayer",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			fmt.Println("Failed to load config: ", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to start transaction relayer: ", err)
		os.Exit(1)
	}
}
