// 接口返回结构
export interface ResponseBodyType {
  code: number;
  msg: string;
  reason: string;
  data?: any;
}

export interface PageInfoType {
  total: number;
  pageNum: number;
  pageSize: number;
}

export type PermissionsType = {
  [key: string]: string;
};


