import { ok, fail, paginate, getPagination, now } from '../common';
import type { MockEntry } from '../plugin';
import { userData, roleData } from '../data';

let rows = userData.map((item) => ({ ...item }));

function findUser(adminId?: number) {
  return rows.find((item) => item.adminId === adminId);
}

const entries: MockEntry[] = [
  {
    url: '/api/admin/user/list',
    method: 'POST',
    response: (req) => {
      const { body } = req;
      const { pageNum, pageSize } = getPagination(req);
      let list = rows;
      if (body.username) list = list.filter((i) => i.username?.includes(body.username));
      if (body.nickname) list = list.filter((i) => i.nickname?.includes(body.nickname));
      if (body.email) list = list.filter((i) => i.email?.includes(body.email));
      const data = paginate(list, pageNum, pageSize);
      return ok({ list: data.list, pageInfo: { total: data.total, pageNum: data.pageNum, pageSize: data.pageSize } });
    },
  },
  {
    url: '/api/admin/user/info',
    method: 'POST',
    response: ({ body }) => {
      const user = findUser(body.adminId);
      if (!user) return fail('账号不存在');
      const roleIds = user.roles?.map((r) => r.roleId) || [];
      return ok({ ...user, roleIds });
    },
  },
  {
    url: '/api/admin/user/add',
    method: 'POST',
    response: ({ body }) => {
      if (rows.some((i) => i.username === body.username)) {
        return fail('账号已存在');
      }
      const adminId = rows.length + 1;
      rows.push({
        adminId,
        username: body.username,
        nickname: body.nickname || '',
        email: body.email || '',
        avatar: '',
        loginTotal: 0,
        lastLoginIp: '',
        lastLoginTime: '',
        isEnabled: true,
        createdAt: now(),
        updatedAt: now(),
        roles: (body.roleIds || []).map((id: number) => {
          const role = roleData.find((r) => r.id === id);
          return { roleId: id, roleName: role?.name || '' };
        }),
      });
      return ok(null, '创建成功');
    },
  },
  {
    url: '/api/admin/user/edit',
    method: 'POST',
    response: ({ body }) => {
      const user = findUser(body.adminId);
      if (!user) return fail('账号不存在');
      if (body.nickname !== undefined) user.nickname = body.nickname;
      if (body.email !== undefined) user.email = body.email;
      user.updatedAt = now();
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/user/delete',
    method: 'POST',
    response: ({ body }) => {
      const user = findUser(body.adminId);
      if (!user) return fail('账号不存在');
      if (user.username === 'admin') return fail('admin 账号不允许删除');
      if (user.isEnabled) return fail('启用中的账号不允许删除，请先禁用');
      rows = rows.filter((i) => i.adminId !== body.adminId);
      return ok(null, '删除成功');
    },
  },
  {
    url: '/api/admin/user/enable',
    method: 'POST',
    response: ({ body }) => {
      const user = findUser(body.adminId);
      if (!user) return fail('账号不存在');
      user.isEnabled = body.enabled;
      user.updatedAt = now();
      return ok(null, body.enabled ? '启用成功' : '禁用成功');
    },
  },
  {
    url: '/api/admin/user/edit_pwd',
    method: 'POST',
    response: ({ body }) => {
      const user = findUser(body.adminId);
      if (!user) return fail('账号不存在');
      return ok(null, '重置成功');
    },
  },
  {
    url: '/api/admin/user/bind_roles',
    method: 'POST',
    response: ({ body }) => {
      const user = findUser(body.adminId);
      if (!user) return fail('账号不存在');
      user.roles = (body.roleIds || []).map((id: number) => {
        const role = roleData.find((r) => r.id === id);
        return { roleId: id, roleName: role?.name || '' };
      });
      user.updatedAt = now();
      return ok(null, '绑定成功');
    },
  },
];

export default entries;
