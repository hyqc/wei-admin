/** 业务成功状态码 */
export const SUCCESS = 0;

/** 业务错误码 */
export const ErrorCode = {
  /** 请求失败 */
  Error: 1,
  /** 请求参数无效 */
  RequestParamsInvalid: 300003,
  /** 账号或密码错误 */
  AdminAccountPasswordInvalid: 400001,
  /** 账号不存在或已被删除 */
  AdminAccountNotExist: 400002,
  /** 超管账号不允许该操作 */
  AdminSuperAccountNotAllow: 400005,
  /** 验证码错误或已失效 */
  AdminCaptchaInvalid: 400006,
  /** 接口名称已存在 */
  AdminApiNameExist: 500001,
  /** 接口路径已存在 */
  AdminApiPathExist: 500002,
  /** 接口键名已存在 */
  AdminApiKeyExist: 500003,
  /** 权限键名已存在 */
  AdminPermissionKeyExist: 600001,
  /** 该菜单下权限已存在 */
  AdminPermissionExist: 600002,
  /** 菜单名称已存在 */
  AdminMenuNameExist: 700002,
  /** 菜单键名已存在 */
  AdminMenuKeyExist: 700003,
  /** 角色名称已存在 */
  AdminRoleNameExist: 800002,
} as const;
