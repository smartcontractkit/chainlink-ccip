#!/bin/bash

while true; do
  echo "Running ccip_send.ts at $(date)"
  npx ts-node src/ccip_send.ts

  echo "Waiting 60 seconds..."
  sleep 60
done
