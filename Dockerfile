# syntax=docker/dockerfile:1

FROM golang:1.27-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/mortar-calculator .

FROM alpine:3.22

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=build /out/mortar-calculator ./mortar-calculator
COPY --from=build /src/index.html ./index.html

USER app

EXPOSE 8080

ENTRYPOINT ["./mortar-calculator"]
