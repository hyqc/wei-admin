import type { MockEntry } from './plugin';
import accountEntries from './admin/account';
import userEntries from './admin/user';
import roleEntries from './admin/role';
import menuEntries from './admin/menu';
import permissionEntries from './admin/permission';
import apiEntries from './admin/api';
import uploadEntries from './admin/upload';

/** 全部 mock 路由 */
export const mockEntries: MockEntry[] = [
  ...accountEntries,
  ...userEntries,
  ...roleEntries,
  ...menuEntries,
  ...permissionEntries,
  ...apiEntries,
  ...uploadEntries,
];
