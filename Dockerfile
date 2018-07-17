FROM golang:1.10-alpine3.8

COPY ./ /go/src/github.com/mrsmuneton/platform-test
WORKDIR /go/src/github.com/mrsmuneton/platform-test

RUN apk update && apk add git
RUN go get ./
RUN go build

CMD ./platform-test

EXPOSE 8080
