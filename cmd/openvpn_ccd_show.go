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
	openvpnCcdShowCmd.Flags().StringVarP(&CCDcommonName, "CCDcommonName", "c","", "The common name to show")
	openvpnCcdShowCmd.MarkFlagRequired("CCDcommonName")

	OpenvpnCcdCmd.AddCommand(openvpnCcdShowCmd)
}

func showCcdRun(_ *cobra.Command, _ []string) {
	opn := OPNsenseConfig()
	var ccd, err = opn.CcdGet(CCDcommonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
