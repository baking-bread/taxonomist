BINARY_NAME=taxonomist
GOARCH=amd64
HASH_CMD=sha256sum
ifeq ($(OS),Windows_NT)
    HASH_CMD=certutil -hashfile
endif

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

.PHONY: all build test clean lint deps windows linux darwin security

all: deps lint test build

build: security

# Cross-compilation targets with security
windows:
	GOOS=windows GOARCH=$(GOARCH) go build -o bin/$(BINARY_NAME).exe main.go
	cd bin && $(HASH_CMD) $(BINARY_NAME).exe > $(BINARY_NAME).exe.sha256

linux:
	GOOS=linux GOARCH=$(GOARCH) go build -o bin/$(BINARY_NAME) main.go
	cd bin && $(HASH_CMD) $(BINARY_NAME) > $(BINARY_NAME).sha256

darwin:
	GOOS=darwin GOARCH=$(GOARCH) go build -o bin/$(BINARY_NAME) main.go
	cd bin && $(HASH_CMD) $(BINARY_NAME) > $(BINARY_NAME).sha256

test:
	go test -v ./...

clean:
	go clean
	$(RM_DIR) bin
	$(RM_CMD) *.sha256

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

# New security target
security: $(BINARY_PATH)
	cd bin && $(HASH_CMD) $(BINARY_NAME)$(BINARY_SUFFIX) > $(BINARY_NAME)$(BINARY_SUFFIX).sha256

deps:
	go mod download
	go mod tidy

# CI pipeline target that runs everything in sequence
ci: deps lint test build
