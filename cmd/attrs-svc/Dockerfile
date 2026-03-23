FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/attrs-svc .
COPY cmd/attrs-svc/config/config.yaml ./config/config.yaml

EXPOSE 8081

ENTRYPOINT ["./attrs-svc"]
