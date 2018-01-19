## WAT

Implementation of the OPNsense WebAPI to be used on the CLI or as a library - written in GO

## Installation

Its precompiled and has no dependencies, so just download the binary and you are good to go

    curl -Lo opn https://github.com/EugenMayer/opnsense-cli/raw/master/dist/opn
    chmod +x opn

## Usage

   opn --help
   opn ccd create foo "10.10.10.1/24"

## Test instance ?

If you miss yourself a opnsense instance to test agains, why do you just dont start one? :)

   vagrant up

You will able to be connect to this using https://localhost:10443 or using the shell `ssh -P 10022 root@localhost`
User: root / Password: opnsense

## Development

   # install glide, whichever way
   brew install glide

   # fetch the dependencies
   glide install

