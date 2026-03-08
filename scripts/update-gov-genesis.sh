#!/bin/bash

GENESIS="./private/.overloand/config/genesis.json"

echo "Updating gov params in genesis..."

jq '
.app_state.gov.params.max_deposit_period = "30s" |
.app_state.gov.params.voting_period = "30s"
' $GENESIS > genesis.tmp && mv genesis.tmp $GENESIS

echo "Gov params updated:"
jq '.app_state.gov.params | {max_deposit_period, voting_period}' $GENESIS