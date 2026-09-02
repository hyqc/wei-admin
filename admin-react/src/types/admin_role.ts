import type { ReqListBase } from './common';

/** 角色列表请求 */
export interface ReqAdminRoleList extends ReqListBase {
  name?: string;
}

/** 创建角色 */
export interface ReqAdminRoleAdd {
  name: string;
  describe?: string;
}

/** 修改角色 */
export interface ReqAdminRoleEdit {
  id: number;
  name?: string;
  describe?: string;
}

/** 删除角色 */
export interface ReqAdminRoleDelete {
  id?: number;
}

/** 启用禁用角色 */
export interface ReqAdminRoleEnable {
  id?: number;
  enabled?: boolean;
}

/** 角色详情 */
export interface ReqAdminRoleInfo {
  id?: number;
}

/** 角色绑定权限 */
export interface ReqAdminRoleBindPermissions {
  id?: number;
  permissionIds?: number[];
}

/** 角色权限列表 */
export interface ReqAdminRolePermissions {
  id?: number;
}

/** 角色 */
export interface RoleItem {
  /** 角色ID */
  id?: number;
  /** 角色名称 */
  name?: string;
  /** 角色描述 */
  describe?: string;
  /** 创建人ID */
  createAdminId?: number;
  /** 创建人名称 */
  createAdminName?: string;
  /** 是否启用 */
  isEnabled?: boolean;
  createdAt?: string;
  updatedAt?: string;
  /** 是否超级管理员角色（ID 恒为 1，自动拥有全部权限，仅可修改描述） */
  isSuperAdmin?: boolean;
}

/** 角色权限 */
export interface RolePermissionItem {
  /** 角色ID */
  roleId?: number;
  /** 权限ID */
  permissionId?: number;
  /** 权限名称 */
  permissionName?: string;
  /** 权限键名 */
  permissionKey?: string;
  /** 权限类型 */
  permissionType?: string;
  /** 权限类型名称 */
  permissionTypeText?: string;
}
