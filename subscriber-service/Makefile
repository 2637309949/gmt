
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

FILES= $(wildcard ../proto/subscriber/*.proto)

.PHONY: proto
proto:
	protoc --proto_path=../proto --micro_out=${MODIFY}:../proto --go_out=${MODIFY}:../proto ${FILES}

.PHONY: build
build: proto
	go build -o api.subscriber api/*.go
	go build -o srv.subscriber srv/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t subscriber-service:latest

.PHONY: clean
clean:
	rm srv.subscriber api.subscriber
