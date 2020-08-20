FROM golang:1.13.12-alpine3.12

WORKDIR $GOPATH/src/chat
COPY src/api/router .
COPY src/api/app/*.json ./app/
COPY src/api/app/*.html ./app/
RUN apk update
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata
RUN cp router /usr/local/bin/router

CMD ["router"]
