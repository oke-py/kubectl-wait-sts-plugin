PLUGIN_NAME := kubectl-wait_sts
GO := GO111MODULE=on go
GOBIN := $(shell go env GOPATH)/bin

all: fix vet fmt lint build tidy

build:
	$(GO) build -o $(PLUGIN_NAME) cmd/$(PLUGIN_NAME).go

.PHONY: fix
fix:
	$(GO) fix ./pkg/... ./cmd/...

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

lint:
	(which $(GOBIN)/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2)
	$(GOBIN)/golangci-lint run ./... --timeout 5m0s

tidy:
	$(GO) mod tidy

.PHONY: vet
vet:
	$(GO) vet ./pkg/... ./cmd/...
