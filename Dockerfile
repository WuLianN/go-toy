# 指定go编译镜像
FROM golang:latest as build

# 指定go的环境变量
ENV GOPROXY=https://goproxy.cn \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /app

COPY . .

# 编译成可执行二进制文件
RUN go build -o main .

# 使用Alpine镜像作为临时环境来设置时区
FROM alpine:latest as timezone

ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo "$TZ" > /etc/timezone

# 指定新的运行环境，最终的运行会基于这个坏境
FROM scratch as deploy

COPY --from=timezone /etc/localtime /etc/localtime
COPY --from=timezone /etc/timezone /etc/timezone
COPY --from=build /app /
COPY --from=build /app/configs ./configs

CMD ["/main"]
