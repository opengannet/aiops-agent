docker run --detach --name "$(hostname)" \
  --privileged --pid host \
  -v /sys/kernel/debug:/sys/kernel/debug:rw \
  -v /sys/fs/cgroup:/host/sys/fs/cgroup:ro \
  docker.fzyun.io/aiops/node-agent:latest \
  --cgroupfs-root=/host/sys/fs/cgroup \
  --collector-endpoint=https://hawk.ops.fzyun.io \
  --api-key=Bl_pq07V3MF77QhRJFunQMt9BsOuVpPs \
  --scrape-interval=30s
