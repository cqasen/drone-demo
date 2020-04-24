# build
FROM golang:alpine

WORKDIR /go/src/github.com/cqasen/drone-demo
COPY . .
# set go mod proxy
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
CMD ["go","run","main.go"]