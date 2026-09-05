# MEMORY — admin-vue3x 项目长期记忆

## 项目
- `admin-vue3x/`：admin-frontend（React+antd）的 Vue 3.5 复刻版，技术栈 Vue3 + Vite6 + Pinia + ant-design-vue 4 + axios，dev 端口 8000，Mock 插件 `src/mock/`
- `admin-react/`（2026-09-02 新建）：admin-vue3x 的 React 19 复刻版（用户要求 React 最新版重写，功能与 vue3x 完全对齐），技术栈 React 19 + Vite 6 + TS strict + antd 5 + zustand + react-router 7 + axios，dev 端口 8001，直接复用 `src/mock/` 自研 mock 插件；`tsc --noEmit` 0 错误、build 通过（manualChunks 分 react/antd 包）
- 目录 `admin-backend/docs/wei.sql` 是后端权威数据源；`admin-frontend/` 是 React 原版参考


## 权限模型（2026-09-04 方案 A 落地，替代旧"每菜单 view/edit/delete 三权限"）
- **权限点 = 页面操作**；动作类型**固定三类、不可新增、不可编辑**：`view` 查看（读）/ `edit` 编辑（写，新增、重置、绑定等一切写操作都归入此类）/ `delete` 删除。2026-09-04 由 7 类（view/add/edit/enable/bind/reset/delete）收敛而来，用户定调"所有操作只有这三类"
- 权限点 key = `菜单Key + 操作首字母大写`（如 AdminUserAdd / AdminRoleAdd）；语义特例：AdminUserResetPwd、AdminUserBindRoles、AdminRoleBindPermissions
- 每个权限点绑定该操作所需的 1~3 个 API（`admin_permission_api`）；后端鉴权 = 用户所持权限点绑定的 API 集合是否覆盖请求路径（`IsAdminCanAccessPath`）
- DB 权威值（docs/wei.sql 与线上 wei 库一致，2026-09-04 核过）：admin_permission **23 行**（类型分布 **view5 / edit13 / delete5**；2026-09-04 已把 8 行 add+reset+bind 迁移为 edit，**key 一律不动**）、admin_permission_api **46 行**、admin_api **43 行**（snake_case path；含 id43 `adminUser::editPwd`、id44 `adminMenu::show`，无 id42）；已删 `uk_permission(menu_id,type)` 唯一索引；角色仅 1(超管)+5(测试)、各绑 23 权限
- 菜单"权限配置"编辑器：任意行增删、key/name 可编辑，**类型只能从 view/edit/delete 三类中选（Select，不可输入）**，默认模板预填这三类；保存 = `BatchAddPermissions` 全量同步（按 id/key upsert + 删除该菜单未提交的旧权限点 + 级联清 admin_permission_api / admin_role_permission）
- 前端 vue3x + react 按钮码 = 权限点 key（与 React 原版旧 3 码逻辑不再一致，新增按钮也做权限控制）
- Mock（vue3x + react `src/mock/`）已同步：apiData 43 行、permissionData 23 行、typeText 三类映射、add_menu_permissions 全量同步语义；react 根目录需有 `.env.mock`（VITE_USE_MOCK=true）dev:mock 才真正启用
- `AdminAPIKey` 正则：`/^[a-z][a-zA-Z0-9]*::[a-z][a-zA-Z0-9]*$/g`
- **权限管理页（`views/admin/permission`）定位（2026-09-04 评估）**：权限点 CRUD 已与菜单"权限配置"重复，但**"绑定接口"（bind_apis/unbind_api）是它独有的、且是鉴权真正生效的地方**，另有跨菜单全局检索价值 → 不能删页，建议收敛为"查询 + 启停 + 绑定接口"
- **2026-09-04 权限管理页收敛：方案 2 已落地**（不考虑工作量的最优解）。入口收敛到「菜单管理 → 权限配置」：权限点定义 + 接口绑定在一次保存内完成，消除"有权限点无接口→按钮亮但 403"的中间态；权限管理页降级为**只读检索/审计页**（新增"按接口反查权限点"）
- 后端实现要点：`dao.PermissionSyncItem{Permission *model.AdminPermission; ApiIds []int32}`；`BatchAddPermissions` 改为**先唯一键校验 → 再逐行 upsert（显式列 Updates，含 key）→ 再按 ApiIds 全量覆盖绑定（nil 表示不改）→ 最后删未提交行并级联清 api/role 绑定**；事务内必须用 `query.Use(tx)`（原代码在 Transaction 里仍用全局 `query.XXX.WithContext(ctx)`，事务是假的）
- `admin_menu.Delete` 已补级联：删菜单 → 清 `admin_permission_api` + `admin_role_permission` + `admin_permission` → 删菜单（原实现只删权限点，留下孤儿绑定）
- 接口反查：`ReqAdminPermissionList.ApiId` / `ReqFrontAdminPermissionList.ApiId`（dao.handleListReq 用 Join AdminPermissionAPI）；`/admin/menu/permissions` 的 `MenuPermissionItem.ApiIds` 返回已绑接口
- **决策（用户未反对默认）**：后端 `add/edit/delete/bind_apis` 接口**保留**（给"非菜单权限点"留口子），仅前端入口收敛；前端已删 `AddPermissionModal` / `EditPermissionModal` / `BindApisModal`（双端共 6 个文件），权限管理页只剩「详情 + 去配置（跳 /admin/menu）」
- ✅ 已修复：菜单权限配置改 key 不生效（原来是 `clause.OnConflict` + `DoUpdates` 不含 key，且库上只有 `PRIMARY(id)`+`uk_key`）；现在改 key 走 UPDATE 并前置校验 `uk_key` 冲突，冲突返回"权限键名已存在"且整批回滚
- 验证方式（无验证码无法走 HTTP 登录）：临时 `tmpcheck/main.go`（package main + `global.Init()` + 手工 `&gin.Context{Request: req}`）直接调 dao 验证，跑完删除；`serves` 包不能放 `_test.go`（该包自定义 flag 与 `-test.*` 冲突）
- 权限管理页自身权限点存在：id16-19（menu_id=5）AdminPermissionView/Add/Edit/Delete
- 选型理由（备查）：方案 2 胜出的唯一标准是"一个权限点的完整定义应在一次操作内完成"；方案 1 把定义劈成两个页面会产生"有壳无魂"中间态，方案 3 让权限点脱离菜单等于退回旧模型并失去全量同步收益
- **权限 type 字段定位（2026-09-04 最终态：固定三类硬枚举）**：type **不参与鉴权**（`IsAdminCanAccessPath` 只看权限点绑定的 API；前端按钮只认 key）。它只用于 ① 默认模板预填 view/edit/delete ② 生成默认 key（`菜单Key + Type首字母大写`，`typeKeySuffix` 过滤非字母数字）③ `TypeText` 展示 ④ 权限列表筛选。后端**从不校验** type 合法性（`ErrorCode_AdminPermissionTypeInvalid` 600005 定义了却从未 return），实际约束来自前端下拉选项
- ⚠️ **已修 bug**：type 展示后端两处不一致 —— `common.GetPermissionTypeName`（菜单权限配置用）未知类型返回原值，`dao.GetAdminPermissionTypeText`（列表/详情用）未知类型返回**空字符串** → 自定义 type 在权限列表里显示空白。现后者也回退原值
- **2026-09-04 type 收敛落地步骤（备查）**：① 先按"软枚举"改过一版（`GetAdminPermissionTypeText` 回退原值 + 控件改 AutoComplete 可自定义）② 用户随后定调"所有操作只有查看/编辑/删除三类"→ **最终改为固定三类硬枚举**：前端 `DEFAULT_PERMISSION_TYPES` 3 项 + 控件回到 Select（不可输入）、`DEFAULT_PERMISSION_TEMPLATE_TYPES` = 全部 3 项；后端 `app/common/permission_type.go` 只留 View/Edit/Delete（`AdminPermissionEnumItems`/`DefaultItems`/`Map` 均 3 项）、`dao.AdminPermissionTypeTextMap` 只留 all/view/edit/delete；`handleKey` 的 `typeKeySuffix` 保留（防中文/空类型）
- **数据迁移已执行**：`UPDATE admin_permission SET type='edit' WHERE type NOT IN ('view','edit','delete')`（8 行：add×5、reset×1、bind×2 → edit），已同步 `docs/wei.sql` 与双端 `mock/data.ts` 的 type/typeText。**权限点 key 一个都没改**（`AdminUserAdd`、`AdminUserResetPwd` 等仍是原值），所以前端按钮码、角色绑定、API 绑定全部不受影响
- type 备选方案（未采纳）：**B 彻底去 type**（DB 去列 + 迁移 23 行）；**C 建 `admin_permission_type` 字典表 + 管理页**（性价比低）

## 后端服务
- `admin-backend` 编译 `admin_test.exe`，监听 `127.0.0.1:3000`（`main.go` 同级 `.exe`），路由无 `/api` 前缀；前端 vite(8000) 仅在 `VITE_USE_MOCK=false` 时把 `/api` 代理到 3000
- 登录接口 `POST /admin/account/login`（body: username/password），返回 `data.token`；鉴权头 `Authorization: Bearer xxx`
- 重启后端：`Get-Process admin_test | Stop-Process -Force` 后在 `admin-backend` 目录启动 exe
- **启动前提：MySQL 3306 必须可用**，否则进程启动即退出（`init config error: store: mysql 数据源 "wei" 初始化失败`）。本机 3306 由 **phpStudy MySQL 8.0.12** 提供（`D:\phpstudy_pro\Extensions\MySQL8.0.12\bin\mysqld.exe`），**不是** Windows 服务、也不是 Docker 容器（Docker 的 `mysql:8.0.30` 容器与它端口冲突，起不来）
- 健康检查：`GET /healthz`（存活）、`GET /readyz`（含 MySQL/Redis 就绪检查），免鉴权
- 本机 MySQL 直连：`& 'D:\phpstudy_pro\Extensions\MySQL8.0.12\bin\mysql.exe' -h 127.0.0.1 -u root -p123456 -D wei -e "..."`（**参数必须空格分隔**，`-h127.0.0.1` 会被解析成 host "127" 报错）；密码见 `admin-backend/config-dev.yaml` 的 `mysql.sources.wei`
- 三端端口：后端 3000、Vue 前端 8000、React 前端 8001；默认 dev 脚本均连接真实后端（非 mock）

## proto 修改流程（admin-backend）
- **本机没有 `protoc-gen-gotags` 插件**（只有 protoc-gen-go / go-grpc / gogofaster / govalidators / micro）
- ⚠️ 因此**不要用 `protoc --go_out` 重新生成 .pb.go**：生成的 struct 会丢失 `binding:"..."`、`label:"..."` 标签（字段校验规则全部失效），@gotags 只会留在注释里
- 正确做法：改完 `.proto` 后，**手动编辑对应 `.pb.go` 的 struct tag**，格式 `protobuf:"..." json:"..." binding:"..." label:"..."`（若确实重新生成过，务必逐字段把 tag 补回）

## K8s / 多副本部署（2026-09-01 已改造 1-4 点）
- 验证码：改用 `mojocn/base64Captcha`，存储可切换 `captcha.store: memory`（单机默认）/`redis`（集群，需配 `redis.addr`）；**store=redis 但 Redis 不可用时启动失败**；按场景扩展在 `captcha.scenes` 加配置 + 定义 `captcha.Scene`
- 探针：`GET /healthz`（存活）、`GET /readyz`（就绪，含 MySQL、Redis 检查），免鉴权
- 日志：`logger.stdout: true` 时同时写文件与 stdout
- 服务：Shutdown 超时 15s、监听失败 os.Exit(1)、退出关 Redis、`pprof` 默认关闭
- 配置源：命令行 `-cfs=nacos` > 配置 `configSource`（默认 file）
- **未处理（用户暂缓）**：Nacos 热更反射替换 AppConfig 的数据竞争、登录 login_total 读-改-写、密钥硬编码、启动探测 8.8.8.8 取 IP
- P1：登录 `login_total`/`last_login_ip` 是读-改-写无锁（`app/admin/logic/admin_user.go`）；Nacos 热更用反射整体替换 `AppConfig` 有数据竞争
- P2 运维：无 healthz/readyz 探针；日志仅写本地文件不写 stdout；Shutdown 无超时、启动失败不退出；pprof 无鉴权暴露；密钥硬编码在 config.yaml；启动依赖探测 8.8.8.8 取 IP，失败即 os.Exit(1)
- 无状态可水平扩展：JWT 无状态、无本地缓存、无定时任务；redsync / go-micro 注册发现 / gRPC 均未实际使用；go-redis 已引入但未初始化
- 另：前端调用 `/admin/common/upload`，后端本仓库**无实现**

## 登录与验证码
- 登录必须带图片验证码：`POST /admin/account/captcha`（免鉴权）取 `captchaId` + 图片，登录时提交 `captchaId`/`captchaCode`；验证码一次性、忽略大小写、默认 4 位/5 分钟，配置在 `config.yaml` 的 `captcha`
- 验证码存储为**进程内存**（无 Redis），重启后失效；多实例部署需改为 Redis
- **2026-09-01 改用 `mojocn/base64Captcha`**，为易辨识已主动降低干扰：`noiseCount=1`、`showLineOptions=0`（不画干扰线），仅保留少量噪点；后续如需提升安全可开启 `OptionShowHollowLine`
- token：登录签发用 `config.yaml` 的 `jwt.expire`（已设 604800 = 7 天），刷新 token 用 `constant.AdminTokenExpireSeconds`（7 天）
- 前端 token 存 localStorage，key 为 `token`（结构 `{token, expire(毫秒), remember}`）

## protobuf 响应字段恒返回技巧
- protoc-gen-go 给所有字段加 `json:",omitempty"`，后端用 `encoding/json`（`code.JSON`）序列化时空字符串字段会整个消失
- 若要求某字段无值也返回（如菜单 icon），把字段声明为 `optional string` → Go 类型变 `*string`，需显式赋值（BeanCopy 无法转换 string→*string，须手动 `data.Icon = &item.Icon`），空值用 `new(string)`
- 重新生成：`protoc --proto_path=proto/admin_proto --go_out=proto/admin_proto --go_opt=paths=source_relative <file>.proto`

## 图标约定
- 菜单 `icon` 字段存 ant 图标名（如 `SettingOutlined`）；侧边栏 `SiderMenu.vue` 与菜单详情均通过 `@/utils/icon` 的 `getAntIcon(name)` 解析（全量 Outlined 系列）
- 图标选择用公共组件 `@/components/IconSelect.vue`（`v-model:value`，支持搜索）
- 注意：`import * as AntIcons from '@ant-design/icons-vue'` 为全量引入，图标名列表按 `/Outlined$/` 过滤即可，不要用 `typeof === 'object'` 判断图标组件

## 超管语义约定
- 超管账号（adminId=1）与超管角色（roleId=1）自动拥有全部权限，判断用 `constant.IsAdministrator` / `constant.IsAdministratorRole`
- 2026-09-04 起 `admin_role_permission` 已给角色 1 与 5 **各绑满 23 条**；但"超管拥有的权限"类查询（如 `/admin/role/permissions`）仍走 `dao.H.AdminPermission.FindAdministerPermissions` 返回全部启用权限（超管自动全量，不依赖绑定表完整性）
- 超管角色不允许绑定权限（前端隐藏按钮），仅可修改描述

## 用户偏好（交互形态）
- **接口树展示格式（2026-09-04 用户指定，替代旧的"名称（唯一键）"约定）**：保持树形（不要改成表格/标签+下划线）；接口节点 = `接口名 (路径)`，如 `接口列表 (/admin/api/list)`；分组 = `菜单名 (路径前缀)`，如 `接口管理 (/admin/api)`（分组按 path 前两段生成，菜单名由 `/admin/menu/all` 的 path→name 映射补全）。半角括号+空格，不用全角
- 接口"唯一键"输入框只读，只能由接口路径自动生成
- 实现位置（双端一致）：`views/admin/permission/components/common.ts` 的 `buildApiTree(list, menuNameByPath)` + `buildMenuNameByPath(menus)`；调用方是菜单权限配置的 `PermissionApiBindModal`；`admin_menu.path` 与接口 path 前两段天然对齐（/admin/api → 接口管理）

## 调试经验
- playwright 脚本里 **不要用 `element.click()`（evaluate 内原生调用不触发 Vue 事件）**，改用 `page.locator(...).click()`
- antd-vue 4：drawer 没有 `.ant-drawer-wrap`（modal 才有）；drawer 是 fixed 定位，`offsetParent` 恒为 null，判断可见用 `getBoundingClientRect().width > 0`
- antd-vue 4 bordered `a-descriptions` 用表格布局：行是 `.ant-descriptions-row`（th 标签 + td 内容），没有 `.ant-descriptions-item`
- 表格占位行 class 为 `ant-table-measure-row`，取数据行用 `tr:not(.ant-table-measure-row)`
- **PowerShell 变量在 `execute_command` 之间不持久**：登录取 token 与后续请求必须写在同一条命令内，否则 header 为空 → 后端返回 200003
- PowerShell 处理中文 JSON 易出编码问题：请求体中文用 `\uXXXX` 转义 + `[System.Text.Encoding]::UTF8.GetBytes`，或测试改用 ASCII 名称
- `Invoke-RestMethod` 可用（-Headers 传 hashtable）；`curl.exe` 登录响应经 `ConvertFrom-Json` 取 token 会拿到空字符串，可用正则提取
- **ant-design-vue 4 关闭 modal 用 `display:none`，不加 `ant-modal-wrap-hidden` 类**。playwright 判断弹窗关闭应查 `w.style.display !== 'none'`
- PowerShell 下 playwright-cli 中文文本选择器有 GBK 编码问题，用 CSS 选择器或 eval 定位
- Mock 登录密码校验为明文 `admin/123456`（SQL 密码是 bcrypt 无法比对）
- antd 表格有 fixed 列时首列(索引0)必须也 fixed，否则 console 警告；操作列宽度按"按钮数 × ~70px"估算防溢出
- 权限树（角色详情/绑定权限）默认只展开第一级（模型层）
