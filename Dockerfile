## 编译镜像
FROM golang:alpine  AS builder
ARG app_env
ENV APP_ENV $app_env
ENV GO111MODULE=on
ENV CGO_ENABLED 0
ENV GOOS=linux
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

#RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

## 运行镜像
FROM scratch as prod
COPY --from=builder /go/release/app /
CMD ["go","run","main.go"]