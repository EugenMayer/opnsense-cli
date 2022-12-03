package cmd

import (
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/spf13/cobra"
	"log"
)

var unboundHostOverrideShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a Unbound Host override",
	Run:   hostOverrideShowRun,
}

func init() {
	unboundHostOverrideShowCmd.Flags().StringVarP(&HOSTOVERRIDEhost, "host", "", "", "the host part")
	unboundHostOverrideShowCmd.Flags().StringVarP(&HOSTOVERRIDEdomain, "domain", "", "", "the domain part")

	_ = unboundHostOverrideShowCmd.MarkFlagRequired("host")
	_ = unboundHostOverrideShowCmd.MarkFlagRequired("domain")

	unboundHostOverrideCmd.AddCommand(unboundHostOverrideShowCmd)
}

func hostOverrideShowRun(_ *cobra.Command, _ []string) {
	var openvpnApi = unbound.UnboundApi{OPNsense: GetOPNsenseApi()}

	var ccd, err = openvpnApi.HostEntryGetByFQDN(HOSTOVERRIDEhost, HOSTOVERRIDEdomain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
