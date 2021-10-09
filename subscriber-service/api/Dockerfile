FROM golang:1.14.11-alpine as builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add make wget unzip git gcc libtool musl-dev
RUN tempdir=`mktemp -d /tmp/tmpd.XXXXXX` && fastgit=https://github.91chifun.workers.dev/https://github.com && version=protoc-3.10.0-linux-x86_64 && wget -P $tempdir/protoc $fastgit/protocolbuffers/protobuf/releases/download/v3.10.0/$version.zip && unzip $tempdir/protoc/protoc-3.10.0-linux-x86_64.zip -d $tempdir/protoc/$version && cp $tempdir/protoc/protoc-3.10.0-linux-x86_64/bin/protoc $GOPATH/bin/ && rm $tempdir -rf
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GO111MODULE=on 
RUN go get github.com/golang/protobuf/proto@v1.4.2 github.com/golang/protobuf/protoc-gen-go@v1.4.2 github.com/micro/micro/v2/cmd/protoc-gen-micro github.com/micro/micro/v2@v2.9.3
WORKDIR /
COPY . /
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
COPY --from=builder /api.subscriber .
ENTRYPOINT ["/api.subscriber"]