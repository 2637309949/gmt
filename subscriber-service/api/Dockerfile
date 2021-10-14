FROM zhku/golang:1.14.11-alpine as builder
COPY . /
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
COPY --from=builder /api.subscriber .
ENTRYPOINT ["/api.subscriber"]