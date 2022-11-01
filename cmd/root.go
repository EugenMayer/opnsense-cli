package cmd

import (
	"fmt"
	opnsenseapi "github.com/eugenmayer/opnsense-cli/opnsense/api"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "opnsense",
	Short: "OPNsense cli to operate with a opnsense API",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Activate for verbose")
}

func GetOPNsenseApi() *opnsenseapi.OPNsense {
	connection, err := opnsenseapi.ConfigureFromEnv()
	if err != nil {
		log.Fatal()
	}
	return connection
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
