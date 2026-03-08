#!/bin/bash

set -e

CHAIN_ID="overloan-1"
HOME_DIR="./private/.overloand"
KEYRING="test"

echo "Cleaning previous chain..."
rm -rf $HOME_DIR

echo "Initializing chain..."
./build/simd init validator \
--chain-id $CHAIN_ID \
--home $HOME_DIR

echo "Configuring CLI..."
./build/simd config chain-id $CHAIN_ID --home $HOME_DIR
./build/simd config keyring-backend $KEYRING --home $HOME_DIR

echo "Creating keys..."
./build/simd keys add validator --keyring-backend $KEYRING --home $HOME_DIR
./build/simd keys add borrower --keyring-backend $KEYRING --home $HOME_DIR
./build/simd keys add laz --keyring-backend $KEYRING --home $HOME_DIR
./build/simd keys add omnibus --keyring-backend $KEYRING --home $HOME_DIR

echo "Adding genesis accounts..."

for acc in validator borrower laz omnibus
do
./build/simd genesis add-genesis-account \
$(./build/simd keys show $acc -a --keyring-backend $KEYRING --home $HOME_DIR) \
100000000stake \
--home $HOME_DIR
done

echo "Creating validator gentx..."

./build/simd genesis gentx validator 50000000stake \
--chain-id $CHAIN_ID \
--keyring-backend $KEYRING \
--home $HOME_DIR

echo "Collecting gentx..."

./build/simd genesis collect-gentxs \
--home $HOME_DIR

echo "Dev chain initialized successfully."