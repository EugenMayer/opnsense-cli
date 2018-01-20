package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"net"
)

var unboundHostEntryCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an Unbound Host entry",
	Run: HostEntryCreateRun,
}

func init() {
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYhost, "host", "","", "the host part")
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYdomain, "domain", "","", "the domain part")
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYip, "ip", "","", "the ip4 address like 10.10.10.10")
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYrr, "rr", "","A", "Record type, defaults to A")
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYmxprio, "mxprio", "","", "MX Prio a number, defaults to empty")
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYmx, "mx", "","", "MX Host, defauts empty")
	unboundHostEntryCreateCmd.Flags().StringVarP(&HOSTENTRYdescription, "description", "","", "Entry description")
	unboundHostEntryCreateCmd.MarkFlagRequired("host")
	unboundHostEntryCreateCmd.MarkFlagRequired("domain")
	unboundHostEntryCreateCmd.MarkFlagRequired("ip")

	OpenvpnCcdCmd.AddCommand(unboundHostEntryCreateCmd)
}

func HostEntryCreateRun(cmd *cobra.Command, args []string) {
	var unboundApi = unbound.UnboundApi{GetOPNsenseApi() }
	hostEntry := unbound.HostEntry{
		Host:     HOSTENTRYhost,
		Domain:  HOSTENTRYdomain,
		Ip: HOSTENTRYip,
		Mxprio:   HOSTENTRYmxprio,
		Mx:  HOSTENTRYmx,
		Description:  HOSTENTRYdescription,
	}

	if net.ParseIP(hostEntry.Ip) == nil {
		log.Fatal(fmt.Sprintf("your IP is invalid: %s", hostEntry.Ip))
	}

	var uuid, err = unboundApi.HostEntryCreate(hostEntry, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}
