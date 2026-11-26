#!/usr/bin/env bash
# Prints one "<program dir> <library name>" pair per line, for every program in
# programs/. The library name is what `anchor build` names the .so and the IDL.
set -euo pipefail

CONTRACTS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

for manifest in "$CONTRACTS_DIR"/programs/*/Cargo.toml; do
  dir="$(basename "$(dirname "$manifest")")"
  lib="$(awk '
    /^\[lib\]/ { in_lib = 1; next }
    /^\[/      { in_lib = 0 }
    in_lib && /^name[[:space:]]*=/ { gsub(/[" ]/, "", $0); sub(/^name=/, "", $0); print; exit }
  ' "$manifest")"
  if [ -z "$lib" ]; then
    echo "error: no [lib] name in $manifest" >&2
    exit 1
  fi
  echo "$dir $lib"
done
