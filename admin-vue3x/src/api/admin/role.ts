import { post } from '@/api/request';
import type { PageInfoType } from '@/api/types';
import type { RoleItem, RolePermissionItem } from '@/types/admin_role';
import type {
  ReqAdminRoleList,
  ReqAdminRoleAdd,
  ReqAdminRoleEdit,
  ReqAdminRoleDelete,
  ReqAdminRoleEnable,
  ReqAdminRoleInfo,
  ReqAdminRoleBindPermissions,
  ReqAdminRolePermissions,
} from '@/types/admin_role';

/** 角色列表返回 */
export interface ResponseAdminRoleListType {
  list: RoleItem[];
  pageInfo: PageInfoType;
}

/** 角色详情返回 */
export interface ResponseAdminRoleInfoType extends RoleItem {
  permissionIds?: number[];
}

/** 角色权限列表返回 */
export interface ResponseAdminRolePermissionsType {
  list: RolePermissionItem[];
}

/** 角色列表 */
export const getAdminRoleList = (body: ReqAdminRoleList) =>
  post<ResponseAdminRoleListType>('/admin/role/list', body);

/** 角色详情 */
export const getAdminRoleInfo = (body: ReqAdminRoleInfo) =>
  post<ResponseAdminRoleInfoType>('/admin/role/info', body);

/** 创建角色 */
export const addAdminRole = (body: ReqAdminRoleAdd) => post<null>('/admin/role/add', body);

/** 修改角色 */
export const editAdminRole = (body: ReqAdminRoleEdit) => post<null>('/admin/role/edit', body);

/** 删除角色 */
export const deleteAdminRole = (body: ReqAdminRoleDelete) => post<null>('/admin/role/delete', body);

/** 启用禁用角色 */
export const enableAdminRole = (body: ReqAdminRoleEnable) => post<null>('/admin/role/enable', body);

/** 全部角色列表 */
export const getAdminRoleAll = (body?: { name?: string }) =>
  post<RoleItem[]>('/admin/role/all', body || {});

/** 角色绑定权限 */
export const bindAdminRolePermissions = (body: ReqAdminRoleBindPermissions) =>
  post<null>('/admin/role/bind_permissions', body);

/** 角色权限列表 */
export const getAdminRolePermissions = (body: ReqAdminRolePermissions) =>
  post<ResponseAdminRolePermissionsType>('/admin/role/permissions', body);
