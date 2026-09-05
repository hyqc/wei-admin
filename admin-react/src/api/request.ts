import axios, { type AxiosError, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { message } from 'antd';
import { SUCCESS, IsAuthForbiddenCode, IsAuthTokenInvalidCode } from './code';
import { IsIgnoreAuthApi } from './config';
import type { ResponseBodyType } from './types';
import { GetLoginToken, IsLoginPage, Logout } from '@/utils/common';

const instance = axios.create({
  baseURL: '/api',
  timeout: 30000,
});

// 请求拦截器：非登录页时注入 token
instance.interceptors.request.use((config) => {
  const token = GetLoginToken();
  if (token && !IsLoginPage()) {
    config.headers.set('Authorization', `Bearer ${token.token}`);
  }
  return config;
});

/** 请求错误：携带后端业务错误码 */
export type RequestError = Error & { code?: number };

// 响应拦截器：统一处理业务码
instance.interceptors.response.use(
  (response) => {
    const resp = response.data as ResponseBodyType;
    if (resp.code !== SUCCESS) {
      // 无权限：仅提示，不退出登录
      if (!IsAuthForbiddenCode(resp.code) && IsAuthTokenInvalidCode(resp.code)) {
        message.error('登录状态已过期，请重新登录');
        Logout();
      } else {
        message.error(resp.msg || resp.message || '请求失败');
      }
      // 带上业务错误码，便于调用方区分“无权限”与“令牌失效”
      const err: RequestError = new Error(resp.msg || resp.message || '请求失败');
      err.code = resp.code;
      return Promise.reject(err);
    }
    // 业务成功后直接返回数据体，由 request 函数断言类型
    return resp as unknown as AxiosResponse;
  },
  (error: AxiosError) => {
    if (error.response?.status === 401 && !IsIgnoreAuthApi(error.config?.url)) {
      // 免鉴权接口不强制登出：后端已放行，401 通常来自业务层判定
      message.error('登录状态已过期，请重新登录');
      Logout();
    } else {
      const resp = error.response?.data as ResponseBodyType | undefined;
      message.error(resp?.msg || resp?.message || error.message || '网络请求失败');
    }
    return Promise.reject(error);
  },
);

/** 统一请求方法（默认 POST，与后端接口保持一致） */
export function request<T = unknown>(url: string, options?: AxiosRequestConfig) {
  return instance.request<T>({ url, method: 'POST', ...options }) as unknown as Promise<ResponseBodyType<T>>;
}

/** GET 请求 */
export function get<T = unknown>(url: string, params?: Record<string, unknown>, options?: AxiosRequestConfig) {
  return instance.get(url, { params, ...options }) as unknown as Promise<ResponseBodyType<T>>;
}

/** POST 请求 */
export function post<T = unknown>(url: string, data?: unknown, options?: AxiosRequestConfig) {
  return instance.post(url, data, options) as unknown as Promise<ResponseBodyType<T>>;
}

/** DELETE 请求 */
export function remove<T = unknown>(url: string, params?: Record<string, unknown>, options?: AxiosRequestConfig) {
  return instance.delete(url, { params, ...options }) as unknown as Promise<ResponseBodyType<T>>;
}

export default { get, post, remove, request };
