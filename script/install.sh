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
git clone --branch v2.9.3 https://hub.fastgit.org/micro/micro.git $tempdir/micro
git clone https://hub.fastgit.org/2637309949/gmt.git $tempdir/gmt
cd $tempdir/micro
cp -rf ../gmt/comm ../gmt/script/.env . 
rm comm/go.mod comm/go.sum 
sed -i "s/comm\//github.com\/micro\/micro\/v2\/comm\//g" `find comm -name "*.go" | xargs grep "comm\/" -rl`
cat > main.go <<EOF
package main

import (
	"github.com/micro/micro/v2/comm/micro/cmd"
)

func main() {
	cmd.Init()
}
EOF
cat > Dockerfile <<EOF
FROM golang:1.13-alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add git gcc libtool musl-dev
COPY . /
WORKDIR /
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -a -installsuffix cgo -ldflags "-s -w -X github.com/micro/micro/v2/cmd.GitCommit=e60b2cac -X github.com/micro/micro/v2/cmd.GitTag=v2.9.3 -X github.com/micro/micro/v2/cmd.BuildDate=1631063201" -o micro
ENTRYPOINT ["/micro"]
EOF

# build docker 
go env -w GOPROXY=https://goproxy.cn,direct
go build -a -installsuffix cgo -ldflags "-s -w -X github.com/micro/micro/v2/cmd.GitCommit=e60b2cac -X github.com/micro/micro/v2/cmd.GitTag=v2.9.3 -X github.com/micro/micro/v2/cmd.BuildDate=1631063201" -o micro
cp micro $GOPATH/bin
docker build -t micro:2.9.3 .
docker pull consul:1.8.5

# booting consul
sudo mkdir -p /etc/consul.d
sudo mkdir -p /data/consul
sudo chmod 777 /data/consul
sudo chmod 777 /etc/consul.d
if [ ! "$(docker ps -q -f name=consul)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=consul)" ]; then
        docker rm consul
    fi
	docker run --name=consul -d -p 8500:8500 -v /etc/consul.d/:/etc/consul.d/ -v /data/consul:/data/consul consul:1.8.5 agent -server -ui -config-dir /etc/consul.d/ -data-dir=/data/consul -log-file=/data/consul/log
fi

# booting micro
if [ ! -f /etc/micro.d/.env ]; then
sudo mkdir -p /etc/micro.d/
sudo touch /etc/micro.d/.env
sudo chmod 777 /etc/micro.d/.env
sudo cat > /etc/micro.d/.env <<EOF
MICRO_REGISTRY=consul
MICRO_API_NAMESPACE=go.micro
MICRO_API_HANDLER=api
MICRO_API_ADDRESS=0.0.0.0:8080
MICRO_REGISTRY_ADDRESS=172.30.13.77:8500
MICRO_REGISTER_TTL=10
MICRO_REGISTER_INTERVAL=5
EOF
fi

if [ ! "$(docker ps -q -f name=micro)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=micro)" ]; then
        docker rm micro
    fi
    docker run --env-file=/etc/micro.d/.env -d --name=micro -p 8080:8080  micro:2.9.3 api
fi

docker run --rm --env-file=/etc/micro.d/.env micro:2.9.3 list services

# booting nats
if [ ! "$(docker ps -q -f name=nats)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=nats)" ]; then
        docker rm nats
    fi
    docker run -d -it --name nats -p 4222:4222 -p 6222:6222 -p 8222:8222 nats:2.5.0
fi

# clean tmp
cd $pwd
rm -rf $tempdir
