import { post } from '@/api/request';
import type {
  ReqAdminUploadList,
  ReqAdminUploadDelete,
  ReqAdminUploadFile,
  ResponseAdminUploadListType,
  UploadItem,
} from '@/types/admin_upload';

/** 上传记录列表 */
export const getAdminUploadList = (body: ReqAdminUploadList) =>
  post<ResponseAdminUploadListType>('/admin/upload/list', body);

/** 上传文件（multipart/form-data：file + uploadGroup） */
export const uploadAdminFile = (data: ReqAdminUploadFile) => {
  const formData = new FormData();
  formData.append('file', data.file);
  formData.append('uploadGroup', data.uploadGroup);
  return post<UploadItem>('/admin/upload/upload', formData);
};

/** 删除上传记录（同时删除存储上的文件） */
export const deleteAdminUpload = (body: ReqAdminUploadDelete) => post<null>('/admin/upload/delete', body);
