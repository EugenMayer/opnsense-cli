package cmd

import (
	"github.com/spf13/cobra"
)

var (
	HOSTOVERRIDEhost        string
	HOSTOVERRIDEdomain      string
	HOSTOVERRIDEuuid        string
	HOSTOVERRIDEip          string
	HOSTOVERRIDErr          string
	HOSTOVERRIDEmxprio      string
	HOSTOVERRIDEmx          string
	HOSTOVERRIDEdescription string
	// HOSTOVERRIDEnabled 0 for disabled, 1 for enabled
	HOSTOVERRIDEnabled string
)

var unboundHostOverrideCmd = &cobra.Command{
	Use:   "hostoverride",
	Short: "Manage Unbound Host Override entries",
}

func init() {
	unboundCmd.AddCommand(unboundHostOverrideCmd)
}
