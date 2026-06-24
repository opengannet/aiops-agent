
VERSION="$(git rev-parse --short HEAD)-aiops-l7fallback-$(date -u +%Y%m%d%H%M%S)"

docker build \
  --build-arg VERSION="$VERSION" \
  -t docker.fzyun.io/aiops/node-agent:latest \
  .
