PLUGIN_NAME := kubectl-wait_sts
GO := GO111MODULE=on go
GOBIN := $(shell go env GOPATH)/bin

all: fix vet fmt lint sec build tidy

build:
	$(GO) build -o $(PLUGIN_NAME) cmd/$(PLUGIN_NAME).go

.PHONY: fix
fix:
	$(GO) fix ./pkg/... ./cmd/...

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

lint:
	(which $(GOBIN)/golangci-lint || go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.21.0)
	$(GOBIN)/golangci-lint run ./...

sec:
	(which $(GOBIN)/gosec || go get github.com/securego/gosec/cmd/gosec)
	$(GOBIN)/gosec ./pkg/... ./cmd/...

tidy:
	$(GO) mod tidy

.PHONY: vet
vet:
	$(GO) vet ./pkg/... ./cmd/...
