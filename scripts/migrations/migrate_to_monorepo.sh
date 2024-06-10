#!/usr/bin/env bash

set -ex

CHAINLINK_CODE_DIR="${1}"

help() {
  echo "missing argument: $0 <absolute path to CHAINLINK_CODE_DIR>"
  echo "CHAINLINK_CODE_DIR is a directory where you checkout clone chainlink repositories locally"
}

if [ $# -ne 1 ]; then
  help
  exit 1
fi

pushd "$CHAINLINK_CODE_DIR"

if [  ! -d "${CHAINLINK_CODE_DIR}/crib" ]; then
  echo "crib repo is already cloned"
  git clone org-25111032@github.com:smartcontractkit/crib.git
fi

# Move .env files
if [ -f "${CHAINLINK_CODE_DIR}/chainlink/crib/.env" ]; then
  old_file="${CHAINLINK_CODE_DIR}/chainlink/crib/.env"
  new_file="${CHAINLINK_CODE_DIR}/crib/core/.env"
  echo -e "# The path to a directory with chainlink repositories\nCHAINLINK_CODE_DIR=\"${CHAINLINK_CODE_DIR}\"\n\n$(cat "$old_file")" > "$new_file"
fi

if [ -f "${CHAINLINK_CODE_DIR}/ccip/crib/.env" ]; then
  old_file="${CHAINLINK_CODE_DIR}/ccip/crib/.env"
  new_file="${CHAINLINK_CODE_DIR}/crib/ccip/.env"
  echo -e "# The path to a directory with chainlink repositories\nCHAINLINK_CODE_DIR=\"${CHAINLINK_CODE_DIR}\"\n\n$(cat "$old_file")" > "$new_file"
fi

popd
