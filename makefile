## Example makefile for Go with fuzz test

## This is not using go.mod
# .EXPORT_ALL_VARIABLES:
# GOPATH=$(shell cd)
# GO111MODULE=off

default: run

help:
	@echo Usage:
	@echo make help  - this help
	@echo make qc    - staticcheck  
	@echo make lint  - start gchk.cmd running all my linter and error checkors with easy to see fail
	@echo make run   - run with 
	@echo make test  - run tests 
	@echo make bench - run bench
	@echo make build - build with full lint and staticcheck
	@echo make clean - delete *.exe

qc: 
	@-golint .
#	@-staticcheck .
	@echo -----------------------------------------------------------------

lint: 
	-@gchk
	@echo -----------------------------------------------------------------

# 	lint:qc
# 	-go vet .
# 	-golangci-lint.exe run .
# 	-revive . 
# 	-errcheck .
	
test: lint
	-go test ./...
	@echo -----------------------------------------------------------------

bench: qc 
	go test -benchmem -run=. -bench=. -benchtime=20s
	@echo -----------------------------------------------------------------
	go test -fuzz=./... -fuzztime=20s 
	@echo -----------------------------------------------------------------
	
build: lint test 
	go build -gcflags="-m=2" --ldflags="-s -w -race" -trimpath .
	@echo -----------------------------------------------------------------

run: qc
	go run .
	@echo -----------------------------------------------------------------

.PHONY : clean
clean :
	@-rm *.exe      
