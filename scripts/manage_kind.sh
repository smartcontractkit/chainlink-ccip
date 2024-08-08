#!/bin/bash
set -o errexit

CLUSTER_NAME="crib-cluster"
REGISTRY_NAME="kind-registry"
REGISTRY_PORT="5001"

# Function to create a kind cluster with two nodes
create_kind_cluster() {
	if kind get clusters | grep -q "^${CLUSTER_NAME}$"; then
		echo "Kind cluster named ${CLUSTER_NAME} already exists."
		return
	fi

	echo "Creating a kind cluster named ${CLUSTER_NAME} with two nodes..."
	cat <<EOF | kind create cluster --name ${CLUSTER_NAME} --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry]
    config_path = "/etc/containerd/certs.d"
nodes:
  - role: control-plane
    kubeadmConfigPatches:
    - |
      kind: InitConfiguration
      nodeRegistration:
        kubeletExtraArgs:
          node-labels: "ingress-ready=true"
    extraPortMappings:
    - containerPort: 80
      hostPort: 80
      protocol: TCP
    - containerPort: 443
      hostPort: 443
      protocol: TCP
  - role: worker
  - role: worker
EOF
	echo "Kind cluster created successfully."
}

# Function to configure kubectl context
configure_kubectl_context() {
	# Set the kubectl context
	kubectl config use-context kind-${CLUSTER_NAME}

	# Check if the context was set successfully
	CURRENT_CONTEXT=$(kubectl config current-context)
	if [ "$CURRENT_CONTEXT" == "kind-${CLUSTER_NAME}" ]; then
		echo "Kubectl context configured successfully to $CURRENT_CONTEXT."
	else
		echo "Failed to configure kubectl context. Current context is $CURRENT_CONTEXT."
		exit 1
	fi
}

# Function to create a Docker registry
create_docker_registry() {
	if [ "$(docker inspect -f '{{.State.Running}}' "${REGISTRY_NAME}" 2>/dev/null || true)" = 'true' ]; then
		echo "Docker registry named ${REGISTRY_NAME} is already running."
		return
	fi

	echo "Creating a local Docker registry..."
	docker run \
		-d --restart=always -p "127.0.0.1:${REGISTRY_PORT}:5000" --network bridge --name "${REGISTRY_NAME}" \
		registry:2
	echo "Docker registry created successfully."
}

# Function to add the registry config to the nodes
configure_registry_on_nodes() {
	REGISTRY_DIR="/etc/containerd/certs.d/localhost:${REGISTRY_PORT}"
	echo "Configuring registry on the cluster nodes..."
	for node in $(kind get nodes --name ${CLUSTER_NAME}); do
		docker exec "${node}" mkdir -p "${REGISTRY_DIR}"
		cat <<EOF | docker exec -i "${node}" cp /dev/stdin "${REGISTRY_DIR}/hosts.toml"
[host."http://${REGISTRY_NAME}:5000"]
EOF
	done
	echo "Registry configured on the cluster nodes successfully."
}

# Function to connect the registry to the cluster network
connect_registry_to_network() {
	if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${REGISTRY_NAME}")" != 'null' ]; then
		echo "Docker registry ${REGISTRY_NAME} is already connected to the cluster network."
		return
	fi

	echo "Connecting registry to the cluster network..."
	docker network connect "kind" "${REGISTRY_NAME}"
	echo "Registry connected to the cluster network successfully."
}

# Function to document the local registry in Kubernetes
document_local_registry() {
	echo "Documenting the local registry in Kubernetes..."
	cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${REGISTRY_PORT}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF
	echo "Local registry documented in Kubernetes successfully."
}

# Function to install Prometheus CRDs if they are not already installed
install_prometheus_crds() {
	local release_name="prometheus-crds"
	local namespace="default" # Update to your namespace if necessary

	echo "Checking for existing Helm release: $release_name..."

	# Check if the Helm release exists
	if helm list -n "$namespace" | grep -q "^$release_name"; then
		echo "Helm release $release_name already exists."
		return
	fi

	echo "Adding the prometheus-community Helm repository..."
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update

	echo "Installing Prometheus CRDs using kube-prometheus-stack Helm chart..."
	helm install "$release_name" prometheus-community/prometheus-operator-crds -n "$namespace"

	echo "Prometheus CRDs installed successfully."
}

# Function to delete the kind cluster if it exists
delete_kind_cluster() {
	if kind get clusters | grep -q "^${CLUSTER_NAME}$"; then
		echo "Deleting existing kind cluster named ${CLUSTER_NAME}..."
		kind delete cluster --name ${CLUSTER_NAME}
		echo "Kind cluster deleted successfully."
	else
		echo "Kind cluster named ${CLUSTER_NAME} does not exist."
	fi
}

# Function to delete the Docker registry if it exists
delete_docker_registry() {
	if [ "$(docker inspect -f '{{.State.Running}}' "${REGISTRY_NAME}" 2>/dev/null || true)" = 'true' ]; then
		echo "Deleting Docker registry named ${REGISTRY_NAME}..."
		docker stop ${REGISTRY_NAME}
		docker rm ${REGISTRY_NAME}
		echo "Docker registry deleted successfully."
	else
		echo "Docker registry named ${REGISTRY_NAME} is not running."
	fi
}

# Function to install NGINX ingress controller if it is not already installed
install_ingress_controller() {
	local ingress_namespace="ingress-nginx"
	local release_name="nginx-ingress"
	local chart_repo="https://kubernetes.github.io/ingress-nginx"
	local chart_name="ingress-nginx/ingress-nginx"

	echo "Checking for existing NGINX ingress controller..."

	# Check if the NGINX ingress controller is already installed
	if helm list -n "$ingress_namespace" | grep -q "$release_name"; then
		echo "NGINX ingress controller is already installed."
		return
	fi

	echo "Adding Helm repository for NGINX ingress controller..."
	helm repo add ingress-nginx "$chart_repo"
	helm repo update

	echo "Installing NGINX ingress controller using Helm..."
	helm install "$release_name" "$chart_name" --namespace "$ingress_namespace" --create-namespace

	echo "NGINX ingress controller installed and is now ready."
}

# Main script execution
main() {
	action=${1:-deploy}

	if [ "${action}" = "purge" ]; then
		delete_kind_cluster
		delete_docker_registry
		echo "Purged the existing cluster and Docker registry."
		exit 0
	elif [ "${action}" != "deploy" ]; then
		echo "Invalid action: ${action}. Use 'deploy', 'purge', or 'dns'."
		exit 1
	fi

	echo "Starting Kubernetes cluster deployment using kind..."

	create_docker_registry
	create_kind_cluster
	configure_registry_on_nodes
	connect_registry_to_network
	configure_kubectl_context
	document_local_registry
	install_prometheus_crds
	install_ingress_controller

	echo "Kubernetes cluster deployed and configured successfully"
}

# Call the main function
main "$@"
