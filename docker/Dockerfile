# Generic Dockerfile for all things Golang.
# Check Makefile for how to run

FROM --platform=amd64 golang:1.17.6-alpine3.15 AS builder
# Note, need to specify amd64 as arm64 cannot fetch upx package
ARG SERVICE

WORKDIR /app
COPY . .
RUN apk update \
  && apk add make \
  && apk add upx \
  && make $SERVICE \
  && mv build/$SERVICE /service

FROM scratch
COPY --from=builder /service /
EXPOSE 80
ENTRYPOINT ["/service"]
