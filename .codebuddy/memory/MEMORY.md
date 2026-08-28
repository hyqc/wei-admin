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

## 调试经验
- **ant-design-vue 4 关闭 modal 用 `display:none`，不加 `ant-modal-wrap-hidden` 类**。playwright 判断弹窗关闭应查 `w.style.display !== 'none'`
- PowerShell 下 playwright-cli 中文文本选择器有 GBK 编码问题，用 CSS 选择器或 eval 定位
- Mock 登录密码校验为明文 `admin/123456`（SQL 密码是 bcrypt 无法比对）
- antd 表格有 fixed 列时首列(索引0)必须也 fixed，否则 console 警告；操作列宽度按"按钮数 × ~70px"估算防溢出
- 权限树（角色详情/绑定权限）默认只展开第一级（模型层）
