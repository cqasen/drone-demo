## 编译镜像
FROM golang:alpine  AS builder
ARG app_env
ENV APP_ENV $app_env
ENV GO111MODULE=on
ENV CGO_ENABLED 0
ENV GOOS=linux
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/

WORKDIR /go/release/app
COPY . .
RUN go mod download
CMD ["go","run","main.go"]