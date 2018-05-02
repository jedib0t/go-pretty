.PHONY: all dep test

default: build

all: dep test bench

tools:
	go get golang.org/x/lint/golint
	go get github.com/fzipp/gocyclo
	go get github.com/golang/dep/cmd/dep
#	go get github.com/golang/mock/gomock
#	go install github.com/golang/mock/mockgen

bench:
	go test -bench=. -benchmem

build:
	go run demo/demo.go

cyclo:
	gocyclo -over 10 ./*/*.go

dep:
	dep ensure

lint:
	golint $(shell go list ./...)

test: lint vet cyclo
	go test -cover $(shell go list ./...)

vet:
	go vet $(shell go list ./...)
