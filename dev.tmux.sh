#!/usr/bin/env bash

# Kill existing session if it exists
tmux kill-session -t blog-services 2>/dev/null

# Create new session
tmux new-session -s blog-services -n services -d

tmux split-window -v
tmux split-window -v

# Start consul and wait for 5s
tmux select-pane -t 0
tmux send-keys "docker compose up -d consul" C-m
sleep 2

# Start posts service
tmux send-keys "cd posts && docker compose up -d" C-m
sleep 2
tmux send-keys "air" C-m

# Start gateway
tmux select-pane -t 1
tmux send-keys "cd gateway" C-m
tmux send-keys "air" C-m

# Start comments
tmux select-pane -t 2
tmux send-keys "cd comments && docker compose up -d" C-m
sleep 2
tmux send-keys "air" C-m

# Set the layout
tmux select-layout tiled

# Enable mouse support
tmux set -g mouse on

# Attach to the session
tmux attach-session -t blog-services