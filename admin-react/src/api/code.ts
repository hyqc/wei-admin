/** 业务成功状态码 */
export const SUCCESS = 0;

/** 鉴权相关错误码（与后端 code_proto.ErrorCode 保持一致） */
export const AuthErrorCode = {
  /** 鉴权失败（未登录或令牌已过期） */
  AuthTokenFailed: 200001,
  /** BearerToken 无效 */
  AuthTokenInvalid: 200002,
  /** 登录令牌检查失败 */
  AuthTokenInspectInvalid: 200003,
  /** Token 信息无效 */
  AuthTokenInfoInvalid: 200004,
  /** 没有访问权限 */
  AuthTokenForbidden: 200005,
} as const;

/** 是否令牌失效类错误（需要重新登录） */
export const IsAuthTokenInvalidCode = (code?: number) =>
  code === AuthErrorCode.AuthTokenFailed ||
  code === AuthErrorCode.AuthTokenInvalid ||
  code === AuthErrorCode.AuthTokenInspectInvalid ||
  code === AuthErrorCode.AuthTokenInfoInvalid;

/** 是否“没有访问权限”（仅提示，不退出登录） */
export const IsAuthForbiddenCode = (code?: number) => code === AuthErrorCode.AuthTokenForbidden;

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
