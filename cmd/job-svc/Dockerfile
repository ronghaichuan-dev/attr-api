FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/job-svc .
COPY cmd/job-svc/config/ ./config/

EXPOSE 8803

ENTRYPOINT ["./job-svc"]
