## WAT

Implementation of the OPNsense WebAPI to be used on the CLI or as a library - written in GO

## Installation

Its precompiled and has no dependencies, so just download the binary and you are good to go

    curl -Lo opn https://github.com/EugenMayer/opnsense-cli/raw/master/dist/opn
    chmod +x opn

You need to create a .env (dotenv) for the secrets, or expose them into your ENV using `export`

    OPN_URL=https://localhost:10443
    OPN_APIKEY=5GWbPwKfXVLzgJnewKuu1IPw2HS7s510jKHmTM+rLA1y9VfEFE57yj/kJiWbXREB0EgpBK48u4gnyign
    OPN_APISECRET=EtpPVbiCBdtvG5VDlYJQfLu7Qck2hRffoLi2vb73arn5bKzxEbGdti8+iZetgc9eHABJy6XYG6/UsW/1
       
## Usage: cli

    opn --help

    # openvpn CCDs ( client specific overrides )
    opn openvpn ccd --help
    opn openvpn ccd create -c foo --tunnel "10.10.10.1/24"
    opn openvpn ccd update -c foo --tunnel "11.11.11.1/24"
    opn openvpn ccd show -c foo
    opn openvpn ccd rm -c foo --tunnel "10.10.10.1/24"
    opn openvpn ccd list

    # unbound host entries
    opn unbound hostentry create --host foo --domain bar.tld --ip 10.10.10.1
    opn unbound hostentry update --host foo --domain bar.tld --ip 10.10.10.2
    opn unbound hostentry show --host foo --domain bar.tld
    opn unbound hostentry rm --host foo --domain bar.tld
    opn unbound hostentry list

## Usage: GoLang library

    import (
      opn_unbound "github.com/eugenmayer/opnsense-cli/opnsense/api/unbound"
      opn_api "github.com/eugenmayer/opnsense-cli/opnsense/api"
    )
     
    func create_host_entry() error {
        var opnUnboundConnection opn_unbound.UnboundApi

        if opnConnection, opnErr := opn_api.ConfiugreFromEnv(); opnErr != nil {
            return errors.New(fmt.Sprintf("Error getting OPNsense connection: %s", opnErr))
        } else {
            opnUnboundConnection = opn_unbound.UnboundApi{opnConnection}
        }
    
        var dnsHostEntry = opn_unbound.HostEntry{
            Host:   "test,
            Domain: "foo.tld,
            Ip:     "10.10.10.1",
        }
        
        _, _ := opnUnboundConnection.HostEntryCreateOrUpdate(dnsHostEntry)
    }
## Test instance ?

If you miss yourself a opnsense instance to test agains, why do you just dont start one? :)

    vagrant up opnsense

You will able to be connect to this using https://localhost:10443 or using the shell `ssh -p 10022 root@localhost`
User: `root` / Password: `opnsense`

## Development

    # install glide, whichever way
    brew install glide

    # fetch the dependencies
    glide install
    
    # building
    go build -o dist/opn opn.go

