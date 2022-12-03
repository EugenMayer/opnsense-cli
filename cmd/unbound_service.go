package cmd

import (
	"github.com/spf13/cobra"
)

var unboundServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage Unbound service",
}

func init() {
	unboundCmd.AddCommand(unboundServiceCmd)
}
