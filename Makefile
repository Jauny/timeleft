.PHONY: all
all: build test

.PHONY: build
build:
	go build main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean ./...
