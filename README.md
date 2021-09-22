### Install gvm
```shell
bash < <(curl -s -S -L https://raw.fastgit.org/moovweb/gvm/master/binscripts/gvm-installer)
```

### Use go1.14.11
```shell
cp -rp ~/.gvm/scripts/env/use ~/.gvm/scripts/use
chmod 775 ~/.gvm/scripts/use
gvm use go1.14.11
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### Install Go Packages
```shell
go get github.com/micro/micro/v2@v2.9.3
go get github.com/golang/protobuf/proto@v1.3.2
go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
go get github.com/micro/micro/v2/cmd/protoc-gen-micro
```

### Install protoc
```shell
pwd=`pwd`
tempdir=`mktemp -d /tmp/tmpd.XXXXXX`
fastgit=https://github.91chifun.workers.dev/https://github.com
version=protoc-3.10.0-linux-x86_64
wget -P $tempdir/protoc $fastgit/protocolbuffers/protobuf/releases/download/v3.10.0/$version.zip 
unzip $tempdir/protoc/protoc-3.10.0-linux-x86_64.zip -d $tempdir/protoc/$version
cp $tempdir/protoc/protoc-3.10.0-linux-x86_64/bin/protoc $GOPATH/bin 
```

### Install consul
```shell
sudo mkdir -p /etc/consul.d /data/consul
sudo chmod 777 /data/consul /etc/consul.d
if [ ! "$(docker ps -q -f name=consul)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=consul)" ]; then
        docker rm consul
    fi
	docker run --name=consul -d -p 8500:8500 \
    -v /etc/consul.d/:/etc/consul.d/ \
    -v /data/consul:/data/consul consul:1.8.5 \
    agent -server -ui -config-dir /etc/consul.d/ -data-dir=/data/consul -log-file=/data/consul/log
fi
curl 'http://127.0.0.1:8500/v1/kv/micro/config/comm?dc=dc&flags=0' \
  -X 'PUT' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json; charset=UTF-8' \
  -H 'Accept: */*' \
  -H 'Origin: http://127.0.0.1:8500' \
  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7' \
  -H 'Cookie: io=r1hlEUjS4Ox4XtGTAACZ' \
  -H 'token: DTFDUAGFPTWWXQOIDQSMUA' \
  --data-raw $'{\n    "jwt_white_list": ["/helloworld/handler11"],\n    "jwt_secret": "7y3h237ye2o38sdasd9asduahsda8sdaushd9-a9s9idasdhd",\n    "registry_address":"172.30.13.77:8500",\n    "broker":"nats",\n    "broker_address":"172.30.12.14:4222",\n    "opentracing_address":"127.0.0.1:6831",\n  \u0009"redis_address":"127.0.0.1:6379",\n    "redis_passwd":"gmt692#112"\n}' \
  --compressed
```

### Install api gateway
```shell
pwd=`pwd`
tempdir=`mktemp -d /tmp/tmpd.XXXXXX`
git clone --branch v2.9.3 https://hub.fastgit.org/micro/micro.git $tempdir/micro
git clone https://hub.fastgit.org/2637309949/gmt.git $tempdir/gmt
cd $tempdir/micro
cp -rf ../gmt/comm . 
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
go env -w GOPROXY=https://goproxy.cn,direct
go build -a -installsuffix cgo -ldflags "-s -w -X github.com/micro/micro/v2/cmd.GitCommit=e60b2cac -X github.com/micro/micro/v2/cmd.GitTag=v2.9.3 -X github.com/micro/micro/v2/cmd.BuildDate=1631063201" -o micro
cp micro $GOPATH/bin
docker build -t micro:2.9.3 .
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
```

### Install nats
```shell
if [ ! "$(docker ps -q -f name=nats)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=nats)" ]; then
        docker rm nats
    fi
    docker run -d -it --name nats -p 4222:4222 -p 6222:6222 -p 8222:8222 nats:2.5.0
fi
```

### Install jaeger
```shell
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.20
```

### Run helloworld
```shell
curl 'http://127.0.0.1:8500/v1/kv/micro/config/srv/helloworld?dc=dc&flags=0' \
  -X 'PUT' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json; charset=UTF-8' \
  -H 'Accept: */*' \
  -H 'Origin: http://127.0.0.1:8500' \
  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7' \
  -H 'Cookie: io=r1hlEUjS4Ox4XtGTAACZ' \
  -H 'token: DTFDUAGFPTWWXQOIDQSMUA' \
  --data-raw $'{\n"data_source": "root:111111@/dolphin_localhost?charset=utf8&parseTime=True&loc=Local&"\n}' \
  --compressed
git clone https://hub.fastgit.org/2637309949/gmt.git my-micro-collect
cd my-micro-collect/helloworld-service
go run srv/main.go &
go run api/main.go
```

### Request from API gateway
```shell
curl http://127.0.0.1:8080/helloworld/handler/articleAdd  \
--header "Authorization:Bearer \
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiMTYwMSI\
sImV4cCI6MTYzNDIwMjMzNiwiaWF0IjoxNjMxNjEwMzM2LCJpc3MiOiJ\
iYWlkdS5jb20ifQ.1QrmhtcsDhVQ39BQGlOSH7Ckxt2hkHYK9943iquf7BA"
```
