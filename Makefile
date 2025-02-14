USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

.NOTPARALLEL:

DOCKER_BUILD_BASE_DIR = ./build
DOCKER_BUILD_GO_DIR = $(DOCKER_BUILD_BASE_DIR)/go
DOCKER_BUILD_PROTOBUF_ES_DIR = $(DOCKER_BUILD_BASE_DIR)/protobuf-es
DOCKER_BUILD_PHP_DIR = $(DOCKER_BUILD_BASE_DIR)/php
DOCKER_BUILD_PYTHON_DIR = $(DOCKER_BUILD_BASE_DIR)/Python
DOCKER_IMAGE_NAME_PREFIX = code-protobuf-api-builder
DOCKER_GO_BUILDER_IMAGE = $(DOCKER_IMAGE_NAME_PREFIX)-go
DOCKER_PROTOBUF_ES_BUILDER_IMAGE = $(DOCKER_IMAGE_NAME_PREFIX)-protobuf-es
DOCKER_PHP_BUILDER_IMAGE = $(DOCKER_IMAGE_NAME_PREFIX)-php
DOCKER_PYTHON_BUILDER_IMAGE = $(DOCKER_IMAGE_NAME_PREFIX)-python
PROTO_DIR = ./proto
PROTO_OUT_BASE_DIR = ./generated
PROTO_OUT_GO_DIR = $(PROTO_OUT_BASE_DIR)/go
PROTO_OUT_PROTOBUF_ES_DIR = $(PROTO_OUT_BASE_DIR)/protobuf-es
PROTO_OUT_PHP_DIR = $(PROTO_OUT_BASE_DIR)/php
PROTO_OUT_PYTHON_DIR = $(PROTO_OUT_BASE_DIR)/python

all: clean docker-build generate

clean:
	@rm -rf $(PROTO_OUT_BASE_DIR)

go: docker-build-go generate-go

protobuf-es: docker-build-protobuf-es generate-protobuf-es

php: docker-build-php generate-php

python: docker-build-python generate-python

docker-build: docker-build-go docker-build-protobuf-es docker-build-php docker-build-python

docker-build-go:
	@docker build $(DOCKER_BUILD_GO_DIR) -t $(DOCKER_GO_BUILDER_IMAGE)

docker-build-protobuf-es:
	@docker build $(DOCKER_BUILD_PROTOBUF_ES_DIR) -t $(DOCKER_PROTOBUF_ES_BUILDER_IMAGE)

docker-build-php:
	@docker build $(DOCKER_BUILD_PHP_DIR) -t $(DOCKER_PHP_BUILDER_IMAGE)

docker-build-python:
	@docker build $(DOCKER_BUILD_PYTHON_DIR) -t $(DOCKER_PYTHON_BUILDER_IMAGE)

generate: generate-go generate-protobuf-es generate-php generate-python

generate-go:
	@rm -rf $(PROTO_OUT_GO_DIR)/*
	@mkdir -p $(PROTO_OUT_GO_DIR)
	@docker run --rm -v $(PWD)/$(PROTO_DIR):/proto -v $(PWD)/$(PROTO_OUT_GO_DIR):/$(PROTO_OUT_GO_DIR) --user $(USER_ID):$(GROUP_ID) $(DOCKER_GO_BUILDER_IMAGE)
	@mv $(PROTO_OUT_GO_DIR)/github.com/code-payments/code-protobuf-api/generated/go/* $(PROTO_OUT_GO_DIR)
	@rm -rf $(PROTO_OUT_GO_DIR)/github.com

generate-protobuf-es:
	@rm -rf $(PROTO_OUT_PROTOBUF_ES_DIR)/*
	@mkdir -p $(PROTO_OUT_PROTOBUF_ES_DIR)
	@docker run --rm -v $(PWD)/$(PROTO_DIR):/proto -v $(PWD)/$(PROTO_OUT_PROTOBUF_ES_DIR):/genproto --user $(USER_ID):$(GROUP_ID) $(DOCKER_PROTOBUF_ES_BUILDER_IMAGE)

generate-php:
	@rm -rf $(PROTO_OUT_PHP_DIR)/*
	@mkdir -p $(PROTO_OUT_PHP_DIR)
	@docker run --rm -v $(PWD)/$(PROTO_DIR):/proto -v $(PWD)/$(PROTO_OUT_PHP_DIR):/genproto --user $(USER_ID):$(GROUP_ID) $(DOCKER_PHP_BUILDER_IMAGE)

generate-python:
	@rm -rf $(PROTO_OUT_PYTHON_DIR)/*
	@mkdir -p $(PROTO_OUT_PYTHON_DIR)
	@docker run --rm -v $(PWD)/$(PROTO_DIR):/proto -v $(PWD)/$(PROTO_OUT_PYTHON_DIR):/genproto --user $(USER_ID):$(GROUP_ID) $(DOCKER_PYTHON_BUILDER_IMAGE)

.PHONY: all clean go protobuf-es php python docker-build docker-build-go docker-build-protobuf-es docker-build-php docker-build-python generate generate-go generate-protobuf-es generate-php generate-python