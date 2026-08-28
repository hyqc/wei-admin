import { post } from '@/api/request';
import type { PageInfoType } from '@/api/types';
import type { AdminUserListItem } from '@/types/common';
import type {
  ReqAdminUserList,
  ReqAdminUserAdd,
  ReqAdminUserEdit,
  ReqAdminUserDelete,
  ReqAdminUserEnable,
  ReqAdminUserInfo,
  ReqAdminUserBindRoles,
  ReqAdminUserPasswordReset,
} from '@/types/admin_user';

/** 账号列表返回 */
export interface ResponseAdminUserListType {
  list: AdminUserListItem[];
  pageInfo: PageInfoType;
}

/** 账号详情返回 */
export interface ResponseAdminUserInfoType extends AdminUserListItem {
  roleIds?: number[];
}

/** 账号列表 */
export const getAdminUserList = (body: ReqAdminUserList) =>
  post<ResponseAdminUserListType>('/admin/user/list', body);

/** 账号详情 */
export const getAdminUserInfo = (body: ReqAdminUserInfo) =>
  post<ResponseAdminUserInfoType>('/admin/user/info', body);

/** 创建账号 */
export const addAdminUser = (body: ReqAdminUserAdd) => post<null>('/admin/user/add', body);

/** 修改账号 */
export const editAdminUser = (body: ReqAdminUserEdit) => post<null>('/admin/user/edit', body);

/** 删除账号 */
export const deleteAdminUser = (body: ReqAdminUserDelete) => post<null>('/admin/user/delete', body);

/** 启用禁用账号 */
export const enableAdminUser = (body: ReqAdminUserEnable) => post<null>('/admin/user/enable', body);

/** 重置密码 */
export const resetAdminUserPassword = (body: ReqAdminUserPasswordReset) =>
  post<null>('/admin/user/edit_pwd', body);

/** 账号绑定角色 */
export const bindAdminUserRoles = (body: ReqAdminUserBindRoles) =>
  post<null>('/admin/user/bind_roles', body);
