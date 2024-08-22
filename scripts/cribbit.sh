#!/usr/bin/env bash

set -euo pipefail

#############################
#                __________
#               < CRIBbit! >
#                ----------
#      _    _    /
#     (o)--(o)  /
#    /.______.\
#    \________/
#   ./        \.
#  ( .        , )
#   \ \_\\//_/ /
#    ~~  ~~  ~~
#
# Initialize your CRIB
# environment.
#############################

# Get the root of the Git repository
repo_root=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")

# If CRIB runs in CI most of the following checks should be skipped
: "${CRIB_CI_ENV:=false}"

# Source the shared functions file
# shellcheck disable=SC1091
source "${repo_root}/scripts/lib/shared_functions.sh"

# Initialize variables
provider=""
default_provider="aws"
DEVSPACE_NAMESPACE="${1:-}"

if [[ $CRIB_CI_ENV != "true" ]]; then
	# Prompt the user for the provider name if not set via environment variable
	if [ -z "$provider" ]; then
		read -r -p "Enter the provider name (supported are 'aws' and 'kind', default is 'aws'): " user_input
		provider=${user_input:-$default_provider}
	else
		echo "Using PROVIDER environment variable: $PROVIDER"
	fi

	if ! [[ $provider == "aws" || $provider == "kind" ]]; then
		echo "Error: Provider is not supported."
		exit 1
	fi
	export PROVIDER=${provider}

	# Check if the DEVSPACE_NAMESPACE environment variable is set
	if [ -n "$DEVSPACE_NAMESPACE" ]; then
		namespace_name=$DEVSPACE_NAMESPACE
		echo "Using namespace name from DEVSPACE_NAMESPACE environment variable: $namespace_name"
	elif [ "$provider" == "kind" ]; then
		namespace_name="crib-local"
		echo "Since the provider is 'kind', the suggested namespace name is 'crib-local'."
	else
		# Otherwise, ask the user for the namespace name
		read -r -p "Enter the namespace name, it should be in format crib-<your-username>: " user_input
		namespace_name=${user_input:-}
	fi

	##
	# Deploy Kind cluster
	##
	if [ "$PROVIDER" = "kind" ]; then
		# Execute the manage_kind.sh script
		"${repo_root}/scripts/manage_kind.sh"

		echo "Configured sucessfully"
		export SETUP_EKS_CONFIG=false
	fi
fi

# Use the first argument if provided, otherwise fall back to namespace_name
DEVSPACE_NAMESPACE="${1:-${namespace_name}}"
if [[ -z ${DEVSPACE_NAMESPACE} ]]; then
	echo "Usage: $0 <DEVSPACE_NAMESPACE>"
	exit 1
fi

if ! check_namespace_prefix "${DEVSPACE_NAMESPACE}"; then
	exit 1
fi

if [[ $CRIB_CI_ENV != "true" ]]; then
	# Automatically determine the directory name from the current working directory
	PRODUCT_DIR=$(basename "$(pwd)")

	# Path to the .env file
	env_file="${repo_root}/deployments/${PRODUCT_DIR}/.env"

	# Check if the .env file exists
	if [[ -f ${env_file} ]]; then
		echo "Info: Found ${env_file}."
	else
		echo "Error: '.env' file not found at ${env_file}."
		read -r -p "CRIB deployment requires several environment variables. Since you don’t have a custom '.env' file set up, would you like to use the predefined '.env' file instead? (yes/no): " choice
		if [[ $choice == "yes" || $choice == "y" ]]; then
			cp "${env_file}.example" "${env_file}"
			echo "Info: Copied ${env_file}.example to ${env_file} path."
		else
			echo "Error: Exiting without sourcing any '.env' file."
			exit 1
		fi
	fi

	# Source the .env file at the end if it exists
	# shellcheck disable=SC1090
	source "${env_file}"

	# List of required environment variables
	required_vars=(
		"AWS_ACCOUNT_ID"
		"DEVSPACE_IMAGE"
		"HOME"
		"PROVIDER"
	)

	missing_vars=0 # Counter for missing variables

	for var in "${required_vars[@]}"; do
		if [[ -z ${!var:-} ]]; then # If variable is unset or empty
			echo "Error: Environment variable ${var} is not set."
			missing_vars=$((missing_vars + 1))
		fi
	done

	# Exit with an error if any variables were missing
	if [[ $missing_vars -ne 0 ]]; then
		echo "Error: Total missing environment variables: $missing_vars"
		exit 1
	fi
fi

##
# Setup AWS Profile
##
path_aws_config="$HOME/.aws/config"
if [[ ${SETUP_AWS_PROFILE:-} != "false" ]]; then

	aws_profile_name="staging-crib"

	if grep -q "$aws_profile_name" "$path_aws_config"; then
		echo "Info: Skip updating ${path_aws_config}. Profile already set: ${aws_profile_name}"
	else
		# List of required environment variables
		required_aws_vars=(
			"AWS_ACCOUNT_ID"
			"AWS_REGION"
			# Should be the short name and not the full IAM role ARN.
			"AWS_SSO_ROLE_NAME"
			# The AWS SSO start URL, e.g. https://<org name>.awsapps.com/start
			"AWS_SSO_START_URL"
		)
		missing_aws_vars=0 # Counter for missing variables
		for var in "${required_aws_vars[@]}"; do
			if [[ -z ${!var:-} ]]; then # If variable is unset or empty
				echo "Error: Environment variable ${var} is not set."
				missing_aws_vars=$((missing_aws_vars + 1))
			fi
		done

		# Exit with an error if any variables were missing
		if [[ $missing_aws_vars -ne 0 ]]; then
			echo "Error: Total missing environment variables: $missing_aws_vars"
			exit 1
		fi

		cat <<EOF >>"$path_aws_config"
[profile $aws_profile_name]
region=${AWS_REGION}
sso_start_url=${AWS_SSO_START_URL}
sso_region=${AWS_REGION}
sso_account_id=${AWS_ACCOUNT_ID}
sso_role_name=${AWS_SSO_ROLE_NAME}
EOF
		echo "Info: ${path_aws_config} modified. Added profile: ${aws_profile_name}"
	fi

	echo "Info: Setting AWS Profile env var: AWS_PROFILE=${aws_profile_name}"
	export AWS_PROFILE=${aws_profile_name}

	if aws sts get-caller-identity >/dev/null 2>&1; then
		echo "Info: AWS credentials working."
	else
		echo "Info: AWS credentials not detected. Attempting to login through SSO."
		aws sso login
	fi

	# Check again and fail this time if not successful
	if ! aws sts get-caller-identity >/dev/null 2>&1; then
		echo "Error: AWS credentials still not detected. Exiting."
		exit 1
	fi
else
	echo "Info: The ENV variable SETUP_AWS_PROFILE is set to false, skipping the AWS profile setup."
fi

##
# Setup EKS KUBECONFIG
##

# Set env var SETUP_EKS_CONFIG=false to skip EKS config.
if [[ ${SETUP_EKS_CONFIG:-} != "false" ]]; then
	path_kubeconfig="${KUBECONFIG:-$HOME/.kube/config}"
	eks_cluster_name="${CRIB_EKS_CLUSTER_NAME:-main-stage-cluster}"
	eks_alias_name="${CRIB_EKS_ALIAS_NAME:-main-stage-cluster-crib}"

	if [[ ! -f ${path_kubeconfig} ]] || ! grep -q "name: ${eks_alias_name}" "${path_kubeconfig}"; then
		echo "Info: KUBECONFIG file (${path_kubeconfig}) not found or alias (${eks_alias_name}) not found. Attempting to update kubeconfig."
		aws eks update-kubeconfig \
			--name "${eks_cluster_name}" \
			--alias "${eks_alias_name}" \
			--region "${AWS_REGION}"
	else
		echo "Info: Alias '${eks_alias_name}' already exists in kubeconfig. No update needed."
		echo "Info: Setting kubernetes context to: ${eks_alias_name}"
		kubectl config use-context "${eks_alias_name}"
	fi

	##
	# Check Docker Daemon
	##

	if docker info >/dev/null 2>&1; then
		echo "Info: Docker daemon is running, authorizing registry"
	else
		echo "Error: Docker daemon is not running. Exiting."
		exit 1
	fi
fi

##
# AWS ECR Login
##

# Function to extract the host URI of the ECR registry from OCI URI
extract_ecr_host_uri() {
	local ecr_uri="${CHAINLINK_HELM_REGISTRY_URI:-}"

	# Regex to capture the ECR host URI
	if [[ $ecr_uri =~ oci:\/\/([0-9]+\.dkr\.ecr\.[a-zA-Z0-9-]+\.amazonaws\.com) ]]; then
		echo "${BASH_REMATCH[1]}"
	else
		echo "No valid ECR host URI found in the URI."
		# Print instructions for configuring environment variables
		echo "Configure CHAINLINK_HELM_REGISTRY_URI env var in your .env config"
		exit 1
	fi
}

# Set env var CRIB_SKIP_DOCKER_ECR_LOGIN=true to skip Docker ECR login.
skip_docker_ecr_login=${CRIB_SKIP_DOCKER_ECR_LOGIN:-}

# Check if the string starts with AWS ACCOUNT_ID
if is_custom_image "${DEVSPACE_IMAGE}"; then
	echo "DEVSPACE_IMAGE is set to use pre-built image $DEVSPACE_IMAGE, building from source is disabled, skipping ECR Login"
	skip_docker_ecr_login=true
fi

if [[ -n ${skip_docker_ecr_login:-} ]]; then
	echo "Info: Skipping ECR login."
else
	aws_account_id_ecr_registry=$(echo "${DEVSPACE_IMAGE}" | cut -d'.' -f1)
	echo "Info: Logging docker into AWS ECR registry."
	aws ecr get-login-password \
		--region "${AWS_REGION}" |
		docker login --username AWS \
			--password-stdin "${aws_account_id_ecr_registry}.dkr.ecr.${AWS_REGION}.amazonaws.com"
fi

##
# Helm ECR Login
##

# Set env var CRIB_SKIP_HELM_ECR_LOGIN=true to skip Helm ECR login.
skip_helm_ecr_login=${CRIB_SKIP_HELM_ECR_LOGIN:-}
if [[ -n ${skip_helm_ecr_login:-} ]]; then
	echo "Info: Skipping Helm ECR login."
else
	echo "Info: Logging helm into AWS ECR registry."
	helm_registry_uri=$(extract_ecr_host_uri)
	aws ecr get-login-password --region "${AWS_REGION}" |
		helm registry login "$helm_registry_uri" --username AWS --password-stdin
fi

##
# Setup DevSpace
##
devspace use namespace "${DEVSPACE_NAMESPACE}"
