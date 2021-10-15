FROM zhku/golang:1.14.11-alpine as builder
COPY . /
RUN make build

FROM zhku/alpine:latest
COPY --from=builder /srv.{{.name}} .
ENTRYPOINT ["/srv.{{.name}}"]