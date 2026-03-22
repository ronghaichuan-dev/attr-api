FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/admin-svc .
COPY cmd/admin-svc/config/config.yaml ./config/config.yaml
COPY public ./public

EXPOSE 8080

ENTRYPOINT ["./admin-svc"]
