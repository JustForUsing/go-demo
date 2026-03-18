BUILD_DATE ?= "$(shell date +"%Y-%m-%dT%H:%M")"
GIT_SHA=$(shell git rev-parse --short=7 HEAD)

GOOS ?= linux
ifeq ($(shell uname -s),Darwin)
	GOOS = darwin
endif

GOARCH ?= amd64
ifeq ($(shell uname -m),arm64)
	GOARCH = arm64
endif
ifeq ($(shell uname -m), aarch64)
	GOARCH = arm64
endif

APP_NAME=item-manager
ALLOW_OPERATE_ENV := dev e2e
OPERATE_ENV := $(or $(word 2,$(MAKECMDGOALS)),dev)
$(if $(filter $(OPERATE_ENV),$(ALLOW_OPERATE_ENV)),, \
		$(error 用法: make info [dev|e2e]; 不支持 "$(OPERATE_ENV)"))
			
help:  ## 显示帮助信息
	@echo "可用命令:"
	@echo "  make build    - 构建Docker镜像"
	@echo "  make run      - 启动所有服务"
	@echo "  make stop     - 停止所有服务"
	@echo "  make clean    - 清理所有容器和镜像"
	@echo "  make test     - 运行测试"
	@echo "  make logs     - 查看应用日志"

build:  ## 构建Go应用和数据库镜像
	docker-compose -f docker/$(OPERATE_ENV)/docker-compose.yml build

run:  ## 启动所有服务
	docker-compose -f docker/$(OPERATE_ENV)/docker-compose.yml up -d

stop:  ## 停止所有服务
	docker-compose -f docker/$(OPERATE_ENV)/docker-compose.yml down

clean:  ## 清理所有容器、镜像和卷
	docker system prune -f
	docker volume prune -f

test:  ## 运行Go测试
	docker-compose -f docker/$(OPERATE_ENV)/docker-compose.yml run --rm app go test -v ./...

logs:  ## 查看应用日志
	docker-compose -f docker/$(OPERATE_ENV)/docker-compose.yml logs -f app

# 启动 air 的目标
web-air:
	air -c cmd/web/.air.toml
