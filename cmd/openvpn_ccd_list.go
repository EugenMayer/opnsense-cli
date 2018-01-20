package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/openvpn"
)

var openvpnCcdListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all OpenVPN CCD entries",
	Run:   ccdListRun,
}

func init() {
	OpenvpnCcdCmd.AddCommand(openvpnCcdListCmd)
}

func ccdListRun(cmd *cobra.Command, args []string) {
	var openvpnApi = openvpn.OpenVpnApi{GetOPNsenseApi() }

	var ccds, err = openvpnApi.CcdList()
	if err != nil {
		log.Fatal(err)
	}
	for _, ccd := range ccds {
		fmt.Println(ccd)
	}
}
