# build
FROM golang:1.13-alpine

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/app
COPY . .

# set go mod proxy
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
CMD ["go","run","app.go"]