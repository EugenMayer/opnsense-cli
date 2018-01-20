package cmd

import (
	"github.com/spf13/cobra"
)

var (
	commonName string
	tunnel string
	tunnel6 string
	local string
	local6 string
	remote string
	remote6 string
	pushRest bool
	block bool
)

var openvpnCcdCmd = &cobra.Command{
	Use:   "ccd",
	Short: "Manage OpenVPN CCD entries",
}

func init() {
	openvpnCmd.AddCommand(openvpnCcdCmd)
}
