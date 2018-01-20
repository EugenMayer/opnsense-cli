package cmd

import (
	"github.com/spf13/cobra"
)

var openvpnCmd = &cobra.Command{
	Use:   "openvpn",
	Short: "Manage OpenVPN using the OPNsense API",
}

func init() {
	RootCmd.AddCommand(openvpnCmd)
}
