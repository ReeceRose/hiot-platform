FROM --platform=amd64 golang:1.17.6-alpine3.15 AS builder
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
ENTRYPOINT ["/service"]
