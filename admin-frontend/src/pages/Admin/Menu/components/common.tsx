import { ResponseAdminMenuListItemType } from '@/services/apis/admin/menu';
import { AdminMenuKey } from '@/services/common/pattern';

export const INPUT_NAME_LENGTH_MAX = 50;
export const FORM_RULES: any = {
  name: [{ required: true, type: 'string', message: '请输入菜单名称!' }],
  path: [{ required: true, type: 'string', message: '请输入菜单路由!' }],
  key: [
    { required: true, type: 'string', message: '请输入菜单键名!' },
    {
      required: true,
      pattern: AdminMenuKey,
      message: '请按照大驼峰法命名唯一键',
    },
  ],
  parentId: [{ required: true, type: 'number', message: '缺少父级菜单ID' }],
  sort: [{ required: true, type: 'number', message: '排序值为不小于0的整数' }],
};

/**
 *
 * @param item 生成铺开的菜单单元
 * @param level
 * @returns
 */
export function menuSpreadTreeItem(item: ResponseAdminMenuListItemType, level: number) {
  const data = {
    ...item,
    level: level,
    hasChildren: item.children !== undefined && item.children.length > 0,
  };
  data.children = [];
  return data;
}

/**
 * 递归的生成铺开的菜单父子有序列表
 * @param result
 * @param data
 * @param level
 * @param menuId
 * @returns
 */
export function menuSpreadTree(
  result: ResponseAdminMenuListItemType[],
  data: ResponseAdminMenuListItemType[],
  level: number,
  menuId?: number,
): ResponseAdminMenuListItemType[] {
  data?.map((item: ResponseAdminMenuListItemType) => {
    if (menuId !== undefined && item.menuId === menuId) {
      return;
    }
    result.push(menuSpreadTreeItem(item, level));
    if (item.children !== undefined && item.children.length > 0) {
      menuSpreadTree(result, item.children, level + 1, menuId);
    }
  });
  return result;
}

/**
 * 输出铺开后的菜单父子有序数据
 * @param data
 * @param menuId
 * @returns
 */
export function makeMenuSpreadTreeData(
  data: ResponseAdminMenuListItemType[],
  menuId?: number,
): ResponseAdminMenuListItemType[] {
  const result: ResponseAdminMenuListItemType[] = [];
  return menuSpreadTree(result, data, 0, menuId);
}

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
