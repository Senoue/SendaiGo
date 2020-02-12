FROM golang:1.13.6-alpine3.10

WORKDIR $GOPATH/src/api
COPY main.go .
RUN apk add git
RUN go get github.com/go-sql-driver/mysql
RUN go build . && \
    cp api /usr/local/bin/api

CMD ["api"]