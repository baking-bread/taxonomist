BINARY_NAME=taxonomist
GOARCH=amd64

# OS specific variables
ifeq ($(OS),Windows_NT)
    BINARY_SUFFIX=.exe
    RM_CMD=del /F /Q
    RM_DIR=rd /S /Q
    BINARY_PATH=bin\$(BINARY_NAME)$(BINARY_SUFFIX)
else
    BINARY_SUFFIX=
    RM_CMD=rm -f
    RM_DIR=rm -rf
    BINARY_PATH=bin/$(BINARY_NAME)$(BINARY_SUFFIX)
endif

.PHONY: all build test clean lint deps windows linux darwin

all: deps lint test build

build:
	go build -o $(BINARY_PATH) main.go

# Cross-compilation targets
windows:
	GOOS=windows GOARCH=$(GOARCH) go build -o bin/$(BINARY_NAME).exe main.go

linux:
	GOOS=linux GOARCH=$(GOARCH) go build -o bin/$(BINARY_NAME) main.go

darwin:
	GOOS=darwin GOARCH=$(GOARCH) go build -o bin/$(BINARY_NAME) main.go

test:
	go test -v ./...

clean:
	go clean
	$(RM_DIR) bin

lint:
	go vet ./...
ifeq ($(OS),Windows_NT)
	@where golangci-lint >nul 2>nul && (\
		golangci-lint run \
	) || (\
		echo golangci-lint not installed \
	)
else
	@which golangci-lint >/dev/null 2>&1 && (\
		golangci-lint run \
	) || (\
		echo "golangci-lint not installed" \
	)
endif

deps:
	go mod download
	go mod tidy

# CI pipeline target that runs everything in sequence
ci: deps lint test build
