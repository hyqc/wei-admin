import { AdminAPIKey, AdminRouterPath } from '@/services/common/pattern';

export const DEFAULT_PAGE_NO = 1;
export const DEFAULT_PAGE_SIZE = 10;
export const DEFAULT_RULES: any = {
  name: [{ required: true, type: 'string', message: '请输入接口资源名称!' }],
  key: [
    { required: true, type: 'string', message: '请输入接口资源唯一键!' },
    {
      pattern: AdminAPIKey,
      message: '格式为:[a-z][a-zA-Z0-9]*::[a-z][a-zA-Z0-9]*，示例：adminUser::add',
    },
  ],
  path: [{ required: true, pattern: AdminRouterPath, message: '请输入接口资源名称!' }],
};

/**
 * 把路径转为键名
 * @param path
 * @returns
 */
export function path2UpperCamelCase(path: string) {
  const names = path
    ?.split('/')
    .filter((name) => name.length > 0)
    .map((name) => name);

  return names
    .map((name: string, index: number) => {
      if (index === 0) {
        return name[0].toLowerCase() + name.substring(1);
      } else if (index === names.length - 1) {
        return '::' + name[0].toLowerCase() + name.substring(1);
      } else {
        return name[0].toUpperCase() + name.substring(1);
      }
    })
    .join('');
}
