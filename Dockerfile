# 使用 golang:alpine 构建两个服务
FROM golang:alpine AS builder

# 设置构建环境变量
ENV CGO_ENABLED 0
ENV GOOS linux

# 安装必要的工具
RUN apk update --no-cache && apk add --no-cache tzdata

# 构建 processor 服务
WORKDIR /processor-build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY enhanced enhanced
COPY enhanced/processor/etc /app/processor/etc

# 替换数据库连接字符串
ARG DB_CONNECTION_STRING="mongodb://root:password@mongodb-host"
RUN echo ${DB_CONNECTION_STRING}
RUN sed -i 's|DataSource:.*|DataSource: "'${DB_CONNECTION_STRING}'"|' /app/processor/etc/processor.yaml

RUN go mod tidy
RUN go build -ldflags="-s -w" -o /app/processor/processor enhanced/processor/processor.go

# 构建 http 服务
WORKDIR /http-build
ADD go.mod .
ADD go.sum .
COPY enhanced enhanced
COPY enhanced/http/etc /app/http/etc

# 替换数据库连接字符串
RUN sed -i 's|DataSource:.*|DataSource: "'${DB_CONNECTION_STRING}'"|' /app/http/etc/http-api.yaml

RUN go mod tidy
RUN go build -ldflags="-s -w" -o /app/http/http enhanced/http/http.go

# 使用最小化的 scratch 镜像 (不支持 sh)
FROM busybox:glibc

# 复制 CA 证书和时区信息
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 复制 processor 和 http 服务及其配置
COPY --from=builder /app/processor /app/processor
COPY --from=builder /app/processor/etc /app/processor/etc
COPY --from=builder /app/http /app/http
COPY --from=builder /app/http/etc /app/http/etc

# 添加启动脚本
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh
RUN sed -i 's/\r$//' start.sh

RUN chmod +x /app/processor/processor /app/http/http

EXPOSE 80

# 使用脚本作为容器入口点
CMD ["/app/start.sh"]