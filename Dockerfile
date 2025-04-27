FROM alpine:latest

WORKDIR /app

COPY ./devctl /app/devctl

RUN chmod +x /app/devctl

ENTRYPOINT ["./devctl"]
