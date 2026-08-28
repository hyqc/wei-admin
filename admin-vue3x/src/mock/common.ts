import type { MockRequest } from './plugin';

/** 成功响应 */
export const ok = (data: unknown = null, msg = '操作成功') => ({
  code: 0,
  msg,
  type: 'SUCCESS',
  data,
});

/** 失败响应 */
export const fail = (msg = '操作失败', data: unknown = null) => ({
  code: 500,
  msg,
  type: 'Error',
  data,
});

/** 提取分页参数（兼容 base 嵌套与平铺两种格式） */
export function getPagination(req: MockRequest) {
  const { body } = req;
  const base = body.base || {};
  const pageNum = Number(body.pageNum ?? base.pageNum ?? 1);
  const pageSize = Number(body.pageSize ?? base.pageSize ?? 10);
  return { pageNum, pageSize };
}

/** 对数组进行分页 */
export function paginate(list: any[], pageNum: number, pageSize: number) {
  const start = (pageNum - 1) * pageSize;
  return {
    list: list.slice(start, start + pageSize),
    total: list.length,
    pageNum,
    pageSize,
  };
}

/** 当前时间字符串 */
export const now = () => new Date().toLocaleString('zh-CN', { hour12: false });
