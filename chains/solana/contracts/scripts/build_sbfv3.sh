#!/usr/bin/env bash
# Builds every program for sBPFv3 into target/deploy, replacing the sBPFv0 binaries
# that `anchor build` leaves there. Anchor 0.29 cannot do this itself: its CLI and
# docker image ship a cargo-build-sbf without `--arch`. Keep using `anchor build`
# for the IDLs; they are generated from the sources and do not depend on the binary.
#
# Needs cargo-build-sbf 4.2 or later (`cargo install cargo-build-sbf --locked`).
# With an older one, set SBF_TOOLS_VERSION=v1.57 to use platform-tools that emit a
# valid sBPFv3; older platform-tools produce an outdated format.
#
# The release build is scripts/verify_build.sh.
set -euo pipefail

CONTRACTS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$CONTRACTS_DIR"

OUT_DIR="${SBF_OUT_DIR:-$CONTRACTS_DIR/target/deploy}"
args=(--arch v3 --sbf-out-dir "$OUT_DIR")

if [ -n "${SBF_TOOLS_VERSION:-}" ]; then
  args+=(--tools-version "$SBF_TOOLS_VERSION")
else
  version="$(cargo build-sbf --version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)"
  major="${version%%.*}"
  minor="${version#*.}"
  minor="${minor%%.*}"
  if [ "$major" -lt 4 ] || { [ "$major" -eq 4 ] && [ "$minor" -lt 2 ]; }; then
    echo "error: cargo-build-sbf $version is too old to build sBPFv3." >&2
    echo "Install 4.2 or later, or set SBF_TOOLS_VERSION=v1.57." >&2
    exit 1
  fi
fi

# One build per program. A single build of the whole workspace unifies the
# `no-entrypoint` feature that programs enable on each other for CPI, and silently
# produces empty binaries for the programs that are depended on.
count=0
while read -r dir lib; do
  echo "=== $lib"
  cargo build-sbf --manifest-path "programs/$dir/Cargo.toml" "${args[@]}" -- --locked
  count=$((count + 1))
done < <(bash "$CONTRACTS_DIR/scripts/list_programs.sh")

echo "Built $count sBPFv3 programs into $OUT_DIR"
