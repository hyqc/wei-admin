import type { ReqListBase } from './common';

/** 登录请求 */
export interface ReqLogin {
  username: string;
  password: string;
  remember?: boolean;
}

/** 登录返回数据 */
export interface LoginRespData {
  token: string;
  expire: number;
  data: AdminInfo;
}

/** 当前账号信息 */
export interface AdminInfo {
  adminId: number;
  username: string;
  nickname: string;
  email: string;
  avatar: string;
  /** 登录 token */
  token?: string;
  /** token 有效期（秒） */
  expire?: number;
  /** 登录次数 */
  loginTotal?: number;
  /** 上次登录IP */
  lastLoginIp?: string;
  /** 上次登录时间 */
  lastLoginTime?: string;
  /** 是否启用 */
  isEnabled?: boolean;
  /** 创建时间 */
  createdAt?: string;
  /** 更新时间 */
  updatedAt?: string;
  /** 角色列表 */
  roles?: { roleId?: number; roleName?: string }[];
  /** 角色名称列表 */
  rolesName?: string[];
  /** 菜单 map（key 为菜单键） */
  menus?: Record<string, any>;
  /** 权限 map（key 为权限键） */
  permissions?: Record<string, string>;
}

/** 修改账号信息 */
export interface ReqAccountEdit {
  adminId?: number;
  nickname: string;
  email: string;
  avatar?: string;
}

/** 修改密码 */
export interface ReqAccountPasswordEdit {
  oldPassword: string;
  newPassword: string;
  confirmPassword: string;
}

/** 账号列表请求 */
export interface ReqAdminAccountList extends ReqListBase {
  username?: string;
  nickname?: string;
}
