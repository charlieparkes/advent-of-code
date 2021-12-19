.DEFAULT_GOAL := all

.PHONY: all
all: test run

.PHONY: build
build:
	@go build -o adventofcode

adventofcode: build

.PHONY: run
run: build
	@./adventofcode 2021 all

.PHONY: test
test:
	go test ./... | grep -v 'no test files' | grep -v 'no tests to run'