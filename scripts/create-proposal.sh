#!/bin/bash

set -e

CHAIN_ID="overloan-1"
HOME_DIR="./private/.overloand"
KEYRING="test"

echo "🔍 Getting addresses..."

VALIDATOR=$(./build/simd keys show validator -a --keyring-backend $KEYRING --home $HOME_DIR)
LAZ=$(./build/simd keys show laz -a --keyring-backend $KEYRING --home $HOME_DIR)
OMNIBUS=$(./build/simd keys show omnibus -a --keyring-backend $KEYRING --home $HOME_DIR)

AUTHORITY=$(./build/simd query auth module-account gov -o json | jq -r '.account.value.address')

echo "Validator: $VALIDATOR"
echo "LAZ: $LAZ"
echo "Omnibus: $OMNIBUS"
echo "Gov authority: $AUTHORITY"

echo "📝 Generating proposal JSON..."

mkdir -p proposals

cat > proposals/update_loan_params.json <<EOF
{
  "title": "Update Loan Params",
  "summary": "Set laz and omnibus authorities",
  "deposit": "10000000stake",
  "messages": [
    {
      "@type": "/cosmos.loan.v1.MsgUpdateParams",
      "authority": "$AUTHORITY",
      "params": {
        "settlement_denom": "stake",
        "min_loan_amount": "100",
        "max_loan_amount": "1000000",
        "max_tenor_months": "12",
        "laz_authorities": [
          "$LAZ"
        ],
        "omnibus_authorities": [
          "$OMNIBUS"
        ]
      }
    }
  ]
}
EOF

echo "📤 Submitting proposal..."

./build/simd tx gov submit-proposal proposals/update_loan_params.json \
--from validator \
--chain-id $CHAIN_ID \
--keyring-backend $KEYRING \
--home $HOME_DIR \
--gas auto --gas-adjustment 1.3 \
-y

echo "⏳ Waiting 3 seconds..."
sleep 3

PROPOSAL_ID=$(./build/simd query gov proposals -o json | jq -r '.proposals[-1].id')

echo "🗳 Voting YES on proposal $PROPOSAL_ID..."

./build/simd tx gov vote $PROPOSAL_ID yes \
--from validator \
--chain-id $CHAIN_ID \
--keyring-backend $KEYRING \
--home $HOME_DIR \
-y

echo "✅ Proposal submitted and voted"