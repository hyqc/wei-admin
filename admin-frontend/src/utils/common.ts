import * as IconMap from '@ant-design/icons';
import { MenuDataItem } from '@umijs/route-utils';
import { parse } from 'query-string';
import { stringify } from 'querystring';
import React from 'react';
import { history } from 'umi';

const iconMap: Map<string, any> = new Set(Object.entries(IconMap))
/**
 * 合并远程菜单和本地菜单
 * @param result
 * @param defaultMenuData
 * @param remoteMenus
 * @param childrenKey
 * @returns
 */
export const HandleRemoteMenuIntoLocal = (
  result: MenuDataItem[],
  defaultMenuData: MenuDataItem[],
  remoteMenus: any,
  childrenKey: string,
): MenuDataItem[] => {
  defaultMenuData?.forEach((item) => {
    let tmpItem: MenuDataItem = { ...item };
    if (tmpItem.key && remoteMenus[tmpItem.key]) {
      //远程返回的菜单中包含，则表面有访问权限
      tmpItem.access = true;
      tmpItem.hideInMenu = !!remoteMenus[tmpItem.key].hideInMenu;
    }
    if (tmpItem.icon !== undefined && tmpItem.icon.length > 0 && iconMap.get(tmpItem.icon) !== undefined) {
      tmpItem.icon = React.createElement(iconMap.get(tmpItem.icon));
    }
    result.push(tmpItem);
    if (item[childrenKey] !== undefined && item[childrenKey].length > 0) {
      tmpItem[childrenKey] = [];
      HandleRemoteMenuIntoLocal(tmpItem[childrenKey], item[childrenKey], remoteMenus, childrenKey);
    }
  });
  return result;
};

/**
 * 用于组件./src/components/PageContainer/Container中判断菜单的访问权限，
 * config.ts中启用access，远程菜单设置access 不生效，暂未找到官方解决方案，暂用在页面容器组件中
 * 根据当前访问的路由来判断是否拥有访问权限
 */
export type MenusMapType = {
  [key: string]: {
    access: boolean;
  };
};

/**
 * 将菜单转成map结构
 * @param result
 * @param menus
 * @param childrenKey
 * @returns
 */
export const HandleMenusToMap = (
  result: MenusMapType,
  menus: MenuDataItem[],
  childrenKey: string,
): MenusMapType => {
  menus.forEach((item) => {
    if (item.path !== undefined && item.path !== '') {
      // 用户判断访问权限
      result[item.path] = { access: item.access };
      if (item[childrenKey] !== undefined && item[childrenKey].length > 0) {
        HandleMenusToMap(result, item[childrenKey], childrenKey);
      }
    }
  });
  return result;
};

/**
 * 当前用户是否已登录
 * @param initialState
 * @returns
 */
export function IsLogin() {
  // eslint-disable-next-line @typescript-eslint/no-use-before-define
  const tokenInfo: TokenType | undefined = GetLoginToken();
  const token = tokenInfo !== undefined && tokenInfo !== null ? tokenInfo.token : '';
  return token.length > 0;
}

/**
 * 登录token存储格式
 */
export type TokenType = {
  token: string;
  expire: number;
  remember: boolean;
};

/**
 * 获取登录token
 * @returns
 */
export const GetLoginToken = (): TokenType | undefined => {
  const storeToken = localStorage.getItem(LocalStorageTokenKey);
  try {
    const now = new Date().getTime();
    const obj: TokenType = storeToken ? JSON.parse(storeToken) : {};
    if (obj.token !== undefined && obj.token.length > 0) {
      if (now >= obj.expire) {
        // 过期
        // eslint-disable-next-line @typescript-eslint/no-use-before-define
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
 * 设置登录token
 */
/**
 *
 * @param token 登录token
 * @param expire 过期时间秒
 * @param remember 是否自动刷新token
 */
export const SetLoginToken = (token: string, expire: number, remember: boolean): void => {
  const tokenInfo: TokenType = { token, expire: expire * 1000, remember };
  console.log('============', tokenInfo)
  // 设置token
  localStorage.setItem(LocalStorageTokenKey, JSON.stringify(tokenInfo));
};

/**
 * 退出
 */
export const Logout = (): void => {
  localStorage.removeItem(LocalStorageTokenKey);
  const query = parse(history.location.search);
  const { search, pathname } = history.location;
  const { redirect } = query;
  // eslint-disable-next-line @typescript-eslint/no-use-before-define
  if (!IsLongPage() && !redirect) {
    history.replace({
      pathname: LoginPath,
      search: stringify({
        redirect: pathname + search,
      }),
    });
  }
};

/**
 * 是否是登录页
 */
export const IsLongPage = (): boolean => {
  return location.pathname === LoginPath;
};

export const first2Upcase = (s: string): string => {
  return s[0].toUpperCase() + s.substring(1).toLowerCase();
};
