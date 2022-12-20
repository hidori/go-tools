ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine AS builder

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    openssh-client git build-base \
    && update-ca-certificates

COPY . /go/src/app
WORKDIR /go/src/app

RUN make mod/download \
    && go build -ldflags "-s -w" -o /tmp/genprop -a ./cmd/genprop/main.go

FROM alpine:latest AS executor

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates

COPY --from=builder /tmp/genprop /usr/local/bin/genprop
