# MEMORY — wei-admin 项目长期记忆

## 项目与技术栈
- `admin-vue3x/`：admin-frontend（React+antd 原版参考）的 Vue3.5 复刻版；Vue3 + Vite6 + Pinia + ant-design-vue4 + axios，dev 端口 **8000**，mock 插件 `src/mock/`
- `admin-react/`：vue3x 的 React19 复刻版（功能对齐）；React19 + Vite6 + TS strict + antd5 + zustand + react-router7 + axios，dev 端口 **8001**，复用 `src/mock/`；`tsc --noEmit` 0 错、build 通过
- `admin-backend/docs/wei.sql` = 后端权威数据源（表结构/权限数据，改表/权限后必须同步）；权限数据也需同步双端 `mock/data.ts`

## 权限模型（2026-09-04 方案 A 落地）
- 权限点 = 页面操作；type 固定三类**硬枚举 view/edit/delete**（不参与鉴权）；key = `菜单Key + 操作大写`（特例：AdminUserResetPwd / AdminUserBindRoles / AdminRoleBindPermissions）
- 鉴权 = 用户权限点绑定的 API 集合是否覆盖请求路径（admin_permission_api，每点绑 1~3 个）；前端按钮码 = 权限点 key
- 权威值（docs/wei.sql 与线上一致）：admin_permission **25 行**（2026-09-05 加个人中心 AccountView/Edit = id33/34；type 分布 view/edit/delete）、admin_permission_api 46、admin_api 43（snake_case path）；角色 1 绑 25、角色 5（查看）按需；已删 `uk_permission(menu_id,type)`
- 菜单「权限配置」= 权限点全量定义入口（保存 = `BatchAddPermissions`：唯一键预检 → 逐行 upsert(含 key) → 按 ApiIds 覆盖绑定 → 删未提交行级联清 api/role 绑定；**事务内必须 `query.Use(tx)`**，否则是假事务）；权限管理页已降级为只读检索/审计页（详情+接口反查+去配置）
- type 字段仅用于模板预填/key 生成/TypeText 展示/筛选，后端从不校验其值；type 迁移 8 行 add/reset/bind→edit 已执行（key 未动）
- 菜单「权限配置」面板（PermissionsSaveModal 双端）权限名称列**只读**：名称跟随动作类型自动生成「菜单名+动作类型」（type 一切换即重算，全行统一逻辑）；已有权限点语义名（查看账号/重置密码等）原样回显**不覆盖**（用户明确不做统一标准化）；唯一键仍可手改（改后 auto=false 不再随 type 联动 key）
- 后端 add/edit/delete/bind_apis 接口保留（给非菜单权限点留口子）；菜单删除已级联清权限
- 后端验证方式（HTTP 登录需验证码不便自动化）：临时 `admin-backend/tmpcheck/main.go`（package main + `global.Init()` + 手工 `&gin.Context{}` 调 dao/logic），跑完删除

## 后端服务与环境
- `admin_test.exe`（admin-backend 目录）监听 **127.0.0.1:3000**，路由无 /api 前缀；运行加载 `./config.yaml`（Env 空），`-env dev` 才读 config-dev.yaml
- 重启：`Get-Process admin_test | Stop-Process -Force` → 在 admin-backend 目录启动 exe（日志 `> f:\work\wei-admin\be.log 2>&1`）
- 启动前提 **MySQL 3306 可用**：由 phpStudy MySQL8.0.12 提供（`D:\phpstudy_pro\Extensions\MySQL8.0.12\bin\mysqld.exe`，非服务非 Docker）；`GET /healthz`、`/readyz`（含 DB/Redis）免鉴权
- MySQL 直连：`& 'D:\phpstudy_pro\Extensions\MySQL8.0.12\bin\mysql.exe' -h 127.0.0.1 -u root -p123456 -D wei -e "..."`（**参数必须空格分隔**）；密码在 config-dev.yaml
- 前端 vite proxy（双端，非 mock 时）：`/api` → 3000（**rewrite 去 /api 前缀**）、`/upload` → 3000（**不 rewrite**，后端本地存储相对路径展示用）

## proto / .pb.go 契约（2026-09-05 修正，替换旧结论）
- 机制：标签源写在 `.proto` 字段的 `// @gotags: label:"..." binding:"..."` 注释里；protoc 生成 `.pb.go` 后再跑 `protoc-go-inject-tag -input=xxx.pb.go` 注入 struct tag
- 本机工具齐备：`F:\go\bin` 有 protoc.exe / protoc-gen-go.exe / protoc-go-inject-tag.exe（及 gogofaster/govalidators/micro/go-grpc）
- ✅ **可直接 `protoc --go_out` 重生成**，binding/label 不会丢（前提：生成后补跑 inject-tag）。校验依据：proto 与 pb.go 的 `@gotags` 数量逐文件一致（admin_user 15/15、admin_menu 32/32、admin_upload 2/2）
- 生成命令：`protoc --proto_path=proto/admin_proto --go_out=proto/admin_proto --go_opt=paths=source_relative <file>.proto`，随后 `protoc-go-inject-tag -input=<file>.pb.go`
- ⚠️ 已作废（勿再引用）：早期"npm grpc-tools 无原生二进制 → 只能手改 pb.go、上传模块未用 proto 用 types 契约"。上传模块契约就是 admin_upload.proto/.pb.go
- 细节：protoc-gen-go 所有字段带 `json:",omitempty"`，需空值恒返回的字段声明 `optional string`（→ `*string`，BeanCopy 无法转，须手动赋值/`new(string)`）

## 上传模块（2026-09-05 新增）
- `pkg/storage`：Driver 接口（Name/Put/Delete/URL）+ Register 工厂；已注册 local（默认）/ aliyun / qcloud / s3；新存储只实现+注册即可
- 配置 `config.yaml` 的 `upload` 段（`pkg/config/types.go` 定义）：driver / domain / max_size(MB) / allowed_exts / local{root,prefix} / aliyun / qcloud / s3
- **默认不限制文件类型**：`allowed_exts: []`（代码逻辑 `len>0` 才校验，空即不限；2026-09-05 起 config.yaml 与 docs/config.yaml 均留空，注释说明如需限制再列）
- 目录结构 `{分组}/{年}/{月}/{yyyyMMdd_HHmmss}_{随机8}{ext}`；根目录 ./upload；router 注册 Static(prefix,root) 免鉴权
- 分组规范化为 path 前缀（avatar 上传分组 `/account/`，存储 `account/...`）；**必须先对原始输入校验 `..`、`\`、`//` 再 `path.Clean`**（否则 `../` 被静默规整变合法）
- 表 `admin_upload`（driver/upload_group/object_key/url/original_name/new_name/ext/mime/size/md5/upload_date…）；接口 `POST /admin/upload/{list,upload,delete}`（upload = multipart：file + uploadGroup）；权限点 AdminUploadView/Edit/Delete
- 前端双端 `views/admin/upload/`（列表+上传弹窗，分组必填）+ `api/admin/upload.ts`；上传成功返回 `data.url`（post 返回 ResponseBodyType<T>）

## 菜单 / 路由约定
- 侧边栏以 `router/menu.ts` 的 **localMenuData 为骨架**：后端 menus 只覆盖 hideInMenu/icon，localMenuData 里没有的 key 不显示；顺序 = 数组顺序，「首页」由 SiderMenu 硬编码置顶，想紧跟首页就把菜单放第一位
- 新增菜单四件套：① 后端菜单记录 ② 后端权限点 + 角色绑定（超管走 IsAdministrator 全量，但角色 1 也按约定绑满）③ 前端 localMenuData 双端各加 ④ 路由 + 页面组件；**mock 的 menuTreeData/permissionData 也要同步**，否则侧边栏/按钮消失
- 非必填字段 binding 必须 `omitempty,url` 式（go-playground 的 url/email 对空字符串报错）；ant-design-vue 校验须 `ref=formRef` + 提交 `await formRef.value?.validate()`（只写 rules 不生效）

## 登录 / 验证码 / token
- 登录需图片验证码：`POST /admin/account/captcha` 取 captchaId+图片，提交 captchaId/captchaCode；一次性、忽略大小写、进程内存（重启失效）、默认 4 位/5 分钟
- token：jwt.expire 604800（7 天）；前端存 localStorage `token` = {token, expire(毫秒), remember}
- Mock 登录：admin / 123456（明文比对，SQL 是 bcrypt 无法比对）；mock 服务固定端口启动用 `--mode mock --port 8010 --strictPort`

## 用户偏好（交互形态）
- 接口树展示：接口节点 = `接口名 (路径)`；分组 = `菜单名 (路径前缀)`（半角括号+空格）；实现 `views/admin/permission/components/common.ts` 的 `buildApiTree/buildMenuNameByPath`；接口"唯一键"输入框只读（由路径自动生成）
- 菜单 icon 存 ant 图标名，用 `@/utils/icon` 的 `getAntIcon(name)` 解析（全量 Outlined 系列）；图标选择 `@/components/IconSelect.vue`（v-model:value）
- 个人中心在首页后（localMenuData 第一位 key Account path /account）

## 调试经验
- **execute_command 传 PowerShell 命令时 `$` 变量/`$_` 会被吞** → 复杂逻辑写临时 `.ps1` 再 `powershell -File` 执行；PowerShell 变量不跨命令持久（登录取 token 与后续请求须同命令）
- playwright：浏览器下载被重置 → 用本机 Chrome `chromium.launch({ executablePath: 'C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe' })`；包要 `npm install --no-save playwright@1.63.0` 装在脚本所在目录；脚本后台跑：`start /min cmd /c "cd /d <dir> && chcp 65001 >nul && node x.mjs > log 2>&1"`；每个断言实时 console.log；不用 evaluate 里 element.click()
- antd-vue4 细节：modal 关闭用 display:none（无 ant-modal-wrap-hidden）；descriptions 行 = `.ant-descriptions-row`（无 item class）；表格数据行用 `tr:not(.ant-table-measure-row)`；fixed 列首列必须也 fixed
- 中文/GBK：PowerShell 处理中文 JSON 用 utf8 bytes 或 ASCII 名称；curl.exe 登录 token 用正则提取（ConvertFrom-Json 会拿到空串）
- 后端校验必填先看 proto binding tag；排查前端"rules 不生效"搜 validate( 差集
