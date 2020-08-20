FROM google/cloud-sdk:alpine

LABEL maintainer="Senoue <hiro@senoue.net>"

WORKDIR /go

ENV GOPATH=/go \
    PATH=/go/bin:/usr/local/go/bin:$PATH

ARG GOLANG_VERSION=1.13.10

RUN apk update && \
    apk add make gcc git musl-dev && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    \
    gcloud components install \
    app-engine-go && \
    \
    curl -o go.tgz -sSL "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz && \
    \
    rm -rf src/*

ENTRYPOINT bash