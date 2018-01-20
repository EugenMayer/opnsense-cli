package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api"
	"strconv"
)

var openvpnCcdCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an OpenVPN CCD entry",
	Run: CcdCreateRun,
}

func init() {
	openvpnCcdCreateCmd.Flags().StringVarP(&commonName, "commonName", "c","", "The common name to show")
	openvpnCcdCreateCmd.Flags().StringVarP(&tunnel, "tunnel", "t","", "cidr for your tunnel network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&tunnel6, "tunnel6", "","", "cidr for your tunnel6 (ipv6)")
	openvpnCcdCreateCmd.Flags().StringVarP(&local, "local", "l","", "cidr for your local network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&local6, "local6", "","", "cidr for your local6 network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&remote, "remote", "r","", "cidr for your remote  network 10.10.10.5/24")
	openvpnCcdCreateCmd.Flags().StringVarP(&remote6, "remote6", "","", "cidr for your remote6 (ipv6)")
	openvpnCcdCreateCmd.Flags().BoolVarP(&pushRest, "pushRest", "p",false, "push a reset on the client, default is false")
	openvpnCcdCreateCmd.Flags().BoolVarP(&block, "block", "b",false, "block client, default is false")
	openvpnCcdCreateCmd.MarkFlagRequired("commonName")

	openvpnCcdCmd.AddCommand(openvpnCcdCreateCmd)
}

func CcdCreateRun(cmd *cobra.Command, args []string) {
	opn := OPNsenseConfig()

	ccd := api.Ccd{
		CommonName:     commonName,
		TunnelNetwork:  tunnel,
		TunnelNetwork6: tunnel6,
		LocalNetwork:   local,
		LocalNetwork6:  local6,
		RemoteNetwork:  remote,
		RemoteNetwork6: remote6,
		PushReset:      strconv.Itoa(boolToInt(pushRest)),
		Block:          strconv.Itoa(boolToInt(block)),
	}

	var uuid, err = opn.CcdCreate(ccd, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}
