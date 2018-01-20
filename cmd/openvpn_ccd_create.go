package cmd

import (
	"github.com/spf13/cobra"
)

var openvpnCcdCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create OpenVPN CCD entries",
}

func init() {
	openvpnCcdCmd.AddCommand(openvpnCcdCreateCmd)
}
