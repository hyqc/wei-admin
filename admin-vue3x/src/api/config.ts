/** 接口基础路径 */
export const BaseAPI = '/api';

/** 登录 token 存储键 */
export const LocalStorageTokenKey = 'token';

/** 登录页路径 */
export const LoginPath = '/login';

/** 首页路径 */
export const HomePath = '/home';

/** 默认分页大小 */
export const DefaultPageSize = 10;

/** 默认分页信息 */
export const DEFAULT_PAGE_INFO = {
  total: 0,
  pageSize: DefaultPageSize,
  pageNum: 1,
};

/** 分页大小选项 */
export const DefaultPageArray = [1, 5, 10, 50, 100];

/** 默认弹窗宽度 */
export const DefaultModalWidth = 800;

/** 默认抽屉宽度 */
export const DefaultDrawerWidth = 600;

/** 超级管理员 ID */
export const AdminId = 1;

/** 是否超级管理员 */
export const IsSuperAdmin = (adminId?: number) => adminId === AdminId;

/** 输入框样式 */
export const INPUT_STYLE = { fontWeight: 400, color: 'black' };
