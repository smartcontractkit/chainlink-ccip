#!/bin/bash

set -euo pipefail

# Check if both parameters are provided, otherwise use default values
namespace="${DEVSPACE_NAMESPACE}"
ttl="${1:-72h}"
overwrite="${2:-}"

# Check if the namespace name is provided
if [ -z "$namespace" ]; then
    echo "Error: ENV variable DEVSPACE_NAMESPACE is not set."
    exit 1
fi

# Function to print the success message
print_info() {
    echo -e "
*****************************************************************************
 *
  *    Namespace ${namespace} will be deleted in ${ttl}
   *   To extend the TTL for e.g. 96 hours, run:
  *    devspace run ttl 96h
 *
*****************************************************************************
"
}

# Determine if overwrite is required
overwrite_flag=""
if [ "$overwrite" == "--overwrite" ]; then
    overwrite_flag="--overwrite"
fi

# Apply the label
echo "Setting cleanup.kyverno.io/ttl: $ttl on namespace $namespace"
if kubectl label namespace "$namespace" cleanup.kyverno.io/ttl="$ttl" $overwrite_flag; then
    echo "Successfully set cleanup.kyverno.io/ttl: $ttl label on namespace $namespace"
    print_info
    exit 0
else
    echo "Failed to set cleanup.kyverno.io/ttl: $ttl label on namespace $namespace"
    exit 1
fi
