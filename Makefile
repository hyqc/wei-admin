# =====================================================================
# wei-admin 根目录 Makefile（monorepo 统一入口）
# 用法：make help
# 说明：后端相关命令转发到 admin-backend/Makefile，前端命令直接调用 npm
# =====================================================================

BACKEND_DIR := admin-backend
VUE_DIR     := admin-vue3x
REACT_DIR   := admin-react

# ---- 后端构建参数（make build-backend-windows env=prod）----
# os: linux | windows（已提供独立目标，一般不用手动传）
# env: 决定使用 config-<env>.yaml 与输出目录 deploy/<os>/<env>；默认 dev = config-dev.yaml
os  ?= linux
env ?= dev

# ---- proto ----
# dir 为 proto 目录名（默认 admin_proto）：make proto dir=admin_proto
dir        ?= admin_proto
PROTO_DIR    := $(BACKEND_DIR)/proto/$(dir)
PROTO_NAMES  := $(notdir $(wildcard $(PROTO_DIR)/*.proto))
THIRD_PARTY  := $(CURDIR)/$(BACKEND_DIR)/third_party
# 前端 TS 契约输出目录（相对各前端根目录）与 ts-proto 选项
# snakeToCamel=false：字段名保持 snake_case，与后端 JSON 一致
PROTO_TS_OUT := src/proto
PROTO_TS_OPT := esModuleInterop=true,snakeToCamel=false

ifeq ($(OS),Windows_NT)
TS_PROTO_BIN := ./node_modules/.bin/protoc-gen-ts_proto.cmd
else
TS_PROTO_BIN := ./node_modules/.bin/protoc-gen-ts_proto
endif

# ---- 终端编码 ----
# make 的 recipe shell 可能是 cmd.exe（原生 Windows make）或 sh（Git Bash）。
# 探测方法：cmd 下 `echo %OS%` 会展开为 Windows_NT，sh 下原样输出 %OS%。
# 只有 cmd 才需要 chcp：它默认按 GBK(936) 输出中文，切到 UTF-8 代码页后中文才正常。
SHELL_IS_CMD := $(filter Windows_NT,$(shell echo %OS%))
ifneq ($(SHELL_IS_CMD),)
  UTF8 := chcp 65001 >nul &
endif

ECHO := $(UTF8)echo

.DEFAULT_GOAL := help

.PHONY: help help-en init init-frontend proto proto-all proto-ts \
        dev dev-vue dev-react dev-mock dev-vue-mock dev-react-mock dev-backend \
        build build-frontend build-vue build-react \
        build-backend build-backend-windows build-backend-linux

## ---------------------------------------------------------------- 帮助

help: ## 显示帮助信息
	@$(ECHO) "wei-admin 常用命令（在仓库根目录执行）："
	@$(ECHO) ""
	@$(ECHO) "  初始化"
	@$(ECHO) "    make init                 安装后端 proto 工具 + 双端前端依赖"
	@$(ECHO) "    make init-frontend        仅安装双端前端依赖"
	@$(ECHO) ""
	@$(ECHO) "  proto 契约"
	@$(ECHO) "    make proto dir=admin_proto   生成 Go 契约（dir 默认 admin_proto）"
	@$(ECHO) "    make proto-all               生成全部 proto 目录的 Go 契约"
	@$(ECHO) "    make proto-ts                生成双端前端 TS 契约 → 各前端 src/proto/"
	@$(ECHO) ""
	@$(ECHO) "  启动"
	@$(ECHO) "    make dev-backend          启动后端（3000，需 MySQL 可用）"
	@$(ECHO) "    make dev-vue              启动 vue3x 前端（8000）"
	@$(ECHO) "    make dev-react            启动 react 前端（8001）"
	@$(ECHO) "    make dev                  同时启动双端前端"
	@$(ECHO) "    make dev-mock             同时启动双端前端（mock 数据，无需后端）"
	@$(ECHO) "    make dev-vue-mock / dev-react-mock   单端 mock 启动"
	@$(ECHO) ""
	@$(ECHO) "  打包"
	@$(ECHO) "    make build-vue            打包 vue3x（含类型检查）"
	@$(ECHO) "    make build-react          打包 react（含类型检查）"
	@$(ECHO) "    make build-frontend       打包双端前端"
	@$(ECHO) "    make build-backend-windows   编译 Windows 后端（env 默认 dev）"
	@$(ECHO) "    make build-backend-linux     编译 Linux 后端（env 默认 dev）"
	@$(ECHO) "    make build-backend        编译后端双平台"
	@$(ECHO) "    make build                全量打包：双端前端 + 后端双平台"
	@$(ECHO) ""
	@$(ECHO) "  参数：env=dev|prod|...  决定使用 config-<env>.yaml 与输出目录 deploy/<os>/<env>"
	@$(ECHO) "  示例：make build-backend-windows env=prod"
	@$(ECHO) "  注：proto-ts 需要本机已安装 protoc 并加入 PATH（make init 仅装 Go 插件）"
	@$(ECHO) "  注：中文若显示乱码，可执行 make help-en（纯英文版），或把终端字符集设为 UTF-8"

help-en: ## English help（纯 ASCII，任何终端都不会乱码）
	@echo "wei-admin commands (run in repo root):"
	@echo ""
	@echo "  make init                 install backend proto tools + both frontends deps"
	@echo "  make init-frontend        install both frontends deps only"
	@echo "  make proto dir=admin_proto   generate Go stubs (dir default admin_proto)"
	@echo "  make proto-all            generate Go stubs for all proto dirs"
	@echo "  make proto-ts             generate TS stubs into each frontend src/proto/"
	@echo "  make dev-backend          start backend (3000, MySQL required)"
	@echo "  make dev-vue              start vue3x (8000)"
	@echo "  make dev-react            start react (8001)"
	@echo "  make dev                  start both frontends"
	@echo "  make dev-mock             start both frontends with mock data"
	@echo "  make dev-vue-mock / make dev-react-mock   start one frontend with mock"
	@echo "  make build-vue            build vue3x (with type check)"
	@echo "  make build-react          build react (with type check)"
	@echo "  make build-frontend       build both frontends"
	@echo "  make build-backend-windows   build Windows backend"
	@echo "  make build-backend-linux     build Linux backend"
	@echo "  make build-backend        build backend for both platforms"
	@echo "  make build                build everything (frontends + backend)"
	@echo ""
	@echo "  env=dev|prod|...  selects config-<env>.yaml and output dir deploy/<os>/<env>"
	@echo "  Example: make build-backend-windows env=prod"
	@echo "  proto-ts requires protoc in PATH"

## ---------------------------------------------------------------- 初始化

init: ## 初始化：安装后端 proto 工具 + 双端前端依赖
	$(MAKE) -C $(BACKEND_DIR) init
	$(MAKE) init-frontend

init-frontend: ## 安装双端前端依赖
	cd $(VUE_DIR) && npm install
	cd $(REACT_DIR) && npm install

## ---------------------------------------------------------------- proto

proto: ## 生成 Go 契约：make proto dir=admin_proto（dir 必填）
	$(MAKE) -C $(BACKEND_DIR) proto dir=$(dir)

proto-all: ## 生成全部 proto 目录的 Go 契约
	$(MAKE) -C $(BACKEND_DIR) proto_all

proto-ts: ## 生成双端前端 TS 契约（输出到各前端 src/proto/，按需安装 ts-proto）
ifeq ($(strip $(PROTO_NAMES)),)
	@$(ECHO) "未找到 $(PROTO_DIR)/*.proto，请检查 dir 参数（当前 dir=$(dir)）"
	@exit 1
else
	cd $(VUE_DIR) && npm install --no-save ts-proto && \
	  protoc --proto_path=$(CURDIR)/$(PROTO_DIR) --proto_path=$(THIRD_PARTY) \
	    --plugin=protoc-gen-ts_proto=$(TS_PROTO_BIN) \
	    --ts_proto_out=$(PROTO_TS_OUT) --ts_proto_opt=$(PROTO_TS_OPT) $(PROTO_NAMES)
	cd $(REACT_DIR) && npm install --no-save ts-proto && \
	  protoc --proto_path=$(CURDIR)/$(PROTO_DIR) --proto_path=$(THIRD_PARTY) \
	    --plugin=protoc-gen-ts_proto=$(TS_PROTO_BIN) \
	    --ts_proto_out=$(PROTO_TS_OUT) --ts_proto_opt=$(PROTO_TS_OPT) $(PROTO_NAMES)
endif

## ---------------------------------------------------------------- 启动

dev-backend: ## 启动后端（3000，需 MySQL 可用）
	cd $(BACKEND_DIR) && go run main.go

dev-vue: ## 启动 vue3x 前端（8000）
	cd $(VUE_DIR) && npm run dev

dev-react: ## 启动 react 前端（8001）
	cd $(REACT_DIR) && npm run dev

dev: ## 同时启动双端前端（8000 / 8001）
	$(MAKE) -j 2 dev-vue dev-react

dev-vue-mock: ## 启动 vue3x 前端（mock 数据，无需后端）
	cd $(VUE_DIR) && npm run dev:mock

dev-react-mock: ## 启动 react 前端（mock 数据，无需后端）
	cd $(REACT_DIR) && npm run dev:mock

dev-mock: ## 同时启动双端前端（mock 数据）
	$(MAKE) -j 2 dev-vue-mock dev-react-mock

## ---------------------------------------------------------------- 打包

build-vue: ## 打包 vue3x（含类型检查，产物 dist/）
	cd $(VUE_DIR) && npm run build

build-react: ## 打包 react（含类型检查，产物 dist/）
	cd $(REACT_DIR) && npm run build

build-frontend: build-vue build-react ## 打包双端前端

build-backend-windows: ## 编译 Windows 后端 → admin-backend/deploy/windows/<env>/admin.exe
	$(MAKE) -C $(BACKEND_DIR) build os=windows env=$(env)

build-backend-linux: ## 编译 Linux 后端 → admin-backend/deploy/linux/<env>/admin
	$(MAKE) -C $(BACKEND_DIR) build os=linux env=$(env)

build-backend: build-backend-linux build-backend-windows ## 编译后端双平台（linux + windows）

build: build-frontend build-backend ## 全量打包：双端前端 + 后端双平台
