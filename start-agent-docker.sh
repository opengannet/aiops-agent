#!/bin/bash
# Start coroot-node-agent via Docker
# Usage: ./start-agent-docker.sh

set -euo pipefail

IMAGE_NAME="coroot-node-agent"
CONTAINER_NAME="coroot-node-agent"

# Build image
echo "Building Docker image..."
docker build -t "$IMAGE_NAME" "$(cd "$(dirname "$0")" && pwd)"

# Stop and remove existing container
docker stop "$CONTAINER_NAME" 2>/dev/null || true
docker rm "$CONTAINER_NAME" 2>/dev/null || true

# Run container
docker run -d --name "$CONTAINER_NAME" \
    --network host \
    --pid host \
    -v /sys/fs/cgroup:/sys/fs/cgroup:ro \
    -v /sys/kernel/debug:/sys/kernel/debug:ro \
    -v /var/log:/var/log:ro \
    "$IMAGE_NAME" \
    --listen=0.0.0.0:80 \
    --disable-gpu-monitoring \
    --collector-endpoint=http://172.19.66.239:8080 \
    --api-key=EKB3M6PrR7WoouV_P5VHawVOIwgJhWR9

echo "Container started. Logs:"
docker logs -f "$CONTAINER_NAME"
