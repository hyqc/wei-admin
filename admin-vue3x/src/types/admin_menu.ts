import type { ReqListBase } from './common';

/** 菜单列表请求 */
export interface ReqAdminMenuList extends ReqListBase {
  key?: string;
  path?: string;
  name?: string;
  parentId?: number;
}

/** 创建/修改菜单 */
export interface ReqAdminMenuAdd {
  id?: number;
  parentId?: number;
  path?: string;
  name: string;
  key: string;
  describe?: string;
  icon?: string;
  sort?: number;
  redirect?: string;
  component?: string;
  hideChildrenInMenu?: boolean;
  hideInMenu?: boolean;
  enabled?: boolean;
  /** 表单使用的字段：是否隐藏菜单 */
  isHideInMenu?: boolean;
  /** 表单使用的字段：是否隐藏子菜单 */
  isHideChildrenInMenu?: boolean;
  /** 表单使用的字段：是否启用 */
  isEnabled?: boolean;
}

/** 启用禁用菜单 */
export interface ReqAdminMenuEnable {
  menuId?: number;
  enabled?: boolean;
}

/** 是否显示菜单 */
export interface ReqAdminMenuShow {
  menuId?: number;
  show?: boolean;
}

/** 删除菜单 */
export interface ReqAdminMenuDelete {
  menuId?: number;
}

/** 菜单详情 */
export interface ReqAdminMenuInfo {
  menuId?: number;
}

/** 菜单权限列表 */
export interface ReqAdminMenuPermissions {
  menuId?: number;
}

/** 页面菜单列表 */
export interface ReqAdminMenuPages {
  all?: boolean;
}

/** 页面模块权限列表 */
export interface ReqAdminMenuMode {}

/** 菜单 */
export interface MenuItem {
  /** 菜单唯一键 */
  key?: string;
  /** 菜单路由 */
  path?: string;
  /** 菜单名称 */
  name?: string;
  /** 菜单图标 */
  icon?: string;
  component?: string;
  authority?: string;
  hideInMenu?: boolean;
  hideChildrenInMenu?: boolean;
  routes?: MenuItem[];
}

/** 有效菜单树 */
export interface MenuTreeItem {
  /** 菜单层级 */
  level?: number;
  /** 菜单自增ID */
  id?: number;
  /** 菜单唯一键 */
  key?: string;
  /** 菜单名称 */
  name?: string;
  /** 父级菜单ID */
  parentId?: number;
  /** 菜单描述 */
  describe?: string;
  /** 菜单路径 */
  path?: string;
  /** 重定向地址 */
  redirect?: string;
  /** 组件名称 */
  component?: string;
  /** 菜单排序 */
  sort?: number;
  /** 菜单图标 */
  icon?: string;
  /** 是否在菜单中隐藏子菜单 */
  hideChildrenInMenu?: boolean;
  /** 是否隐藏菜单 */
  hideInMenu?: boolean;
  /** 是否启用 */
  enabled?: boolean;
  createTime?: number;
  modifyTime?: number;
  children?: MenuTreeItem[];
}

/** 菜单模型 */
export interface AdminMenuModel {
  id?: number;
  parentId?: number;
  path?: string;
  name?: string;
  key?: string;
  describe?: string;
  icon?: string;
  sort?: number;
  redirect?: string;
  component?: string;
  isHideInMenu?: boolean;
  isHideChildrenInMenu?: boolean;
  isEnabled?: boolean;
  createdAt?: string;
  updatedAt?: string;
}

/** 菜单模块 */
export interface MenuModeItem {
  /** 模块菜单ID */
  modelId?: number;
  /** 模块菜单名称 */
  modelName?: string;
  /** 模块下面的页面列表 */
  pages?: MenuPageItem[];
}

/** 模块页面 */
export interface MenuPageItem {
  /** 页面菜单ID */
  pageId?: number;
  /** 页面菜单名称 */
  pageName?: string;
  /** 页面菜单权限 */
  permissions?: MenuPagePermissions[];
}

/** 模块页面权限 */
export interface MenuPagePermissions {
  /** 权限ID */
  permissionId?: number;
  /** 权限名称 */
  permissionName?: string;
  /** 权限类型 */
  permissionType?: string;
  /** 权限类型名称 */
  permissionTypeName?: string;
}

/** 菜单权限 */
export interface MenuPermissionItem {
  id?: number;
  name?: string;
  key?: string;
  type?: string;
  typeName?: string;
  describe?: string;
  enabled?: boolean;
  menuId?: number;
}
