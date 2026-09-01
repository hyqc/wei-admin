/** 分段转小驼峰：admin-user → adminUser，User → user */
function toLowerCamel(str: string): string {
  const segments = str.split(/[^a-zA-Z0-9]+/).filter(Boolean);
  if (segments.length === 0) return '';
  return segments
    .map((seg, index) => {
      const lower = seg.toLowerCase();
      if (index === 0) return lower.charAt(0).toLowerCase() + lower.slice(1);
      return lower.charAt(0).toUpperCase() + lower.slice(1);
    })
    .join('');
}

/**
 * 由接口路径生成唯一键：最后一个 `/` 后的部分作为 `::` 后的动作名，其余部分按小驼峰拼接成模块名
 * 例：/admin/user/list → adminUser::list；/admin/role/bindPermissions → adminRole::bindPermissions
 */
export function generateApiKeyByPath(path?: string): string {
  if (!path) return '';
  const parts = path
    .split('/')
    .map((item) => item.trim())
    .filter(Boolean);
  if (parts.length === 0) return '';
  // 动作名保留原有大小写（如 bindPermissions），仅首字母小写以匹配键名规则
  const rawAction = parts[parts.length - 1];
  const action = rawAction.charAt(0).toLowerCase() + rawAction.slice(1);
  const module = parts
    .slice(0, -1)
    .map(toLowerCamel)
    .map((seg, index) => (index === 0 ? seg : seg.charAt(0).toUpperCase() + seg.slice(1)))
    .join('');
  return module ? `${module}::${action}` : action;
}
