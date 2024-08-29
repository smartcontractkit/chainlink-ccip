#!/usr/bin/env bash

set -euo pipefail

# Get the root of the Git repository.
repo_root=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")

# Source the shared functions file.
# shellcheck disable=SC1091
source "${repo_root}/scripts/lib/shared_functions.sh"

if is_custom_image "${DEVSPACE_IMAGE}"; then
	echo "DEVSPACE_IMAGE var was set to $DEVSPACE_IMAGE, which is a non standard image, use --skip-build and -o <image_tag> options to skip the build entirely"
	exit 1
fi

if [[ -z ${CHAINLINK_REPO_DIR:-} ]]; then echo "Error: the CHAINLINK_REPO_DIR environment variable is not set."; fi
if check_repo_exists "${CHAINLINK_REPO_DIR}"; then
	echo "Info: Repository exists at ${CHAINLINK_REPO_DIR}."
else
	echo "Error: repository does not exist at ${CHAINLINK_REPO_DIR}."
	exit 1
fi
