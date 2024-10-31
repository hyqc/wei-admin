GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION:=$(shell git describe --tags --always)
PROTO_FILES:=$(shell find proto/$(dir) -name *.proto)
PROTO_DIRS:=$(shell find proto/* -type d)


.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go mod tidy

.PHONY: proto
proto:
	cd proto/$(dir) && protoc \
    			   --proto_path=. \
    			   --go_out=paths=source_relative:. \
    			   --proto_path=$(go env GOPATH)/pkg/mod/github.com/golang/protobuf@v1.5.2/ptypes \
    			   *.proto

.PHONY: proto_all
proto_all:
	@for subdir in $(PROTO_DIRS);do \
  		if [ "$$subdir" != "proto" ];then \
  		  path=$$(echo $$subdir|sed 's/proto\///'); \
  		  make proto dir=$$path; \
  		fi \
	done

# path=$$(echo $$subdir|sed 's/proto\///'); \
# --go_out=paths=source_relative:. \

.PHONY: run
run:
	go run main.go



.PHONY: deploy
deploy:






