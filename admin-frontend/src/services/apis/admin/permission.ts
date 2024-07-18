// adminPermissions 管理员管理接口
import { request } from 'umi';
import { APIAdminPermissions } from './api';
import type { ResponseBodyType, ResponseListType } from '../types';
import { ResponseAdminAPIAllItemType } from './resource';

/************************************************************/
/**
 * 全部权限列表
 * @returns
 */
export async function adminPermissionAll() {
  return request<ResponseBodyType>(APIAdminPermissions.all.url, {
    method: APIAdminPermissions.all.method,
    data: {},
  });
}

/************************************************************/
/**
 * 权限分页列表
 */
export type RequestAdminPermissionListParamsType = {
  name?: string;
  key?: string;
  menuId?: number;
  enabled?: boolean;
  type?: string;
  pageSize?: number;
  pageNum?: number;
};

export type ResponseAdminMenuPermissionsItemType = {
  menuId: number;
  id?: number;
  key: string;
  type: string;
  typeText: string;
  name: string;
  apis: ResponseAdminAPIAllItemType[];
  enabled: boolean;
  describe?: string;
};

export type ResponseAdminPermissionListItemType = {
  menuId: number;
  menuName?: string;
  menuPath?: string;
  id: number;
  key: string;
  name: string;
  type: string;
  typeText: string;
  describe?: string;
  enabled: boolean;
  enabledText?: string;
  createTime: string;
  modifyTime: string;
  apis: ResponseAdminAPIAllItemType[];
};

export async function adminPermissionList(params?: RequestAdminPermissionListParamsType) {
  return request<ResponseListType>(APIAdminPermissions.list.url, {
    method: APIAdminPermissions.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 给菜单创建权限
 */
export type RequestAdminPermissionAddForMenuParamsType = {
  menuId: number;
  permissions: ResponseAdminMenuPermissionsItemType[];
};

export async function adminAddMenuPermission(params: RequestAdminPermissionAddForMenuParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.addMenuPermissions.url, {
    method: APIAdminPermissions.addMenuPermissions.method,
    data: params,
  });
}
/************************************************************/
/**
 * 添加权限
 */
export type RequestAdminPermissionAddParamsType = {
  name: string; // 权限唯一名称
  key: string; // 权限的唯一键
  type: string; // 权限类型
  menuId: number; // 菜单ID
  describe: string; // 权限描述
  enabled: boolean; // 权限状态
};

export async function adminPermissionAdd(params: RequestAdminPermissionAddParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.add.url, {
    method: APIAdminPermissions.add.method,
    data: params,
  });
}

/************************************************************/
/**
 * 详情
 */
export type RequestAdminPermissionDetailParamsType = {
  id: number;
};

export type ResponseAdminPermissionDetailType = {
  menuId: number;
  menuName?: string;
  menuPath: string;
  id: number;
  key: string;
  name: string;
  type: string;
  typeText: string;
  describe?: string;
  enabled: boolean;
  enabledText?: string;
  createTime: string;
  modifyTime: string;
  apis: ResponseAdminAPIAllItemType[];
};
export async function adminPermissionDetail(params: RequestAdminPermissionDetailParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.detail.url, {
    method: APIAdminPermissions.detail.method,
    data: params,
  });
}

/************************************************************/
/**
 * 编辑
 */
export type RequestAdminPermissionEditParamsType = {
  id: number;
  name?: string;
  key?: string;
  describe?: string;
  enabled: boolean;
  type: string;
};

export async function adminPermissionEdit(params: RequestAdminPermissionEditParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.edit.url, {
    method: APIAdminPermissions.edit.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export type RequestAdminPermissionEnableParamsType = {
  id: number;
  enabled: boolean;
};

export async function adminPermissionEnable(params: RequestAdminPermissionEnableParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.enable.url, {
    method: APIAdminPermissions.enable.method,
    data: params,
  });
}

/************************************************************/
/**
 * 删除
 */
export type RequestAdminPermissionDeleteParamsType = {
  id: number;
  enabled: boolean;
};

export async function adminPermissionDelete(params: RequestAdminPermissionDeleteParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.delete.url, {
    method: APIAdminPermissions.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 解绑接口
 */
export type RequestAdminPermissionUnbindApiParamsType = {
  id: number;
  apiId: number;
};

export async function adminPermissionUnbindApi(params: RequestAdminPermissionUnbindApiParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.unbind.url, {
    method: APIAdminPermissions.unbind.method,
    data: params,
  });
}

/************************************************************/
/**
 * 绑定接口
 */
export type RequestAdminPermissionBindApiParamsType = {
  permissionId: number;
  apiIds: number;
};

export async function adminPermissionBindApi(params: RequestAdminPermissionBindApiParamsType) {
  return request<ResponseBodyType>(APIAdminPermissions.bindApis.url, {
    method: APIAdminPermissions.bindApis.method,
    data: params,
  });
}
