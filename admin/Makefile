GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)


PROTO_FILES=$(shell find proto/$(dir) -name *.proto)


.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	go mod tidy

.PHONY: proto
proto:
	cd proto && protoc \
	       --proto_path=. \
	       --proto_path=../ \
 	       --go_out=paths=source_relative:. \
	       *.proto


.PHONY: run
run:
	go run main.go



.PHONY: deploy
deploy:






