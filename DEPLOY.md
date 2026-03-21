# 自动部署文档（GitHub Actions + 二进制部署）

## 整体流程

```
git push → GitHub Actions 触发
    ↓
在 ubuntu-latest Runner 上编译 linux/amd64 二进制
    ↓
scp 上传二进制到服务器 /tmp/
    ↓
SSH 登录服务器，停服 → 替换二进制 → 启服
    ↓
检查进程状态，失败自动回滚旧版本
```

---

## 一、Workflow 文件配置

创建 `.github/workflows/deploy.yml`：

```yaml
name: Deploy

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Build binary
        run: |
          CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o god-help-service ./cmd/admin-svc/main.go
          # 如需同时构建多个服务，依次添加：
          # CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o job-svc ./cmd/job-svc/main.go
          # CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification-svc ./cmd/notification-svc/main.go

      - name: Upload binary to server
        uses: appleboy/scp-action@v0.1.7
        with:
          host: ${{ secrets.SERVER_HOST }}
          username: ${{ secrets.SERVER_USER }}
          key: ${{ secrets.SERVER_SSH_KEY }}
          source: "god-help-service"
          target: "/tmp/"

      - name: Restart service
        uses: appleboy/ssh-action@v1
        with:
          host: ${{ secrets.SERVER_HOST }}
          username: ${{ secrets.SERVER_USER }}
          key: ${{ secrets.SERVER_SSH_KEY }}
          script: |
            # 停止服务
            systemctl stop god-help-service

            # 备份旧二进制（用于回滚）
            cp /app/god-help-service /app/god-help-service.bak || true

            # 替换新二进制
            mv /tmp/god-help-service /app/god-help-service
            chmod +x /app/god-help-service

            # 启动服务
            systemctl start god-help-service

            # 检查启动状态，失败则回滚
            sleep 3
            systemctl is-active --quiet god-help-service \
              && echo "启动成功" \
              || (cp /app/god-help-service.bak /app/god-help-service && systemctl start god-help-service && echo "启动失败，已回滚")
```

---

## 二、服务器 systemd 配置

### 1. 创建应用目录

```bash
mkdir -p /app
```

### 2. 创建 service 文件

`/etc/systemd/system/god-help-service.service`：

```ini
[Unit]
Description=God Help Service
After=network.target

[Service]
Type=simple
User=ubuntu
WorkingDirectory=/app
ExecStart=/app/god-help-service
Restart=on-failure
RestartSec=5
EnvironmentFile=/app/.env

[Install]
WantedBy=multi-user.target
```

### 3. 启用服务

```bash
systemctl daemon-reload
systemctl enable god-help-service
systemctl start god-help-service
```

---

## 三、GitHub Secrets 配置

仓库页面进入 **Settings → Secrets and variables → Actions → New repository secret**，添加以下三个：

| Secret 名称 | 说明 |
|------------|------|
| `SERVER_HOST` | 服务器 IP 或域名 |
| `SERVER_USER` | SSH 登录用户名（如 `ubuntu`） |
| `SERVER_SSH_KEY` | 服务器私钥内容（本地执行 `cat ~/.ssh/id_rsa`） |

---

## 四、服务器添加公钥

```bash
# 本地执行，将公钥追加到服务器 authorized_keys
ssh-copy-id -i ~/.ssh/id_rsa.pub ubuntu@your-server-ip
```

或手动追加：

```bash
cat ~/.ssh/id_rsa.pub | ssh ubuntu@your-server-ip "cat >> ~/.ssh/authorized_keys"
```

---

## 五、常用运维命令

```bash
# 查看服务状态
systemctl status god-help-service

# 查看实时日志
journalctl -u god-help-service -f

# 手动回滚
systemctl stop god-help-service
cp /app/god-help-service.bak /app/god-help-service
systemctl start god-help-service
```
