#!/usr/bin/env bash

set -e

# Delete all .go files (anchor-go v1.0.0 generates all needed types directly)
find ./gobindings/latest -name "*.go" -type f -delete

function generate_bindings() {
  local idl_path_str="$1"
  IFS='/' read -r -a idl_path <<< "${idl_path_str}"
  IFS='.' read -r -a idl_name <<< "${idl_path[3]}"
  anchor-go -idl "${idl_path_str}" -output ./gobindings/latest/"${idl_name}" -no-go-mod
}

# Generate bindings for all IDLs (including vendor)
for idl_path_str in "contracts/target/idl"/*
do
  generate_bindings "${idl_path_str}"
done

for idl_path_str in "contracts/target/vendor"/*.json
do
  generate_bindings "${idl_path_str}"
done

go fmt ./...
