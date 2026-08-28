import type { ReqListBase } from './common';

/** 账号列表请求 */
export interface ReqAdminUserList extends ReqListBase {
  username?: string;
  nickname?: string;
  email?: string;
}

/** 创建账号 */
export interface ReqAdminUserAdd {
  username: string;
  password: string;
  nickname?: string;
  email?: string;
  roleIds?: number[];
}

/** 修改账号 */
export interface ReqAdminUserEdit {
  adminId: number;
  nickname?: string;
  email?: string;
}

/** 删除账号 */
export interface ReqAdminUserDelete {
  adminId?: number;
}

/** 启用禁用账号 */
export interface ReqAdminUserEnable {
  adminId?: number;
  enabled?: boolean;
}

/** 账号详情 */
export interface ReqAdminUserInfo {
  adminId?: number;
}

/** 账号绑定角色 */
export interface ReqAdminUserBindRoles {
  adminId?: number;
  roleIds?: number[];
}

/** 重置密码 */
export interface ReqAdminUserPasswordReset {
  adminId?: number;
  password: string;
}
