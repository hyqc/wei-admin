import { request } from 'umi';
import { ResponseBodyType } from '../types';
import { APIAccount, APICommon } from './api';
import { MenuDataItem } from '@ant-design/pro-components';
import { ReqAccountEdit, ReqAccountPasswordEdit, ReqLogin, RespLoginData } from '@/proto/admin_ts/admin_account';

export type MenusRemoteItem = {
  [key: string]: MenuDataItem;
};

export async function login(params: ReqLogin)  {
  return request<ResponseBodyType>(APIAccount.login.url, {
    method: APIAccount.login.method,
    data: params,
  });
}

export async function logout(params?: any) {
  return request<ResponseBodyType>(APIAccount.logout.url, {
    method: APIAccount.logout.method,
    data: params,
  });
}

export async function currentAdminInfo(refreshToken?: boolean) {
  return request<ResponseBodyType>(APIAccount.info.url, {
    method: APIAccount.info.method,
    data: {
      refreshToken: refreshToken || false,
    },
  });
}

// 1:image
export type RequestUploadFileParamsType = {
  fileType: number;
  file: File;
};

export async function upload(data: RequestUploadFileParamsType) {
  return request<ResponseBodyType>(APICommon.upload.url, {
    method: APICommon.upload.method,
    data: data,
  });
}

export async function currentAdminEdit(params?: ReqAccountEdit) {
  return request<ResponseBodyType>(APIAccount.edit.url, {
    method: APIAccount.edit.method,
    data: params,
  });
}

export async function currentAdminEditPassword(params?: ReqAccountPasswordEdit) {
  return request<ResponseBodyType>(APIAccount.password.url, {
    method: APIAccount.password.method,
    data: params,
  });
}
