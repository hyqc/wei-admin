[toc]

# 默认编程规则

### 接口路径定义规则

1. 统一在 src/services/apis/admin/api.ts 中定义
2. 按照一个模块定义，以管理员管理页面示例

```javascript
export const APIAdminUsers = {
  list: {
    url: '/admin/user/list',
    method: POST,
  },
  add: {
    url: '/admin/user/add',
    method: POST,
  },
  delete: {
    url: '/admin/user/delete',
    method: POST,
  },
  edit: {
    url: '/admin/user/edit',
    method: POST,
  },
  enable: {
    url: '/admin/user/enable',
    method: POST,
  },
};
```

3. 导出的接口模块名称必须以 API 开头，大驼峰法命名，以管理员管理页面示例：

```javascript
export const APIAdminUsers = {
  list: {
    url: '/admin/user/list',
    method: POST,
  },
  add: {
    url: '/admin/user/add',
    method: POST,
  },
  delete: {
    url: '/admin/user/delete',
    method: POST,
  },
  edit: {
    url: '/admin/user/edit',
    method: POST,
  },
  enable: {
    url: '/admin/user/enable',
    method: POST,
  },
};
```

### 接口请求方法的请求参数类型

1. 请求参数以 **Request** 开头，以 **ParamsType** 结束。示例：

```javascript
export type RequestLoginParamsType = {
  username: string,
  password: string,
};
```

看名称就可以知道是登录请求的请求参数类型

### 接口请求方法的返回参数 data 键部分类型

1. 请求参数以 **Response** 开头，以 **DataType** 结束。示例：

```javascript
export type ResponseLoginDataType = {
  token: string,
  expire: number,
  admin: any,
  menus: any,
};
```

看名称就可以知道是登录请求的返回结果的 **data** 键的类型
