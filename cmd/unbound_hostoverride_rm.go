package cmd

import (
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/spf13/cobra"
	"log"
)

var unboundHostOverrideRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a Unbound host override entry",
	Run:   hostEntryRmRun,
}

func init() {
	unboundHostOverrideRmCmd.Flags().StringVarP(&HOSTOVERRIDEhost, "host", "", "", "lookup entry by the host / domain the host part")
	unboundHostOverrideRmCmd.Flags().StringVarP(&HOSTOVERRIDEdomain, "domain", "", "", "lookup entry by the host / domainthe domain part")
	unboundHostOverrideRmCmd.Flags().StringVarP(&HOSTOVERRIDEuuid, "uuid", "", "", "delete by uuid explicitly")

	unboundHostOverrideCmd.AddCommand(unboundHostOverrideRmCmd)
}

func hostEntryRmRun(_ *cobra.Command, _ []string) {
	var openvpnApi = unbound.UnboundApi{OPNsense: GetOPNsenseApi()}

	if HOSTOVERRIDEhost == "" || HOSTOVERRIDEdomain == "" && HOSTOVERRIDEuuid == "" {
		log.Fatal("Please either set host and domain or set uuid")
	}

	if HOSTOVERRIDEhost != "" && HOSTOVERRIDEdomain != "" {
		var entry, err = openvpnApi.HostEntryGetByFQDN(HOSTOVERRIDEhost, HOSTOVERRIDEdomain)
		if err != nil {
			log.Fatal(err.Error())
		}
		HOSTOVERRIDEuuid = entry.Uuid
		log.Printf("Found Uuid %s for FQDM %s.%s", HOSTOVERRIDEuuid, HOSTOVERRIDEhost, HOSTOVERRIDEdomain)
	}

	var err = openvpnApi.HostEntryRemove(HOSTOVERRIDEuuid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted successfully")
}
