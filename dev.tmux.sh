#!/usr/bin/env bash

# Kill existing session if it exists
tmux kill-session -t microservices 2>/dev/null

# Create new session
tmux new-session -s microservices -n services -d

# Split the window into panes
tmux split-window -h
tmux split-window -v
tmux select-pane -t 0
tmux split-window -v

# Setup consul in the first pane
tmux select-pane -t 0
tmux send-keys "docker compose up consul" C-m

# Wait for consul to start
sleep 5

# Setup posts service with Air in the second pane
tmux select-pane -t 1
tmux send-keys "cd posts" C-m
tmux send-keys "air" C-m

# Setup gateway with Air in the third pane
tmux select-pane -t 2
tmux send-keys "cd gateway" C-m
tmux send-keys "air" C-m

# Reserve the fourth pane for commands/testing
tmux select-pane -t 3
tmux send-keys "echo 'Development environment ready. Use this pane for testing.'" C-m

# Set the layout
tmux select-layout tiled

# Attach to the session
tmux attach-session -t microservices