#!/bin/sh

EVM_CHAIN=ethereum

## For Testing chain
COSMOS_CHAIN=umee
COSMOS_DENOM=uumee
## For Cudos chain
CUDOS_CHAIN=cudos
CUDOS_DENOM=acudos
AXELAR_CHAIN=axelarnet
EVM_DENOM=uusdc
DIR="$(dirname "$0")"
GATEWAY_ID="0x67bf05Bd35EB69a5aAD899B0E6480DCD7d74e23b"
TOKEN_ADDRESS="0x06175B42522310bce7F7B263C8C95e3Bc6568C0D"
TX_HASH="0x7efc707d6794332649ede05f44193985bc8ffa53eaa317d92fccc79bf3cccdd1"

echo "#### 1. Adding EVM chain ####"
sh "${DIR}/steps/01-add-evm-chain.sh" ${EVM_CHAIN}

echo "\n#### 2. Adding Cosmos chain ####"
sh "${DIR}/steps/02-add-cosmos-chain.sh" ${COSMOS_CHAIN}

echo "\n#### 3. Register Broadcaster ####"
sh "${DIR}/steps/03-register-broadcaster.sh"

echo "\n#### 4. Activate EVM Chains ####"
sh "${DIR}/steps/04-activate-evm-chain.sh" ${EVM_CHAIN}

echo "\n#### 5. Activate Cosmos Chains ####"
sh "${DIR}/steps/05-activate-cosmos-chain.sh" ${COSMOS_CHAIN}

echo "\n#### 6. Register Cosmos native asset ####"
sh "${DIR}/steps/09-register-native-asset.sh" ${COSMOS_CHAIN} ${COSMOS_DENOM}

echo "\n#### 7. Set gateway id ####"
sh "${DIR}/steps/06-setgateway.sh" ${EVM_CHAIN} ${GATEWAY_ID}

echo "\n#### 8. Register new token to Ethereum ####"
sh "${DIR}/steps/07-create-deploy-token.sh" ${TOKEN_ADDRESS}

echo "\n#### 9. Confirm deploy token ####"
sh "${DIR}/steps/08-confirm-deploy-token.sh" ${TX_HASH}

echo "\n#### 10. Register Cosmos support aUSDC asset ####"
sh "${DIR}/steps/foregin-asset.sh" ${COSMOS_CHAIN} ${EVM_DENOM}

echo "\n#### 11. Register Axelar support aUSDC asset ####"
sh "${DIR}/steps/foregin-asset.sh" ${AXELAR_CHAIN} ${EVM_DENOM}

