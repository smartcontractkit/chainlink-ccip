#!/usr/bin/env bash

set -e

function generate_bindings() {
  local idl_path_str="$1"
  IFS='/' read -r -a idl_path <<< "${idl_path_str}"
  IFS='.' read -r -a idl_name <<< "${idl_path[3]}"
  anchor-go -src "${idl_path_str}" -dst ./gobindings/latest/"${idl_name}" -codec borsh
}


for idl_path_str in "contracts/target/idl"/*
do
  generate_bindings "${idl_path_str}"
done
for idl_path_str in "contracts/target/vendor"/*.json
do
  generate_bindings "${idl_path_str}"
done

go fmt ./...
