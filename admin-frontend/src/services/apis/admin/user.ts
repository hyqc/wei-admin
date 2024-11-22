// adminUsers 管理员管理接口
import { request } from 'umi';
import { APIAdminUsers } from './api';
import { ResponseBodyType } from '../types';
import { ReqAdminUserInfo, ReqAdminUserList } from '@/proto/admin_ts/admin_user';
import { ReqAdminUserEdit } from '@/proto/admin_ts/admin_user';
import { ReqAdminUserEditPassword } from '@/proto/admin_ts/admin_user';


export async function adminUserList(params?: ReqAdminUserList) {
  return request<ResponseBodyType>(APIAdminUsers.list.url, {
    method: APIAdminUsers.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 新增
 */
export type RequestAdminUserAddParamsType = {
  username: string; // 管理员名称（账号）
  enabled: boolean;
  passwrod: string;
  confirmPassword: string;
  nickname?: string; // 管理员昵称
  avatar?: string;
  email?: string;
};

export async function adminUserAdd(params: RequestAdminUserAddParamsType) {
  return request<ResponseBodyType>(APIAdminUsers.add.url, {
    method: APIAdminUsers.add.method,
    data: params,
  });
}


export async function adminUserDetail(params: ReqAdminUserInfo) {
  return request<ResponseBodyType>(APIAdminUsers.info.url, {
    method: APIAdminUsers.info.method,
    data: params,
  });
}

export async function adminUserEdit(params: ReqAdminUserEdit) {
  return request<ResponseBodyType>(APIAdminUsers.edit.url, {
    method: APIAdminUsers.edit.method,
    data: params,
  });
}

export async function adminUserEditPwd(params: ReqAdminUserEditPassword) {
  return request<ResponseBodyType>(APIAdminUsers.editPwd.url, {
    method: APIAdminUsers.editPwd.method,
    data: params,
  });
}

/************************************************************/
/**
 * 彻底删除用户
 */
export type RequestAdminUserDeleteParamsType = {
  adminId: number;
  enabled: boolean;
};

export async function adminUserDelete(params: RequestAdminUserDeleteParamsType) {
  return request<ResponseBodyType>(APIAdminUsers.delete.url, {
    method: APIAdminUsers.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 绑定角色
 */
export type RequestAdminUserAssignRolesParamsType = {
  adminId: number;
  roleIds: number[];
};

export async function adminUserBindRoles(params: RequestAdminUserAssignRolesParamsType) {
  return request<ResponseBodyType>(APIAdminUsers.bindRoles.url, {
    method: APIAdminUsers.bindRoles.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export type RequestAdminUserEnableParamsType = {
  adminId: number;
  enabled: boolean;
};

export async function adminUserEnable(params: RequestAdminUserEnableParamsType) {
  return request<ResponseBodyType>(APIAdminUsers.enable.url, {
    method: APIAdminUsers.enable.method,
    data: params,
  });
}
