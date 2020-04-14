package cmd

import (
	"github.com/gogolok/finance-gogo/pkg/sheet"

	"github.com/spf13/cobra"
)

func newCmdFetch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Download ETF transactions",
		Long:  `Download ETF transactions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sheet.Data()
		},
	}

	return cmd
}
