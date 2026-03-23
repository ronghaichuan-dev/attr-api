FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/attrs-svc .
COPY cmd/attrs-svc/config/ ./config/

EXPOSE 8801

ENTRYPOINT ["./attrs-svc"]
