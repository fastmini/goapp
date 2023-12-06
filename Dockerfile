FROM golang:1.18-alpine3.14 AS builder

WORKDIR /build
RUN adduser -u 10001 -D api-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o goApp

FROM alpine:3.14
LABEL maintainer="AN <751815097@qq.com>" version="1.0" license="MIT"

ENV TZ="Asia/Shanghai"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add ttf-dejavu libuuid tzdata ca-certificates wget fontconfig bash curl \
    && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && rm -rf /var/cache/apk/*

WORKDIR /app
COPY .env.online .env
COPY --from=builder /build/goApp ./
ENTRYPOINT ["/app/goApp"]
