#!/usr/bin/env bash

set -euo pipefail

# Get the root of the Git repository.
repo_root=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")

# Source the shared functions file.
# shellcheck disable=SC1091
source "${repo_root}/scripts/lib/shared_functions.sh"

if ! check_namespace_prefix "${DEVSPACE_NAMESPACE}"; then
	exit 1
fi
