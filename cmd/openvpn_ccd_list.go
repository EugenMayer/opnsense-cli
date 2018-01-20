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
	openvpnCcdCmd.AddCommand(openvpnCcdListCmd)
}

func ccdListRun(cmd *cobra.Command, args []string) {
	opn := OPNsenseConfig()
	var ccd, err = opn.CcdList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
