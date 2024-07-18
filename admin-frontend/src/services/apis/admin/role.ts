// adminRoles 管理员角色接口
import { request } from 'umi';
import { APIAdminRoles } from './api';
import { ResponseBodyType, ResponseListType } from '../types';

/************************************************************/
/**
 * 列表
 */
export type RequestAdminRoleListParamsType = {
  roleName?: string;
  enabled?: number;
  pageNum?: number;
  pageSize?: number;
  sortField?: string;
  sortType?: string;
};

export type ResponseAdminRoleListItemType = {
  roleId: number;
  roleName: string;
  enabled: boolean;
  enabledText?: string;
  createTime: string;
  modifyTime: string;
  [key: string]: any;
};

export async function adminRoleList(params?: RequestAdminRoleListParamsType) {
  return request<ResponseListType>(APIAdminRoles.list.url, {
    method: APIAdminRoles.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 新增
 */
export type RequestAdminRoleAddParamsType = {
  name: string;
  enabled: boolean;
  describe?: string;
};

export async function adminRoleAdd(params: RequestAdminRoleAddParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.add.url, {
    method: APIAdminRoles.add.method,
    data: params,
  });
}

/************************************************************/
/**
 * 详情
 */
export type RequestAdminRoleIDetailParamsType = {
  id: number;
};

export type ResponseAdminRoleDetailType = {
  id: number;
  name: string;
  permissionIds?: number[];
  enabled?: boolean;
  enabledText?: string;
  createTime?: string;
  modifyTime?: string;
  [key: string]: any;
};

export async function adminRoleDetail(params: RequestAdminRoleIDetailParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.detail.url, {
    method: APIAdminRoles.detail.method,
    data: params,
  });
}

/************************************************************/
/**
 * 编辑
 */
export type RequestAdminRoleEditParamsType = {
  id: number;
  name?: string;
  describe?: string;
  enabled?: boolean;
  permissionIds?: string;
};

export async function adminRoleEdit(params: RequestAdminRoleEditParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.edit.url, {
    method: APIAdminRoles.edit.method,
    data: params,
  });
}

/************************************************************/
/**
 * 全部
 */
export type RequestAdminRoleAllParamsType = {
  name?: string; // 角色名称
};

export type ResponseAdminRoleAllItemType = {
  roleId: number;
  roleName: string;
};

export async function adminRoleAll(params?: RequestAdminRoleAllParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.all.url, {
    method: APIAdminRoles.all.method,
    data: params,
  });
}

/************************************************************/
/**
 * 删除
 */

export type RequestAdminRoleDeleteParamsType = {
  id: number;
  enabled: boolean;
};

export async function adminRoleDelete(params: RequestAdminRoleDeleteParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.delete.url, {
    method: APIAdminRoles.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export type RequestAdminRoleEnableParamsType = {
  id: number;
  enabled: boolean;
};

export async function adminRoleEnable(params: RequestAdminRoleEnableParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.enable.url, {
    method: APIAdminRoles.enable.method,
    data: params,
  });
}

/************************************************************/
/**
 * 角色分配权限
 */
export type RequestAdminRoleBindPermissionsParamsType = {
  id: number;
  permissionIds: number[];
};

export async function adminRoleBind(params: RequestAdminRoleBindPermissionsParamsType) {
  return request<ResponseBodyType>(APIAdminRoles.bindPermissions.url, {
    method: APIAdminRoles.bindPermissions.method,
    data: params,
  });
}
