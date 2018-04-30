FROM golang:1.10-alpine

RUN apk --update add git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/oxequa/realize

WORKDIR /go/src/go-rest-starter/go/src/starter
