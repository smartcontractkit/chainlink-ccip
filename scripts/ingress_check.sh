#!/usr/bin/env bash

set -euo pipefail

###
# To be invoked by `devspace` after a successful DevSpace deploy via a hook.
###

# Check if DEVSPACE_HOOK_KUBE_NAMESPACE is set
if [[ -z ${DEVSPACE_HOOK_KUBE_NAMESPACE:-} ]]; then
	echo "Error: DEVSPACE_HOOK_KUBE_NAMESPACE is not set. Make sure to run from devspace."
	exit 1
fi

# Check if INGRESS_NAME is provided
INGRESS_NAME="${1:-}"
if [[ -z ${INGRESS_NAME} ]]; then
	echo "Usage: $0 INGRESS_NAME"
	exit 1
fi

# Initialize variables
max_retries=10
sleep_duration_retry=10     # 10 seconds
sleep_duration_propagate=60 # 60 seconds
timeout=$((60 * 2))         # 2 minutes
elapsed=0                   # Track the elapsed time
ingress_suffix=""

# Fetch the ingress class from the Ingress object
INGRESS_CLASS=$(kubectl get ingress "${INGRESS_NAME}" -n "${DEVSPACE_HOOK_KUBE_NAMESPACE}" \
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
	echo "Error: Unsupported ingress class '${INGRESS_CLASS}'. Supported classes are 'alb' and 'nginx'."
	exit 1
	;;
esac

echo "Ingress suffix: ${ingress_suffix}"

# Loop until conditions are met or we reach max retries or timeout
for ((i = 1; i <= max_retries && elapsed <= timeout; i++)); do
	ingress_hostname=$(kubectl get ingress "${INGRESS_NAME}" -n "${DEVSPACE_HOOK_KUBE_NAMESPACE}" \
		-o jsonpath='{.status.loadBalancer.ingress[0].hostname}')

	# Check if ingress_hostname matches the expected value
	if [[ ${ingress_hostname} == *"${ingress_suffix}" ]]; then
		echo "#############################################################"
		echo "# Ingress hostnames:"
		echo "#############################################################"
		devspace run ingress-hosts
		echo

		# If the ingress class is nginx, skip the sleep for propagation
		if [[ ${INGRESS_CLASS} == "alb" ]]; then
			echo "Sleeping for ${sleep_duration_propagate} seconds to allow DNS records to propagate... (Use CTRL+C to safely skip this step.)"
			sleep $sleep_duration_propagate
			echo "...done. NOTE: If you have an issue with the DNS records, try to reset your local and/or VPN DNS cache."
		fi

		if [[ ${INGRESS_CLASS} == "nginx" && ${PROVIDER} == "kind" ]]; then
			namespace="${DEVSPACE_NAMESPACE}"

			if [ -z "$namespace" ]; then
				echo "Error: DEVSPACE_NAMESPACE environment variable is not set."
				return 1
			fi

			echo "Configuring hosts file for namespace $namespace..."

			# Fetch all ingress hostnames from the specified namespace
			ingress_ips=$(kubectl get ingress -n "$namespace" -o jsonpath='{.items[*].spec.rules[*].host}' | tr ' ' '\n')

			if [ -z "$ingress_ips" ]; then
				echo "No ingress hostnames found in namespace $namespace."
				return 1
			fi

			# Define the IP address to use
			ingress_ip="127.0.0.1"

			# Add each hostname to the /etc/hosts file
			for host in $ingress_ips; do
				# Check if the hostname is already in /etc/hosts
				if grep -q "$host" /etc/hosts; then
					echo "Hostname $host already exists in /etc/hosts."
				else
					echo "$ingress_ip $host" | sudo tee -a /etc/hosts >/dev/null
					echo "Added $host to /etc/hosts."
				fi
			done

			echo "Hosts file configured successfully."
		fi

		exit 0
	else
		echo "Attempt $i: Waiting for the ingress to be created..."
		sleep $sleep_duration_retry
		((elapsed += sleep_duration_retry))
	fi
done

# If we reached here, it means we hit the retry limit or the timeout
echo "Error: Ingress was not successfully created within the given constraints."
exit 1
