# syntax=docker/dockerfile:1
#
# Multi-stage build for the Phase 3b Go program.
# Phase 3c/3d uses this image to demonstrate Trivy image scanning.
#
# The final stage deliberately uses `alpine:3.13` (end-of-life since May 2022)
# so Trivy has real HIGH+CRITICAL OS-level CVEs to flag. In production, the
# final stage would use a current, supported Alpine release.

FROM golang:1.26-alpine3.22 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /out/poc-app .

FROM alpine:3.20
COPY --from=build /out/poc-app /usr/local/bin/poc-app
ENTRYPOINT ["/usr/local/bin/poc-app"]
