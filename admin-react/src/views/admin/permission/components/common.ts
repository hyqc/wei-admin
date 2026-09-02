import { AdminPerssionKey } from '@/api/pattern';

export type PermissionTypesItemType = {
  key: string;
  name: string;
};

/** 权限类型 */
export const DEFAULT_PERMISSION_TYPES: PermissionTypesItemType[] = [
  { key: 'view', name: '查看' },
  { key: 'edit', name: '编辑' },
  { key: 'delete', name: '删除' },
];

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
 * 根据路径与权限类型生成权限唯一键
 * @param path 菜单路径
 * @param type 权限类型
 */
export function handleKey(path: string, type: string) {
  return path2UpperCamelCase(path) + (type[0].toUpperCase() + type.substring(1));
}
