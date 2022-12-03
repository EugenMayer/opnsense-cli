package cmd

import (
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/spf13/cobra"
	"log"
)

var unboundStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of your unbound server",
	Run:   unboundRStatusRun,
}

func init() {
	unboundServiceCmd.AddCommand(unboundStatusCmd)
}

func unboundRStatusRun(_ *cobra.Command, _ []string) {
	var unboundApi = unbound.UnboundApi{OPNsense: GetOPNsenseApi()}

	var status, err = unboundApi.ServiceStatus()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Service status: %s", status)
}
