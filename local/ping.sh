#!/bin/bash

while true; do
  ping -c5 -t10 $PINGHOST
  if [ $? -ne 0 ]; then
    traceroute $PINGHOST
  fi
  sleep 5;
done
