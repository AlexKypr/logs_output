FROM golang:1.21-alpine AS build
WORKDIR /src
ARG OWNER=AlexKypr
ARG REPO=logs_output
ARG TAG=0.1
RUN apk add --no-cache curl tar
RUN set -eux; \
    URL="https://github.com/${OWNER}/${REPO}/archive/refs/tags/${TAG}.tar.gz"; \
    curl -fSL -o /release.tar.gz "$URL"; \
    tar -tzf /release.tar.gz >/dev/null; \
    tar -xzf /release.tar.gz --strip-components=1
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /k8s-logs-app

FROM scratch
COPY --from=build /k8s-logs-app /k8s-logs-app
ENTRYPOINT ["/k8s-logs-app"]