#!/usr/bin/env bash

# Kill existing session if it exists
tmux kill-session -t blog-services 2>/dev/null

# Create new session
tmux new-session -s blog-services -n services -d

tmux split-window -v
tmux split-window -h
tmux split-window -v
tmux split-window -h

# Start consul and wait for 5s
tmux select-pane -t 0
tmux send-keys "docker compose up -d consul" C-m
sleep 2

# Start posts service
tmux send-keys "cd posts" C-m
tmux send-keys "air" C-m

# Start gateway
tmux select-pane -t 1
tmux send-keys "cd gateway" C-m
tmux send-keys "air" C-m

# Start comments
tmux select-pane -t 2
tmux send-keys "cd comments" C-m
tmux send-keys "air" C-m

# Start likes
tmux select-pane -t 3
tmux send-keys "cd likes" C-m
tmux send-keys "npm run watch" C-m

# Start authentication
tmux select-pane -t 4
tmux send-keys "cd authentication" C-m
tmux send-keys "npm run watch" C-m

# Set the layout
tmux select-layout tiled

# Enable mouse support
tmux set -g mouse on

# Attach to the session
tmux attach-session -t blog-services