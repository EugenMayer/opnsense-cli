package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api"
	"strconv"
)

var openvpnCcdUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an OpenVPN CCD entry",
	Run: CcdUpdateRun,
}

func init() {
	openvpnCcdUpdateCmd.Flags().StringVarP(&commonName, "commonName", "c","", "The common name to show")
	openvpnCcdUpdateCmd.Flags().StringVarP(&tunnel, "tunnel", "t","", "cidr for your tunnel network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&tunnel6, "tunnel6", "","", "cidr for your tunnel6 (ipv6)")
	openvpnCcdUpdateCmd.Flags().StringVarP(&local, "local", "l","", "cidr for your local network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&local6, "local6", "","", "cidr for your local6 network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&remote, "remote", "r","", "cidr for your remote  network 10.10.10.5/24")
	openvpnCcdUpdateCmd.Flags().StringVarP(&remote6, "remote6", "","", "cidr for your remote6 (ipv6)")
	openvpnCcdUpdateCmd.Flags().BoolVarP(&pushRest, "pushRest", "p",false, "push a reset on the client, default is false")
	openvpnCcdUpdateCmd.Flags().BoolVarP(&block, "block", "b",false, "block client, default is false")
	openvpnCcdUpdateCmd.MarkFlagRequired("commonName")

	openvpnCcdCmd.AddCommand(openvpnCcdUpdateCmd)
}

func CcdUpdateRun(cmd *cobra.Command, args []string) {
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

	var uuid, err = opn.CcdCreate(ccd, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}