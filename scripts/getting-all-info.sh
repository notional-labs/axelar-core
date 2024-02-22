#!/bin/bash

CHAIN_ID=axelar
HOME=.axelar
DEFAULT_KEYS_FLAGS="--keyring-backend test --home ${HOME}"
DIR="$(dirname "$0")"
COSMOS_CHAIN=umee
# Gán kết quả của lệnh docker exec vào biến RESULT
OWNER_VAL_ADDRESS=$(axelard keys show owner -a --bech val ${DEFAULT_KEYS_FLAGS})
OWNER_ADDRESS=$(axelard keys show owner ${DEFAULT_KEYS_FLAGS})

# In ra giá trị của biến RESULT
echo "owner validator address: $OWNER_VAL_ADDRESS"
echo "owner address: $OWNER_ADDRESS"

SUPPORTED_CHAINS=$(axelard q nexus chains)
echo "Supported chains: $SUPPORTED_CHAINS"

ETHEREUM_CHAINS=$(axelard q nexus assets ethereum)
echo "Ethereum chain assets: $ETHEREUM_CHAINS"

echo "Cosmos chain assets: $(axelard q nexus assets "${COSMOS_CHAIN}")"


#sh "$DIR/06-setgateway.sh $1 $2"






