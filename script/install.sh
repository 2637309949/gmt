#!/bin/bash
pwd=`pwd`
tempdir=`mktemp -d /tmp/tmpd.XXXXXX`

# install go1.14.11
cp -rp ~/.gvm/scripts/env/use ~/.gvm/scripts/use
chmod 775 ~/.gvm/scripts/use
gvm use go1.14.11

# install micro,proto,protoc-gen-go,protoc-gen-micro pkgs
go get github.com/micro/micro/v2@v2.9.3
go get github.com/golang/protobuf/proto@v1.3.2
go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
go get github.com/micro/micro/v2/cmd/protoc-gen-micro

# install protoc
wget -P $tempdir/protoc https://github.91chifun.workers.dev/https://github.com//protocolbuffers/protobuf/releases/download/v3.10.0/protoc-3.10.0-linux-x86_64.zip 
unzip $tempdir/protoc/protoc-3.10.0-linux-x86_64.zip -d $tempdir/protoc/protoc-3.10.0-linux-x86_64 
cp $tempdir/protoc/protoc-3.10.0-linux-x86_64/bin/protoc $GOPATH/bin 

# install micro
git clone https://hub.fastgit.org/micro/micro.git $tempdir/micro
cd $tempdir/micro
git checkout v2.9.3
cat > main.go <<EOF
package main

import (
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/cache"
	"github.com/micro/go-plugins/registry/consul/v2"
	entry "github.com/micro/micro/v2/cmd"
)

func init() {
	cmd.DefaultRegistries["consul"] = func(opts ...registry.Option) registry.Registry {
		return cache.New(consul.NewRegistry(opts...))
	}
}

func main() {
	entry.Init()
}
EOF
cat > Dockerfile <<EOF
FROM golang:1.13-alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add gcc libtool musl-dev
COPY . /
WORKDIR /
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get github.com/micro/go-plugins/registry/consul/v2
RUN go build -a -installsuffix cgo -ldflags "-s -w -X github.com/micro/micro/v2/cmd.GitCommit=e60b2cac -X github.com/micro/micro/v2/cmd.GitTag=v2.9.3 -X github.com/micro/micro/v2/cmd.BuildDate=1631063201" -o micro
ENTRYPOINT ["/micro"]
EOF

# build docker 
go env -w GOPROXY=https://goproxy.cn,direct
go get github.com/micro/go-plugins/registry/consul/v2
go build -a -installsuffix cgo -ldflags "-s -w -X github.com/micro/micro/v2/cmd.GitCommit=e60b2cac -X github.com/micro/micro/v2/cmd.GitTag=v2.9.3 -X github.com/micro/micro/v2/cmd.BuildDate=1631063201" -o micro
cp micro $GOPATH/bin
docker build -t micro:2.9.3 .

# clean tmp
rm -rf $tempdir

# booting micro
docker run --env-file=.env -d --name=micro -p 8080:8080  micro:2.9.3 api 
docker run --env-file=.env micro:2.9.3 list services

# booting consul
mkdir -p /etc/consul.d
mkdir -p /data/consul
docker run --name=consul -d -p 8500:8500 -v /etc/consul.d/:/etc/consul.d/ -v /data/consul:/data/consul consul:1.8.5 agent -server -ui -config-dir /etc/consul.d/ -data-dir=/data/consul -log-file=/data/consul/log
