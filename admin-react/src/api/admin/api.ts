import { post } from '@/api/request';
import type { PageInfoType } from '@/api/types';
import type { AdminApiItem } from '@/types/common';
import type {
  ReqAdminApiList,
  ReqAdminApiAdd,
  ReqAdminApiEdit,
  ReqAdminApiDelete,
  ReqAdminApiEnable,
  ReqAdminApiInfo,
} from '@/types/admin_api';

/** 接口列表返回 */
export interface ResponseAdminApiListType {
  list: AdminApiItem[];
  pageInfo: PageInfoType;
}

/** 接口列表 */
export const getAdminApiList = (body: ReqAdminApiList) => post<ResponseAdminApiListType>('/admin/api/list', body);

/** 接口详情 */
export const getAdminApiInfo = (body: ReqAdminApiInfo) => post<AdminApiItem>('/admin/api/info', body);

/** 创建接口 */
export const addAdminApi = (body: ReqAdminApiAdd) => post<null>('/admin/api/add', body);

/** 修改接口 */
export const editAdminApi = (body: ReqAdminApiEdit) => post<null>('/admin/api/edit', body);

/** 删除接口 */
export const deleteAdminApi = (body: ReqAdminApiDelete) => post<null>('/admin/api/delete', body);

/** 启用禁用接口 */
export const enableAdminApi = (body: ReqAdminApiEnable) => post<null>('/admin/api/enable', body);

/** 全部接口列表 */
export const getAdminApiAll = () => post<AdminApiItem[]>('/admin/api/all', {});
