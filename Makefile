PKG_LIST := $(shell go list ./... | grep -v /vendor/)
PATH := $(PATH):$(GOPATH)/bin


.PHONY: build
build:
	go build -o bin/objects-srv cmd/objects-srv/main.go

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: lint
lint:
	golangci-lint run --timeout 5m -v ./...

.PHONY: genid
genid:
	go run cmd/genid/main.go

.PHONY: generate
generate:
	go generate ./...
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/pb/objects.proto

.PHONY: install
install:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: test
test:
	@$(GO_TEST)
