# MEMORY — admin-vue3x 项目长期记忆

## 项目
- `admin-vue3x/`：admin-frontend（React+antd）的 Vue 3.5 复刻版，技术栈 Vue3 + Vite6 + Pinia + ant-design-vue 4 + axios，dev 端口 8000，Mock 插件 `src/mock/`
- 目录 `admin-backend/docs/wei.sql` 是后端权威数据源；`admin-frontend/` 是 React 原版参考

## Mock 数据权威约定（2026-08-28 起）
- `src/mock/data.ts` 数据以 `admin-backend/docs/wei.sql` 为准：
  - 接口 key 为**双冒号**格式（`adminUser::list`），path 无 `/api` 前缀（`/admin/user/list`），共 41 条
  - 权限仅 view/edit/delete 三种类型（无 add），共 15 条；权限-接口关联（apiIds）对应 `admin_permission_api` 表
  - 菜单 8 条（含 id 7/8 添加菜单/编辑菜单隐藏菜单）；角色仅"管理员"；用户 5 条
- 前端按钮权限码：只存在 `AdminXxxView/Edit/Delete`；**新增按钮不做权限控制**（无 Add 权限码，与 React 原版一致）
- `AdminAPIKey` 正则：`/^[a-z][a-zA-Z0-9]*::[a-z][a-zA-Z0-9]*$/g`

## 后端服务
- `admin-backend` 编译 `admin_test.exe`，监听 `127.0.0.1:3000`（`main.go` 同级 `.exe`），路由无 `/api` 前缀；前端 vite(8000) 仅在 `VITE_USE_MOCK=false` 时把 `/api` 代理到 3000
- 登录接口 `POST /admin/account/login`（body: username/password），返回 `data.token`；鉴权头 `Authorization: Bearer xxx`
- 重启后端：`Get-Process admin_test | Stop-Process -Force` 后在 `admin-backend` 目录启动 exe

## proto 修改流程（admin-backend）
- 改 `.proto` 后用 `F:\go\bin\protoc.exe --proto_path=proto/admin_proto --go_out=proto/admin_proto --go_opt=paths=source_relative <file>.proto` 重新生成；本机 protoc-gen-go v1.36.11 与现有 pb.go 一致，可安全覆盖

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
- 但 `admin_role_permission` 绑定表**并未给超管角色全量绑定**（实际只有 6 条），因此涉及"超管拥有的权限"的查询（如 `/admin/role/permissions`）必须走 `dao.H.AdminPermission.FindAdministerPermissions` 返回全部启用权限，不能只查绑定表
- 超管角色不允许绑定权限（前端隐藏按钮），仅可修改描述

## 用户偏好（交互形态）
- 角色详情"绑定权限"、权限"绑定接口"弹窗**保持树形展示**，接口节点用 `名称（唯一键）` 纯文本，不要改成表格/标签+下划线
- 接口"唯一键"输入框只读，只能由接口路径自动生成

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
