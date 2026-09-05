/** 请求成功状态码 */
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

/** 业务错误状态码 */
export const ErrorCode = {
  SUCCESS,
};
