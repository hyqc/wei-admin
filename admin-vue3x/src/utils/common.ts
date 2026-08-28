import { LocalStorageTokenKey, LoginPath } from '@/api/config';

/** 登录 token 存储结构 */
export type TokenType = {
  token: string;
  expire: number;
  remember: boolean;
};

/** 菜单 map 结构 */
export type MenusMapType = {
  [key: string]: {
    access: boolean;
  };
};

/**
 * 获取登录 token
 */
export const GetLoginToken = (): TokenType | undefined => {
  const storeToken = localStorage.getItem(LocalStorageTokenKey);
  try {
    const now = new Date().getTime();
    const obj: TokenType = storeToken ? JSON.parse(storeToken) : {};
    if (obj.token !== undefined && obj.token.length > 0) {
      if (now >= obj.expire) {
        // 过期
        Logout();
        return undefined;
      }
      return obj;
    }
    return undefined;
  } catch (e) {
    console.log(e);
    return undefined;
  }
};

/**
 * 设置登录 token
 * @param token 登录token
 * @param expire 过期时间秒
 * @param remember 是否自动刷新token
 */
export const SetLoginToken = (token: string, expire: number, remember: boolean): void => {
  const tokenInfo: TokenType = { token, expire: expire * 1000, remember };
  localStorage.setItem(LocalStorageTokenKey, JSON.stringify(tokenInfo));
};

/** 当前用户是否已登录 */
export function IsLogin() {
  const tokenInfo: TokenType | undefined = GetLoginToken();
  const token = tokenInfo !== undefined && tokenInfo !== null ? tokenInfo.token : '';
  return token.length > 0;
}

/** 是否是登录页 */
export const IsLoginPage = (): boolean => window.location.pathname === LoginPath;

/**
 * 退出
 */
export const Logout = (): void => {
  localStorage.removeItem(LocalStorageTokenKey);
  const query = new URLSearchParams(window.location.search);
  const { pathname, search } = window.location;
  const redirect = query.get('redirect');
  if (!IsLoginPage() && !redirect) {
    window.location.href = `${LoginPath}?redirect=${encodeURIComponent(pathname + search)}`;
  }
};

/** 首字母转大写 */
export const first2Upcase = (s: string): string => {
  if (!s) return '';
  return s[0].toUpperCase() + s.substring(1).toLowerCase();
};

/** 把路径转为键名（驼峰） */
export function path2UpperCamelCase(path: string) {
  return path
    ?.split('/')
    .filter((name) => name.length > 0)
    .map((name) => name[0].toUpperCase() + name.substring(1))
    .join('');
}
