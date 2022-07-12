FROM golang:1.18-alpine

WORKDIR /projects/starter
ENV CGO_ENABLED=0

RUN apk --update add git
RUN GO111MODULE=off go get github.com/oxequa/realize
