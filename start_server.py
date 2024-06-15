#!/usr/bin/env python

import argparse
import os
from re import sub
import subprocess

DEFAULT_RAFT_IDS = {"A", "B", "C"}

parser = argparse.ArgumentParser(
    os.path.basename(__file__),
    description="script for startiting the Tuple Space Server",
    epilog=f"./{os.path.basename(__file__)} --raft-id A --build --client-port 60000 --client-addr 127.0.0.1"
)

parser.add_argument(
    "--raft-id", "-r",
    help=f"Unique identifier for raft server. If different from those ({','.join(DEFAULT_RAFT_IDS)}), the server will be dinacally added to the cluster",
    metavar="id",
    required=True,
    type=str
)

parser.add_argument(
    "--build", "-b",
    help="Build the docker container before running it",
    action="store_true",
)

parser.add_argument(
    "--client-addr",
    metavar="addr",
    type=str,
    help="Address for the CLIENT interface. This is the address to where the client program should connect to. Defaults to 127.0.0.1"
)
parser.add_argument(
    "--client-port",
    metavar="port",
    type=int,
    help="Port for the CLIENT interface. This is the port to where the client program should connect to. Defaults to 0 (i.e: random)"
)

parser.add_argument(
    "--print-docker-run",
    help="Print the command used to run docker",
    action="store_true"
)

args = parser.parse_args()
container_tag = "computacao-distribuida-tuple-spaces"


if args.build:
    subprocess.run(f"docker build -t {container_tag} {os.path.dirname(__file__)}", shell=True, check=True)


env_vars = []
for name in ("raft_id", "client_port", "client_addr"):
    if getattr(args, name) is not None:
        env_vars.append((name, getattr(args, name)))

cmd = f"docker run {' '.join([f'-e {name}={value}' for name, value in env_vars])} -it --rm --network=host {container_tag}"

if args.print_docker_run:
    print(cmd)

subprocess.run(f"docker run {' '.join([f'-e {name}={value}' for name, value in env_vars])} -it --rm --network=host {container_tag}", shell=True)
