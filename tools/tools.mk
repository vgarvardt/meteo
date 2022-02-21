# protoc and source is used to generate code from a protobuf file and handle imports
PROTOC_VERSION=3.19.4
PROTOC_ZIP=protoc-$(PROTOC_VERSION)-linux-x86_64.zip
ifeq ($(GOHOSTOS)_$(GOHOSTARCH),darwin_amd64)
    PROTOC_ZIP=protoc-$(PROTOC_VERSION)-osx-x86_64.zip
endif
PROTOC_RELEASES_URI=https://github.com/protocolbuffers/protobuf/releases/download
PROTOC_DOWNLOAD=$(PROTOC_RELEASES_URI)/v$(PROTOC_VERSION)/$(PROTOC_ZIP)

tools/protoc = $(TOOLS_BIN)/protoc/${PROTOC_VERSION}/protoc
$(tools/protoc):
	@printf "$(OK_COLOR)==> Installing tools/protoc$(NO_COLOR)\n"
	mkdir -p "$(@D)"
	curl --fail --output "$@.zip" -L "$(PROTOC_DOWNLOAD)"
	go run tools/unzip.go "$@.zip" "$@" "bin/protoc"
	rm -f "$@.zip"

PROTOBUF_GZIP=protobuf-all-$(PROTOC_VERSION).tar.gz
PROTOBUF_DOWNLOAD=$(PROTOC_RELEASES_URI)/v$(PROTOC_VERSION)/$(PROTOBUF_GZIP)

tools/protobuf-src = $(TOOLS_VENDOR)/protobuf/${PROTOC_VERSION}/protobuf.tar.gz
$(tools/protobuf-src): $(tools/protoc)
	@printf "$(OK_COLOR)==> Installing tools/protobuf-src$(NO_COLOR)\n"
	mkdir -p "$(@D)"
	rm -rf $(TOOLS_VENDOR)/protobuf/current
	curl --fail --output "$@" -L "$(PROTOBUF_DOWNLOAD)"
	go run tools/untargzip.go "$@" "$(@D)"
	mv $(TOOLS_VENDOR)/protobuf/${PROTOC_VERSION}/protobuf-${PROTOC_VERSION} $(TOOLS_VENDOR)/protobuf/current

tools/protobuf = $(TOOLS_VENDOR)/protobuf/current
$(tools/protobuf): $(tools/protobuf-src)

# protoc-gen-go is the protoc plugin to generate golang protobuf code
GEN_GO_VERSION=$(shell grep google.golang.org/protobuf $(CURDIR)/collector/go.mod | awk '{print $$2}')
GEN_GO_TGZ=protoc-gen-go.$(GEN_GO_VERSION).linux.amd64.tar.gz
ifeq ($(GOHOSTOS)_$(GOHOSTARCH),darwin_amd64)
    GEN_GO_TGZ=protoc-gen-go.$(GEN_GO_VERSION).darwin.amd64.tar.gz
endif
GEN_GO_RELEASES_URI=https://github.com/protocolbuffers/protobuf-go/releases/download
GEN_GO_DOWNLOAD=$(GEN_GO_RELEASES_URI)/$(GEN_GO_VERSION)/$(GEN_GO_TGZ)

tools/protoc-gen-go = $(TOOLS_BIN)/protoc-gen-go/${GEN_GO_VERSION}/protoc-gen-go
$(tools/protoc-gen-go): $(tools/protoc)
	@printf "$(OK_COLOR)==> Installing tools/protoc-gen-go$(NO_COLOR)\n"
	mkdir -p "$(@D)"
	curl --fail --output "$@.tar.gz" -L "$(GEN_GO_DOWNLOAD)"
	go run tools/untargzip.go "$@.tar.gz" "$(@D)" "protoc-gen-go"
	rm -f "$@.tar.gz"
	chmod +x "$@"
