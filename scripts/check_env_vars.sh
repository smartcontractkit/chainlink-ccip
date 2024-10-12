#!/usr/bin/env bash

# Default value for product type
product="${1:-}"

# List of required environment variables for CORE
required_vars_common=(
	"IS_CRIB"
	"CHAINLINK_CODE_DIR"
	"DEVSPACE_IMAGE"
	"DEVSPACE_INGRESS_CIDRS"
	"DEVSPACE_INGRESS_BASE_DOMAIN"
	"DEVSPACE_INGRESS_CERT_ARN"
	"DEVSPACE_K8S_POD_WAIT_TIMEOUT"
)

# Function to check environment variables
check_vars() {
	local vars=("$@")
	local missing_vars=0
	for var in "${vars[@]}"; do

		if [ -z "${!var}" ]; then # If variable is unset or empty
			echo "Error: Environment variable ${var} is not set."
			missing_vars=$((missing_vars + 1))
		fi
	done
	return $missing_vars
}

# Initialize the missing_vars counter
missing_vars_total=0

# Check each variable for CORE
check_vars "${required_vars_common[@]}"
missing_vars_total=$((missing_vars_total + $?))

# Check each variable for product CORE
if [[ $product == "" ]]; then
	required_vars_core=(
		"CHAINLINK_CLUSTER_HELM_CHART_URI"
	)
	check_vars "${required_vars_core[@]}"
	missing_vars_total=$((missing_vars_total + $?))
fi

# Check each variable for product CCIP
if [[ $product == "ccip" ]]; then
	required_vars_ccip=(
		"CHAINLINK_HELM_REGISTRY_URI"
	)
	check_vars "${required_vars_ccip[@]}"
	missing_vars_total=$((missing_vars_total + $?))
fi

# Check for keystone specific profiles
if [[ ${DEVSPACE_PROFILE} == "keystone" ]]; then
	required_vars_keystone=(
		"KEYSTONE_ETH_WS_URL"
		"KEYSTONE_ETH_HTTP_URL"
		"KEYSTONE_ACCOUNT_KEY"
	)
	check_vars "${required_vars_keystone[@]}"
	missing_vars_total=$((missing_vars_total + $?))
fi

# Exit with an error if any variables were missing
if [[ $missing_vars_total -ne 0 ]]; then
	echo "Total missing environment variables: $missing_vars_total"
	echo 'To fix it, add missing variables in the ".env" file.'
	echo 'You can find an example of the .env config in the ".env.example"'
	exit 1
else
	echo "All required environment variables are set."
fi
