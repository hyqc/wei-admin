import { post } from '@/api/request';
import type {
  MenuTreeItem,
  AdminMenuModel,
  MenuModeItem,
  MenuPermissionItem,
  ReqAdminMenuList,
  ReqAdminMenuAdd,
  ReqAdminMenuEnable,
  ReqAdminMenuShow,
  ReqAdminMenuDelete,
  ReqAdminMenuInfo,
  ReqAdminMenuPermissions,
  ReqAdminMenuPages,
  ReqAdminMenuMode,
} from '@/types/admin_menu';

/** 菜单树返回 */
export interface ResponseAdminMenuTreeType {
  list: MenuTreeItem[];
}

/** 菜单权限返回 */
export interface ResponseAdminMenuPermissionsType {
  menu?: AdminMenuModel;
  permissions?: MenuPermissionItem[];
}

/** 页面菜单返回 */
export type ResponseAdminMenuPagesType = MenuTreeItem[];

/** 模块权限返回 */
export interface ResponseAdminMenuModeType {
  modes: MenuModeItem[];
}

/** 菜单树 */
export const getAdminMenuTree = () => post<ResponseAdminMenuTreeType>('/admin/menu/tree', {});

/** 菜单列表 */
export const getAdminMenuList = (body: ReqAdminMenuList) => post<{ list: MenuTreeItem[] }>('/admin/menu/list', body);

/** 菜单详情 */
export const getAdminMenuInfo = (body: ReqAdminMenuInfo) => post<AdminMenuModel>('/admin/menu/info', body);

/** 创建菜单 */
export const addAdminMenu = (body: ReqAdminMenuAdd) => post<null>('/admin/menu/add', body);

/** 修改菜单 */
export const editAdminMenu = (body: ReqAdminMenuAdd) => post<null>('/admin/menu/edit', body);

/** 删除菜单 */
export const deleteAdminMenu = (body: ReqAdminMenuDelete) => post<null>('/admin/menu/delete', body);

/** 启用禁用菜单 */
export const enableAdminMenu = (body: ReqAdminMenuEnable) => post<null>('/admin/menu/enable', body);

/** 是否显示菜单 */
export const showAdminMenu = (body: ReqAdminMenuShow) => post<null>('/admin/menu/show', body);

/** 全部菜单 */
export const getAdminMenuAll = () => post<MenuTreeItem[]>('/admin/menu/all', {});

/** 菜单权限列表 */
export const getAdminMenuPermissions = (body: ReqAdminMenuPermissions) =>
  post<ResponseAdminMenuPermissionsType>('/admin/menu/permissions', body);

/** 页面菜单列表 */
export const getAdminMenuPages = (body: ReqAdminMenuPages) =>
  post<ResponseAdminMenuPagesType>('/admin/menu/pages', body);

/** 页面模块权限列表 */
export const getAdminMenuMode = (body: ReqAdminMenuMode) => post<ResponseAdminMenuModeType>('/admin/menu/modes', body);
