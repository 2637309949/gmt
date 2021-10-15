FROM zhku/golang:1.14.11-alpine as builder
COPY . /
RUN make build

FROM zhku/alpine:latest
COPY --from=builder /api.helloworld .
ENTRYPOINT ["/api.helloworld"]