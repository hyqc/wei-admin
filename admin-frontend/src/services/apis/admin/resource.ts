// adminAPIs 管理员接口资源接口
import { request } from 'umi';
import { APIAdminAPIResources as APIAdminAPIs } from './api';
import { ResponseBodyType, ResponseListType } from '../types';

/************************************************************/
/**
 * 列表
 */
export type RequestAdminAPIListParamsType = {
  path?: string;
  key?: string;
  anme?: string;
  enabled?: number;
  pageNum?: number;
  pageSize?: number;
  sortField?: string;
  sortType?: string;
};

export type ResponseAdminAPIListItemType = {
  id: number;
  key: string;
  path: string;
  name: string;
  enabled: boolean;
  enabledText?: string;
  createTime: string;
  modifyTime: string;
};

export async function adminAPIList(params?: RequestAdminAPIListParamsType) {
  return request<ResponseListType>(APIAdminAPIs.list.url, {
    method: APIAdminAPIs.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 新增
 */
export type RequestAdminAPIAddParamsType = {
  path: string;
  name: string;
  key: string;
  describe: string;
  enabled: boolean;
};

export async function adminAPIAdd(params: RequestAdminAPIAddParamsType) {
  return request<ResponseBodyType>(APIAdminAPIs.add.url, {
    method: APIAdminAPIs.add.method,
    data: params,
  });
}

/************************************************************/
/**
 * 详情
 */
export type RequestAdminAPIDetailParamsType = {
  id: number;
};

export type ResponseAdminAPIDetailType = {
  id: number;
  path: string;
  name: string;
  key: string;
  describe: string;
  enabled: boolean;
  createTime: string;
  modifyTime: string;
  [key: string]: any;
};

export async function adminAPIDetail(params: RequestAdminAPIDetailParamsType) {
  return request<ResponseBodyType>(APIAdminAPIs.detail.url, {
    method: APIAdminAPIs.detail.method,
    data: params,
  });
}

/************************************************************/
/**
 * 编辑
 */
export type RequestAdminAPIEditParamsType = {
  id: number;
  enabled: boolean;
  path?: string;
  name?: string;
  key?: string;
  describe?: string;
};

export async function adminAPIEdit(params: RequestAdminAPIEditParamsType) {
  return request<ResponseBodyType>(APIAdminAPIs.edit.url, {
    method: APIAdminAPIs.edit.method,
    data: params,
  });
}

/************************************************************/
/**
 * 全部
 */
export type RequestAdminAPIAllParamsType = {
  name?: string; // 接口资源名称
  path?: string; // 接口资源名称
  key?: string; // 接口资源名称
};

export type ResponseAdminAPIAllItemType = {
  permissionId: number;
  id: number;
  name: string;
  path?: string;
  key?: string;
};

export async function adminAPIAll(params?: RequestAdminAPIAllParamsType) {
  return request<ResponseBodyType>(APIAdminAPIs.all.url, {
    method: APIAdminAPIs.all.method,
    data: params,
  });
}

/************************************************************/
/**
 * 删除
 */

export type RequestAdminAPIDeleteParamsType = {
  id: number;
  enabled: boolean;
};

export async function adminAPIDelete(params: RequestAdminAPIDeleteParamsType) {
  return request<ResponseBodyType>(APIAdminAPIs.delete.url, {
    method: APIAdminAPIs.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export type RequestAdminAPIEnableParamsType = {
  id: number;
  enabled: boolean;
};

export async function adminAPIEnable(params: RequestAdminAPIEnableParamsType) {
  return request<ResponseBodyType>(APIAdminAPIs.enable.url, {
    method: APIAdminAPIs.enable.method,
    data: params,
  });
}
