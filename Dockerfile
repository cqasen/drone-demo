# build
FROM golang:alpine

WORKDIR /go/src/app
COPY . .

# set go mod proxy
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
CMD ["go","run","app.go"]