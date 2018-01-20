package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/openvpn"
)

var openvpnCcdRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a OpenVPN CCD entry",
	Run:  showCcdRun,
}

func init() {
	openvpnCcdRmCmd.Flags().StringVarP(&CCDcommonName, "commonName", "c","", "The common name to show")
	openvpnCcdRmCmd.MarkFlagRequired("commonName")

	OpenvpnCcdCmd.AddCommand(openvpnCcdRmCmd)
}

func showCcdRun(_ *cobra.Command, _ []string) {
	var openvpnApi = openvpn.OpenVpnApi{GetOPNsenseApi() }

	var ccd, err = openvpnApi.CcdRemove(CCDcommonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
