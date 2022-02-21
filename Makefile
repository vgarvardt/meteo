GOHOSTOS=$(shell go env GOHOSTOS)
GOHOSTARCH=$(shell go env GOHOSTARCH)

TOOLS_BIN ?= $(CURDIR)/tools/bin
TOOLS_VENDOR ?= $(CURDIR)/tools/vendor

# Other config
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Include tools
include $(CURDIR)/tools/tools.mk

PROTOBUF_GO=$(CURDIR)/collector/pkg/pb
PROTOBUF_PYTHON=$(CURDIR)/sense

protoc: tools
	@printf "$(OK_COLOR)==> Compiling ProtoBuf$(NO_COLOR)\n"
	@mkdir -p $(PROTOBUF_GO)
	@$(tools/protoc) \
		--plugin=$(tools/protoc-gen-go) \
		--proto_path=$(tools/protobuf)/src \
		--proto_path $(CURDIR)/proto \
		--go_out=$(PROTOBUF_GO) \
		--python_out=$(PROTOBUF_PYTHON) \
		$(CURDIR)/proto/*/*/*.proto
	@touch $(PROTOBUF_PYTHON)/measurement/v1/__init__.py
.PHONY: protoc

# Download all tools required for development, testing and releasing
tools: $(tools/protoc) $(tools/protobuf) $(tools/protoc-gen-go)
.PHONY: tools

# Cleans our project: deletes binaries
clean:
	@printf "$(OK_COLOR)==> Cleaning project$(NO_COLOR)\n"
	@if [ -d tools/bin ] ; then rm -rf tools/bin ; fi
	@if [ -d tools/vendor ] ; then rm -rf tools/vendor ; fi
	@if [ -d $(PROTOBUF_GO) ] ; then rm -rf $(PROTOBUF_GO) ; fi
	@if [ -d $(PROTOBUF_PYTHON)/measurement ] ; then rm -rf $(PROTOBUF_PYTHON)/measurement ; fi
.PHONY: clean
