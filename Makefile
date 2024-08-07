EXECUTABLE=memento-mori
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN_AMD=$(EXECUTABLE)_darwin_amd64
DARWIN_ARM=$(EXECUTABLE)_darwin_arm64
#VERSION=$(shell git describe --tags --always --long --dirty)
VERSION="0.0.1"
.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	go test ./...

build: windows linux darwin-amd darwin-arm ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin-amd: $(DARWIN_AMD) ## Build for Darwin (macOS) amd64

darwin-arm: $(DARWIN_ARM) ## Build for Darwin (macOS) arm64

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

$(DARWIN_AMD):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

$(DARWIN_ARM):
	env GOOS=darwin GOARCH=arm64 go build -v -o $(DARWIN_ARM) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN) $(DARWIN_ARM)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'