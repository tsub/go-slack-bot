PROJECT  = go-slack-bot
SRC      ?= $(shell go list ./... | grep -v vendor)
TESTARGS ?= -v

deps:
	dep ensure
.PHONY: deps

build:
	go build -o build/$(PROJECT)
.PHONY: build

run: build
	build/$(PROJECT)
.PHONY: run

test:
	go test $(SRC) $(TESTARGS)
.PHONY: test

fmt:
	go fmt $(SRC)
.PHONY: fmt

vet:
	go vet $(SRC)
.PHONY: vet
