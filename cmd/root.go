package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the root Cobra command
var RootCmd = &cobra.Command{
	Use:   "finance-gogo",
	Short: "finance-gogo gives an overview over your ETF assets",
	Long:  `finance-gogo gives an overview over your ETF assets.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(newCmdFetch())
}
