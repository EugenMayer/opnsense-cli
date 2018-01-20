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
	Run: CcdUpdateRun,
}

func init() {
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDcommonName, "CCDcommonName", "c","", "The common name to show")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDtunnel, "CCDtunnel", "t","", "cidr for your CCDtunnel network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDtunnel6, "CCDtunnel6", "","", "cidr for your CCDtunnel6 (ipv6)")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDlocal, "CCDlocal", "l","", "cidr for your CCDlocal network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDlocal6, "CCDlocal6", "","", "cidr for your CCDlocal6 network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDremote, "CCDremote", "r","", "cidr for your CCDremote  network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&CCDremote6, "CCDremote6", "","", "cidr for your CCDremote6 (ipv6)")
	openvpnCcdUpdateCmd.Flags().BoolVarP(&CCDpushRest, "CCDpushRest", "p",false, "push a reset on the client, default is false")
	openvpnCcdUpdateCmd.Flags().BoolVarP(&CCDblock, "CCDblock", "b",false, "CCDblock client, default is false")
	openvpnCcdUpdateCmd.MarkFlagRequired("CCDcommonName")

	OpenvpnCcdCmd.AddCommand(openvpnCcdUpdateCmd)
}

func CcdUpdateRun(cmd *cobra.Command, args []string) {
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