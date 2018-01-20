package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	opnsenseapi "github.com/eugenmayer/opnsense-cli/opnsense/api"
)

var rootCmd = &cobra.Command{
	Use:   "opnsense",
	Short: "OPNsense cli to operate with a opsense API",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Activate for verbose")
}

func OPNsenseConfig() *opnsenseapi.OPNsense {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf("Error with the dotenv environment: %s", err))
	}

	if _, isset := os.LookupEnv("OPN_URL"); !isset {
		log.Fatal(fmt.Println("Please set the OPN_URL to your opnsense opnUrl like https://myopnsense:10443"))
	}

	if _, isset := os.LookupEnv("OPN_APIKEY"); !isset {
		log.Fatal(fmt.Println("Please set OPN_APIKEY to your opnsense api apiKey"))
	}

	if _, isset := os.LookupEnv("OPN_APISECRET"); !isset {
		log.Fatal(fmt.Println("Please set OPN_APISECRET to your opnsense api apiSecret"))
	}

	var parsedUrl, err = url.Parse(os.Getenv("OPN_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return &opnsenseapi.OPNsense{
		BaseUrl:     *parsedUrl,
		ApiKey:      os.Getenv("OPN_APIKEY"),
		ApiSecret:   os.Getenv("OPN_APISECRET"),
		NoSslVerify: os.Getenv("OPN_NOSSLVERIFY") == "1",
	}
}