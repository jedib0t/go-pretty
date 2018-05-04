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

profile:
	rm -fr .profile
	# profile list
	mkdir -p .profile/list
	go build -o .profile/list/list.exe list/profile/profile.go
	(cd .profile/list && ./list.exe)
	(cd .profile/list && go tool pprof -pdf list.exe cpu.pprof > cpu.pdf)
	(cd .profile/list && go tool pprof -pdf list.exe mem.pprof > mem.pdf)
	# profile table
	mkdir -p .profile/table
	go build -o .profile/table/table.exe table/profile/profile.go
	(cd .profile/table && ./table.exe)
	(cd .profile/table && go tool pprof -pdf table.exe cpu.pprof > cpu.pdf)
	(cd .profile/table && go tool pprof -pdf table.exe mem.pprof > mem.pdf)

test: lint vet cyclo
	go test -cover $(shell go list ./...)

vet:
	go vet $(shell go list ./...)
