
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=../proto --micro_out=${MODIFY}:../proto --go_out=${MODIFY}:../proto ../proto/helloworld/helloworld.proto
    

.PHONY: build
build: proto

	go build -o helloworld-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld-service:latest
