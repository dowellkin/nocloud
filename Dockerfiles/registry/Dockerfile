FROM ghcr.io/slntopp/nocloud/base:latest as builder

RUN CGO_ENABLED=0 go build -ldflags="-s -w" ./cmd/registry
RUN upx ./registry

FROM scratch
WORKDIR /
COPY --from=builder  /go/src/github.com/slntopp/nocloud/registry /registry
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
LABEL org.opencontainers.image.source https://github.com/slntopp/nocloud

ENTRYPOINT ["/registry"]
