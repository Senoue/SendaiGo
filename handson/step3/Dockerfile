FROM golang:1.13.6-alpine3.10

WORKDIR $GOPATH/src/api
COPY step3/main.go .
COPY step3/form.html .
COPY sendaigo-sql.json . 

RUN mkdir /cloudsql && \
    go mod init && \ 
    go build . && \
cp api /usr/local/bin/api

CMD ["api"]                                            