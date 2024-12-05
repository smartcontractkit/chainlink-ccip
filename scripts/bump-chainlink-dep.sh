#!/usr/bin/env bash

set -xeuo pipefail

# This script is work in progress, it can potentially be replace by the tool for Chad
# More details here: https://smartcontract-it.atlassian.net/browse/CRIB-574

help() {
	echo "Must provide 1 argument: $0 <chainlink repo ref>"
}
if [ $# -ne 1 ]; then
	help
	exit 1
fi
# Init variables
ref="${1}"

echo "updating chainlink refs in sdk module"

pushd ./sdk
go get "github.com/smartcontractkit/chainlink/deployment@${ref}"
go mod tidy

# todo: update chainlink/v2 references to use the same ref as deployment package
# sed -E -i.bak "s|(github\.com/smartcontractkit/chainlink/v2 v2\.[0-9]+\.[0-9]+-)[a-zA-Z0-9_-]+|\1$ref|g" "go.mod"
go mod tidy
popd

echo "updating chainlink refs in ccip-v2-scripts"

pushd ./dependencies/ccip-v2-scripts
go get "github.com/smartcontractkit/chainlink/deployment@${ref}"
go mod tidy
# todo: Use sed to replace the dynamic suffix
# update chainlink/v2 references to use the same ref as deployment package
# sed -E -i.bak "s|(github\.com/smartcontractkit/chainlink/v2 v2\.[0-9]+\.[0-9]+-)[a-zA-Z0-9_-]+|\1$ref|g" "go.mod"
go mod tidy
popd
