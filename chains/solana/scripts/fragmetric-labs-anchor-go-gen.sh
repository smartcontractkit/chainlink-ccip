#!/usr/bin/env bash

# Script run from the anchor-go-gen repo

set -e

for idl_path_str in "../chainlink-ccip/chains/solana/contracts/target/idl"/*
do
  IFS='/' read -r -a idl_path <<< "${idl_path_str}"
  IFS='.' read -r -a idl_name <<< "${idl_path[7]}"
  # skip if ccip_offramp
  if [[ "${idl_name[0]}" == "ccip_offramp" ]]; then
    continue
  fi

  ./solana-anchor-go -src "${idl_path_str}" -dst=../chainlink-ccip/chains/solana/gobindings/"${idl_name}" -codec borsh
done

# go fmt ./...

