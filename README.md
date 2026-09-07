# wei-admin 管理后台

管理后台前后端一体仓库（monorepo）：1 个 Go 后端 + 2 套功能对齐的前端（Vue3 复刻版、React 复刻版，维护时需双端同步改动）。

## 目录与端口

| 目录 | 说明 | 技术栈 | dev 端口 |
|---|---|---|---|
| `admin-backend/` | Go 后端 | Go + Gin + GORM(gorm.io/gen) + JWT | 3000 |
| `admin-vue3x/` | 前端（参考实现） | Vue3.5 + Vite6 + ant-design-vue4 + Pinia | 8000 |
| `admin-react/` | 前端（功能对齐） | React19 + Vite6 + antd5 + zustand + react-router7 | 8001 |
| `Makefile`（根目录） | 统一命令入口：启动/打包/proto/构建 | `make help` 查看全部 | — |

后端路由**无 `/api` 前缀**（如 `POST /admin/menu/list`）；前端通过 vite proxy 的 `/api` 前缀转发并 rewrite 去掉。

## 快速开始

### 1. 环境依赖
- Go 1.21+（模块名 `admin`）、Node 18+
- MySQL 8（必需，安装方式不限：本机安装、Docker、云端实例均可）、Redis（可选，`redis` 配置段为空则跳过初始化）

### 2. 初始化数据库
1. 建库 `wei`（utf8mb4），导入权威建表脚本：`admin-backend/docs/wei.sql`（表结构 + 菜单/权限/角色种子数据）。
2. 修改 `admin-backend/config.yaml` 的 `store.mysql.sources.wei`（账号/密码/端口/库名），与本地一致。

### 3. 启动后端
```bash
make dev-backend        # 等价：cd admin-backend && go run main.go（默认加载 config.yaml）
# 构建产物：make build-backend-windows / make build-backend-linux（输出 admin-backend/deploy/<os>/<env>/）
```
- 校验：`GET /healthz`、`GET /readyz` 返回 ok（readyz 同时探测 MySQL）。
- 换环境：`-env dev` 读 `config-dev.yaml`；`-cfs nacos` 切 Nacos 配置中心（`configSource` 也可配置）。
- 停止/重启：结束占用 3000 端口的进程后重新启动即可（Linux/macOS：`lsof -ti:3000 | xargs kill`；Windows：`netstat -ano | findstr :3000` 查 PID 后结束）。

### 4. 启动前端
```bash
make init-frontend  # 首次：安装双端依赖（等价各端 npm install）
make dev            # 同时启动双端前端（vue3x 8000 / react 8001）
make dev-vue        # 只启动 vue3x；make dev-react 只启动 react（8001）
make dev-mock       # 双端 mock 模式（无需后端）
# 等价原始写法：cd admin-vue3x && npm run dev | npm run dev:mock
```
- 登录：mock 模式账号 `admin / 123456`；真实后端走登录页图片验证码，账号密码以库中 `admin_user` 为准。
- 根目录 `Makefile` 是统一命令入口，全部命令见 `make help`。

## 配置说明

### 后端 `admin-backend/config.yaml`
| 配置段 | 关键项 | 说明 |
|---|---|---|
| `server` | `port` `debug` `pprof` | 监听端口；debug 记录验证码明文等；pprof 生产勿开 |
| `logger` | `level` `stdout` `json` | 日志级别/双写 |
| `jwt` | `expire` `ignores` `public` `private` | token 秒数；免鉴权路由（method + paths，支持 `/xxx/*` 前缀）；RSA 公/私钥为 **base64(PEM)**，生成方式见下文 |
| `store.mysql` | `sources.wei` | 数据源（默认库名 `wei`）；未配置的类型不会初始化 |
| `redis` | `sources` | 为空则不启用 |
| `captcha` | — | 登录验证码（进程内存、重启失效） |
| `upload` | `driver` `domain` `max_size` `allowed_exts` `local` `aliyun/qcloud/s3` | 文件上传：驱动切换（local/aliyun/qcloud/s3）、返回链接域名前缀（本地单机填 `http://127.0.0.1:3000`）、单文件 MB 上限（0 不限）、扩展名白名单（留空=不限制）、本地根目录/静态前缀 |

生成 JWT 密钥对（base64）：
```bash
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in private.pem -out public.pem
# 将两个 PEM 文件内容分别 base64 后填入 jwt.public / jwt.private
```
> 密钥不要提交到仓库。Nacos 模式：configSource=nacos 并填 `nacos` 段。

### 前端（双端一致）
- `.env` / `.env.mock`：`VITE_USE_MOCK=true|false`（`npm run dev:mock` 读 mock 环境）。
- `vite.config.ts` 的 `server.proxy`：`/api → http://127.0.0.1:3000`（**rewrite 去掉 /api**）、`/upload → 3000`（**不 rewrite**）。若后端换端口，两处 target 同步改。

## 架构与实现方案

### 后端分层
请求链路：`app/router`（注册路由，见 `app/router/admin.go`，`/admin` 分组下按模块挂 controller 方法）→ `app/middleware`（global 通用中间件、auth 鉴权）→ `app/admin/controller`（解析/校验/响应）→ `app/admin/logic`（业务与事务）→ `app/admin/dao`（数据访问）→ `gen/model`、`gen/query`（gorm.io/gen 生成的表模型与查询器）。

- 数据访问：**写 dao 不直接拼 SQL**——`gen/model/*.gen.go` 定义表结构，`gen/query/*.gen.go` 提供类型安全查询器；`app/admin/dao/` 内手写业务查询方法并注册到容器 `dao.H`（`dao/common.go` 的 `AdminDao`）。
- 消息契约：请求/响应结构定义在 `proto/admin_proto/*.proto`，字段以 `// @gotags: label:"..." binding:"..."` 注释声明参数标签；`make proto dir=admin_proto` 生成 `.pb.go` 并注入 struct tag。校验规则以注入后的 tag 为准。
- 启动装配：`global` 包按 `initConfigCall` 顺序初始化 config/logger/database/jwt/server/redis/captcha，并持有全局单例（`global.AppDB`、`global.AppAuth` 等）。
- 存储：`pkg/storage` 提供 `Driver` 接口（Name/Put/Delete/URL）与注册工厂，local/aliyun/qcloud/s3 四种驱动可选，新增存储只实现接口并注册即可（见「新增功能 D」）。

### 权限模型（核心，理解后再做菜单/接口/权限相关功能）
```
admin_user —— admin_user_role —— admin_role —— admin_role_permission —— admin_permission —— admin_permission_api —— admin_api
  账号          账号-角色             角色          角色-权限点            权限点（=页面操作）        权限点-接口              接口(记录 path)
```
- **鉴权** = 中间件校验 JWT 后，取该账号权限点**绑定的全部 API 路径**，判断是否覆盖当前请求路径（`IsAdminCanAccessPath`）。超管（is_admin=1）全放行。
- **权限点**（`admin_permission`）= 页面上一个可授权操作，`key` 与前端按钮权限码一一对应（命名 `菜单Key+操作大写`，如 `AdminUserEdit`；特例：`AdminUserResetPwd`/`AdminUserBindRoles`/`AdminRoleBindPermissions`）。
- **type 固定三类硬枚举**：`view`（查看/读）/ `edit`（编辑/写，含新增、绑定、重置等一切写操作）/ `delete`（删除）。type 仅用于模板预填、key 生成与展示，后端不校验其值。
- 权限点全量定义入口是**菜单管理 → 权限配置**（保存 = 整组 upsert + 按所选接口覆盖绑定 + 删除未提交行）；「权限管理」页是只读审计页。
- 权威数据三处同步：`admin-backend/docs/wei.sql` ↔ 双端 `src/mock/data.ts`（menuTreeData/permissionData）↔ 运行库。**改权限/菜单后必须同步，否则 mock 下侧边栏/按钮消失。**

### 前端约定（双端对称）
- 侧边栏以 `src/router/menu.ts` 的 `localMenuData` 为骨架：本地没有的菜单 key 不渲染；后端仅覆盖 `hideInMenu`/`icon`。顺序 = 数组顺序（「首页」硬编码置顶）。
- 按钮权限码：vue 用 `v-permission="'AdminUserEdit'"`（`directives/permission.ts`，基于 `userStore.hasPermission(key)`）；react 用 `<Authorization permission="AdminUserEdit">`。key 即后端权限点 key。
- API 封装：`src/api/admin/*.ts`（按模块，对应后端 `/admin/xxx`）；mock 时由 `src/mock/` 拦截（`mock/admin/` 下同名文件 + `mock/data.ts` 业务数据），请求经 `src/api/request.ts`。

## 新手指南：如何新增功能

### A. 新增一个菜单页面（只读/查看型）
1. **后端建菜单**：在库 `admin_menu` 插记录（key/name/path/parent 结构参考已有菜单），并同步到 `docs/wei.sql`。
2. **建权限点**：接口管理/权限相关入口或直接 SQL 加 `admin_permission`（菜单 id + type=view + key=`菜单Key+View` + name），超管无需绑定即可用，但要给普通角色绑（`admin_role_permission`）。
3. **前端菜单**：双端 `localMenuData` 各加一项（key/path/name/icon，icon 用 ant 图标名如 `SettingOutlined`）。
4. **路由 + 页面**：`src/router` 加路由（vue 懒加载），`src/views/` 下建页面组件。
5. **mock 同步**：双端 `mock/data.ts` 的 `menuTreeData` 与 `permissionData` 补对应项（否则 dev:mock 下不显示）。
6. 重启后端、刷新前端验证。

### B. 在既有页面上加一个受控操作（按钮）
1. 后端加接口：controller 方法 → 路由注册到 `app/router/admin.go`（如 `menu.POST("/xxx", ...)`）。
2. **接口登记**：在「接口管理」页新增 `admin_api` 记录（path 必须与注册的路由一致，如 `/admin/menu/xxx`，绑定到相应权限点）。
3. 前端按钮包权限码（vue `v-permission` / react `<Authorization>`）；写操作统一挂到该菜单的 **edit** 权限点（在菜单权限配置面板勾选接口）。
4. 双端 mock 的 `mock/admin/*.ts` 同步处理该 URL。

### C. 新增一个业务表（完整 CRUD 模块）
1. `docs/wei.sql` 建表。
2. 生成 `gen/model`、`gen/query` 的 `.gen.go`（参照同目录既有文件，或 gorm.io/gen 工具生成）。
3. `proto/admin_proto` 加消息定义（字段写 `// @gotags` 注释）→ `make proto dir=admin_proto`。
4. `app/admin/dao/` 写 dao（interface + 实现，`query2.Xxx` 取查询器），在 `dao/common.go` 的 `AdminDao` 容器登记。
5. `app/admin/logic/` 写业务（**事务内用 `query.Use(tx)` 传事务上下文，勿用全局 query**）→ `controller/` 写 handler。
6. `app/router/admin.go` 注册路由组。
7. 前端三件套：`api/admin/*.ts`、`mock/admin/*.ts` + `mock/data.ts`、页面组件；受控操作见 B。
8. 走「接口管理登记 + 菜单权限配置绑定」完成鉴权闭环。

### D. 新增对象存储
在 `pkg/storage` 实现 `Driver` 接口（Name/Put/Delete/URL）并 `Register`，在 `config.yaml` 的 `upload` 段补对应驱动配置，切 `driver` 即生效（参考 aliyun/qcloud/s3 实现）。

### E. 新增接口必须同步的三处（防呆）
后端路由、`admin_api` 记录（path 精确匹配）、双端 mock（若走 mock）。前两点漏了会导致接口不受权限管控或 403。

## 常用命令

全部在**仓库根目录**执行，详见 `make help`。

| 场景 | 命令 |
|---|---|
| 安装依赖 | `make init`（后端 proto 工具 + 双端前端依赖）；只装前端：`make init-frontend` |
| 启动后端（3000） | `make dev-backend` |
| 启动双端前端 | `make dev`（vue3x 8000 + react 8001 并发）；单端 `make dev-vue` / `make dev-react` |
| 启动双端（mock） | `make dev-mock`；单端 `make dev-vue-mock` / `make dev-react-mock` |
| 打包双端前端 | `make build-frontend`（含类型检查）；单端 `make build-vue` / `make build-react` |
| 编译后端 Windows | `make build-backend-windows env=prod` → `admin-backend/deploy/windows/prod/admin.exe` |
| 编译后端 Linux | `make build-backend-linux env=prod` → `admin-backend/deploy/linux/prod/admin` |
| 编译后端双平台 / 全量 | `make build-backend` / `make build`（双端前端 + 后端双平台） |
| 生成 Go 契约 | `make proto dir=admin_proto`（需 protoc + protoc-gen-go + protoc-go-inject-tag，安装见 `make init`） |
| 生成前端 TS 契约 | `make proto-ts`（ts-proto，输出到各前端 `src/proto/`，按需 `npm install --no-save ts-proto`） |
| 直连 MySQL 调试 | 本机：`mysql -h 127.0.0.1 -u root -p -D wei -e "SELECT ..."`；Docker：`docker exec -i <容器名> mysql -uroot -p<密码> wei -e "SELECT ..."` |

## 约定与踩坑速查
- 非必填字段的 binding 必须是 `omitempty,xxx`（go-playground 的 url/email 对空字符串直接报错）。
- 前端 antd-vue4 表单：只有 `:rules` 不生效，须 `ref="formRef"` 且提交前 `formRef.value?.validate()`（react 端用 `validateFields`/`onFinish`）。
- 上传目录结构 `{分组}/{年}/{月}/{yyyyMMdd_HHmmss}_{随机8}{ext}`；分组校验必须先校验原始输入中的 `..`、`\`、`//` 再 `path.Clean`。
- 本仓库附带运行日志（`*.log`）、截图等杂项文件在根目录，可忽略；`docs/wei.sql` 是唯一权威 schema/seed 来源。
