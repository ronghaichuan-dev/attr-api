# 构建阶段
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git tzdata

# 复制go.mod和go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o job-svc ./cmd/job-svc/

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add tzdata ca-certificates && \
    adduser -D -u 1000 appuser

WORKDIR /app

# 复制编译好的二进制文件
COPY --from=builder /app/job-svc .
COPY --from=builder /app/cmd/job-svc/config ./config

# 创建日志目录
RUN mkdir -p tmp/logs && chown -R appuser:appuser .

# 切换用户
USER appuser

# 暴露端口
EXPOSE 8803

# 启动应用
CMD ["./job-svc"]
