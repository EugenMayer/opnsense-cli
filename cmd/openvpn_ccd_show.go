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
	openvpnCcdShowCmd.Flags().StringVarP(&commonName, "commonName", "c","", "The common name to show")
	openvpnCcdShowCmd.MarkFlagRequired("commonName")

	openvpnCcdCmd.AddCommand(openvpnCcdShowCmd)
}

func showCcdRun(cmd *cobra.Command, args []string) {
	opn := OPNsenseConfig()
	var ccd, err = opn.CcdGet(commonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ccd)
}
