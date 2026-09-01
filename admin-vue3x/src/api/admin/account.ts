import { post } from '@/api/request';
import type { ResponseBodyType } from '@/api/types';
import type {
  ReqLogin,
  AdminInfo,
  ReqAccountEdit,
  ReqAccountPasswordEdit,
  CaptchaRespData,
} from '@/types/admin_account';

/** 登录 */
export const login = (body: ReqLogin) =>
  post<AdminInfo>('/admin/account/login', body) as Promise<ResponseBodyType<AdminInfo>>;

/** 登录图片验证码（位数与尺寸由后端配置决定） */
export const getCaptcha = () =>
  post<CaptchaRespData>('/admin/account/captcha') as Promise<ResponseBodyType<CaptchaRespData>>;

/** 退出登录 */
export const logout = (params?: unknown) => post<null>('/admin/account/logout', params);

/** 获取当前账号信息 */
export const currentAdminInfo = (refreshToken = false) =>
  post<AdminInfo>('/admin/account/info', { refreshToken });

/** 修改账号信息 */
export const currentAdminEdit = (body: ReqAccountEdit) => post<null>('/admin/account/edit', body);

/** 修改密码 */
export const currentAdminEditPassword = (body: ReqAccountPasswordEdit) =>
  post<null>('/admin/account/password', body);

/** 上传文件（1: image） */
export const upload = (data: { fileType: number; file: File }) => {
  const formData = new FormData();
  formData.append('file', data.file);
  formData.append('fileType', String(data.fileType));
  return post<null>('/admin/common/upload', formData);
};
