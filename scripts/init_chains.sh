#!/usr/bin/env bash

tmux new -s osmosis -d ./scripts/node_start/runnode_osmosis.sh 
tmux new -s feeabs -d ./scripts/node_start/runnode_custom.sh 

echo "Osmosisd and feeabs is up and run!"