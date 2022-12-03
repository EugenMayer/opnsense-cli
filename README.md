[![build](https://github.com/EugenMayer/opnsense-cli/actions/workflows/build.yml/badge.svg)](https://github.com/EugenMayer/opnsense-cli/actions/workflows/build.yml)

## WAT

Implementation of the OPNsense WebAPI to be used on the CLI or as a library - written in GO

Compatible with OPNsense 22.1 and higher.

## Installation

Its precompiled and has no dependencies, so just download the binary, and you are good to go

    Pick a release from https://github.com/EugenMayer/opnsense-cli/releases
    chmod +x opn-*

You need to create a .env (dotenv) for the secrets, or expose them into your ENV using `export`

`   OPN_URL=https://localhost:10443
    OPN_APIKEY=5GWbPwKfXVLzgJnewKuu1IPw2HS7s510jKHmTM+rLA1y9VfEFE57yj/kJiWbXREB0EgpBK48u4gnyign
    OPN_APISECRET=EtpPVbiCBdtvG5VDlYJQfLu7Qck2hRffoLi2vb73arn5bKzxEbGdti8+iZetgc9eHABJy6XYG6/UsW/1`
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

    # unbound service
    opn unbound service restart
    opn unbound service reconfigure
    opn unbound service status

**HINT**: Right now, as of 22.7, you have to run `reconfigure` everytime you change a hostentry on `unbound`, so e.g.

    # this will yet not show up when you run a DNS query
    opn unbound hostoverride create --host foo --domain bar.tld --ip 10.10.10.1
    # after that, it will
    opn unbound service reconfigure

This might change in later releases.

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

## Development

    # building
    make build

## Contributions

If you like to add a new command implementing OPNsense API reference https://docs.opnsense.org/development/api.html#introduction - open a PR and iam happy to add it. Be bold. 
