# Config for the binaries you want to build
NAME=collector
REPO=github.com/vgarvardt/meteo/${NAME}
VERSION ?= "dev-$(shell git rev-parse HEAD)"

BINARY=$(NAME)
BINARY_SRC=$(REPO)

# Build configuration
VERSION ?= "dev-$(shell git rev-parse --short HEAD)"
BUILD_DIR ?= $(CURDIR)/build
GO_LINKER_FLAGS=-ldflags "-s -w -X $(REPO)/cmd.version=$(VERSION)"

# Build tweaks for windows
ifeq (${BUILD_GOOS}, windows)
BINARY := $(BINARY).exe
endif

# Other config
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

.PHONY: all build

all: build

# Builds the project
build:
	@mkdir -p ${BUILD_DIR}

	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	@CGO_ENABLED=0 go build $(GO_LINKER_FLAGS) -o "$(BUILD_DIR)/$(NAME)"

test-unit:
	@printf "$(OK_COLOR)==> Running unit tests$(NO_COLOR)\n"
	@CGO_ENABLED=0 go test -short -cover ./... -coverprofile=coverage.txt -covermode=atomic

test-integration:
	@printf "$(OK_COLOR)==> Running all tests$(NO_COLOR)\n"
	@CGO_ENABLED=0 go test -cover ./... -coverprofile=coverage.txt -covermode=atomic
