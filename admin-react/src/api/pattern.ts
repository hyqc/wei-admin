/** 账号名称 */
export const AdminUsername = /^[0-9a-zA-Z]{1,50}$/g;

/** 账号密码 */
export const AdminUserPassword = /^[\S][\s\S]{4,30}[\S]$/g;

/** 菜单键名 */
export const AdminMenuKey = /^([A-Z][a-zA-Z0-9]*)+$/g;

/** 权限键名 */
export const AdminPerssionKey = AdminMenuKey;

/** 路由路径 */
export const AdminRouterPath = /^(\/[a-z][a-zA-Z0-9]*\/?)+$/g;

/** 接口键名 */
export const AdminAPIKey = /^[a-z][a-zA-Z0-9]*::[a-z][a-zA-Z0-9]*$/g;

/** 邮箱 */
export const AdminEmail = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
