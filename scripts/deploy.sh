
#!/bin/bash

# This is a simple deployment script for a Go application.
# It uploads the compiled binary to a remote server and starts the application.

# Server information
REMOTE_USER="jmrashed"
REMOTE_HOST="your_server_ip"
REMOTE_DIR="/path/to/deploy/directory"

# Local binary name (output of the build process)
LOCAL_BINARY="myapp"

# Remote binary name (desired name on the remote server)
REMOTE_BINARY="myapp"

# Copy the binary to the remote server
scp "$LOCAL_BINARY" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR/$REMOTE_BINARY"

# SSH into the remote server and start the application
ssh "$REMOTE_USER@$REMOTE_HOST" "cd $REMOTE_DIR && ./$REMOTE_BINARY"
