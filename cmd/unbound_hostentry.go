package cmd

import (
	"github.com/spf13/cobra"
)

var(

)

var unboundHostEntryCmd = &cobra.Command{
	Use:   "hostentry",
	Short: "Manage Unbound HostEntries entries",
}

func init() {
	unboundCmd.AddCommand(unboundHostEntryCmd)
}
