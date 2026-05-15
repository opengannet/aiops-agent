#!/bin/bash
# Start coroot-node-agent with kernel 4.4 fixes
# Usage: ./start-agent.sh

set -euo pipefail

AGENT_DIR="$(cd "$(dirname "$0")" && pwd)"
AGENT_BIN="$AGENT_DIR/coroot-node-agent"

# Check if binary exists, build if not
if [ ! -f "$AGENT_BIN" ]; then
    echo "Binary not found. Building..."
    cd "$AGENT_DIR"
    export http_proxy=http://172.19.202.167:7897
    export https_proxy=http://172.19.202.167:7897
    export GOPROXY=https://goproxy.cn,direct
    source /etc/bash.bashrc 2>/dev/null || true
    go build -o coroot-node-agent .
    echo "Build complete."
fi

# Kill any existing instance
pkill -f 'coroot-node-agent.*--api-key=' 2>/dev/null || true
sleep 1

# Start the agent
exec "$AGENT_BIN" \
    --listen=0.0.0.0:80 \
    --disable-gpu-monitoring \
    --collector-endpoint=http://172.19.66.239:8080 \
    --api-key=EKB3M6PrR7WoouV_P5VHawVOIwgJhWR9
