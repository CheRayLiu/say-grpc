FROM golang:1.12 AS builder 

RUN go get -u github.com/Sirupsen/logrus
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u golang.org/x/net/context
RUN go get -u google.golang.org/grpc

FROM alpine

RUN apk update && apk add flite
ADD app /app

ENTRYPOINT ["/app"]