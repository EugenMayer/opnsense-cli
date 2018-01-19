package main

import (
	"github.com/alecthomas/kingpin"
	"os"
	"log"
	"fmt"
	"net/url"
	"github.com/EugenMayer/opnsense-cli/opnsense"
	"github.com/joho/godotenv"
)

var (

	ccdCommand     = kingpin.Command("ccd", "Delete an object.")
	createCcdCommand = ccdCommand.Command("create", "Create ccd")
	showCcdCommand = ccdCommand.Command("show", "Show ccd")
	rmCcdCommand = ccdCommand.Command("rm", "Remove ccd")

	createCcdCommanName = createCcdCommand.Arg("commonName", "The common name").String()
	createCcdCidr = createCcdCommand.Arg("cidr", "The IP/netmask in CIDR annotation like 10.10.10.5/24").String()

	showCcdCommanName = showCcdCommand.Arg("commonName", "The common name to show").String()

	rmCcdCommanName = rmCcdCommand.Arg("commonName", "The common name to remove").String()

)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	var command = kingpin.Parse()

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
	var opnUrl = os.Getenv("OPN_URL")
	var apiKey = os.Getenv("OPN_APIKEY")
	var apiSecret = os.Getenv("OPN_APISECRET")

	var parsedUrl, err = url.Parse(opnUrl)
	if err != nil {
		log.Fatal(err)
	}

	var opn = opnsense.OPNsense{
		BaseUrl: *parsedUrl,
		ApiKey: apiKey,
		ApiSecret: apiSecret,
	}

	switch command {
	case "ccd create":
	case "ccd show":
		opn.CcdExists(*showCcdCommanName)
	case "ccd rm":

	}
}
