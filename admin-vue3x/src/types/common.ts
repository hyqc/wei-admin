/** 列表请求基础结构 */
export interface ReqListBase {
  /** 分页偏移量 */
  pageSize?: number;
  /** 页码 */
  pageNum?: number;
  /** 排序字段 */
  sortField?: string;
  /** 排序值 */
  sortType?: string;
  /** 0：全部 1：启用 2：禁用 */
  enabled?: number;
  /** 查询开始时间戳秒 */
  createStartTime?: number;
  /** 查询结束时间戳秒 */
  createEndTime?: number;
}

/** 接口列表返回结构 */
export interface AdminApiItem {
  id?: number;
  path?: string;
  key?: string;
  name?: string;
  isEnabled?: boolean;
  permissionId?: number;
  createdAt?: string;
  updatedAt?: string;
  describe?: string;
}

/** 账号列表返回结构 */
export interface AdminUserListItem {
  /** 管理员ID */
  adminId?: number;
  /** 账号名称 */
  username?: string;
  /** 昵称 */
  nickname?: string;
  /** 邮箱 */
  email?: string;
  /** 头像 */
  avatar?: string;
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
  roles?: AdminUserRoleItem[];
  /** 是否禁用启停按钮 */
  isEnabledButtonDisabled?: boolean;
  /** 是否超级管理员（ID 恒为 1，自动拥有全部权限，不可绑定角色） */
  isSuperAdmin?: boolean;
}

/** 账号角色列表 */
export interface AdminUserRoleItem {
  /** 角色ID */
  roleId?: number;
  roleName?: string;
}

/** 分页信息 */
export interface PageInfoType {
  pageSize: number;
  pageNum: number;
  total: number;
}

/** 权限集合 */
export interface PermissionsType {
  [key: string]: boolean;
}
