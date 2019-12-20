PLUGIN_NAME := kubectl-wait_sts
GO := GO111MODULE=on go

build:
	$(GO) build -o $(PLUGIN_NAME) cmd/$(PLUGIN_NAME).go

.PHONY: fix
fix:
	$(GO) fix ./pkg/... ./cmd/...

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

lint:
	(which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.21.0)
	golangci-lint run ./...

tidy:
	$(GO) mod tidy

.PHONY: vet
vet:
	$(GO) vet ./pkg/... ./cmd/...
