// adminRoles 管理员角色接口
import { request } from 'umi';
import { APIAdminRoles } from './api';
import { ResponseBodyType } from '../types';
import { ReqAdminRoleAdd, ReqAdminRoleAll, ReqAdminRoleBindPermissions, ReqAdminRoleDelete, ReqAdminRoleEdit, ReqAdminRoleEnable, ReqAdminRoleInfo, ReqAdminRoleList, ReqAdminRolePermissions } from '@/proto/admin_ts/admin_role';

//列表
export async function adminRoleList(params?: ReqAdminRoleList) {
  return request<ResponseBodyType>(APIAdminRoles.list.url, {
    method: APIAdminRoles.list.method,
    data: params,
  });
}

//新增
export async function adminRoleAdd(params: ReqAdminRoleAdd) {
  return request<ResponseBodyType>(APIAdminRoles.add.url, {
    method: APIAdminRoles.add.method,
    data: params,
  });
}

//详情
export async function adminRoleDetail(params: ReqAdminRoleInfo) {
  return request<ResponseBodyType>(APIAdminRoles.info.url, {
    method: APIAdminRoles.info.method,
    data: params,
  });
}

//编辑
export async function adminRoleEdit(params: ReqAdminRoleEdit) {
  return request<ResponseBodyType>(APIAdminRoles.edit.url, {
    method: APIAdminRoles.edit.method,
    data: params,
  });
}

//全部
export async function adminRoleAll(params?: ReqAdminRoleAll) {
  return request<ResponseBodyType>(APIAdminRoles.all.url, {
    method: APIAdminRoles.all.method,
    data: params,
  });
}

//删除
export async function adminRoleDelete(params: ReqAdminRoleDelete) {
  return request<ResponseBodyType>(APIAdminRoles.delete.url, {
    method: APIAdminRoles.delete.method,
    data: params,
  });
}

//启用禁用
export async function adminRoleEnable(params: ReqAdminRoleEnable) {
  return request<ResponseBodyType>(APIAdminRoles.enable.url, {
    method: APIAdminRoles.enable.method,
    data: params,
  });
}

//角色分配权限
export async function adminRoleBindPermissions(params: ReqAdminRoleBindPermissions) {
  return request<ResponseBodyType>(APIAdminRoles.bindPermissions.url, {
    method: APIAdminRoles.bindPermissions.method,
    data: params,
  });
}
