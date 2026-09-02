import { ok, fail, paginate, getPagination } from '../common';
import type { MockEntry } from '../plugin';
import { permissionData, apiData } from '../data';
import type { MockPermission } from '../data';

let rows = permissionData.map((item) => ({ ...item })) as MockPermission[];

function buildListItem(p: MockPermission) {
  return {
    id: p.id,
    menuId: p.menuId,
    menuName: p.menuName,
    menuPath: p.menuPath,
    key: p.key,
    name: p.name,
    type: p.type,
    typeText: p.typeText,
    describe: p.describe,
    enabled: p.enabled,
    apis: apiData.filter((a) => p.apiIds.includes(a.id!)),
  };
}

const entries: MockEntry[] = [
  {
    url: '/api/admin/permission/list',
    method: 'POST',
    response: (req) => {
      const { body } = req;
      const { pageNum, pageSize } = getPagination(req);
      let list = rows;
      if (body.menuId) list = list.filter((i) => i.menuId === body.menuId);
      if (body.key) list = list.filter((i) => i.key?.includes(body.key));
      if (body.name) list = list.filter((i) => i.name?.includes(body.name));
      if (body.type) list = list.filter((i) => i.type === body.type);
      const data = paginate(list, pageNum, pageSize);
      return ok({
        list: data.list.map(buildListItem),
        pageInfo: { total: data.total, pageNum: data.pageNum, pageSize: data.pageSize },
      });
    },
  },
  {
    url: '/api/admin/permission/info',
    method: 'POST',
    response: ({ body }) => {
      const p = rows.find((i) => i.id === body.id);
      if (!p) return fail('权限不存在');
      return ok({
        ...buildListItem(p),
        permissionId: p.id,
        apiIds: p.apiIds,
      });
    },
  },
  {
    url: '/api/admin/permission/add',
    method: 'POST',
    response: ({ body }) => {
      const id = rows.length + 1;
      rows.push({
        id,
        menuId: body.menuId,
        menuName: body.menuName || '',
        menuPath: body.menuPath || '',
        key: body.key,
        name: body.name,
        type: body.type || 'view',
        typeText: body.type === 'view' ? '查看' : body.type === 'edit' ? '编辑' : body.type === 'delete' ? '删除' : '添加',
        describe: body.describe || '',
        enabled: body.enabled ?? true,
        apiIds: [],
      });
      return ok(null, '创建成功');
    },
  },
  {
    url: '/api/admin/permission/add_menu_permissions',
    method: 'POST',
    response: ({ body }) => {
      const list = Array.isArray(body) ? body : (body as any).permissions || [];
      list.forEach((item: any) => {
        rows.push({
          id: rows.length + 1,
          menuId: item.menuId,
          menuName: item.menuName || '',
          menuPath: item.menuPath || '',
          key: item.key,
          name: item.name,
          type: item.type || 'view',
          typeText: item.type === 'view' ? '查看' : item.type === 'edit' ? '编辑' : '删除',
          describe: item.describe || '',
          enabled: item.enabled ?? true,
          apiIds: [],
        });
      });
      return ok(null, '创建成功');
    },
  },
  {
    url: '/api/admin/permission/edit',
    method: 'POST',
    response: ({ body }) => {
      const p = rows.find((i) => i.id === body.id);
      if (!p) return fail('权限不存在');
      Object.assign(p, {
        menuId: body.menuId ?? p.menuId,
        key: body.key,
        name: body.name,
        type: body.type ?? p.type,
        describe: body.describe,
        enabled: body.enabled ?? p.enabled,
      });
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/permission/delete',
    method: 'POST',
    response: ({ body }) => {
      rows = rows.filter((i) => i.id !== body.id);
      return ok(null, '删除成功');
    },
  },
  {
    url: '/api/admin/permission/enable',
    method: 'POST',
    response: ({ body }) => {
      const p = rows.find((i) => i.id === body.id);
      if (!p) return fail('权限不存在');
      p.enabled = body.enabled;
      return ok(null, body.enabled ? '启用成功' : '禁用成功');
    },
  },
  {
    url: '/api/admin/permission/all',
    method: 'POST',
    response: () => ok(rows.map(buildListItem)),
  },
  {
    url: '/api/admin/permission/bind_apis',
    method: 'POST',
    response: ({ body }) => {
      const p = rows.find((i) => i.id === body.permissionId);
      if (!p) return fail('权限不存在');
      p.apiIds = body.apiIds || [];
      return ok(null, '绑定成功');
    },
  },
  {
    url: '/api/admin/permission/unbind_api',
    method: 'POST',
    response: ({ body }) => {
      const p = rows.find((i) => i.id === body.permissionId);
      if (!p) return fail('权限不存在');
      p.apiIds = p.apiIds.filter((id) => id !== body.apiId);
      return ok(null, '解绑成功');
    },
  },
];

export default entries;
