import type { ReqListBase } from './common';

/** 接口列表请求 */
export interface ReqAdminApiList extends ReqListBase {
  key?: string;
  name?: string;
  path?: string;
}

/** 创建接口 */
export interface ReqAdminApiAdd {
  key: string;
  name: string;
  path: string;
  describe?: string;
}

/** 修改接口 */
export interface ReqAdminApiEdit {
  id: number;
  key?: string;
  name?: string;
  path?: string;
  describe?: string;
}

/** 删除接口 */
export interface ReqAdminApiDelete {
  id?: number;
}

/** 启用禁用接口 */
export interface ReqAdminApiEnable {
  id?: number;
  enabled?: boolean;
}

/** 接口详情 */
export interface ReqAdminApiInfo {
  id?: number;
}

/** 接口资源 */
export interface AdminApiResourceItem {
  id?: number;
  key?: string;
  name?: string;
  path?: string;
  describe?: string;
  isEnabled?: boolean;
  createdAt?: string;
  updatedAt?: string;
}
