#!/usr/bin/env bash

go run main.go -chain-overrides-dir="./../../values/chain-overrides/" \
	-chain-overrides-file="example.yaml" \
	-chains-count="4" \
	-provider="aws" \
	-product="ccip"
