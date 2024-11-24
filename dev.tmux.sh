#!/usr/bin/env bash

# Kill existing session if it exists
tmux kill-session -t blog-services 2>/dev/null

# Create new session
tmux new-session -s blog-services -n services -d

# Split the window into 3 panes vertically
tmux split-window -v
tmux split-window -v

# Start all docker services in detached mode
tmux select-pane -t 0
tmux send-keys "docker compose up -d consul" C-m
# Wait for consul to start
sleep 5
tmux send-keys "cd posts && docker compose up -d" C-m
# Wait for postgres to start
sleep 5
tmux send-keys "air" C-m

# Setup gateway with Air in the second pane
tmux select-pane -t 1
tmux send-keys "cd gateway" C-m
tmux send-keys "air" C-m

# Setup third pane for commands/testing
tmux select-pane -t 2
tmux send-keys "echo 'Development environment ready. Use this pane for testing.'" C-m

# Set the layout
tmux select-layout even-vertical

# Enable mouse support
tmux set -g mouse on

# Attach to the session
tmux attach-session -t blog-services