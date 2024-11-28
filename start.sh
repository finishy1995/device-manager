#!/bin/sh

# 启动 processor 服务
/app/processor/processor -f /app/processor/etc/processor.yaml &

# 启动 http 服务
/app/http/http -f /app/http/etc/http-api.yaml &

# 等待所有子进程的退出
wait