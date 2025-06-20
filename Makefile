# Build variables

GO := go
GOLANGCI_VER := 1.61.0

BUILD_DIR ?= build

GO_SERVICE_NAME=redis-in-go
BINARY_NAME=${GO_SERVICE_NAME}

GO_MAIN_FILE=cmd/main.go
GO_MAIN_DIR=$(dir ${GO_MAIN_FILE})

LOG_FILE_NAME="${GO_SERVICE_NAME}.log"

export CGO_ENABLED ?= 0

clean: stop ## Clean all binaries
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> cleaning up..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	rm -rf ${BUILD_DIR}

build: clean ## Build all binaries
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> building..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	@mkdir -p ${BUILD_DIR}
	# -cp -rf configs/config.json ${BUILD_DIR}/
	$(GO) build -trimpath -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${BINARY_NAME} ${GO_MAIN_DIR}/*.go

debug-build: clean ## Build all binaries
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> building debug build..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	@mkdir -p ${BUILD_DIR}
	# -cp -rf configs/config.json ${BUILD_DIR}/
	CGO_ENABLED=0 $(GO) build -trimpath -gcflags "all=-N -l" -o ${BUILD_DIR}/${BINARY_NAME}-debug ${GO_MAIN_DIR}/*.go

run: stop clean tidy lint build ## Build and run the application
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> running..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	cd build; ./${BINARY_NAME} > ${LOG_FILE_NAME} 2>&1 & echo "$$!" > ./process.pid

debug: stop clean tidy lint debug-build ## Build and run the application
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> running in debug mode..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	cd build; ./${BINARY_NAME} -debug > ${LOG_FILE_NAME} 2>&1 & echo "$$!" > ./process.pid

stop: 
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> stopping..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	-kill -9 `cat ${BUILD_DIR}/process.pid` 
## leading '-' ignores if command fails

generate-mocks: 
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> generating mocks..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	mockery --all --keeptree

coverage:
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> running coverage..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	$(GO) test -cover ./...

lint:
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> running linter..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	@# This assumes GOPATH/bin is in $PATH -- if not, the target will fail.
	@# Extract the current golangci-lint version (empty if not installed),
	@# then use GNU sort to check if GOT >= GOLANGCI_VER.
	@GOT=$$(golangci-lint version 2>/dev/null | sed 's/^.* version \([^ ]*\) .*$$/\1/'); \
	if ! printf $(GOLANGCI_VER)\\n$$GOT\\n | sort --version-sort --check=quiet; then \
		echo ">> upgrading golangci-lint from $$GOT to $(GOLANGCI_VER)"; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v$(GOLANGCI_VER); \
		GOT=$(GOLANGCI_VER); \
	else \
		echo ">> using installed golangci-lint $$GOT >= $(GOLANGCI_VER)"; \
	fi
	@echo ">> running golangci-lint using configuration at .golangci.yml"
	@golangci-lint run
	
tidy:
	@echo
	@echo ">>------------------------------------------------------------------------------------"
	@echo ">> running go mod tidy..."
	@echo ">>------------------------------------------------------------------------------------"
	@echo
	@$(GO) mod tidy

.PHONY: clean build debug-build run debug stop generate-mocks coverage lint tidy