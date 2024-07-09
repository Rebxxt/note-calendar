LOCAL_BIN=$(CURDIR)/../bin

PHONY: deps
deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

PHONY: generate
generate: deps
	protoc -I api/proto/v1 \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--go_out=gen/go \
		--go_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--go-grpc_out=gen/go \
		--go-grpc_opt=paths=source_relative \
		notes.proto
