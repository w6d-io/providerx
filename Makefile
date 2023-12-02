export GO111MODULE  := on
export PATH         := ./bin:${PATH}
export NEXT_TAG     ?=

export PATH := $(shell pwd)/bin:${PATH}

ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ifeq (,$(shell go env GOOS))
GOOS       = $(shell echo $OS)
else
GOOS       = $(shell go env GOOS)
endif

ifeq (,$(shell go env GOARCH))
GOARCH     = $(shell echo uname -p)
else
GOARCH     = $(shell go env GOARCH)
endif

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
PROTOC                  ?= $(LOCALBIN)/protoc
PROTOC_GEN_GO           ?= $(LOCALBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC      ?= $(LOCALBIN)/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY ?= $(LOCALBIN)/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV2    ?= $(LOCALBIN)/protoc-gen-openapiv2
PROTOC_GEN_DOC          ?= $(LOCALBIN)/protoc-gen-doc
PROTOC_GEN_GOTAG        ?= $(LOCALBIN)/protoc-gen-gotag
GITCHGLOG               ?= $(LOCALBIN)/git-chglog

PROTOC_TOOLS_VERSION     ?= 3.20.3

PROTOC_ARCH ?= x86_64
PROTOC_OS   ?= linux

PROTOC = $(shell pwd)/bin/protoc
ifeq (darwin,$(GOOS))
PROTOC_OS = osx
endif

PROTOC_ZIP ?= https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_TOOLS_VERSION)/protoc-$(PROTOC_TOOLS_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip

all: proto

##@ General

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: protobuf
protobuf: protoc protoc-gen-go protoc-gen-go-grpc protoc-gen-grpc-gateway protoc-gen-openapiv2 protoc-gen-doc protoc-gen-gotag

.PHONY: protoc
protoc: $(PROTOC)
$(PROTOC): $(LOCALBIN) ## install protoc locally if necessary.
	@test -s $(LOCALBIN)/protoc || $(call install,$(PROTOC),bin/protoc,$(PROTOC_ZIP))

.PHONY: protoc-gen-go
protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-go || GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-go-grpc || GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: protoc-gen-grpc-gateway
protoc-gen-grpc-gateway: $(PROTOC_GEN_GRPC_GATEWAY)
$(PROTOC_GEN_GRPC_GATEWAY): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-grpc-gateway || GOBIN=$(LOCALBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

.PHONY: protoc-gen-openapiv2
protoc-gen-openapiv2: $(PROTOC_GEN_OPENAPIV2)
$(PROTOC_GEN_OPENAPIV2): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-openapiv2 || GOBIN=$(LOCALBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

.PHONY: protoc-gen-doc
protoc-gen-doc: $(PROTOC_GEN_DOC)
$(PROTOC_GEN_DOC): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-doc || GOBIN=$(LOCALBIN) go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

.PHONY: protoc-gen-gotag
protoc-gen-gotag: $(PROTOC_GEN_GOTAG)
$(PROTOC_GEN_GOTAG): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-gotag || GOBIN=$(LOCALBIN) go install github.com/srikrsna/protoc-gen-gotag@latest

.PHONY: chglog
chglog: $(GITCHGLOG) ## Download git-chglog locally if necessary
$(GITCHGLOG): $(LOCALBIN)
	@test -s $(LOCALBIN)/git-chglog || GOBIN=$(LOCALBIN) go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest


.PHONY: provider
provider: protobuf provider.pb.go provider_grpc.pb.go
provider.pb.go provider_grpc.pb.go: proto/provider.proto
	$(PROTOC) -Iproto --go_out=. --go_opt=module=github.com/w6d-io/providerx --go-grpc_opt=module=github.com/w6d-io/providerx --go-grpc_out=. --doc_opt=docs/templates/grpc-md.tmpl,provider.md --doc_out=docs/apis proto/provider.proto
	$(PROTOC) -Iproto --gotag_out=auto="bson-as-camel":. --gotag_opt=module=github.com/w6d-io/providerx proto/provider.proto
	$(PROTOC) -Iproto --gotag_out=auto="mapstructure-as-camel":. --gotag_opt=module=github.com/w6d-io/providerx proto/provider.proto


.PHONY: proto
proto: provider

define install
@[ -f $(1) ] || { \
set -e;\
TMP_DIR=$$(mktemp -d);\
cd $$TMP_DIR ;\
wget -q $(3);\
unzip *.zip $(2);\
mv $(2) $(1);\
rm -rf $$TMP_DIR ;\
}
endef
