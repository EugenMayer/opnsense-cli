package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/openvpn"
)

var openvpnCcdRegenerateCmd = &cobra.Command{
	Use:   "show",
	Short: "Regenerate all OpenVPN CCD entries",
	Run:   ccdRegenerateRun,
}

func init() {
	OpenvpnCcdCmd.AddCommand(openvpnCcdRegenerateCmd)
}

func ccdRegenerateRun(_ *cobra.Command, _ []string) {
	var openvpnApi = openvpn.OpenVpnApi{GetOPNsenseApi() }

	var err = openvpnApi.CcdRegenrate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("finished")
}
