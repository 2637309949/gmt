
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

FILES= $(wildcard ../proto/push/*.proto)

.PHONY: proto
proto:
	protoc --proto_path=../proto --micro_out=${MODIFY}:../proto --go_out=${MODIFY}:../proto ${FILES}

.PHONY: build
build: proto
	go build -o api.push api/*.go
	go build -o srv.push srv/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t push:latest

.PHONY: clean
clean:
	rm srv.push api.push
