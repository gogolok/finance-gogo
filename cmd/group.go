package cmd

import (
	"github.com/gogolok/finance-gogo/pkg/sheet"

	"github.com/spf13/cobra"
)

func newCmdGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Fetches and groups ETF transactions",
		Long:  `Fetches and groups ETF transactions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sheet.GroupDataAndOutput()
		},
	}

	return cmd
}
