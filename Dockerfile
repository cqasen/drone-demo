## 编译镜像
FROM golang:1.13-alpine  AS builder
ENV GO111MODULE=on
ENV CGO_ENABLED 0
ENV GOOS=linux
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/"

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go



## 运行镜像
FROM scratch as prod
COPY --from=builder /go/release/app /
CMD ["/app"]
