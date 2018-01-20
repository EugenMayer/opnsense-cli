package cmd

import (
	"github.com/spf13/cobra"
)

var (
	CCDcommonName string
	CCDtunnel     string
	CCDtunnel6    string
	CCDlocal      string
	CCDlocal6     string
	CCDremote     string
	CCDremote6    string
	CCDpushRest   bool
	CCDblock      bool
)

var OpenvpnCcdCmd = &cobra.Command{
	Use:   "ccd",
	Short: "Manage OpenVPN CCD entries",
}

func init() {
	openvpnCmd.AddCommand(OpenvpnCcdCmd)
}
