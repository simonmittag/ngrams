.PHONY: all build test clean

all: clean test build

build:
	go build $(LDFLAGS) -o $(BINARY_NAME) ./...

test:
	go test -v ./...

clean:
	go clean
