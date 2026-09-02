import { post } from '@/api/request';
import type { PageInfoType } from '@/api/types';
import type {
  PermissionListItem,
  AdminPermissionInfo,
  ReqAdminPermissionList,
  ReqAdminPermissionAdd,
  ReqAdminPermissionEdit,
  ReqAdminPermissionInfo,
  ReqAdminPermissionBindApis,
} from '@/types/admin_permission';

/** 权限列表返回 */
export interface ResponseAdminPermissionListType {
  list: PermissionListItem[];
  pageInfo: PageInfoType;
}

/** 权限详情返回 */
export interface ResponseAdminPermissionInfoType extends AdminPermissionInfo {
  apiIds?: number[];
}

/** 权限列表 */
export const getAdminPermissionList = (body: ReqAdminPermissionList) =>
  post<ResponseAdminPermissionListType>('/admin/permission/list', body);

/** 权限详情 */
export const getAdminPermissionInfo = (body: ReqAdminPermissionInfo) =>
  post<ResponseAdminPermissionInfoType>('/admin/permission/info', body);

/** 创建权限 */
export const addAdminPermission = (body: ReqAdminPermissionAdd) => post<null>('/admin/permission/add', body);

/** 权限指定菜单批量创建权限 */
export const addAdminMenuPermissions = (body: ReqAdminPermissionAdd[]) =>
  post<null>('/admin/permission/add_menu_permissions', body);

/** 修改权限 */
export const editAdminPermission = (body: ReqAdminPermissionEdit) => post<null>('/admin/permission/edit', body);

/** 删除权限 */
export const deleteAdminPermission = (body: { id?: number }) => post<null>('/admin/permission/delete', body);

/** 启用禁用权限 */
export const enableAdminPermission = (body: { id?: number; enabled?: boolean }) =>
  post<null>('/admin/permission/enable', body);

/** 全部权限 */
export const getAdminPermissionAll = () => post<AdminPermissionInfo[]>('/admin/permission/all', {});

/** 权限绑定接口 */
export const bindAdminPermissionApis = (body: ReqAdminPermissionBindApis) =>
  post<null>('/admin/permission/bind_apis', body);

/** 权限解绑接口 */
export const unbindAdminPermissionApi = (body: { permissionId?: number; apiId?: number }) =>
  post<null>('/admin/permission/unbind_api', body);
