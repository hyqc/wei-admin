// 接口返回结构
export interface ResponseBodyType {
  code: number;
  message: string;
  type?: string;
  data?: any;
}

export interface PageInfoType {
  total: number;
  pageNum: number;
  pageSize: number;
}

// 列表数据 data 部分
export interface ResponseListDataType extends PageInfoType {
  list: any[];
}

// 列表数据 data 部分
export interface ResponseDetailDataType {
  [key: string]: any;
}

// 列表返回结构
export interface ResponseListType extends ResponseBodyType {
  code: number;
  message: string;
  type?: string;
  data: ResponseListDataType;
}

// 详情返回结构
export interface ResponseDetailType extends ResponseBodyType {
  code: number;
  message: string;
  type?: string;
  data: ResponseDetailDataType;
}
