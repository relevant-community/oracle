#!/usr/bin/env bash
rm -rf ~/.oracle

oracled init test --chain-id=oracle

name=$1
token=$2
stake=$3

# oracled config output json
# oracled config indent true
# oracled config trust-node true
# oracled config chain-id oracle
# oracled config keyring-backend test

oracled keys add $name
oracled add-genesis-account $(oracled keys show "$name" -a) $token,$stake
oracled gentx $name $stake --chain-id oracle

