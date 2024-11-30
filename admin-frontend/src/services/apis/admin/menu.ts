// adminMenus 菜单管理接口
import { request } from 'umi';
import { APIAdminMenus } from './api';
import type { ResponseBodyType } from '../types';
import { ReqAdminMenuAdd, ReqAdminMenuDelete, ReqAdminMenuEdit, ReqAdminMenuEnable, ReqAdminMenuInfo, ReqAdminMenuList, ReqAdminMenuPages, ReqAdminMenuPermissions, ReqAdminMenuShow, ReqMenuList } from '@/proto/admin_ts/admin_menu';

/************************************************************/
/**
 * 树形菜单列表全
 */
export async function adminMenuTree() {
  return request<ResponseBodyType>(APIAdminMenus.tree.url, {
    method: APIAdminMenus.tree.method,
    data: {},
  });
}

/************************************************************/
/**
 * 列表
 */
export async function adminMenuList(params?: ReqAdminMenuList) {
  return request<ResponseBodyType>(APIAdminMenus.list.url, {
    method: APIAdminMenus.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 新增
 */
export async function adminMenuAdd(params: ReqAdminMenuAdd) {
  return request<ResponseBodyType>(APIAdminMenus.add.url, {
    method: APIAdminMenus.add.method,
    data: params,
  });
}

/************************************************************/
/**
 * 详情
 */
export async function adminMenuDetail(params: ReqAdminMenuInfo) {
  return request<ResponseBodyType>(APIAdminMenus.info.url, {
    method: APIAdminMenus.info.method,
    data: params,
  });
}
/************************************************************/
/**
 * 编辑
 */
export async function adminMenuEdit(params: ReqAdminMenuEdit) {
  return request<ResponseBodyType>(APIAdminMenus.edit.url, {
    method: APIAdminMenus.edit.method,
    data: params,
  });
}

/************************************************************/
/**
 * 删除
 */
export async function adminMenuDelete(params: ReqAdminMenuDelete) {
  return request<ResponseBodyType>(APIAdminMenus.delete.url, {
    method: APIAdminMenus.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export async function adminMenuEnable(params: ReqAdminMenuEnable) {
  return request<ResponseBodyType>(APIAdminMenus.enable.url, {
    method: APIAdminMenus.enable.method,
    data: params,
  });
}

/************************************************************/
/**
 * 菜单显示隐藏
 * @returns 
 */
export async function adminMenuShow(params: ReqAdminMenuShow) {
  return request<ResponseBodyType>(APIAdminMenus.show.url, {
    method: APIAdminMenus.show.method,
    data: params,
  });
}

/************************************************************/
/**
 * 菜单的权限列表
 */

export async function adminMenuPermissions(params: ReqAdminMenuPermissions) {
  return request<ResponseBodyType>(APIAdminMenus.permissions.url, {
    method: APIAdminMenus.permissions.method,
    data: params,
  });
}

/************************************************************/
/**
 * 权限的菜单列表
 */
export async function adminMenuPages(params?: ReqAdminMenuPages) {
  return request<ResponseBodyType>(APIAdminMenus.pages.url, {
    method: APIAdminMenus.pages.method,
    data: {
      all: params?.all || false,
    },
  });
}

/************************************************************/
/**
 * 菜单模块
 * @returns 
 */
export async function adminMenuMode() {
  return request<ResponseBodyType>(APIAdminMenus.mode.url, {
    method: APIAdminMenus.mode.method,
    data: {},
  });
}
