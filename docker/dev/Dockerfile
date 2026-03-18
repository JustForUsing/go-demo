# 构建阶段
FROM golang:1.25.5-alpine AS builder

# 安装必要的工具
RUN apk add --no-cache git

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 复制源代码
COPY .. .

# 构建应用
RUN go build -o web ./cmd/web/main.go

# 运行阶段
FROM alpine:latest

# 安装时区和CA证书
RUN apk --no-cache add tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 创建非root用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder --chown=appuser:appgroup /app/web .

# 切换用户
USER appuser

# 暴露端口
EXPOSE 8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:80/health || exit 1

# 启动应用
ENTRYPOINT ["/web"]