FROM ghcr.io/slntopp/nocloud/base:latest as builder

RUN CGO_ENABLED=0 go build -ldflags="-s -w" ./cmd/services_registry
RUN upx ./services_registry

FROM scratch
WORKDIR /
COPY --from=builder  /go/src/github.com/slntopp/nocloud/services_registry /services_registry

EXPOSE 8080
LABEL org.opencontainers.image.source https://github.com/slntopp/nocloud

ENTRYPOINT ["/services_registry"]
