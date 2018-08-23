.PHONY: all dep profile test

default: test

all: dep test bench

tools:
	go get github.com/fzipp/gocyclo
	go get github.com/golang/dep/cmd/dep
	go get golang.org/x/lint/golint

bench:
	go test -bench=. -benchmem

cyclo:
	gocyclo -over 13 ./*/*.go

demo-list:
	go run cmd/demo-list/demo.go

demo-progress:
	go run cmd/demo-progress/demo.go

demo-table:
	go run cmd/demo-table/demo.go

dep:
	dep ensure

fmt:
	go fmt $(shell go list ./...)

lint:
	golint -set_exit_status $(shell go list ./...)

profile:
	sh profile.sh

test: fmt lint vet cyclo
	go test -cover -coverprofile=.coverprofile $(shell go list ./...)

vet:
	go vet $(shell go list ./...)
