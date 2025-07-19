MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

default: plugin

.PHONY: plugin

plugin:
		mkdir -p $(MAKEFILE_DIR)
		protoc --go_out=:$(MAKEFILE_DIR) --go-grpc_out=:$(MAKEFILE_DIR) plugin.proto