FROM golang:latest

ADD . /go/src/github.com/joubertredrat-tests/smartmei-dev-test-2k20

WORKDIR /go/src/github.com/joubertredrat-tests/smartmei-dev-test-2k20

ENTRYPOINT go run crawler.go exchange.go main.go resolver.go struct.go

EXPOSE 8000