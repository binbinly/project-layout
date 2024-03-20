FROM golang:1.20-alpine as builder

# 作者
LABEL maintainer="app"

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 执行镜像的工作目录
WORKDIR /go/src/app

# 将目录拷贝到镜像里
COPY . .

RUN go build -o app .

# 引入alphine最小linux镜像
FROM alpine

# 解决时区问题: unknown time zone Asia/Shanghai
RUN apk update && apk add tzdata

WORKDIR /app
# 复制生成的可执行命令和一些配置文件
COPY --from=builder  /go/src/app/app .
COPY --from=builder  /go/src/app/configs/prod ./configs

# 开放http ws端口
EXPOSE 9050 9060

# 启动执行命令
ENTRYPOINT ["/app/app"]

# 1. build image: docker build -t app:latest -f Dockerfile .
# 2. start: docker run --rm -it -p 9050:9050 -p 9060:9060 app:latest server
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统
