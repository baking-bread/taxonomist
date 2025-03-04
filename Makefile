BINARY_NAME=taxonomist
GOARCH=amd64
HASH_CMD=sha256sum
ifeq ($(OS),Windows_NT)
    HASH_CMD=certutil -hashfile
	HASH_ALGORITHM=sha256
endif

# OS specific variables
ifeq ($(OS),Windows_NT)
    BINARY_SUFFIX=.exe
    RM_CMD=del /F /Q
    RM_DIR=rd /S /Q
    BINARY_PATH=bin\$(BINARY_NAME)$(BINARY_SUFFIX)
	CHECK_CMD=where
	ECHO_NULL=>nul 2>nul
else
    BINARY_SUFFIX=
    RM_CMD=rm -f
    RM_DIR=rm -rf
    BINARY_PATH=bin/$(BINARY_NAME)$(BINARY_SUFFIX)
	CHECK_CMD=which
	ECHO_NULL=>/dev/null 2>&1
endif



.PHONY: all test build lint security clean

all: build lint test scan

build:
	go build -o bin/$(BINARY_NAME)$(BINARY_SUFFIX) main.go
	cd bin && $(HASH_CMD) $(BINARY_NAME)$(BINARY_SUFFIX) $(HASH_ALGORITHM) > $(BINARY_NAME)$(BINARY_SUFFIX).sha256

test:
	go test -v ./...

clean:
	go clean
	$(RM_DIR) bin
	$(RM_CMD) *.sha256

lint:
	go vet ./...
	$(CHECK_CMD) golangci-lint $(ECHO_NULL) && golangci-lint run || echo "golangci-lint not installed"

scan:
	$(CHECK_CMD) gosec $(ECHO_NULL) && gosec -no-fail -terse -out=./bin/sonarqube.gosec.json ./... || echo "gosec not installed"
	$(CHECK_CMD) trivy $(ECHO_NULL) && trivy fs --scanners vuln . || echo "Trivy not installed"
