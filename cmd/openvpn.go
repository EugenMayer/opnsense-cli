package cmd

import (
	//"os"
	//"log"
	//"fmt"
	//"net/url"
	//"github.com/EugenMayer/opnsense-cli/opnsense/api"
	//"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var openvpnCmd = &cobra.Command{
	Use:   "openvpn",
	Short: "Manage OpenVPN using the OPNsense API",
}

func init() {
	rootCmd.AddCommand(openvpnCmd)
}
//
//var (
//	openvpnCommand = kingpin.Command("openvpn", "OpenVPN operations.")
//
//	ccdCommand       = openvpnCommand.Command("ccd", "Client Specific Overides for openvpn ( CCD ).")
//	createCcdCommand = ccdCommand.Command("create", "Create ccd")
//	updateCcdCommand = ccdCommand.Command("update", "Update ccd")
//	showCcdCommand   = ccdCommand.Command("show", "Show ccd")
//	rmCcdCommand     = ccdCommand.Command("rm", "Remove ccd")
//
//	createCcdCommanName = createCcdCommand.Arg("commonName", "The common name").Required().String()
//	createTunnel4Cidr   = createCcdCommand.Flag("tunnel", "cidr for your tunnel network 10.10.10.5/24").Default("").String()
//	createTunnel6Cidr   = createCcdCommand.Flag("tunnel6", "cidr for your tunnel v6 network").Default("").String()
//	createLocal4Cidr    = createCcdCommand.Flag("local", "cidr for your local v4 network 10.10.10.5/24").Default("").String()
//	createLocal6Cidr    = createCcdCommand.Flag("local6", "cidr for your local v6 network ").Default("").String()
//	createRemote4Cidr   = createCcdCommand.Flag("remote", "cidr for your remote v4 network 10.10.10.5/24").Default("").String()
//	createRemote6Cidr   = createCcdCommand.Flag("remote6", "cidr for your remote v6 network").Default("").String()
//	createPushReset     = createCcdCommand.Flag("pushReset", "push a reset on the client, either 1 for true, default is 0").Default("0").String()
//	createBlock         = createCcdCommand.Flag("block", "block client, either 1 for true, default is 0").Default("0").String()
//
//	updateCcdCommanName = updateCcdCommand.Arg("commonName", "The common-name to be updated").Required().String()
//	updateTunnel4Cidr   = updateCcdCommand.Flag("tunnel", "cidr for your tunnel network 10.10.10.5/24").Default("").String()
//	updateTunnel6Cidr   = updateCcdCommand.Flag("tunnel6", "cidr for your tunnel v6 network").Default("").String()
//	updateLocal4Cidr    = updateCcdCommand.Flag("local", "cidr for your local v4 network 10.10.10.5/24").Default("").String()
//	updateLocal6Cidr    = updateCcdCommand.Flag("local6", "cidr for your local v6 network ").Default("").String()
//	updateRemote4Cidr   = updateCcdCommand.Flag("remote", "cidr for your remote v4 network 10.10.10.5/24").Default("").String()
//	updateRemote6Cidr   = updateCcdCommand.Flag("remote6", "cidr for your remote v6 network").Default("").String()
//	updatePushReset     = updateCcdCommand.Flag("pushReset", "push a reset on the client, either 1 for true, default is 0").Default("0").String()
//	updateBlock         = updateCcdCommand.Flag("block", "block client, either 1 for true, default is 0").Default("0").String()
//
//	showCcdCommanName = showCcdCommand.Arg("commonName", "The common name to show").String()
//
//	rmCcdCommanName = rmCcdCommand.Arg("commonName", "The common name to remove").String()
//)
//

//	case "openvpn ccd update":
//		ccd := opnsense.Ccd{
//			CommonName:     *updateCcdCommanName,
//			TunnelNetwork:  *updateTunnel4Cidr,
//			TunnelNetwork6: *updateTunnel6Cidr,
//			LocalNetwork:   *updateLocal4Cidr,
//			LocalNetwork6:  *updateLocal6Cidr,
//			RemoteNetwork:  *updateRemote4Cidr,
//			RemoteNetwork6: *updateRemote6Cidr,
//			Block:          *updateBlock,
//			PushReset:      *updatePushReset,
//		}
//
//		var uuid, err = opn.CcdCreate(ccd, true)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(uuid)
//	case "openvpn ccd show":
//		var ccd, err = opn.CcdGet(*showCcdCommanName)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(ccd)
//	case "openvpn ccd rm":
//		var uuid, err = opn.CcdRemove(*rmCcdCommanName)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(uuid)
//	}
//}
