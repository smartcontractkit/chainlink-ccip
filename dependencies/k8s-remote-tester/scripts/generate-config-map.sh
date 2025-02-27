#!/bin/bash

set -x

ENV_STATE_DIR="../../deployments/ccip-v2/.tmp"

# Check if the directory exists
if [ ! -d "$ENV_STATE_DIR" ]; then
	echo "Error: Directory $ENV_STATE_DIR does not exist."
	exit 1
fi

# Generate the ConfigMap YAML file
CONFIGMAP_NAME="k8s-remote-tester-crib-env-state-cm"
OUTPUT_FILE="./generated-manifests/configmap.yaml"

mkdir -p ./generated-manifests

echo "Creating ConfigMap $CONFIGMAP_NAME with JSON files from $JSON_DIR..."

# Create the ConfigMap with all .json files in the directory
kubectl create configmap $CONFIGMAP_NAME \
	--from-file=$ENV_STATE_DIR \
	-o yaml --dry-run=client >$OUTPUT_FILE

echo "ConfigMap YAML file created: $OUTPUT_FILE"
