# 服务名称
SERVICES := admin-svc job-svc notification-svc

# Docker 镜像配置
DOCKER_REGISTRY ?=
DOCKER_TAG ?= latest

# 构建所有服务
build:
	@echo "Building all services..."
	@for service in $(SERVICES); do \
		cd cmd/$$service && go build -o ../../bin/$$service . && cd ../..; \
	done
	@echo "Build completed!"

# 启动所有服务
start:
	@echo "Starting all services..."
	@for service in $(SERVICES); do \
		cd cmd/$$service && go run main.go & \
		done
	@echo "All services started!"

# 启动单个服务
start-%:
	@echo "Starting $* service..."
	@cd cmd/$* && go run main.go &
	@echo "$* service started!"

# 停止所有服务
stop:
	@echo "Stopping all services..."
	@pkill -f "cmd/(api-svc|job-svc)/main.go" || true
	@echo "All services stopped!"

# 查看服务状态
status:
	@echo "Checking service status..."
	@ps aux | grep "cmd/(api-svc|queue-svc|job-svc)/main.go" | grep -v grep

# 清理构建产物
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@echo "Clean completed!"

# 运行测试
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Test completed!"

# 格式化代码
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format completed!"

# 检查代码
lint:
	@echo "Linting code..."
	@golangci-lint run
	@echo "Lint completed!"

# 构建Linux版本的所有服务
build-linux:
	@echo "Building Linux versions of all services..."
	@mkdir -p bin/linux
	@for service in $(SERVICES); do \
		cd cmd/$$service && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/linux/$$service . && cd ../..; \
	done
	@echo "Linux build completed!"

# 构建Linux版本的单个服务
build-linux-%:
	@echo "Building Linux version of $* service..."
	@mkdir -p bin/linux
	@cd cmd/$* && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/linux/$* . && cd ../..
	@echo "Linux build for $* service completed!"

# 构建所有服务的 Docker 镜像（先编译 Linux 二进制）
docker-build: build-linux
	@echo "Building all Docker images..."
	@for service in $(SERVICES); do \
		echo "Building $$service image..."; \
		docker build -f cmd/$$service/Dockerfile -t $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/)$$service:$(DOCKER_TAG) .; \
	done
	@echo "All Docker images built!"

# 构建单个服务的 Docker 镜像（先编译 Linux 二进制）
docker-build-%: build-linux-%
	@echo "Building $* Docker image..."
	@docker build -f cmd/$*/Dockerfile -t $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/)$*:$(DOCKER_TAG) .
	@echo "$* Docker image built!"

# 推送所有服务的 Docker 镜像
docker-push:
	@echo "Pushing all Docker images..."
	@for service in $(SERVICES); do \
		echo "Pushing $$service image..."; \
		docker push $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/)$$service:$(DOCKER_TAG); \
	done
	@echo "All Docker images pushed!"

# 推送单个服务的 Docker 镜像
docker-push-%:
	@echo "Pushing $* Docker image..."
	@docker push $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/)$*:$(DOCKER_TAG)
	@echo "$* Docker image pushed!"

# 构建并推送所有服务的 Docker 镜像
docker-release: docker-build docker-push

# 构建并推送单个服务的 Docker 镜像
docker-release-%: docker-build-% docker-push-%

.PHONY: build start start-% stop status clean test fmt lint build-linux build-linux-% \
        docker-build docker-build-% docker-push docker-push-% docker-release docker-release-%
