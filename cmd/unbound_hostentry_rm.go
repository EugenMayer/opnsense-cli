package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
)

var unboundHostEntryRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a Unbound HostEntry entry",
	Run:   HostEntryRmRun,
}

func init() {
	unboundHostEntryRmCmd.Flags().StringVarP(&HOSTENTRYhost, "host", "","", "the host part")
	unboundHostEntryRmCmd.Flags().StringVarP(&HOSTENTRYdomain, "domain", "","", "the domain part")

	unboundHostEntryRmCmd.MarkFlagRequired("host")
	unboundHostEntryRmCmd.MarkFlagRequired("domain")

	unboundCmd.AddCommand(unboundHostEntryRmCmd)
}

func HostEntryRmRun(_ *cobra.Command, _ []string) {
	var openvpnApi = unbound.UnboundApi{GetOPNsenseApi() }

	var ccd, err = openvpnApi.HostEntryRemove(HOSTENTRYhost, HOSTENTRYdomain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
