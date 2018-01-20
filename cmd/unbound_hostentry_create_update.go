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
	Run: HostEntryCreateUpdateRun,
}

var unboundHostEntryUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an Unbound Host entry",
	Run: HostEntryCreateUpdateRun,
}

func init() {
	setupUnboundHostEntryCreateUpdateCommand(unboundHostEntryCreateCmd)
	setupUnboundHostEntryCreateUpdateCommand(unboundHostEntryUpdateCmd)
	unboundCmd.AddCommand(unboundHostEntryCreateCmd)
	unboundCmd.AddCommand(unboundHostEntryUpdateCmd)
}

func setupUnboundHostEntryCreateUpdateCommand(command *cobra.Command) {
	command.Flags().StringVarP(&HOSTENTRYhost, "host", "","", "the host part")
	command.Flags().StringVarP(&HOSTENTRYdomain, "domain", "","", "the domain part")
	command.Flags().StringVarP(&HOSTENTRYip, "ip", "","", "the ip4 address like 10.10.10.10")
	command.Flags().StringVarP(&HOSTENTRYrr, "rr", "","A", "Record type, defaults to A")
	command.Flags().StringVarP(&HOSTENTRYmxprio, "mxprio", "","", "MX Prio a number, defaults to empty")
	command.Flags().StringVarP(&HOSTENTRYmx, "mx", "","", "MX Host, defauts empty")
	command.Flags().StringVarP(&HOSTENTRYdescription, "description", "","", "Entry description")
	command.MarkFlagRequired("host")
	command.MarkFlagRequired("domain")
	command.MarkFlagRequired("ip")
}

func HostEntryCreateUpdateRun(cmd *cobra.Command, args []string) {
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

	var uuid, err = unboundApi.HostEntryCreateOrUpdate(hostEntry)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}
