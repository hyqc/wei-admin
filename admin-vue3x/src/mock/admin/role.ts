import { ok, fail, paginate, getPagination, now } from '../common';
import type { MockEntry } from '../plugin';
import { roleData, permissionData } from '../data';

let rows = roleData.map((item) => ({ ...item }));

function findRole(id?: number) {
  return rows.find((item) => item.id === id);
}

const entries: MockEntry[] = [
  {
    url: '/api/admin/role/list',
    method: 'POST',
    response: (req) => {
      const { body } = req;
      const { pageNum, pageSize } = getPagination(req);
      let list = rows;
      if (body.name) list = list.filter((i) => i.name?.includes(body.name));
      const data = paginate(list, pageNum, pageSize);
      return ok({ list: data.list, pageInfo: { total: data.total, pageNum: data.pageNum, pageSize: data.pageSize } });
    },
  },
  {
    url: '/api/admin/role/info',
    method: 'POST',
    response: ({ body }) => {
      const role = findRole(body.id);
      if (!role) return fail('角色不存在');
      return ok({ ...role });
    },
  },
  {
    url: '/api/admin/role/add',
    method: 'POST',
    response: ({ body }) => {
      rows.push({
        id: rows.length + 1,
        name: body.name,
        describe: body.describe || '',
        createAdminId: 1,
        createAdminName: 'admin',
        isEnabled: true,
        createdAt: now(),
        updatedAt: now(),
      });
      return ok(null, '创建成功');
    },
  },
  {
    url: '/api/admin/role/edit',
    method: 'POST',
    response: ({ body }) => {
      const role = findRole(body.id);
      if (!role) return fail('角色不存在');
      if (role.id === 1) return fail('管理员角色不允许编辑');
      if (body.name !== undefined) role.name = body.name;
      if (body.describe !== undefined) role.describe = body.describe;
      role.updatedAt = now();
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/role/delete',
    method: 'POST',
    response: ({ body }) => {
      rows = rows.filter((i) => i.id !== body.id);
      return ok(null, '删除成功');
    },
  },
  {
    url: '/api/admin/role/enable',
    method: 'POST',
    response: ({ body }) => {
      const role = findRole(body.id);
      if (!role) return fail('角色不存在');
      role.isEnabled = body.enabled;
      role.updatedAt = now();
      return ok(null, body.enabled ? '启用成功' : '禁用成功');
    },
  },
  {
    url: '/api/admin/role/all',
    method: 'POST',
    response: () => ok(rows),
  },
  {
    url: '/api/admin/role/bind_permissions',
    method: 'POST',
    response: ({ body }) => {
      const role = findRole(body.id);
      if (!role) return fail('角色不存在');
      if (role.id === 1) return fail('管理员角色不允许绑定权限');
      return ok(null, '绑定成功');
    },
  },
  {
    url: '/api/admin/role/permissions',
    method: 'POST',
    response: ({ body }) => {
      const role = findRole(body.id);
      if (!role) return fail('角色不存在');
      const list = permissionData
        .filter((p) => p.menuId !== undefined)
        .map((p) => ({
          roleId: body.id,
          permissionId: p.id,
          permissionName: p.name,
          permissionKey: p.key,
          permissionType: p.type,
          permissionTypeText: p.typeText,
        }));
      return ok({ list });
    },
  },
];

export default entries;
