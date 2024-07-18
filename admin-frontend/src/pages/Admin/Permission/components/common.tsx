import { AdminPerssionKey } from '@/services/common/pattern';

export type PermissionTypesItemType = {
  key: string;
  name: string;
};
export const DEFAULT_PERMISSION_TYPES: PermissionTypesItemType[] = [
  { key: 'view', name: '查看' },
  { key: 'edit', name: '编辑' },
  { key: 'delete', name: '删除' },
];

export const DEFAULT_RULES: any = {
  key: [{ required: true, pattern: AdminPerssionKey, message: '请按照驼峰法命名' }],
  name: [
    { required: true, type: 'string', message: '请添加权限名称' },
    {
      type: 'string',
      max: '50',
      message: '名称长度不能超过50个字符',
    },
  ],
  describe: [{ required: false, type: 'string', message: '请添加权限名称' }],
};

export const PERMIDDION_RULES = {
  view: DEFAULT_RULES,
  edit: DEFAULT_RULES,
  delete: DEFAULT_RULES,
};

/**
 * 把路径转为键名
 * @param path
 * @returns
 */
export function path2UpperCamelCase(path: string) {
  return path
    ?.split('/')
    .filter((name) => name.length > 0)
    .map((name) => name[0].toUpperCase() + name.substring(1))
    .join('');
}
