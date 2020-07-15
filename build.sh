#!/bin/bash

# 开启
export GO111MODULE=on
# 1.13 之后才支持多个地址，之前版本只支持一个
export GOPROXY=https://mirrors.aliyun.com/goproxy,direct
# 1.13 开始支持，配置私有 module，不去校验 checksum
export GOPRIVATE=*.corp.example.com,rsc.io/private,github.com/gin-project

if [ -f "go.mod" ]; then
  rm -rf go.mod
fi

go mod init drampix-utils && go mod tidy -v && go mod verify && go mod vendor -v
