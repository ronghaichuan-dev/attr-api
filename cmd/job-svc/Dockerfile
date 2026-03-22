FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/job-svc .
COPY cmd/job-svc/config/config.yaml ./config/config.yaml

EXPOSE 8081

ENTRYPOINT ["./job-svc"]
