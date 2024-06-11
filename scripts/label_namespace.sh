#!/bin/bash

set -euo pipefail

# Check if both parameters are provided, otherwise use default values
namespace="${1:-}"
ttl="${2:-72h}"

# Check if the namespace name is provided
if [ -z "$namespace" ]; then
    echo "Error: Namespace name is not provided."
    exit 1
fi

# Function to print the success message
print_info() {
    echo -e "
*****************************************************************************
 *
  *    Namespace ${namespace} will be deleted in ${ttl}
   *   To extend the TTL for e.g. 72 hours, run:
  *    devspace run ttl ${namespace} 72h
 *
*****************************************************************************
"
}

# Check if the cleanup.kyverno.io/ttl label is set on the namespace
ttl_label=$(kubectl get namespace "$namespace" -o jsonpath='{.metadata.labels.cleanup\.kyverno\.io/ttl}' || echo "")

# If the label is not set, apply it
if [ -z "$ttl_label" ]; then
    echo "Setting cleanup.kyverno.io/ttl: $ttl on namespace $namespace"
    if kubectl label namespace "$namespace" cleanup.kyverno.io/ttl="$ttl"; then
        echo "Successfully set cleanup.kyverno.io/ttl: $ttl label on namespace $namespace"
        print_info
        exit 0
    else
        echo "Failed to set cleanup.kyverno.io/ttl: $ttl label on namespace $namespace"
        print_info
        exit 1
    fi
else
    echo "Namespace $namespace already has cleanup.kyverno.io/ttl set to $ttl_label"
    print_info
    exit 0
fi
