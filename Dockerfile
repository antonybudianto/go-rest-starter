FROM golang:1.11-alpine

RUN apk --update add git

# RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/oxequa/realize

ENV CGO_ENABLED=0

WORKDIR /projects/starter
