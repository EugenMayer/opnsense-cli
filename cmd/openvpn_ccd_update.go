package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"strconv"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/openvpn"
)

var openvpnCcdUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an OpenVPN CCD entry",
	Run:   ccdUpdateRun,
}

func init() {
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDcommonName, "commonName", "c","", "The common name to show")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDtunnel, "tunnel", "t","", "cidr for your CCDtunnel network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDtunnel6, "tunnel6", "","", "cidr for your CCDtunnel6 (ipv6)")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDlocal, "local", "l","", "cidr for your CCDlocal network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDlocal6, "local6", "","", "cidr for your CCDlocal6 network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDremote, "remote", "r","", "cidr for your CCDremote  network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDremote6, "remote6", "","", "cidr for your CCDremote6 (ipv6)")
	openvpnCcdUpdateCmd.Flags().BoolVarP(&CCDpushRest, "pushRest", "p",false, "push a reset on the client, default is false")
	openvpnCcdUpdateCmd.Flags().BoolVarP(&CCDblock, "block", "b",false, "block client, default is false")
	openvpnCcdUpdateCmd.MarkFlagRequired("commonName")

	OpenvpnCcdCmd.AddCommand(openvpnCcdUpdateCmd)
}

func ccdUpdateRun(_ *cobra.Command, _ []string) {
	var openvpnApi = openvpn.OpenVpnApi{GetOPNsenseApi() }

	ccd := openvpn.Ccd{
		CommonName:     CCDcommonName,
		TunnelNetwork:  CCDtunnel,
		TunnelNetwork6: CCDtunnel6,
		LocalNetwork:   CCDlocal,
		LocalNetwork6:  CCDlocal6,
		RemoteNetwork:  CCDremote,
		RemoteNetwork6: CCDremote6,
		PushReset:      strconv.Itoa(BoolToInt(CCDpushRest)),
		Block:          strconv.Itoa(BoolToInt(CCDblock)),
	}

	var uuid, err = openvpnApi.CcdCreate(ccd, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}