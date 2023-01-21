# tear down
tmux kill-session -t server

# init
tmux new -s server -d

tmux new -s client -d 

tmux join-pane -h -s client -t server

tmux send-keys -t server.0 "make server" Enter

tmux send-keys -t server.1 "make client" Enter

tmux attach -t server

# keep bash
bash

