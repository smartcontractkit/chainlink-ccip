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
	local DEVSPACE_NAMESPACE="${DEVSPACE_NAMESPACE:-}"
	local DEVSPACE_IMAGE="${1:-}"

	# Regular expression to match the desired image patterns
	if [[ $DEVSPACE_NAMESPACE == "crib-local" && $DEVSPACE_IMAGE =~ ^localhost:5001\/chainlink-*[a-z]*-devspace(:[a-zA-Z0-9._-]+)?$ ]]; then
		return 1
	elif [[ $DEVSPACE_IMAGE =~ ^323150190480\.dkr\.ecr\.us-west-2\.amazonaws\.com\/chainlink-*[a-z]*-devspace(:[a-zA-Z0-9._-]+)?$ ]]; then
		return 1
	else
		return 0
	fi
}

# Checks if the required repository dir exists and if it is a git repository.
function check_repo_exists() {
	local repo_dir="${1:-}"
	
	# Check if it's a Git directory
	if [[ -d "${repo_dir}/.git" ]]; then
		return 0
	fi
	
	# If it's not a Git directory, check if it's a Git file (e.g. in case of worktrees / submodules )
	if [[ -f "${repo_dir}/.git" ]]; then
		return 0
	fi
	
	# If neither a Git directory nor a Git file, return failure
	return 1
}
