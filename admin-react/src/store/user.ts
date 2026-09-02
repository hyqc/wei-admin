import { create } from 'zustand';
import type { AdminInfo, ReqLogin } from '@/types/admin_account';
import { login, logout as logoutApi, currentAdminInfo } from '@/api/admin/account';
import { GetLoginToken, SetLoginToken, Logout } from '@/utils/common';

/** 用户状态 */
interface UserState {
  /** 当前用户信息 */
  userInfo: AdminInfo | null;
  /** 菜单 map（后端返回，key 为菜单键） */
  menus: Record<string, unknown>;
  /** 权限 map（后端返回，key 为权限键） */
  permissions: Record<string, string>;
  /** 是否已登录（读取本地 token） */
  isLogin: () => boolean;
  /** 设置当前用户信息 */
  setCurrentUser: (data: AdminInfo) => void;
  /** 登录 */
  loginAsync: (params: ReqLogin) => Promise<void>;
  /** 获取当前用户信息（刷新页面后用 token 恢复菜单/权限） */
  fetchUserInfo: () => Promise<AdminInfo | null>;
  /** 退出登录 */
  logoutAsync: () => Promise<void>;
  /** 校验权限键是否存在 */
  hasPermission: (key?: string) => boolean;
}

export const useUserStore = create<UserState>((set, get) => ({
  userInfo: null,
  menus: {},
  permissions: {},

  isLogin: () => !!GetLoginToken(),

  setCurrentUser: (data) =>
    set({
      userInfo: data,
      menus: data.menus ?? {},
      permissions: data.permissions ?? {},
    }),

  loginAsync: async (params) => {
    const res = await login(params);
    const data = res.data;
    SetLoginToken(data.token ?? '', data.expire ?? 0, params.remember ?? false);
    get().setCurrentUser(data);
  },

  fetchUserInfo: async () => {
    const tokenInfo = GetLoginToken();
    if (!tokenInfo?.token) return null;
    const res = await currentAdminInfo(tokenInfo.remember);
    get().setCurrentUser(res.data);
    return res.data;
  },

  logoutAsync: async () => {
    try {
      await logoutApi();
    } catch {
      /* 忽略退出接口异常 */
    }
    Logout();
  },

  hasPermission: (key) => {
    if (!key) return true;
    return Object.prototype.hasOwnProperty.call(get().permissions, key);
  },
}));
