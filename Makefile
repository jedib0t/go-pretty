.PHONY: all profile test

default: test

all: test bench

tools:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.5.1
	go install github.com/rinchsan/gosimports/cmd/gosimports@v0.3.8

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
	go fmt ./...
	gosimports -w .

profile:
	sh profile.sh

test: fmt vet cyclo
	go test -cover -coverprofile=.coverprofile ./...

test-race:
	go run -race ./cmd/demo-progress/demo.go

vet:
	go vet ./...

