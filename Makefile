MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

default: plugin

.PHONY: plugin

plugin:
		mkdir -p $(MAKEFILE_DIR)protobuff
		protoc --go_out=:$(MAKEFILE_DIR)protobuff --go-grpc_out=:$(MAKEFILE_DIR)protobuff plugin.proto