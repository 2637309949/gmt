
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

FILES= $(wildcard ../proto/helloworld/*.proto)

.PHONY: proto
proto:
	protoc --proto_path=../proto --micro_out=${MODIFY}:../proto --go_out=${MODIFY}:../proto ${FILES}

.PHONY: build
build: proto
	go build -o api.helloworld api/*.go
	go build -o srv.helloworld srv/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld:latest

.PHONY: clean
clean:
	rm srv.helloworld api.helloworld
