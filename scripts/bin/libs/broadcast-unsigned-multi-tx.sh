#!/bin/bash

CHAIN_ID=axelar
HOME=.axelar
DEFAULT_KEYS_FLAGS="--keyring-backend test --home ${HOME}"

## Sign unsigned transaction.
axelard tx sign ${HOME}/unsigned_msg.json --from gov1 \
--multisig $(axelard keys show governance -a ${DEFAULT_KEYS_FLAGS}) \
--chain-id $CHAIN_ID ${DEFAULT_KEYS_FLAGS} &> ${HOME}/signed_tx.json
cat ${HOME}/signed_tx.json

## Multisign signed transaction.
axelard tx multisign ${HOME}/unsigned_msg.json governance ${HOME}/signed_tx.json \
--from owner --sign-mode amino-json --chain-id $CHAIN_ID ${DEFAULT_KEYS_FLAGS} &> ${HOME}/tx-ms.json
cat ${HOME}/tx-ms.json

## Broadcast multisigned transaction.
axelard tx broadcast ${HOME}/tx-ms.json ${DEFAULT_KEYS_FLAGS} --sign-mode amino-json
