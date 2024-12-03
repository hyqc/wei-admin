// adminAPIs 管理员接口资源接口
import { request } from 'umi';
import { APIAdminAPIResources as APIAdminAPIs } from './api';
import { ResponseBodyType } from '../types';
import { ReqAdminApiAdd, ReqAdminApiDelete, ReqAdminApiEnable, ReqAdminApiInfo, ReqAdminApiList } from '@/proto/admin_ts/admin_api';

/************************************************************/
/**
 * 列表
 */
export async function adminAPIList(params?: ReqAdminApiList) {
  return request<ResponseBodyType>(APIAdminAPIs.list.url, {
    method: APIAdminAPIs.list.method,
    data: params,
  });
}

/************************************************************/
/**
 * 新增
 */
export async function adminAPIAdd(params: ReqAdminApiAdd) {
  return request<ResponseBodyType>(APIAdminAPIs.add.url, {
    method: APIAdminAPIs.add.method,
    data: params,
  });
}

/************************************************************/
/**
 * 详情
 */
export async function adminAPIDetail(params: ReqAdminApiInfo) {
  return request<ResponseBodyType>(APIAdminAPIs.info.url, {
    method: APIAdminAPIs.info.method,
    data: params,
  });
}

/************************************************************/
/**
 * 编辑
 */
export async function adminAPIEdit(params: ReqAdminApiEnable) {
  return request<ResponseBodyType>(APIAdminAPIs.edit.url, {
    method: APIAdminAPIs.edit.method,
    data: params,
  });
}

/************************************************************/
/**
 * 全部
 */
export async function adminAPIAll(params?: {}) {
  return request<ResponseBodyType>(APIAdminAPIs.all.url, {
    method: APIAdminAPIs.all.method,
    data: params,
  });
}

/************************************************************/
/**
 * 删除
 */
export async function adminAPIDelete(params: ReqAdminApiDelete) {
  return request<ResponseBodyType>(APIAdminAPIs.delete.url, {
    method: APIAdminAPIs.delete.method,
    data: params,
  });
}

/************************************************************/
/**
 * 启用禁用
 */
export async function adminAPIEnable(params: ReqAdminApiEnable) {
  return request<ResponseBodyType>(APIAdminAPIs.enable.url, {
    method: APIAdminAPIs.enable.method,
    data: params,
  });
}
