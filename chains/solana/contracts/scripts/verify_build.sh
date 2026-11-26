#!/usr/bin/env bash
# Release build for the Solana programs: one `solana-verify build --library-name
# <lib>` per program, the documented standard for multi-program repos.
#
# The programs are built for sBPFv3. Once SIMD-0500 is active, the network rejects
# deployments and upgrades of any older format.
#
# solana-verify prints the executable hash of the built program as its last output
# line; those hashes are collected into target/executable-hashes.txt and shipped in
# the release tarball so deployments can be checked against them.
set -euo pipefail

# Pinned because a remote verifier cannot resolve it: it mounts the repository root
# and misses [workspace.metadata.cli] in contracts/Cargo.toml. Keep the tag equal to
# that solana version.
BASE_IMAGE="solanafoundation/solana-verifiable-build:4.3.0"

CONTRACTS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$CONTRACTS_DIR"

HASHES_FILE="$CONTRACTS_DIR/target/executable-hashes.txt"
mkdir -p "$CONTRACTS_DIR/target"
: > "$HASHES_FILE"

programs="$(bash "$CONTRACTS_DIR/scripts/list_programs.sh")"
count="$(printf '%s\n' "$programs" | wc -l | tr -d ' ')"
echo "Building $count programs with solana-verify for sBPFv3..."

while read -r _dir lib; do
  echo "=== $lib"
  log="$(mktemp)"
  if ! solana-verify build --library-name "$lib" --base-image "$BASE_IMAGE" --arch v3 2>&1 | tee "$log"; then
    rm -f "$log"
    echo "error: solana-verify build failed for $lib" >&2
    exit 1
  fi
  hash="$(grep -oE '^[0-9a-f]{64}$' "$log" | tail -1 || true)"
  rm -f "$log"
  if [ -z "$hash" ]; then
    echo "error: solana-verify did not print an executable hash for $lib" >&2
    exit 1
  fi
  artifact="$CONTRACTS_DIR/target/deploy/$lib.so"
  if [ ! -f "$artifact" ]; then
    echo "error: solana-verify did not produce expected artifact $artifact" >&2
    exit 1
  fi
  echo "$hash  $lib.so" >> "$HASHES_FILE"
done <<< "$programs"

echo "Wrote $HASHES_FILE"
