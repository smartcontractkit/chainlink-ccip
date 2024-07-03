#!/usr/bin/env bash

set -euo pipefail

# Check if the namespace name is provided in the environment variable
namespace="${DEVSPACE_NAMESPACE:-}"
if [[ -z $namespace ]]; then
	echo "Error: ENV variable DEVSPACE_NAMESPACE is not set."
	exit 1
fi

ttl="${2:-72h}"
overwrite="${3:-}"

# Function to print the success message
print_info() {
	echo -e "
*****************************************************************************
 *
 *    Namespace ${namespace} will be deleted in ${ttl}
 *    To extend the TTL for e.g. 96 hours, run:
 *    devspace run ttl 96h
 *
*****************************************************************************
"
}

# Function to check if the role binding exists
check_role_binding() {
	for _ in {1..3}; do
		if kubectl get rolebinding "${namespace}-crib-poweruser" -n "$namespace" >/dev/null 2>&1; then
			echo "Role binding ${namespace}-crib-poweruser found in namespace $namespace"
			return
		else
			echo "Role binding ${namespace}-crib-poweruser not found. Retrying in 5 seconds..."
			sleep 5
		fi
	done

	echo "Failed to find role binding ${namespace}-crib-poweruser in namespace $namespace after 3 attempts"
	exit 1
}

# Function to create the namespace
create_namespace() {
	if kubectl get namespace "$namespace" >/dev/null 2>&1; then
		echo "Namespace $namespace already exists."
	else
		echo "Creating namespace $namespace"
		if kubectl create namespace "$namespace"; then
			echo "Successfully created namespace $namespace."
		else
			echo "Failed to create namespace $namespace"
			exit 1
		fi
	fi

	check_role_binding "$namespace"
}

# Function to label the namespace
label_namespace() {
	# Determine if overwrite is required
	local overwrite_flag=""
	if [[ $overwrite == "--overwrite" ]]; then
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
}

# Main script logic to call create or label based on the first argument
if [[ $# -lt 1 ]]; then
	echo "Usage: $0 {create|label} [ttl] [--overwrite]"
	exit 1
fi

action="$1"
shift

case "$action" in
create)
	create_namespace
	;;
label)
	label_namespace "$@"
	;;
*)
	echo "Invalid action: $action"
	echo "Usage: $0 {create|label} [ttl] [--overwrite]"
	exit 1
	;;
esac
