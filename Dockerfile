FROM golang:1.10-alpine3.8

COPY ./ /go/src/github.com/mrsmuneton/platform-test
WORKDIR /go/src/github.com/mrsmuneton/platform-test

RUN apk update && apk add git

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o ./bin/platform-test

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]
