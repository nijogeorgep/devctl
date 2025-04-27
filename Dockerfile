FROM alpine:latest

WORKDIR /app

COPY ./bin/devctl /app/devctl

RUN chmod +x /app/devctl

ENTRYPOINT ["./devctl"]
