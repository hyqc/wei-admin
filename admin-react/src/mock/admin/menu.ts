import { ok, fail, paginate, getPagination } from '../common';
import type { MockEntry } from '../plugin';
import { menuTreeData, permissionData } from '../data';
import type { MenuTreeItem } from '@/types/admin_menu';

let tree = JSON.parse(JSON.stringify(menuTreeData)) as MenuTreeItem[];

function flatten(list: MenuTreeItem[]): MenuTreeItem[] {
  const result: MenuTreeItem[] = [];
  const walk = (items: MenuTreeItem[]) => {
    items.forEach((item) => {
      result.push(item);
      if (item.children?.length) walk(item.children);
    });
  };
  walk(list);
  return result;
}

const entries: MockEntry[] = [
  {
    url: '/api/admin/menu/tree',
    method: 'POST',
    response: () => ok({ list: tree }),
  },
  {
    url: '/api/admin/menu/list',
    method: 'POST',
    response: (req) => {
      const { pageNum, pageSize } = getPagination(req);
      const list = flatten(tree);
      return ok(paginate(list, pageNum, pageSize));
    },
  },
  {
    url: '/api/admin/menu/info',
    method: 'POST',
    response: ({ body }) => {
      const item = flatten(tree).find((i) => i.id === body.menuId);
      if (!item) return fail('菜单不存在');
      return ok({ ...item });
    },
  },
  {
    url: '/api/admin/menu/add',
    method: 'POST',
    response: ({ body }) => {
      const parent = flatten(tree).find((i) => i.id === body.parentId);
      const level = parent ? (parent.level || 0) + 1 : 1;
      const newNode: MenuTreeItem = {
        id: Date.now(),
        level,
        key: body.key,
        name: body.name,
        parentId: body.parentId || 0,
        describe: body.describe || '',
        path: body.path,
        redirect: body.redirect || '/',
        component: body.component || '',
        sort: body.sort ?? 0,
        icon: body.icon || '',
        hideChildrenInMenu: !!body.isHideChildrenInMenu,
        hideInMenu: !!body.isHideInMenu,
        enabled: body.isEnabled ?? true,
        createTime: Date.now() / 1000,
        modifyTime: Date.now() / 1000,
      };
      if (parent) {
        parent.children = parent.children || [];
        parent.children.push(newNode);
      } else {
        tree.push(newNode);
      }
      return ok(null, '创建成功');
    },
  },
  {
    url: '/api/admin/menu/edit',
    method: 'POST',
    response: ({ body }) => {
      const list = flatten(tree);
      const item = list.find((i) => i.id === body.id);
      if (!item) return fail('菜单不存在');
      Object.assign(item, {
        key: body.key,
        name: body.name,
        path: body.path,
        describe: body.describe,
        redirect: body.redirect,
        icon: body.icon,
        sort: body.sort,
        component: body.component,
        hideChildrenInMenu: !!body.isHideChildrenInMenu,
        hideInMenu: !!body.isHideInMenu,
        enabled: body.isEnabled ?? item.enabled,
        modifyTime: Date.now() / 1000,
      });
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/menu/delete',
    method: 'POST',
    response: ({ body }) => {
      const removeById = (list: MenuTreeItem[]): MenuTreeItem[] =>
        list.filter((i) => i.id !== body.menuId && i.children);
      const walk = (list: MenuTreeItem[]) => {
        list.forEach((i) => {
          i.children = removeById(i.children || []);
          if (i.children.length === 0) delete i.children;
        });
      };
      tree = tree.filter((i) => i.id !== body.menuId);
      walk(tree);
      return ok(null, '删除成功');
    },
  },
  {
    url: '/api/admin/menu/enable',
    method: 'POST',
    response: ({ body }) => {
      const item = flatten(tree).find((i) => i.id === body.menuId);
      if (!item) return fail('菜单不存在');
      item.enabled = body.enabled;
      return ok(null, body.enabled ? '启用成功' : '禁用成功');
    },
  },
  {
    url: '/api/admin/menu/show',
    method: 'POST',
    response: ({ body }) => {
      const item = flatten(tree).find((i) => i.id === body.menuId);
      if (!item) return fail('菜单不存在');
      item.hideInMenu = !body.show;
      return ok(null, '操作成功');
    },
  },
  {
    url: '/api/admin/menu/all',
    method: 'POST',
    response: () => ok(tree),
  },
  {
    url: '/api/admin/menu/permissions',
    method: 'POST',
    response: ({ body }) => {
      const menu = flatten(tree).find((i) => i.id === body.menuId);
      if (!menu) return fail('菜单不存在');
      const permissions = permissionData
        .filter((p) => p.menuId === body.menuId)
        .map((p) => ({
          id: p.id,
          name: p.name,
          key: p.key,
          type: p.type,
          typeName: p.typeText,
          describe: p.describe,
          enabled: p.enabled,
          menuId: p.menuId,
        }));
      return ok({ menu, permissions });
    },
  },
  {
    url: '/api/admin/menu/pages',
    method: 'POST',
    response: () => ok(tree),
  },
  {
    url: '/api/admin/menu/modes',
    method: 'POST',
    response: () => {
      const modes = tree.map((m) => ({
        modelId: m.id,
        modelName: m.name,
        pages: (m.children || []).map((child) => ({
          pageId: child.id,
          pageName: child.name,
          permissions: permissionData
            .filter((p) => p.menuId === child.id)
            .map((p) => ({
              permissionId: p.id,
              permissionName: p.name,
              permissionType: p.type,
              permissionTypeName: p.typeText,
            })),
        })),
      }));
      return ok({ modes });
    },
  },
];

export default entries;
