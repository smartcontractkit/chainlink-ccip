#!/usr/bin/env bash

#DEVSPACE_NAMESPACE="crib-local" \
#DEVSPACE_INGRESS_BASE_DOMAIN="crib.local" \
#  go run main.go deploy-home-chain -o .tmp

TMP_DIR="/tmp/ccip-v2/"

DEVSPACE_NAMESPACE="crib-local" \
	DEVSPACE_INGRESS_BASE_DOMAIN="crib.local" \
	go run main.go deploy-home-chain \
	--deploy-home-out "$TMP_DIR"

#DEVSPACE_NAMESPACE="crib-local" \
#	DEVSPACE_INGRESS_BASE_DOMAIN="crib.local" \
#	go run main.go deploy-ccip \
#	--deploy-ccip-in "$TMP_DIR" \
#	--deploy-ccip-out "$TMP_DIR"
