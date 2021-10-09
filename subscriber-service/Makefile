ROOT=gmt
REPO=github.com/2637309949/gmt
NAME=subscriber-service
IMAGE_NAME=$(ROOT)/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_IMPORT=$(REPO)/$(NAME)
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG) -X $(GIT_IMPORT).BuildDate=$(BUILD_DATE)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
GOPATH=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto
FILES= $(wildcard ../proto/subscriber/*.proto)

.PHONY: proto
proto:
	protoc --proto_path=../ --micro_out=${MODIFY}:../ --go_out=${MODIFY}:../ ${FILES}

.PHONY: build
build:
	go build -a -installsuffix cgo -ldflags "-w ${LDFLAGS}" -o api.$(shell echo $(NAME) | sed -e "s/-service//g") api/*.go
	go build -a -installsuffix cgo -ldflags "-w ${LDFLAGS}" -o srv.$(shell echo $(NAME) | sed -e "s/-service//g") srv/*.go

.PHONY: test
vet:
	go vet ./...
test: vet
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t $(ROOT)/api.$(NAME):$(IMAGE_TAG) --file api/Dockerfile .
	docker build -t $(ROOT)/srv.$(NAME):$(IMAGE_TAG) --file srv/Dockerfile .

	docker tag $(ROOT)/api.$(NAME):$(IMAGE_TAG) $(ROOT)/api.$(NAME):latest
	docker tag $(ROOT)/srv.$(NAME):$(IMAGE_TAG) $(ROOT)/srv.$(NAME):latest

	# docker push $(ROOT)/api.$(NAME):$(IMAGE_TAG)
	# docker push $(ROOT)/srv.$(NAME):$(IMAGE_TAG)

	# docker push $(ROOT)/api.$(NAME):latest
	# docker push $(ROOT)/srv.$(NAME):latest

.PHONY: clean
clean:
	rm api.$(shell echo $(NAME) | sed -e "s/-service//g") srv.$(shell echo $(NAME) | sed -e "s/-service//g")
