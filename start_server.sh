#!/bin/env bash
[[ "$1" == "-b" || "$1" == "--build" || "$2" == "-b" || "$2" == "--build" ]] && docker build -t distribuida-tuple-space .

server_id=""
if [[ "$1" == "-b" || "$1" == "--build" ]]; then
    server_id="$2"
else
    server_id="$1"
fi



docker run -e RAFT_ID=$server_id -it --rm --network=host distribuida-tuple-space
