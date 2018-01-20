package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/openvpn"
)

var openvpnCcdShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a OpenVPN CCD entry",
	Run:  CcdRmRun,
}

func init() {
	openvpnCcdShowCmd.Flags().StringVarP(&CCDcommonName, "commonName", "c","", "The common-name to identify the CCD")
	openvpnCcdShowCmd.MarkFlagRequired("commonName")

	OpenvpnCcdCmd.AddCommand(openvpnCcdShowCmd)
}

func CcdRmRun(_ *cobra.Command, _ []string) {
	var openvpnApi = openvpn.OpenVpnApi{GetOPNsenseApi() }

	var ccd, err = openvpnApi.CcdGet(CCDcommonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
