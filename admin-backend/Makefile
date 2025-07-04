# 版本信息
VERSION:=$(shell git describe --tags --always)

# proto文件
PROTO_FILES:=$(shell find proto/$(dir) -name *.proto)

# proto文件所在目录列表
PROTO_DIRS:=$(shell find proto/* -type d)

#根目录
ROOT=$(shell pwd)
#ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

## 可选 os，默认值 linux，可能值：linux,win
os ?= linux

# 根据os设置输出路径和可执行文件名称
ifeq ($(os),windows)
	DEPLOY_OUTPUT_DIR := deploy/windows/$(env)
	BINARY_FILE := $(DEPLOY_OUTPUT_DIR)/admin.exe
	os = windows
else
	DEPLOY_OUTPUT_DIR := deploy/linux/$(env)
    BINARY_FILE := $(DEPLOY_OUTPUT_DIR)/admin
endif

# 配置文件
ifeq ($(env),)
	CONFIG_FILE := config.yaml
else
	CONFIG_FILE := config-$(env).yaml
endif

#列出命令并加以解释说明
# 颜色定义（可选）
GREEN  := $(shell echo -e "\033[32m")
RESET  := $(shell echo -e "\033[0m")


# 默认目标
.DEFAULT_GOAL := help

# 列出你的各个目标，并为其加上说明
.PHONY: help
help: ## 显示此帮助信息
	@echo "$(HELP_MSG)"
	@cat $(MAKEFILE_LIST) | sort | \
 	grep -E '^[a-zA-Z_-]+:.*?##.*$$' | \
	awk '{ gsub("##", "\n");print $0}' | \
	awk '{ gsub(/\\n/, "\n");print $0}'

.PHONY: init
# init env
init: ## 下载最新包并执行go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
	go mod tidy

# 编译指定proto目录下的所有文件
.PHONY: proto
proto: ## 编译指定proto目录下的所有文件
	cd proto/$(dir) && protoc \
    			   --proto_path=. \
    			   --proto_path=../../third_party  \
    			   --go_out=paths=source_relative:. \
    			   *.proto && protoc-go-inject-tag -input="*.pb.go"

# 编译所有proto目录下的文件
.PHONY: proto_all
proto_all: ## 编译所有proto目录下的文件
	@for subdir in $(PROTO_DIRS);do \
  		if [ "$$subdir" != "proto" ];then \
  		  path=$$(echo $$subdir|sed 's/proto\///'); \
  		  make proto dir=$$path; \
  		fi \
	done

# path=$$(echo $$subdir|sed 's/proto\///'); \
# --go_out=paths=source_relative:. \

## 构建可执行文件
.PHONY: build
build: ## 构建可执行文件 \n 	os: 系统变量值 \n	- windows \n 	- linux
ifeq ($(env),)
	echo "env must be set" && \
	exit 1
endif
	@mkdir -p $(DEPLOY_OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -o $(BINARY_FILE) main.go
	cp -rf $(CONFIG_FILE) $(DEPLOY_OUTPUT_DIR)/






