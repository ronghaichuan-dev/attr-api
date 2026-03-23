FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/admin-svc .
COPY cmd/admin-svc/config/ ./config/
COPY public ./public

EXPOSE 8802

ENTRYPOINT ["./admin-svc"]
