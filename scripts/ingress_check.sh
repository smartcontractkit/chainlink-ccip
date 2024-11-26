#!/usr/bin/env bash

set -euo pipefail

###
# To be invoked by `devspace` after a successful DevSpace deploy via a hook.
###

# Check if DEVSPACE_NAMESPACE is set
if [[ -z ${DEVSPACE_NAMESPACE:-} ]]; then
	echo "Error: 'DEVSPACE_NAMESPACE' env variable isn't set. Make sure to run from devspace."
	exit 1
fi

# Initialize variables
max_retries=10
sleep_duration_retry=10     # 10 seconds
sleep_duration_propagate=60 # 60 seconds
timeout=$((60 * 2))         # 2 minutes
elapsed=0                   # Track the elapsed time
namespace=${DEVSPACE_NAMESPACE}

print_hostnames=false

# Parse arguments
for arg in "$@"; do
	case $arg in
	--print-hostnames)
		print_hostnames=true
		shift
		;;
	*) ;;
	esac
done

print_ingress_hostnames() {
	echo "#############################################################"
	echo "###    Ingress hostnames"
	echo "#############################################################"

	jsonpath=""
	if [[ $PROVIDER == "kind" ]]; then
		jsonpath="{range .items[*].spec.rules[*]}{'http://'}{.host}{'\n'}{end}"
	else
		jsonpath="{range .items[*].spec.rules[*]}{'https://'}{.host}{'\n'}{end}"
	fi

	kubectl get ingress -n "${DEVSPACE_NAMESPACE}" -o=jsonpath="$jsonpath"
}

if $print_hostnames; then
	print_ingress_hostnames
	exit 0
fi

# Function to check if a hostname can be resolved
check_hostname_resolution() {
	local hostname="$1"
	local ns_timeout=$2
	local interval=$3

	local start_time
	start_time=$(date +%s)

	while true; do
		if dig "$hostname" +short >/dev/null 2>&1; then
			echo "DNS lookup successful for $hostname"
			return 0
		fi

		local current_time
		current_time=$(date +%s)

		local elapsed_time
		elapsed_time=$((current_time - start_time))
		if [ $elapsed_time -ge "$ns_timeout" ]; then
			echo "Error: DNS lookup failed after $ns_timeout seconds for $hostname."
			return 1
		fi

		# Wait for the specified interval before trying again
		echo "DNS lookup for $hostname failed, retrying in $interval seconds..."
		sleep "$interval"
	done
}

# Get all ingress names in the current namespace
if ! mapfile -t ingresses < <(kubectl get ingress -n "${namespace}" -o jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}' 2>/dev/null); then
	echo "Error: Failed to retrieve ingresses from namespace '${namespace}'."
	exit 1
fi

# Loop through each ingress
for INGRESS_NAME in "${ingresses[@]}"; do
	echo "Validating ingress [${namespace}/${INGRESS_NAME}]..."

	# Fetch the ingress class from the Ingress object
	INGRESS_CLASS=$(kubectl get ingress "${INGRESS_NAME}" -n "${namespace}" \
		-o jsonpath='{.spec.ingressClassName}' 2>/dev/null || echo "")

	# Determine ingress_suffix based on ingress class
	case "${INGRESS_CLASS}" in
	"alb")
		ingress_suffix=".elb.amazonaws.com"
		;;
	"nginx")
		ingress_suffix="localhost"
		;;
	*)
		echo "Error: Unsupported ingress class '${INGRESS_CLASS}' for ingress '${INGRESS_NAME}'. Supported classes are 'alb' and 'nginx'."
		exit 1
		;;
	esac

	# Loop until conditions are met or we reach max retries or timeout
	for ((i = 1; i <= max_retries && elapsed <= timeout; i++)); do
		ingress_hostname=$(kubectl get ingress "${INGRESS_NAME}" -n "${namespace}" \
			-o jsonpath='{.status.loadBalancer.ingress[0].hostname}' 2>/dev/null || echo "")

		# Check if ingress_hostname matches the expected value
		if [[ ${ingress_hostname} == *"${ingress_suffix}" ]]; then

			# Fetch all ingress hostnames from the specified namespace
			ingress_hosts=$(kubectl get ingress "${INGRESS_NAME}" -n "${namespace}" -o jsonpath='{.spec.rules[*].host}' | tr ' ' '\n')

			if [ -z "$ingress_hosts" ]; then
				echo "No ingress hostnames found in namespace."
				exit 1
			fi

			if [[ ${INGRESS_CLASS} == "nginx" ]]; then
				echo "Configuring hosts file for namespace..."

				# Define the IP address to use
				ingress_ip="127.0.0.1"

				# Add each hostname to the /etc/hosts file
				for host in $ingress_hosts; do
					# Check if the hostname is already in /etc/hosts
					if ! grep -q "$host" /etc/hosts; then
						echo "$ingress_ip $host" | sudo tee -a /etc/hosts >/dev/null
						echo "Added $host to /etc/hosts."
					fi
				done

				echo "Hosts file configured successfully."
			fi

			# Check resolution for each hostname
			if [[ ${CRIB_CI_ENV:-false} != "true" ]]; then
				for hostname in $ingress_hosts; do
					if ! check_hostname_resolution "$hostname" $sleep_duration_propagate $sleep_duration_retry; then
						echo "Error: DNS lookup failed for $hostname after $sleep_duration_propagate seconds."
						exit 1
					fi
				done
			fi

			break
		else
			echo "Attempt $i: Waiting for the ingress '${INGRESS_NAME}' to be created..."
			sleep $sleep_duration_retry
			((elapsed += sleep_duration_retry))
		fi
	done

	# If we reached here, it means we hit the retry limit or the timeout
	if [[ $i -gt max_retries ]]; then
		echo "Error: Ingress '${INGRESS_NAME}' was not successfully created within the given constraints."
	fi
done

print_ingress_hostnames

exit 0
