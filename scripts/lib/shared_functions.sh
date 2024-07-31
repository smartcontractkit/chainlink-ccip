#!/usr/bin/env bash

set -euo pipefail

# Bail if $DEVSPACE_NAMESPACE does not begin with a crib- prefix unless an override is set.
function check_namespace_prefix() {
	local DEVSPACE_NAMESPACE="${1:-}"

	if [[ ! ${DEVSPACE_NAMESPACE} =~ ^crib- ]] && [[ -z ${CRIB_IGNORE_NAMESPACE_PREFIX:-} ]]; then
		echo "Error: DEVSPACE_NAMESPACE must begin with 'crib-' prefix unless the CRIB_IGNORE_NAMESPACE_PREFIX env var is set." >&2
		return 1
	fi
}

# Check if DEVSPACE_IMAGE is overridden to use custom image
function is_custom_image() {
	local DEVSPACE_IMAGE="${1:-}"

	if ! [[ $DEVSPACE_IMAGE =~ ^323150190480\.dkr\.ecr\.us-west-2\.amazonaws\.com\/chainlink-?[a-z]*-devspace$ ]]; then
		return 0
	fi
	return 1
}
