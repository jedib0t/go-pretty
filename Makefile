.PHONY: all dep profile test

default: build

all: dep test bench

tools:
	go get github.com/fzipp/gocyclo
	go get github.com/golang/dep/cmd/dep
	go get golang.org/x/lint/golint

bench:
	go test -bench=. -benchmem

build:
	go run cmd/demo/demo.go

cyclo:
	gocyclo -over 13 ./*/*.go

dep:
	dep ensure

lint:
	golint -set_exit_status $(shell go list ./...)

profile:
	sh profile.sh

test: lint vet cyclo
	go test -cover -coverprofile=.coverprofile $(shell go list ./...)

vet:
	go vet $(shell go list ./...)
