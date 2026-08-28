/** 响应体统一结构 */
export interface ResponseBodyType<T = unknown> {
  code: number;
  message?: string;
  msg?: string;
  type?: string;
  data: T;
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
