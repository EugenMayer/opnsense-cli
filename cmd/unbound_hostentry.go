package cmd

import (
	"github.com/spf13/cobra"
)

var(
	HOSTENTRYhost string
	HOSTENTRYdomain string
	HOSTENTRYip string
	HOSTENTRYrr string
	HOSTENTRYmxprio string
	HOSTENTRYmx string
	HOSTENTRYdescription string
)

var unboundHostEntryCmd = &cobra.Command{
	Use:   "hostentry",
	Short: "Manage Unbound HostEntries entries",
}

func init() {
	unboundCmd.AddCommand(unboundHostEntryCmd)
}
