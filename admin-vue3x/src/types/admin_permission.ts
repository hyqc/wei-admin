import type { AdminApiItem, ReqListBase } from './common';

/** 权限列表请求 */
export interface ReqAdminPermissionList extends ReqListBase {
  /** 菜单ID */
  menuId?: number;
  /** 权限唯一标识符 */
  key?: string;
  /** 权限名称 */
  name?: string;
  /** 权限类型 */
  type?: string;
}

/** 创建权限 */
export interface ReqAdminPermissionAdd {
  /** 权限ID */
  id?: number;
  /** 权限对应的菜单ID */
  menuId?: number;
  /** 菜单名称 */
  menuName?: string;
  /** 菜单路由 */
  menuPath?: string;
  /** 权限唯一标识符 */
  key?: string;
  /** 权限名称 */
  name?: string;
  /** 权限描述 */
  describe?: string;
  /** 权限类型 */
  type?: string;
  /** 重定向地址 */
  redirect?: string;
  /** 是否启用 */
  enabled?: boolean;
}

/** 修改权限 */
export interface ReqAdminPermissionEdit {
  id?: number;
  menuId?: number;
  key?: string;
  name?: string;
  describe?: string;
  type?: string;
  redirect?: string;
  enabled?: boolean;
}

/** 权限详情 */
export interface ReqAdminPermissionInfo {
  id?: number;
}

/** 权限绑定接口 */
export interface ReqAdminPermissionBindApis {
  permissionId?: number;
  apiIds?: number[];
}

/** 绑定权限菜单 */
export interface ReqAdminPermissionBindMenu {
  /** 菜单ID */
  menuId?: number;
  /** 菜单对应的权限列表 */
  permissions?: ReqAdminPermissionAdd[];
}

/** 权限列表项 */
export interface PermissionListItem {
  /** 权限ID */
  id?: number;
  /** 权限对应的菜单ID */
  menuId?: number;
  /** 菜单名称 */
  menuName?: string;
  /** 菜单路由 */
  menuPath?: string;
  /** 接口列表 */
  apis?: AdminApiItem[];
  /** 权限唯一标识符 */
  key?: string;
  /** 权限名称 */
  name?: string;
  /** 权限描述 */
  describe?: string;
  /** 权限类型 */
  type?: string;
  /** 权限类型名称 */
  typeText?: string;
  /** 是否启用 */
  isEnabled?: boolean;
}

/** 权限接口列表 */
export interface PermissionApiItem {
  /** 权限ID */
  id?: number;
  /** 菜单ID */
  menuId?: number;
  /** 权限键名 */
  key?: string;
  /** 权限类型 */
  type?: string;
  /** 权限类型名称 */
  typeText?: string;
  /** 权限名称 */
  name?: string;
  /** 接口列表 */
  apis?: AdminApiItem[];
  /** 是否启用 */
  enabled?: boolean;
  /** 权限描述 */
  describe?: string;
}

/** 权限详情 */
export interface AdminPermissionInfo {
  /** 权限ID */
  id?: number;
  /** 权限对应的菜单ID */
  menuId?: number;
  /** 菜单名称 */
  menuName?: string;
  /** 菜单路由 */
  menuPath?: string;
  /** 接口列表 */
  apis?: AdminApiItem[];
  /** 权限唯一标识符 */
  key?: string;
  /** 权限名称 */
  name?: string;
  /** 权限描述 */
  describe?: string;
  /** 权限类型 */
  type?: string;
  /** 权限类型名称 */
  typeText?: string;
}
