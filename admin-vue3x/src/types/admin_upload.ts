import type { PageInfoType } from '@/api/types';

/** 存储驱动：local 本地 | aliyun 阿里云OSS | qcloud 腾讯云COS | s3 亚马逊S3 */
export type UploadDriverType = 'local' | 'aliyun' | 'qcloud' | 's3';

/** 上传记录列表请求 */
export interface ReqAdminUploadList {
  pageNum?: number;
  pageSize?: number;
  createStartTime?: number;
  createEndTime?: number;
  /** 分组（上传路径前缀），按前缀匹配 */
  uploadGroup?: string;
  /** 原始文件名（模糊匹配） */
  originalName?: string;
  /** 扩展名 */
  ext?: string;
  /** 存储驱动 */
  driver?: string;
}

/** 上传记录 */
export interface UploadItem {
  id?: number;
  adminId?: number;
  adminName?: string;
  driver?: string;
  driverText?: string;
  /** 分组（上传路径前缀） */
  uploadGroup?: string;
  /** 存储对象键（相对路径） */
  objectKey?: string;
  /** 访问链接 */
  url?: string;
  originalName?: string;
  newName?: string;
  ext?: string;
  mime?: string;
  /** 文件大小（字节） */
  size?: number;
  sizeText?: string;
  md5?: string;
  /** 上传日期 */
  uploadDate?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 上传记录列表返回 */
export interface ResponseAdminUploadListType {
  list: UploadItem[];
  pageInfo: PageInfoType;
}

/** 删除上传记录请求 */
export interface ReqAdminUploadDelete {
  id?: number;
}

/** 上传文件请求：分组为上传路径前缀，如 /admin/user/ */
export interface ReqAdminUploadFile {
  file: File;
  uploadGroup: string;
}

/** 驱动下拉选项 */
export const UPLOAD_DRIVER_OPTIONS: { value: string; label: string }[] = [
  { value: 'local', label: '本地' },
  { value: 'aliyun', label: '阿里云OSS' },
  { value: 'qcloud', label: '腾讯云COS' },
  { value: 's3', label: '亚马逊S3' },
];
