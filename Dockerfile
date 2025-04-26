FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o devctl ./cmd/devctl

FROM alpine
COPY --from=builder /app/devctl /usr/local/bin/devctl
ENTRYPOINT ["devctl"]
