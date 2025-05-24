#!/bin/bash

while true; do
  for id in sensor01 sensor02 sensor03; do
    curl -sk -X POST https://localhost:8443/device/update \
      -H "Authorization: Bearer supersecrettoken" \
      -H "Content-Type: application/json" \
      -d "{\"id\":\"$id\", \"status\":\"online\"}"
    echo "[$(date)] Sent update for $id"
  done
  sleep 6
done
