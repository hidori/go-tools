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
    && go build -ldflags "-s -w" -o /tmp/bin/genfldnam -a ./cmd/genfldnam/main.go \
    && go build -ldflags "-s -w" -o /tmp/bin/genprop   -a ./cmd/genprop/main.go \
    && go build -ldflags "-s -w" -o /tmp/bin/printast  -a ./cmd/printast/main.go

FROM alpine:latest AS executor

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates

COPY --from=builder /tmp/bin/* /usr/local/bin/
