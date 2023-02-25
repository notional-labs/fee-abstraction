#!/usr/bin/env bash

hermes --config scripts/relayer_hermes/config-testnet.toml keys delete --chain osmo-test-4 --all
hermes --config scripts/relayer_hermes/config-testnet.toml keys add --chain osmo-test-4 --mnemonic-file scripts/relayer_hermes/alice.json

hermes --config scripts/relayer_hermes/config-testnet.toml keys delete --chain feeappd-t1 --all
hermes --config scripts/relayer_hermes/config-testnet.toml keys add --chain feeappd-t1 --mnemonic-file scripts/relayer_hermes/bob.json

hermes --config scripts/relayer_hermes/config-testnet.toml start
