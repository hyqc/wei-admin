import { AdminPerssionKey } from '@/api/pattern';

export type PermissionTypesItemType = {
  key: string;
  name: string;
};

/**
 * 权限动作类型：固定三类，不可新增、不可自定义
 * 任何页面操作本质上只分 查看（读）/ 编辑（写，含新增、重置、绑定等一切写操作）/ 删除
 * 与后端 common.AdminPermissionEnumItems 保持一致
 */
export const DEFAULT_PERMISSION_TYPES: PermissionTypesItemType[] = [
  { key: 'view', name: '查看' },
  { key: 'edit', name: '编辑' },
  { key: 'delete', name: '删除' },
];

/** 菜单权限配置默认模板（首次配置预填，与后端 AdminPermissionEnumDefaultItems 一致） */
export const DEFAULT_PERMISSION_TEMPLATE_TYPES: PermissionTypesItemType[] = [...DEFAULT_PERMISSION_TYPES];

/** 类型下拉选项：只允许从固定三类中选择，不支持自定义输入 */
export const PERMISSION_TYPE_OPTIONS = DEFAULT_PERMISSION_TYPES.map((t) => ({ value: t.key, label: t.name }));

/** 权限表单校验规则 */
export const DEFAULT_RULES: Record<string, any> = {
  key: [{ required: true, pattern: AdminPerssionKey, message: '请按照驼峰法命名' }],
  name: [
    { required: true, type: 'string', message: '请添加权限名称' },
    {
      type: 'string',
      max: 50,
      message: '名称长度不能超过50个字符',
    },
  ],
  describe: [{ required: false, type: 'string', message: '请添加权限描述' }],
};

export const PERMIDDION_RULES = {
  view: DEFAULT_RULES,
  edit: DEFAULT_RULES,
  delete: DEFAULT_RULES,
};

/**
 * 把路径转为键名
 * @param path 菜单路径
 */
export function path2UpperCamelCase(path: string) {
  return path
    ?.split('/')
    .filter((name) => name.length > 0)
    .map((name) => name[0].toUpperCase() + name.substring(1))
    .join('');
}

/**
 * 权限类型转唯一键后缀：仅保留字母数字并首字母大写
 * 自定义类型（export）同样可用；中文等无法提取时返回空串，交由用户手工补全
 */
export function typeKeySuffix(type: string) {
  const cleaned = (type || '').replace(/[^a-zA-Z0-9]/g, '');
  if (!cleaned) return '';
  return cleaned[0].toUpperCase() + cleaned.slice(1);
}

/**
 * 根据路径与权限类型生成权限唯一键
 * @param path 菜单路径
 * @param type 权限类型（内置或自定义）
 */
export function handleKey(path: string, type: string) {
  return path2UpperCamelCase(path) + typeKeySuffix(type);
}

/** 接口树节点 key 前缀：api-{apiId} */
export const API_NODE_PREFIX = 'api-';

/**
 * 接口按 path 模块前缀分组（/admin/user/list → /admin/user），构建两级可勾选树
 * 不依赖权限绑定关系，未被任何权限绑定的接口也能展示
 * 分组标题：菜单名 (路径前缀)，如「接口管理 (/admin/api)」
 * 接口节点：接口名 (路径)，如「接口列表 (/admin/api/list)」
 * @param menuNameByPath 菜单 path → 菜单名，用于让分组标题可读
 */
export function buildApiTree(
  list: { id?: number; path?: string; name?: string; key?: string }[],
  menuNameByPath?: Map<string, string>,
) {
  const groups = new Map<string, { key: string; title: string; children: { key: string; title: string }[] }>();
  for (const api of list || []) {
    const path = api.path || '';
    const seg = path.split('/').filter(Boolean);
    const group = seg.length >= 2 ? `/${seg[0]}/${seg[1]}` : '/其他';
    if (!groups.has(group)) {
      const menuName = menuNameByPath?.get(group);
      groups.set(group, {
        key: `group-${group}`,
        title: menuName ? `${menuName} (${group})` : group,
        children: [],
      });
    }
    groups.get(group)!.children.push({
      key: `${API_NODE_PREFIX}${api.id}`,
      title: `${api.name} (${path})`,
    });
  }
  return [...groups.values()].sort((a, b) => a.title.localeCompare(b.title));
}

/** 菜单列表转 path → 菜单名映射，供接口树分组标题使用 */
export function buildMenuNameByPath(menus: { path?: string; name?: string }[]) {
  const map = new Map<string, string>();
  for (const menu of menus || []) {
    if (menu.path) {
      map.set(menu.path, menu.name || '');
    }
  }
  return map;
}

/** 从勾选的树节点 key 中解析出接口ID列表 */
export function parseCheckedApiIds(keys: (string | number)[]): number[] {
  const ids: number[] = [];
  for (const k of keys || []) {
    if (typeof k === 'string' && k.startsWith(API_NODE_PREFIX)) {
      const id = Number(k.slice(API_NODE_PREFIX.length));
      if (!Number.isNaN(id)) ids.push(id);
    }
  }
  return [...new Set(ids)];
}
