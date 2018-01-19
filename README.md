## WAT

Implementation of the OPNsense WebAPI to be used on the CLI or as a library - written in GO

## Installation

Its precompiled and has no dependencies, so just download the binary and you are good to go

    curl -Lo opn https://github.com/EugenMayer/opnsense-cli/raw/master/dist/opn
    chmod +x opn

You need to create a .env for the secrets

    OPN_URL=https://localhost:10443
    OPN_APIKEY=5GWbPwKfXVLzgJnewKuu1IPw2HS7s510jKHmTM+rLA1y9VfEFE57yj/kJiWbXREB0EgpBK48u4gnyign
    OPN_APISECRET=EtpPVbiCBdtvG5VDlYJQfLu7Qck2hRffoLi2vb73arn5bKzxEbGdti8+iZetgc9eHABJy6XYG6/UsW/1
       
## Usage

    opn --help
    opn openvpn ccd --help
    opn openvpn ccd create foo --tunnel "10.10.10.1/24"
    opn openvpn ccd update foo --tunnel "11.11.11.1/24"
    opn openvpn ccd show foo
    opn openvpn ccd rm foo --tunnel "10.10.10.1/24"

## Test instance ?

If you miss yourself a opnsense instance to test agains, why do you just dont start one? :)

    vagrant up

You will able to be connect to this using https://localhost:10443 or using the shell `ssh -p 10022 root@localhost`
User: `root` / Password: `opnsense`

## Development

    # install glide, whichever way
    brew install glide

    # fetch the dependencies
    glide install

