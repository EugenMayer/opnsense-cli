package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
)

var openvpnCcdListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all OpenVPN CCD entries",
	Run:   ccdListRun,
}

func init() {
	OpenvpnCcdCmd.AddCommand(openvpnCcdListCmd)
}

func ccdListRun(cmd *cobra.Command, args []string) {
	opn := OPNsenseConfig()
	var ccds, err = opn.CcdList()
	if err != nil {
		log.Fatal(err)
	}
	for _, ccd := range ccds {
		fmt.Println(ccd)
	}
}
