[build by our opensource concourse pipeline](https://github.com/EugenMayer/concourse-our-open-pipelines)

## WAT

Implementation of the OPNsense WebAPI to be used on the CLI or as a library - written in GO

## Installation

Its precompiled and has no dependencies, so just download the binary, and you are good to go

    Pick a release from https://github.com/EugenMayer/opnsense-cli/releases
    chmod +x opn-*

You need to create a .env (dotenv) for the secrets, or expose them into your ENV using `export`

    OPN_URL=https://localhost:10443
    OPN_APIKEY=5GWbPwKfXVLzgJnewKuu1IPw2HS7s510jKHmTM+rLA1y9VfEFE57yj/kJiWbXREB0EgpBK48u4gnyign
    OPN_APISECRET=EtpPVbiCBdtvG5VDlYJQfLu7Qck2hRffoLi2vb73arn5bKzxEbGdti8+iZetgc9eHABJy6XYG6/UsW/1
    # if we should not verify SSL while talking to opn, enable that
    #OPN_NOSSLVERIFY=1

## Commands included

Just run 

    opn 

to see a full list. Currently implemented

- managing host overrides for unbound (CRUD + list)


## Usage: cli

    opn --help

    # unbound host DNS entries
    opn unbound hostoverride create --host foo --domain bar.tld --ip 10.10.10.1
    opn unbound hostoverride update --host foo --domain bar.tld --ip 10.10.10.2
    opn unbound hostoverride show --host foo --domain bar.tld
    opn unbound hostoverride rm --host foo --domain bar.tld
    opn unbound hostoverride list

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
    
        var dnsHostEntry = opn_unbound.HostOverride{
            Host:   "test,
            Domain: "foo.tld,
            Ip:     "10.10.10.1",
        }
        
        _, _ := opnUnboundConnection.HostOverrideCreateOrUpdate(dnsHostEntry)
    }
## Test instance ?

If you miss yourself a opnsense instance to test agains, why do you just dont start one? :)

    vagrant up opnsense

You will be able to connect to this using https://localhost:10443 or using the shell `ssh -p 10022 root@localhost`
User: `root` / Password: `opnsense`

## Development

    # building
    make build

