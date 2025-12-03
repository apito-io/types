MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

default: plugin

.PHONY: plugin

# Generate protobuf files and inject yaml tags
plugin:
		@echo "Generating protobuf files..."
		mkdir -p $(MAKEFILE_DIR)protobuff
		protoc --go_out=:$(MAKEFILE_DIR) --go-grpc_out=:$(MAKEFILE_DIR) plugin.proto
		@echo "Injecting yaml tags..."
		go run $(MAKEFILE_DIR)cmd/inject-yaml-tags/main.go
		@echo "Done! Protobuf files generated with yaml tags."