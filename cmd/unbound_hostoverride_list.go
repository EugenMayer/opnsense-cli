package cmd

import (
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/spf13/cobra"
	"log"
)

var unboundHostOverrideListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Unbound Host Override entries",
	Run:   hostEntryListRun,
}

func init() {
	unboundCmd.AddCommand(unboundHostOverrideListCmd)
}

func hostEntryListRun(_ *cobra.Command, _ []string) {
	var unboundApi = unbound.UnboundApi{OPNsense: GetOPNsenseApi()}

	var hostOverrideEntries, err = unboundApi.HostOverrideList()
	if err != nil {
		log.Fatal(err)
	}
	for _, hostEntry := range hostOverrideEntries {
		fmt.Println(hostEntry)
	}
}
