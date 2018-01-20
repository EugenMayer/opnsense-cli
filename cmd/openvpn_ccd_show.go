package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
)

var openvpnCcdShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a OpenVPN CCD entry",
	Run:  showCcdRun,
}

func init() {
	//openvpnCcdCmd.MarkPersistentFlagRequired("commonName")
	openvpnCcdCmd.AddCommand(openvpnCcdShowCmd)
}

func showCcdRun(cmd *cobra.Command, args []string) {

	cmd.MarkPersistentFlagRequired("commonName")
	opn := OPNsenseConfig()
	var ccd, err = opn.CcdGet(commonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
