## Example makefile for Go with fuzz test

## This is not using go.mod
# .EXPORT_ALL_VARIABLES:
# GOPATH=$(shell cd)
# GO111MODULE=off

default: run

test: 
	staticcheck .
	golangci-lint.exe run .
	revive . 
	go vet .
	errcheck .
	
	go test ./...
	go test -benchmem -run=. -bench=. -benchtime=20s
	go test -fuzz=./... -fuzztime=20s 
	
build:  
	go build -gcflags="-m=2" --ldflags="-s -w -race" -trimpath .

run: 
	@staticcheck .
	@go run .

.PHONY : clean
clean :
	@-rm *.exe      