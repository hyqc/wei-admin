// adminPermissions 管理员管理接口
import { request } from 'umi';
import { APIAdminPermissions } from './api';
import type { ResponseBodyType } from '../types';
import { ReqAdminPermissionAdd, ReqAdminPermissionBindApis, ReqAdminPermissionDelete, ReqAdminPermissionEdit, ReqAdminPermissionEnable, ReqAdminPermissionInfo, ReqAdminPermissionList, ReqAdminPermissionUnBindApi } from '@/proto/admin_ts/admin_permission';
import { ReqAdminRoleBindPermissions } from '@/proto/admin_ts/admin_role';

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
export async function adminPermissionList(params?: ReqAdminPermissionList) {
  return request<ResponseBodyType>(APIAdminPermissions.list.url, {
    method: APIAdminPermissions.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 给菜单创建权限
 */
export async function adminAddMenuPermission(params: ReqAdminRoleBindPermissions) {
  return request<ResponseBodyType>(APIAdminPermissions.addMenuPermissions.url, {
    method: APIAdminPermissions.addMenuPermissions.method,
    data: params,
  });
}
/************************************************************/
/**
 * 添加权限
 */
export async function adminPermissionAdd(params: ReqAdminPermissionAdd) {
  return request<ResponseBodyType>(APIAdminPermissions.add.url, {
    method: APIAdminPermissions.add.method,
    data: params,
  });
}

/************************************************************/
/**
 * 详情
 */
export async function adminPermissionDetail(params: ReqAdminPermissionInfo) {
  return request<ResponseBodyType>(APIAdminPermissions.info.url, {
    method: APIAdminPermissions.info.method,
    data: params,
  });
}

/************************************************************/
/**
 * 编辑
 */
export async function adminPermissionEdit(params: ReqAdminPermissionEdit) {
  return request<ResponseBodyType>(APIAdminPermissions.edit.url, {
    method: APIAdminPermissions.edit.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export async function adminPermissionEnable(params: ReqAdminPermissionEnable) {
  return request<ResponseBodyType>(APIAdminPermissions.enable.url, {
    method: APIAdminPermissions.enable.method,
    data: params,
  });
}

/************************************************************/
/**
 * 删除
 */
export async function adminPermissionDelete(params: ReqAdminPermissionDelete) {
  return request<ResponseBodyType>(APIAdminPermissions.delete.url, {
    method: APIAdminPermissions.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 绑定接口列表
 */
export async function adminPermissionBindApis(params: ReqAdminPermissionBindApis) {
  return request<ResponseBodyType>(APIAdminPermissions.bindApis.url, {
    method: APIAdminPermissions.bindApis.method,
    data: params,
  });
}
/************************************************************/
/**
 * 解绑接口
 */
export async function adminPermissionUnBindApi(params: ReqAdminPermissionUnBindApi) {
  return request<ResponseBodyType>(APIAdminPermissions.unbind.url, {
    method: APIAdminPermissions.unbind.method,
    data: params,
  });
}
