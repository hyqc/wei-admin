// adminUsers 管理员管理接口
import { request } from 'umi';
import { APIAdminUsers } from './api';
import { ResponseBodyType } from '../types';
import { ReqAdminUserAdd, ReqAdminUserBindRoles, ReqAdminUserDelete, ReqAdminUserEnabled, ReqAdminUserInfo, ReqAdminUserList } from '@/proto/admin_ts/admin_user';
import { ReqAdminUserEdit } from '@/proto/admin_ts/admin_user';
import { ReqAdminUserEditPassword } from '@/proto/admin_ts/admin_user';

//列表
export async function adminUserList(params?: ReqAdminUserList) {
  return request<ResponseBodyType>(APIAdminUsers.list.url, {
    method: APIAdminUsers.list.method,
    data: params,
  });
}


//新增
export async function adminUserAdd(params: ReqAdminUserAdd) {
  return request<ResponseBodyType>(APIAdminUsers.add.url, {
    method: APIAdminUsers.add.method,
    data: params,
  });
}

//详情
export async function adminUserDetail(params: ReqAdminUserInfo) {
  return request<ResponseBodyType>(APIAdminUsers.info.url, {
    method: APIAdminUsers.info.method,
    data: params,
  });
}

//编辑
export async function adminUserEdit(params: ReqAdminUserEdit) {
  return request<ResponseBodyType>(APIAdminUsers.edit.url, {
    method: APIAdminUsers.edit.method,
    data: params,
  });
}

//修改密码
export async function adminUserEditPwd(params: ReqAdminUserEditPassword) {
  return request<ResponseBodyType>(APIAdminUsers.editPwd.url, {
    method: APIAdminUsers.editPwd.method,
    data: params,
  });
}

//彻底删除用户
export async function adminUserDelete(params: ReqAdminUserDelete) {
  return request<ResponseBodyType>(APIAdminUsers.delete.url, {
    method: APIAdminUsers.delete.method,
    data: params,
  });
}


//绑定角色
export async function adminUserBindRoles(params: ReqAdminUserBindRoles) {
  return request<ResponseBodyType>(APIAdminUsers.bindRoles.url, {
    method: APIAdminUsers.bindRoles.method,
    data: params,
  });
}

//启用禁用
export async function adminUserEnable(params: ReqAdminUserEnabled) {
  return request<ResponseBodyType>(APIAdminUsers.enable.url, {
    method: APIAdminUsers.enable.method,
    data: params,
  });
}
