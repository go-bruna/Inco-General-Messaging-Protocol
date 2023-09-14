#!/bin/bash

set -Eeuo pipefail

# in /config folder

# init node
echo "setuping ..."
/home/ubuntu/Inco-chain/setup.sh

# generate keys
echo "generating keys ..."
./scripts/prepare_volumes_from_fhe_tool.sh /usr/local/bin

# Copy keys to evmos home folder
echo "copying keys to evmos home folder"
EVMOS_NETWORK_KEYS_PATH=/home/ubuntu/.evmosd/zama/keys/network-fhe-keys ./scripts/prepare_validator_ci.sh

# start the node
TRACE=""
LOGLEVEL="info"

fhevm-decryptions-db &

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
evmosd start --home /home/ubuntu/.evmosd --pruning=nothing $TRACE --log_level $LOGLEVEL \
        --minimum-gas-prices=0.0001aevmos \
        --json-rpc.gas-cap=50000000 \
        --json-rpc.api eth,txpool,net,web3 \
        --rpc.laddr "tcp://0.0.0.0:26657"
