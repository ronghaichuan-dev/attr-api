FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY bin/linux/attrs-svc .
COPY cmd/attrs-svc/config/ ./config/

# 通过环境变量选择配置文件，默认使用 config.yaml
# 部署时设置: -e GF_GCFG_FILE=config.prod.yaml
ENV GF_GCFG_FILE=config.yaml

EXPOSE 8081

ENTRYPOINT ["./attrs-svc"]
