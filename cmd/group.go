package cmd

import (
	"github.com/gogolok/finance-gogo/pkg/sheet"

	"github.com/spf13/cobra"
)

func newCmdGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Fetch and group ETF transactions",
		Long:  `Fetch and group ETF transactions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sheet.GroupDataAndOutput()
		},
	}

	return cmd
}
