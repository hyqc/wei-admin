import { ok, fail, paginate, getPagination } from '../common';
import type { MockEntry } from '../plugin';

/** 上传记录（mock：仅保存在内存中） */
interface MockUpload {
  id: number;
  adminId: number;
  adminName: string;
  driver: string;
  driverText: string;
  uploadGroup: string;
  objectKey: string;
  url: string;
  originalName: string;
  newName: string;
  ext: string;
  mime: string;
  size: number;
  sizeText: string;
  md5: string;
  uploadDate: string;
  createdAt: string;
  updatedAt: string;
}

const DRIVER_TEXT: Record<string, string> = {
  local: '本地',
  aliyun: '阿里云OSS',
  qcloud: '腾讯云COS',
  s3: '亚马逊S3',
};

let rows: MockUpload[] = [];
let nextId = 1;

/** 规范化分组：/admin/user/ → admin/user */
/**
 * 解析 multipart 请求体：mock 插件无法把 multipart 解析为 JSON，
 * 这里从原始文本中提取普通字段与文件名（二进制内容无需还原）
 */
function parseMultipart(rawBody: string, contentType: string) {
  const fields: Record<string, string> = { __fileName: '' };
  const match = /boundary=(?:"([^"]+)"|([^;]+))/.exec(contentType || '');
  if (!match) return fields;
  const boundary = (match[1] || match[2] || '').trim();
  if (!boundary) return fields;
  for (const part of rawBody.split(`--${boundary}`)) {
    const idx = part.indexOf('\r\n\r\n');
    if (idx < 0) continue;
    const head = part.slice(0, idx);
    const value = part.slice(idx + 4).replace(/\r\n$/, '');
    const nameMatch = /name="([^"]+)"/.exec(head);
    if (!nameMatch) continue;
    if (nameMatch[1] === 'file') {
      const fileMatch = /filename="([^"]+)"/.exec(head);
      if (fileMatch) fields.__fileName = fileMatch[1];
      continue;
    }
    fields[nameMatch[1]] = value;
  }
  return fields;
}

function normalizeGroup(raw: string) {
  return String(raw || '')
    .split('/')
    .filter(Boolean)
    .join('/');
}

function humanSize(size: number) {
  if (size < 1024) return `${size} B`;
  let value = size / 1024;
  const units = ['KB', 'MB', 'GB'];
  let idx = 0;
  while (value >= 1024 && idx < units.length - 1) {
    value /= 1024;
    idx += 1;
  }
  return `${value.toFixed(2)} ${units[idx]}`;
}

const entries: MockEntry[] = [
  {
    url: '/api/admin/upload/list',
    method: 'POST',
    response: (req) => {
      const { body } = req;
      const { pageNum, pageSize } = getPagination(req);
      let list = rows;
      if (body.uploadGroup) {
        list = list.filter((i) => i.uploadGroup.startsWith(normalizeGroup(body.uploadGroup)));
      }
      if (body.originalName) list = list.filter((i) => i.originalName.includes(body.originalName));
      if (body.ext) list = list.filter((i) => i.ext === String(body.ext).toLowerCase());
      if (body.driver) list = list.filter((i) => i.driver === body.driver);
      const data = paginate(list, pageNum, pageSize);
      return ok({
        list: data.list,
        pageInfo: { total: data.total, pageNum: data.pageNum, pageSize: data.pageSize },
      });
    },
  },
  {
    url: '/api/admin/upload/upload',
    method: 'POST',
    response: (req) => {
      const contentType = String(req.headers?.['content-type'] || '');
      // JSON 请求直接取字段；multipart 请求从原始文本中解析
      const isMultipart = contentType.includes('multipart/form-data');
      const fields = isMultipart ? parseMultipart(req.rawBody || '', contentType) : {};
      const group = normalizeGroup(
        isMultipart ? fields.uploadGroup || '' : (req.body?.uploadGroup as string) || '',
      );
      if (!group) return fail('请填写分组（上传路径前缀）');
      const name =
        (isMultipart ? fields.__fileName : (req.body?.originalName as string)) ||
        `mock-file-${nextId}.png`;
      const ext = name.includes('.') ? (name.split('.').pop() as string).toLowerCase() : '';
      const now = new Date();
      const stamp = `${now.getFullYear()}${String(now.getMonth() + 1).padStart(2, '0')}${String(now.getDate()).padStart(2, '0')}`;
      const ymd = `${stamp.slice(0, 4)}-${stamp.slice(4, 6)}-${stamp.slice(6, 8)}`;
      const newName = `${stamp}_${String(now.getHours()).padStart(2, '0')}${String(now.getMinutes()).padStart(2, '0')}${String(now.getSeconds()).padStart(2, '0')}_${Math.random().toString(16).slice(2, 10)}${ext ? '.' + ext : ''}`;
      const objectKey = [group, String(now.getFullYear()), String(now.getMonth() + 1).padStart(2, '0'), newName].join('/');
      const record: MockUpload = {
        id: nextId++,
        adminId: 1,
        adminName: 'admin',
        driver: 'local',
        driverText: DRIVER_TEXT.local,
        uploadGroup: group,
        objectKey,
        url: `/upload/${objectKey}`,
        originalName: name,
        newName,
        ext,
        mime: 'application/octet-stream',
        size: name.length,
        sizeText: humanSize(name.length),
        md5: '',
        uploadDate: ymd,
        createdAt: ymd,
        updatedAt: ymd,
      };
      rows = [record, ...rows];
      return ok(record, '上传成功');
    },
  },
  {
    url: '/api/admin/upload/delete',
    method: 'POST',
    response: ({ body }) => {
      const before = rows.length;
      rows = rows.filter((i) => i.id !== body.id);
      if (rows.length === before) return fail('上传记录不存在');
      return ok(null, '删除成功');
    },
  },
];

export default entries;
