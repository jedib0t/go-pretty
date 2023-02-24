.PHONY: all profile test

default: test

all: test bench

tools:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.5.1

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

fmt:
	go fmt $(shell go list ./...)

profile:
	sh profile.sh

test: fmt vet cyclo
	go test -cover -coverprofile=.coverprofile $(shell go list ./...)

test-race:
	go run -race ./cmd/demo-progress/demo.go

vet:
	go vet $(shell go list ./...)
