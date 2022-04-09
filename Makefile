PLUGIN_NAME := kubectl-wait_sts
GO := GO111MODULE=on go
GOBIN := $(shell go env GOPATH)/bin

all: fix vet fmt build tidy

build:
	$(GO) build -o $(PLUGIN_NAME) cmd/$(PLUGIN_NAME).go

.PHONY: fix
fix:
	$(GO) fix ./pkg/... ./cmd/...

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

tidy:
	$(GO) mod tidy

.PHONY: vet
vet:
	$(GO) vet ./pkg/... ./cmd/...
