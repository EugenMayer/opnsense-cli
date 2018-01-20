package cmd

import (
	"github.com/spf13/cobra"
)

var unboundCmd = &cobra.Command{
	Use:   "unbound",
	Short: "Manage Unbound using the OPNsense API",
}

func init() {
	RootCmd.AddCommand(unboundCmd)
}
