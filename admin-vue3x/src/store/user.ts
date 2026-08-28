import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { AdminInfo } from '@/types/admin_account';
import type { ReqLogin } from '@/types/admin_account';
import { login, logout as logoutApi, currentAdminInfo } from '@/api/admin/account';
import { GetLoginToken, SetLoginToken, Logout } from '@/utils/common';

/** 用户状态 */
export const useUserStore = defineStore('user', () => {
  /** 当前用户信息 */
  const userInfo = ref<AdminInfo | null>(null);
  /** 菜单 map（后端返回，key 为菜单键） */
  const menus = ref<Record<string, any>>({});
  /** 权限 map（后端返回，key 为权限键） */
  const permissions = ref<Record<string, string>>({});

  const isLogin = computed(() => !!GetLoginToken());

  /** 设置当前用户信息 */
  function setCurrentUser(data: AdminInfo) {
    userInfo.value = data;
    menus.value = data.menus ?? {};
    permissions.value = data.permissions ?? {};
  }

  /** 登录 */
  async function loginAsync(params: ReqLogin) {
    const res = await login(params);
    const data = res.data;
    SetLoginToken(data.token ?? '', data.expire ?? 0, params.remember ?? false);
    setCurrentUser(data);
    return res;
  }

  /** 获取当前用户信息（刷新页面后用 token 恢复菜单/权限） */
  async function fetchUserInfo() {
    const tokenInfo = GetLoginToken();
    if (!tokenInfo?.token) return null;
    const res = await currentAdminInfo(tokenInfo.remember);
    setCurrentUser(res.data);
    return res.data;
  }

  /** 退出登录 */
  async function logoutAsync() {
    try {
      await logoutApi();
    } catch {
      /* 忽略退出接口异常 */
    }
    Logout();
  }

  /** 校验权限键是否存在 */
  function hasPermission(key?: string) {
    if (!key) return true;
    return Object.prototype.hasOwnProperty.call(permissions.value, key);
  }

  return {
    userInfo,
    menus,
    permissions,
    isLogin,
    setCurrentUser,
    loginAsync,
    fetchUserInfo,
    logoutAsync,
    hasPermission,
  };
});
