#!/bin/bash

# todo: this is quick fix, and we should replace it with crib cli, once crib cli support ingress_check

GRPC_SERVER=$1    # The gRPC server address (e.g., localhost:50051)
GRPC_METHOD=$2    # The full gRPC method to query (e.g., package.Service.Method)
TIMEOUT=${3:-300} # Timeout in seconds (default: 300)
INTERVAL=${4:-5}  # Check interval in seconds (default: 5)

START=$(date +%s)

echo "Waiting for gRPC server: $GRPC_SERVER, method: $GRPC_METHOD"
while true; do
	# Try calling the gRPC method
	if grpcurl "$GRPC_SERVER" "$GRPC_METHOD" >/dev/null 2>&1; then
		echo "gRPC server is ready: $GRPC_SERVER"
		exit 0
	fi

	NOW=$(date +%s)
	ELAPSED=$((NOW - START))
	if [ $ELAPSED -ge "$TIMEOUT" ]; then
		echo "Timeout reached! gRPC server is not ready: $GRPC_SERVER"
		exit 1
	fi

	sleep "$INTERVAL"
done
