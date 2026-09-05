import { ok, fail, paginate, getPagination } from '../common';
import type { MockEntry } from '../plugin';
import { permissionData, apiData } from '../data';
import type { MockPermission } from '../data';

/** 权限类型中文名（与前端 DEFAULT_PERMISSION_TYPES / 后端 GetPermissionTypeName 保持一致） */
/** 权限类型中文名（固定三类，与前端 DEFAULT_PERMISSION_TYPES / 后端 AdminPermissionTypeTextMap 一致） */
const TYPE_TEXT_MAP: Record<string, string> = {
  view: '查看',
  edit: '编辑',
  delete: '删除',
};

/** 权限类型中文名；未知（自定义）类型回退为原值，与后端 GetAdminPermissionTypeText 保持一致 */
function typeTextOf(type?: string) {
  return (type && TYPE_TEXT_MAP[type]) || type || '';
}

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
      // 反查：只保留绑定了该接口的权限点
      if (body.apiId) list = list.filter((i) => i.apiIds.includes(body.apiId));
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
        typeText: typeTextOf(body.type),
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
      const menuId = list[0]?.menuId;
      const postedIds = new Set<number>();
      list.forEach((item: any) => {
        if (item.id) postedIds.add(item.id);
      });
      // 全量同步：删除该菜单下未在本次提交列表中的旧权限点（与后端 BatchAddPermissions 一致）
      if (menuId && postedIds.size > 0) {
        rows = rows.filter((r) => r.menuId !== menuId || postedIds.has(r.id!));
      }
      let nextId = rows.reduce((max, r) => Math.max(max, r.id || 0), 0) + 1;
      list.forEach((item: any) => {
        if (item.id) {
          const existing = rows.find((r) => r.id === item.id);
          if (existing) {
            Object.assign(existing, {
              menuId: item.menuId ?? existing.menuId,
              menuName: item.menuName ?? existing.menuName,
              menuPath: item.menuPath ?? existing.menuPath,
              key: item.key,
              name: item.name,
              type: item.type || existing.type,
              typeText: typeTextOf(item.type),
              describe: item.describe ?? existing.describe,
              enabled: item.enabled ?? existing.enabled,
              // 未提交 apiIds 表示不改动接口绑定
              apiIds: item.apiIds ?? existing.apiIds,
            });
            return;
          }
        }
        rows.push({
          id: nextId++,
          menuId: item.menuId,
          menuName: item.menuName || '',
          menuPath: item.menuPath || '',
          key: item.key,
          name: item.name,
          type: item.type || 'view',
          typeText: typeTextOf(item.type),
          describe: item.describe || '',
          enabled: item.enabled ?? true,
          apiIds: item.apiIds || [],
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
        typeText: body.type !== undefined ? typeTextOf(body.type) : p.typeText,
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
