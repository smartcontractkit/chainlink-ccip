#!/bin/bash
set -e

# Source shared functions
# shellcheck disable=SC1091
source "../../scripts/lib/shared_functions.sh"

echo "Keystone profile detected, provisioning Keystone"

build_images --all

CHAIN_ID=${CHAIN_ID:-1337}
KEYSTONE_DIR=$(realpath "${CHAINLINK_REPO_DIR}/core/scripts/keystone")
ARTEFACTS_DIR=$(realpath "${CHAINLINK_REPO_DIR}/core/scripts/keystone/artefacts")

build_keystone() {
	pushd "$KEYSTONE_DIR" >/dev/null
	go build -o keystone main.go
	popd >/dev/null
}

preprovision_keystone() {
	pushd "$KEYSTONE_DIR" >/dev/null
	./keystone provision-keystone \
		--preprovision=true \
		--artefacts="${ARTEFACTS_DIR}" \
		--chainid="$CHAIN_ID"
	popd >/dev/null

	if [ "$(get_flag "skip_provision")" == "false" ]; then
		deploy_app "${ARTEFACTS_DIR}/crib-preprovision.yaml"
	fi
	../../scripts/ingress_check.sh
	kubectl label namespace/"${DEVSPACE_NAMESPACE}" network=crib >/dev/null 2>&1 || true
}

postprovision_keystone() {
	pushd "$KEYSTONE_DIR" >/dev/null
	./keystone provision-keystone \
		--preprovision=false \
		--ethurl="https://${DEVSPACE_NAMESPACE}-geth-${CHAIN_ID}-http.${DEVSPACE_INGRESS_BASE_DOMAIN}" \
		--chainid="$CHAIN_ID" \
		--accountkey="${KEYSTONE_ACCOUNT_KEY}" \
		--artefacts="${ARTEFACTS_DIR}" \
		--clean=false
	popd >/dev/null
	deploy_app "${ARTEFACTS_DIR}/crib-postprovision.yaml"
	restart_app_pods
}

deploy_app() {
	local file_path="$1"
	create_deployments app --from-file="$file_path"
}

restart_app_pods() {
	kubectl delete pod -l "app=app"
	kubectl wait --for=condition=ready pod -l "app=app" --timeout=2m
}

# Main execution flow
build_keystone
preprovision_keystone
postprovision_keystone
