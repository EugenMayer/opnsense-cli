package cmd

import (
	"fmt"
	"github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
	"github.com/spf13/cobra"
	"log"
	"net"
)

var unboundHostOverrideCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an Unbound Host Override",
	Run:   hostOverrideCreateUpdateRun,
}

var unboundHostOverrideUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an Unbound Host Override",
	Run:   hostOverrideCreateUpdateRun,
}

func init() {
	setupUnboundHostOverrideCreateUpdateCommand(unboundHostOverrideCreateCmd)
	setupUnboundHostOverrideCreateUpdateCommand(unboundHostOverrideUpdateCmd)
	unboundHostOverrideCmd.AddCommand(unboundHostOverrideCreateCmd)
	unboundHostOverrideCmd.AddCommand(unboundHostOverrideUpdateCmd)
}

func setupUnboundHostOverrideCreateUpdateCommand(command *cobra.Command) {
	command.Flags().StringVarP(&HOSTOVERRIDEhost, "host", "", "", "the host part")
	command.Flags().StringVarP(&HOSTOVERRIDEdomain, "domain", "", "", "the domain part")
	command.Flags().StringVarP(&HOSTOVERRIDEip, "ip", "", "", "the ip4 address like 10.10.10.10")
	command.Flags().StringVarP(&HOSTOVERRIDErr, "rr", "", "A", "Record type, defaults to A")
	command.Flags().StringVarP(&HOSTOVERRIDEmxprio, "mxprio", "", "", "MX Prio a number, defaults to empty")
	command.Flags().StringVarP(&HOSTOVERRIDEmx, "mx", "", "", "MX Host, defaults to empty")
	command.Flags().StringVarP(&HOSTOVERRIDEdescription, "description", "", "", "Entry description")
	command.Flags().StringVarP(&HOSTOVERRIDEnabled, "enabled", "", "1", "Should the entry be enabled. 1 (true) by default")

	_ = command.MarkFlagRequired("host")
	_ = command.MarkFlagRequired("domain")
	_ = command.MarkFlagRequired("ip")
}

func hostOverrideCreateUpdateRun(_ *cobra.Command, _ []string) {
	var unboundApi = unbound.UnboundApi{OPNsense: GetOPNsenseApi()}
	hostOverride := unbound.HostOverride{
		Host:        HOSTOVERRIDEhost,
		Domain:      HOSTOVERRIDEdomain,
		Ip:          HOSTOVERRIDEip,
		Mxprio:      HOSTOVERRIDEmxprio,
		Mx:          HOSTOVERRIDEmx,
		Description: HOSTOVERRIDEdescription,
		Rr:          HOSTOVERRIDErr,
		Enabled:     HOSTOVERRIDEnabled,
	}

	if net.ParseIP(hostOverride.Ip) == nil {
		log.Fatal(fmt.Sprintf("your IP is invalid: %s", hostOverride.Ip))
	}

	var uuid, err = unboundApi.HostOverrideCreateOrUpdate(hostOverride)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Create or update successfully, uuid: %s", uuid))
}
