#!/bin/bash

while true; do
  echo "Waiting for data..."
  data=$(nc -u -w0 -l 12345)
  echo "$data"
done
