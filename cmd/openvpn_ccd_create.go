package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/openvpn"
	"strconv"
)

var openvpnCcdCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an OpenVPN CCD entry",
	Run: CcdCreateRun,
}

func init() {
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDcommonName, "CCDcommonName", "c","", "The common name to show")
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDtunnel, "CCDtunnel", "t","", "cidr for your CCDtunnel network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDtunnel6, "CCDtunnel6", "","", "cidr for your CCDtunnel6 (ipv6)")
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDlocal, "CCDlocal", "l","", "cidr for your CCDlocal network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDlocal6, "CCDlocal6", "","", "cidr for your CCDlocal6 network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDremote, "CCDremote", "r","", "cidr for your CCDremote  network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&CCDremote6, "CCDremote6", "","", "cidr for your CCDremote6 (ipv6)")
	openvpnCcdCreateCmd.Flags().BoolVarP(&CCDpushRest, "CCDpushRest", "p",false, "push a reset on the client, default is false")
	openvpnCcdCreateCmd.Flags().BoolVarP(&CCDblock, "CCDblock", "b",false, "CCDblock client, default is false")
	openvpnCcdCreateCmd.MarkFlagRequired("CCDcommonName")

	OpenvpnCcdCmd.AddCommand(openvpnCcdCreateCmd)
}

func CcdCreateRun(cmd *cobra.Command, args []string) {
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

	var uuid, err = openvpnApi.CcdCreate(ccd, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}
