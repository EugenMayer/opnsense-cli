package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
)

var unboundHostEntryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Unbound Host entries",
	Run:   hostEntryListRun,
}

func init() {
	unboundCmd.AddCommand(unboundHostEntryListCmd)
}

func hostEntryListRun(_ *cobra.Command, _ []string) {
	var unboundApi = unbound.UnboundApi{GetOPNsenseApi() }

	var hostEntries, err = unboundApi.HostEntryList()
	if err != nil {
		log.Fatal(err)
	}
	for _, hostEntry := range hostEntries {
		fmt.Println(hostEntry)
	}
}
