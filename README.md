# Go-Version

Add GIT version and build timestamp to a project.

## Usage
```Go
import "github.com/gaaf/go-appversion"

main() {
	appversion.Print()
}
```

## Example Makefile
```Makefile
.DEFAULT_GOAL := build

export GOBIN=${CURDIR}/bin/

MODULE:=github.com/gaaf/go-appversion
VERSION:=$(shell git describe --tags --long --always --dirty)
BUILD_DATE:=$(shell date +%FT%T%z)

TARGET?=./...
BUILDFLAGS?=-trimpath
LDFLAGS:=-X $(MODULE).Version=$(VERSION) -X $(MODULE).BuildDate=$(BUILD_DATE)


# No dependency management by make, only by the go tool
.PHONY: build run prod debug clean


# Build the project
build:
	go install -v $(BUILDFLAGS) -ldflags "$(LDFLAGS)" $(TARGET)


# Run with same options as build
run:
	go run -v $(BUILDFLAGS) -ldflags "$(LDFLAGS)" $(TARGET)


# Production builds have symbols stripped
prod:	BUILDFLAGS += -a
prod:	LDFLAGS += -s -w
prod:	build


# Start the delve debugger
debug:
	dlv debug -ldflags "$(LDFLAGS)" $(TARGET)


# Clean project: delete binaries
clean:
	go clean -v $(BUILDFLAGS) -i $(TARGET)
```
