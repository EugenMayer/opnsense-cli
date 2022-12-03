package cmd

import (
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/spf13/cobra"
	"log"
)

var unboundRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the unbound server",
	Run:   unboundRestartRun,
}

func init() {
	unboundServiceCmd.AddCommand(unboundRestartCmd)
}

func unboundRestartRun(_ *cobra.Command, _ []string) {
	var unboundApi = unbound.UnboundApi{OPNsense: GetOPNsenseApi()}

	var err = unboundApi.ServiceRestart()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Service restarted")
}
