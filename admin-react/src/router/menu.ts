import { HomePath } from '@/api/config';

/** 本地菜单配置（与后端返回的菜单 key 对应） */
export interface MenuConfigItem {
  key: string;
  path: string;
  name: string;
  icon?: string;
  hideInMenu?: boolean;
  children?: MenuConfigItem[];
}

/** 非菜单页面的路由标题兜底映射（首页不在本地菜单配置中，需单独配置） */
export const routeTitles: Record<string, string> = {
  [HomePath]: '首页',
  '/login': '登录',
  '/account': '个人中心',
  '/403': '403',
  '/404': '404',
};

/** 按路径取菜单名称：优先本地菜单配置，其次静态映射，最后返回路径本身 */
export function getRouteTitle(path: string): string {
  const find = (list: MenuConfigItem[]): string | undefined => {
    for (const item of list) {
      if (item.path === path) return item.name;
      const hit = item.children && find(item.children);
      if (hit) return hit;
    }
    return undefined;
  };
  return find(localMenuData) ?? routeTitles[path] ?? path;
}

/** 本地菜单配置 */
export const localMenuData: MenuConfigItem[] = [
  {
    key: 'Admin',
    path: '/admin',
    name: '系统设置',
    icon: 'SettingOutlined',
    children: [
      { key: 'AdminUser', path: '/admin/user', name: '账号管理' },
      { key: 'AdminRole', path: '/admin/role', name: '角色管理' },
      { key: 'AdminMenu', path: '/admin/menu', name: '菜单管理' },
      { key: 'AdminPermission', path: '/admin/permission', name: '权限管理' },
      { key: 'AdminApi', path: '/admin/api', name: '接口管理' },
    ],
  },
];

/**
 * 把远程菜单转为本地菜单结构
 * @param result 结果数组
 * @param defaultMenuData 本地默认菜单
 * @param remoteMenus 远程菜单 map（key 为菜单键）
 * @param childrenKey 子菜单键名
 */
export function HandleRemoteMenuIntoLocal(
  result: MenuConfigItem[],
  defaultMenuData: MenuConfigItem[],
  remoteMenus: Record<string, any>,
  childrenKey: string,
): MenuConfigItem[] {
  defaultMenuData?.forEach((item) => {
    const tmpItem: MenuConfigItem = { ...item };
    if (tmpItem.key && remoteMenus[tmpItem.key]) {
      tmpItem.hideInMenu = !!remoteMenus[tmpItem.key].hideInMenu;
      // 图标以后端配置为准，后端未配置时沿用本地默认图标
      const remoteIcon = remoteMenus[tmpItem.key].icon;
      if (remoteIcon !== undefined && remoteIcon !== null) {
        tmpItem.icon = remoteIcon;
      }
      result.push(tmpItem);
      if ((item as any)[childrenKey]?.length > 0) {
        (tmpItem as any)[childrenKey] = [];
        HandleRemoteMenuIntoLocal(
          (tmpItem as any)[childrenKey],
          (item as any)[childrenKey],
          remoteMenus,
          childrenKey,
        );
      }
    }
  });
  return result;
}
