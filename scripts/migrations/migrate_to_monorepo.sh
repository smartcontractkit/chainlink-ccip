#!/usr/bin/env bash

set -e

CHAINLINK_CODE_DIR="${1}"

help() {
  echo "missing argument: $0 <absolute path to CHAINLINK_CODE_DIR>"
  echo "CHAINLINK_CODE_DIR is a directory where you clone/checkout chainlink repositories"
  echo "example $0 ~/code"
}

if [ $# -ne 1 ]; then
  help
  exit 1
fi

pushd "$CHAINLINK_CODE_DIR"

if [  ! -d "${CHAINLINK_CODE_DIR}/crib" ]; then
  echo "cloning crib repo"
  git clone org-25111032@github.com:smartcontractkit/crib.git
fi

# Copy .env files

update_dot_env() {
  local old_file=$1
  local new_file=$2

  echo "Updating $old_file and copying to crib repo"
  if [ -f "$new_file" ]; then
    echo "$new_file already exists, skipping"
  elif [ -f "$old_file" ]; then
    echo "copying and updating ccip/crib/.env"
    echo -e "# The path to a directory with chainlink repositories\nCHAINLINK_CODE_DIR=\"${CHAINLINK_CODE_DIR}\"\n\n$(cat "$old_file")" > "$new_file"
  else
    echo "couldn't find $old_file, skipping"
  fi
}

# Copy env file for chainlink repo
old_file="${CHAINLINK_CODE_DIR}/chainlink/crib/.env"
new_file="${CHAINLINK_CODE_DIR}/crib/core/.env"

update_dot_env "$old_file" "$new_file"

# Copy .env for ccip
old_file="${CHAINLINK_CODE_DIR}/ccip/crib/.env"
new_file="${CHAINLINK_CODE_DIR}/crib/ccip/.env"

update_dot_env "$old_file" "$new_file"

popd
