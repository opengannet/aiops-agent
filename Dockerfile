FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN GOPROXY=https://goproxy.cn,direct go build -o coroot-node-agent .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/coroot-node-agent /usr/local/bin/coroot-node-agent
EXPOSE 80 10300
ENTRYPOINT ["/usr/local/bin/coroot-node-agent"]
