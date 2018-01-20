package cmd

import (
	"github.com/spf13/cobra"
)

var (
	commonName string
)

var openvpnCcdCmd = &cobra.Command{
	Use:   "ccd",
	Short: "Manage OpenVPN CCD entries",
}

func init() {
	openvpnCmd.AddCommand(openvpnCcdCmd)
	openvpnCcdCmd.PersistentFlags().StringVarP(&commonName, "commonName", "c","", "The common name to show")
}
