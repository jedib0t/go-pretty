.PHONY: all dep profile test

default: build

all: dep test bench

tools:
	go get golang.org/x/lint/golint
	go get github.com/fzipp/gocyclo
	go get github.com/golang/dep/cmd/dep

bench:
	go test -bench=. -benchmem

build:
	go run cmd/demo/demo.go

cyclo:
	gocyclo -over 10 ./*/*.go

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
