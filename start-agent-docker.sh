#!/bin/bash
# Start coroot-node-agent via Docker (based on official coroot-node-agent Dockerfile)
# Usage: ./start-agent-docker.sh

set -euo pipefail

IMAGE_NAME="coroot-node-agent"
CONTAINER_NAME="coroot-node-agent"

# Configure Docker daemon proxy (needed for goproxy.cn access)
DOCKER_DIR="/etc/docker"
DAEMON_JSON="$DOCKER_DIR/daemon.json"
mkdir -p "$DOCKER_DIR"
if ! grep -q 'proxy' "$DAEMON_JSON" 2>/dev/null; then
    cat > "$DAEMON_JSON" <<EOF
{
    "proxies": {
        "httpProxy": "http://172.19.202.167:7897",
        "httpsProxy": "http://172.19.202.167:7897"
    }
}
EOF
    systemctl restart docker 2>/dev/null || service docker restart 2>/dev/null || true
fi

# Build image (uses official Dockerfile with CGO_ENABLED=1)
echo "Building Docker image..."
docker build -t "$IMAGE_NAME" "$(cd "$(dirname "$0")" && pwd)"

# Stop and remove existing container
docker stop "$CONTAINER_NAME" 2>/dev/null || true
docker rm "$CONTAINER_NAME" 2>/dev/null || true

# Run container with proper mounts for kernel 4.4 systemd log collection
docker run -d --name "$CONTAINER_NAME" \
    --network host \
    --pid host \
    -v /sys/fs/cgroup:/sys/fs/cgroup:ro \
    -v /sys/kernel/debug:/sys/kernel/debug:ro \
    -v /var/log:/var/log:ro \
    -v /run/systemd/private:/run/systemd/private:ro \
    "$IMAGE_NAME" \
    --listen=0.0.0.0:80 \
    --disable-gpu-monitoring \
    --collector-endpoint=http://172.19.66.239:8080 \
    --api-key=EKB3M6PrR7WoouV_P5VHawVOIwgJhWR9

echo "Container started. Checking logs..."
sleep 3
docker logs "$CONTAINER_NAME"
