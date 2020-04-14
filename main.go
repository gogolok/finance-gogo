package main

import (
	"os"

	"github.com/gogolok/finance-gogo/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
