.PHONY: all
all: build test

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean ./...
