EXECUTABLE=memento-mori
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX_AMD64=$(EXECUTABLE)_linux_amd64
LINUX_ARM64=$(EXECUTABLE)_linux_arm64
DARWIN_AMD64=$(EXECUTABLE)_darwin_amd64
DARWIN_ARM64=$(EXECUTABLE)_darwin_arm64
#VERSION=$(shell git describe --tags --always --long --dirty)
VERSION="0.0.1"
.PHONY: all test clean

all: clean test build ## Build and run tests

test: ## Run unit tests
	go test ./...

build: windows linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux-amd64: $(LINUX_AMD64) ## Build for Linux amd64

linux-arm64: $(LINUX_ARM64) ## Build for Linux arm64

darwin-amd64: $(DARWIN_AMD64) ## Build for Darwin (macOS) amd64

darwin-arm64: $(DARWIN_ARM64) ## Build for Darwin (macOS) arm64

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.Version=$(VERSION)"  ./cmd/main.go

$(LINUX_AMD64):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX_AMD64) -ldflags="-s -w -X main.Version=$(VERSION)"  ./cmd/main.go

$(LINUX_ARM64):
	env GOOS=linux GOARCH=arm64 go build -v -o $(LINUX_ARM64) -ldflags="-s -w -X main.Version=$(VERSION)"  ./cmd/main.go

$(DARWIN_AMD64):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN_AMD64) -ldflags="-s -w -X main.Version=$(VERSION)"  ./cmd/main.go

$(DARWIN_ARM64):
	env GOOS=darwin GOARCH=arm64 go build -v -o $(DARWIN_ARM64) -ldflags="-s -w -X main.Version=$(VERSION)"  ./cmd/main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX_AMD64) ${LINUX_ARM64} $(DARWIN_AMD64) $(DARWIN_ARM64)

help: ## Display available commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'