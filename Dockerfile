FROM golang:1.6.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get github.com/gorilla/mux/... && \
    go get github.com/gorilla/websocket/.. && \
    go get gopkg.in/mgo.v2