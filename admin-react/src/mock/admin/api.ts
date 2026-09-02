import { ok, fail, paginate, getPagination, now } from '../common';
import type { MockEntry } from '../plugin';
import { apiData } from '../data';
import type { AdminApiItem } from '@/types/common';

let rows = apiData.map((item) => ({ ...item })) as AdminApiItem[];

const entries: MockEntry[] = [
  {
    url: '/api/admin/api/list',
    method: 'POST',
    response: (req) => {
      const { body } = req;
      const { pageNum, pageSize } = getPagination(req);
      let list = rows;
      if (body.key) list = list.filter((i) => i.key?.includes(body.key));
      if (body.name) list = list.filter((i) => i.name?.includes(body.name));
      if (body.path) list = list.filter((i) => i.path?.includes(body.path));
      const data = paginate(list, pageNum, pageSize);
      return ok({ list: data.list, pageInfo: { total: data.total, pageNum: data.pageNum, pageSize: data.pageSize } });
    },
  },
  {
    url: '/api/admin/api/info',
    method: 'POST',
    response: ({ body }) => {
      const item = rows.find((i) => i.id === body.id);
      if (!item) return fail('接口不存在');
      return ok({ ...item });
    },
  },
  {
    url: '/api/admin/api/add',
    method: 'POST',
    response: ({ body }) => {
      rows.push({
        id: rows.length + 1,
        key: body.key,
        name: body.name,
        path: body.path,
        describe: body.describe || '',
        isEnabled: true,
        createdAt: now(),
        updatedAt: now(),
      });
      return ok(null, '创建成功');
    },
  },
  {
    url: '/api/admin/api/edit',
    method: 'POST',
    response: ({ body }) => {
      const item = rows.find((i) => i.id === body.id);
      if (!item) return fail('接口不存在');
      Object.assign(item, {
        key: body.key,
        name: body.name,
        path: body.path,
        describe: body.describe,
        updatedAt: now(),
      });
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/api/delete',
    method: 'POST',
    response: ({ body }) => {
      rows = rows.filter((i) => i.id !== body.id);
      return ok(null, '删除成功');
    },
  },
  {
    url: '/api/admin/api/enable',
    method: 'POST',
    response: ({ body }) => {
      const item = rows.find((i) => i.id === body.id);
      if (!item) return fail('接口不存在');
      item.isEnabled = body.enabled;
      return ok(null, body.enabled ? '启用成功' : '禁用成功');
    },
  },
  {
    url: '/api/admin/api/all',
    method: 'POST',
    response: () => ok(rows),
  },
];

export default entries;
