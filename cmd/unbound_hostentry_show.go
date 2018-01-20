package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
)

var unboundHostEntryShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a OpenVPN CCD entry",
	Run:   HostEntryShowRun,
}

func init() {
	unboundHostEntryShowCmd.Flags().StringVarP(&HOSTENTRYhost, "host", "","", "the host part")
	unboundHostEntryShowCmd.Flags().StringVarP(&HOSTENTRYdomain, "domain", "","", "the domain part")

	unboundHostEntryShowCmd.MarkFlagRequired("host")
	unboundHostEntryShowCmd.MarkFlagRequired("domain")

	unboundCmd.AddCommand(unboundHostEntryShowCmd)
}

func HostEntryShowRun(_ *cobra.Command, _ []string) {
	var openvpnApi = unbound.UnboundApi{GetOPNsenseApi() }

	var ccd, err = openvpnApi.HostEntryGet(HOSTENTRYhost, HOSTENTRYdomain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
