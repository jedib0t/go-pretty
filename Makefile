.PHONY: all dep gen test

default: build

all: dep test bench

tools:
	go get -u golang.org/x/lint/golint
	go get -u github.com/fzipp/gocyclo
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

bench:
	go test -bench . -benchmem

build: gen
	go run demo/demo.go

cyclo:
	gocyclo -over 10 ./*.go

dep:
	dep ensure

gen:
	go generate $(shell go list .)

lint:
	golint $(shell go list .)

test: lint vet cyclo
	go test -cover $(shell go list .)

vet:
	go vet $(shell go list .)
