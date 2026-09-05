import type { MenuTreeItem } from '@/types/admin_menu';
import type { RoleItem } from '@/types/admin_role';
import type { AdminUserListItem } from '@/types/common';
import type { AdminApiItem } from '@/types/common';
import type { AdminInfo } from '@/types/admin_account';

/** 菜单树数据 */
export const menuTreeData: MenuTreeItem[] = [
  {
    id: 1,
    level: 1,
    key: 'Admin',
    name: '系统设置',
    parentId: 0,
    describe: '系统设置',
    path: '/admin',
    redirect: '/',
    component: './Admin',
    sort: 0,
    icon: 'SettingOutlined',
    hideChildrenInMenu: false,
    hideInMenu: false,
    enabled: true,
    createTime: 1648605804,
    modifyTime: 1648605804,
    children: [
      {
        id: 2,
        level: 2,
        key: 'AdminUser',
        name: '账号管理',
        parentId: 1,
        describe: '账号列表',
        path: '/admin/user',
        redirect: '/',
        component: './Admin/User',
        sort: 0,
        icon: '',
        hideChildrenInMenu: true,
        hideInMenu: false,
        enabled: true,
        createTime: 1648605804,
        modifyTime: 1648605804,
      },
      {
        id: 3,
        level: 2,
        key: 'AdminRole',
        name: '角色管理',
        parentId: 1,
        describe: '角色列表',
        path: '/admin/role',
        redirect: '/',
        component: './Admin/Role',
        sort: 0,
        icon: '',
        hideChildrenInMenu: true,
        hideInMenu: false,
        enabled: true,
        createTime: 1648605804,
        modifyTime: 1648605804,
      },
      {
        id: 4,
        level: 2,
        key: 'AdminMenu',
        name: '菜单管理',
        parentId: 1,
        describe: '菜单列表',
        path: '/admin/menu',
        redirect: '/',
        component: './Admin/Menu',
        sort: 0,
        icon: '',
        hideChildrenInMenu: true,
        hideInMenu: false,
        enabled: true,
        createTime: 1648605804,
        modifyTime: 1648605804,
        children: [
          {
            id: 7,
            level: 3,
            key: 'AdminMenuAdd',
            name: '添加菜单',
            parentId: 4,
            describe: '添加菜单',
            path: '/admin/menu/add',
            redirect: '/',
            component: '',
            sort: 0,
            icon: '',
            hideChildrenInMenu: true,
            hideInMenu: true,
            enabled: true,
            createTime: 1648605804,
            modifyTime: 1648605804,
          },
          {
            id: 8,
            level: 3,
            key: 'AdminMenuEdit',
            name: '编辑菜单',
            parentId: 4,
            describe: '编辑菜单',
            path: '/admin/menu/edit',
            redirect: '/',
            component: '',
            sort: 0,
            icon: '',
            hideChildrenInMenu: true,
            hideInMenu: true,
            enabled: true,
            createTime: 1648605804,
            modifyTime: 1648605804,
          },
        ],
      },
      {
        id: 5,
        level: 2,
        key: 'AdminPermission',
        name: '权限管理',
        parentId: 1,
        describe: '权限管理',
        path: '/admin/permission',
        redirect: '/',
        component: './Admin/Permission',
        sort: 0,
        icon: '',
        hideChildrenInMenu: true,
        hideInMenu: false,
        enabled: true,
        createTime: 1648605804,
        modifyTime: 1648605804,
      },
      {
        id: 6,
        level: 2,
        key: 'AdminApi',
        name: '接口管理',
        parentId: 1,
        describe: '接口管理',
        path: '/admin/api',
        redirect: '/',
        component: './Admin/Api',
        sort: 0,
        icon: '',
        hideChildrenInMenu: false,
        hideInMenu: false,
        enabled: true,
        createTime: 1648605804,
        modifyTime: 1648605804,
      },
    ],
  },
];

/** 角色数据 */
export const roleData: RoleItem[] = [
  {
    id: 1,
    name: '管理员',
    describe: '拥有全部权限',
    createAdminId: 1,
    createAdminName: 'admin',
    isEnabled: true,
    createdAt: '2022-08-11 02:09:58',
    updatedAt: '2022-08-11 02:09:59',
  },
];

/** 用户数据 */
export const userData: AdminUserListItem[] = [
  {
    adminId: 1,
    username: 'admin',
    nickname: '骑着八戒游天河',
    email: 'ddd@q1.com',
    avatar: '',
    loginTotal: 39,
    lastLoginIp: '127.0.0.1',
    currentLoginIp: '127.0.0.1',
    lastLoginTime: '2024-06-23 11:31:49',
    currentLoginTime: '2024-06-23 11:31:49',
    isEnabled: true,
    createdAt: '2024-06-23 03:31:49',
    updatedAt: '2024-06-23 11:31:49',
    roles: [],
  },
  {
    adminId: 5,
    username: 'test00001',
    nickname: '测试00001',
    email: '',
    avatar: '',
    loginTotal: 0,
    lastLoginIp: '',
    lastLoginTime: '',
    isEnabled: true,
    createdAt: '2023-12-04 19:31:25',
    updatedAt: '2023-12-04 19:31:25',
    roles: [],
  },
  {
    adminId: 7,
    username: 'test00002',
    nickname: '测试00002',
    email: '',
    avatar: '',
    loginTotal: 0,
    lastLoginIp: '',
    lastLoginTime: '',
    isEnabled: true,
    createdAt: '2023-12-04 19:31:26',
    updatedAt: '2023-12-04 19:31:26',
    roles: [],
  },
  {
    adminId: 8,
    username: 'test00003',
    nickname: '测试00003',
    email: '',
    avatar: '',
    loginTotal: 0,
    lastLoginIp: '',
    lastLoginTime: '',
    isEnabled: true,
    createdAt: '2023-12-04 19:31:27',
    updatedAt: '2023-12-04 19:31:27',
    roles: [],
  },
  {
    adminId: 9,
    username: 'test00004',
    nickname: '测试00004',
    email: 'test0004@qq.com',
    avatar: '',
    loginTotal: 19,
    lastLoginIp: '',
    lastLoginTime: '2022-08-11 02:17:29',
    isEnabled: true,
    createdAt: '2023-12-04 19:31:30',
    updatedAt: '2023-12-04 19:31:30',
    roles: [{ roleId: 1, roleName: '管理员' }],
  },
];

/** 接口资源数据 */
export const apiData: AdminApiItem[] = [
  { id: 1, key: 'adminUser::list', name: '账号列表', path: '/admin/user/list', describe: '账号列表', isEnabled: true },
  { id: 2, key: 'adminUser::add', name: '账号创建', path: '/admin/user/add', describe: '账号创建', isEnabled: true },
  { id: 3, key: 'adminUser::detail', name: '账号详情', path: '/admin/user/info', describe: '账号详情', isEnabled: true },
  { id: 4, key: 'adminUser::edit', name: '账号编辑', path: '/admin/user/edit', describe: '账号编辑', isEnabled: true },
  { id: 5, key: 'adminUser::enable', name: '账号启用禁用', path: '/admin/user/enable', describe: '账号启用禁用', isEnabled: true },
  { id: 6, key: 'adminUser::delete', name: '账号删除', path: '/admin/user/delete', describe: '账号删除', isEnabled: true },
  { id: 7, key: 'adminUser::bindRoles', name: '账号绑定角色', path: '/admin/user/bind_roles', describe: '账号绑定角色', isEnabled: true },
  { id: 8, key: 'adminRole::list', name: '角色列表', path: '/admin/role/list', describe: '角色列表', isEnabled: true },
  { id: 9, key: 'adminRole::add', name: '角色创建', path: '/admin/role/add', describe: '', isEnabled: true },
  { id: 10, key: 'adminRole::detail', name: '角色详情', path: '/admin/role/info', describe: '', isEnabled: true },
  { id: 11, key: 'adminRole::edit', name: '角色编辑', path: '/admin/role/edit', describe: '', isEnabled: true },
  { id: 12, key: 'adminRole::enable', name: '角色禁用启用', path: '/admin/role/enable', describe: '', isEnabled: true },
  { id: 13, key: 'adminRole::delete', name: '角色删除', path: '/admin/role/delete', describe: '', isEnabled: true },
  { id: 14, key: 'adminRole::bindPermissions', name: '角色绑定权限', path: '/admin/role/bind_permissions', describe: '', isEnabled: true },
  { id: 15, key: 'adminRole::permissions', name: '角色权限列表', path: '/admin/role/permissions', describe: '', isEnabled: true },
  { id: 16, key: 'adminMenu::tree', name: '菜单树', path: '/admin/menu/tree', describe: '', isEnabled: true },
  { id: 17, key: 'adminMenu::list', name: '菜单列表', path: '/admin/menu/list', describe: '', isEnabled: true },
  { id: 18, key: 'adminMenu::add', name: '菜单创建', path: '/admin/menu/add', describe: '', isEnabled: true },
  { id: 19, key: 'adminMenu::detail', name: '菜单详情', path: '/admin/menu/info', describe: '', isEnabled: true },
  { id: 20, key: 'adminMenu::edit', name: '菜单编辑', path: '/admin/menu/edit', describe: '', isEnabled: true },
  { id: 21, key: 'adminMenu::enable', name: '菜单禁用启用', path: '/admin/menu/enable', describe: '', isEnabled: true },
  { id: 22, key: 'adminMenu::delete', name: '菜单删除', path: '/admin/menu/delete', describe: '', isEnabled: true },
  { id: 23, key: 'adminMenu::permissions', name: '菜单权限', path: '/admin/menu/permissions', describe: '', isEnabled: true },
  { id: 24, key: 'adminMenu::pages', name: '菜单页面列表', path: '/admin/menu/pages', describe: '', isEnabled: true },
  { id: 25, key: 'adminPermission::list', name: '权限列表', path: '/admin/permission/list', describe: '', isEnabled: true },
  { id: 26, key: 'adminPermission::add', name: '权限创建', path: '/admin/permission/add', describe: '', isEnabled: true },
  { id: 27, key: 'adminPermission::detail', name: '权限详情', path: '/admin/permission/info', describe: '', isEnabled: true },
  { id: 28, key: 'adminPermission::edit', name: '权限编辑', path: '/admin/permission/edit', describe: '', isEnabled: true },
  { id: 29, key: 'adminPermission::enable', name: '权限禁用启用', path: '/admin/permission/enable', describe: '', isEnabled: true },
  { id: 30, key: 'adminPermission::delete', name: '权限删除', path: '/admin/permission/delete', describe: '', isEnabled: true },
  { id: 31, key: 'adminPermission::addMenuPermissions', name: '权限指定菜单批量创建权限', path: '/admin/permission/add_menu_permissions', describe: '给指定的菜单创建查看，编辑，删除权限', isEnabled: true },
  { id: 32, key: 'adminPermission::bindApis', name: '权限绑定接口', path: '/admin/permission/bind_apis', describe: '', isEnabled: true },
  { id: 33, key: 'adminApi::list', name: '接口列表', path: '/admin/api/list', describe: '', isEnabled: true },
  { id: 34, key: 'adminApi::add', name: '接口创建', path: '/admin/api/add', describe: '', isEnabled: true },
  { id: 35, key: 'adminApi::detail', name: '接口详情', path: '/admin/api/info', describe: '', isEnabled: true },
  { id: 36, key: 'adminApi::edit', name: '接口编辑', path: '/admin/api/edit', describe: '', isEnabled: true },
  { id: 37, key: 'adminApi::enable', name: '接口禁用启用', path: '/admin/api/enable', describe: '', isEnabled: true },
  { id: 38, key: 'adminApi::delete', name: '接口删除', path: '/admin/api/delete', describe: '', isEnabled: true },
  { id: 39, key: 'adminApi::all', name: '接口全部', path: '/admin/api/all', describe: '全部有效接口列表', isEnabled: true },
  { id: 40, key: 'adminRole::all', name: '角色全部', path: '/admin/role/all', describe: '全部有效的角色列表', isEnabled: true },
  { id: 41, key: 'adminMenu::mode', name: '菜单页面权限列表', path: '/admin/menu/modes', describe: '', isEnabled: true },
  { id: 43, key: 'adminUser::editPwd', name: '账号重置密码', path: '/admin/user/edit_pwd', describe: '账号重置密码', isEnabled: true },
  { id: 44, key: 'adminMenu::show', name: '菜单显示隐藏', path: '/admin/menu/show', describe: '菜单显示隐藏', isEnabled: true },
];

/** 权限数据（与菜单关联） */
export interface MockPermission {
  id: number;
  menuId: number;
  menuName: string;
  menuPath: string;
  key: string;
  name: string;
  type: string;
  typeText: string;
  describe: string;
  enabled: boolean;
  apiIds: number[];
}

export const permissionData: MockPermission[] = [
  { id: 1, menuId: 2, menuName: '账号管理', menuPath: '/admin/user', key: 'AdminUserView', name: '查看账号', type: 'view', typeText: '查看', describe: '查看账号列表与详情', enabled: true, apiIds: [1, 3] },
  { id: 2, menuId: 2, menuName: '账号管理', menuPath: '/admin/user', key: 'AdminUserAdd', name: '新建账号', type: 'edit', typeText: '编辑', describe: '创建账号', enabled: true, apiIds: [2] },
  { id: 3, menuId: 2, menuName: '账号管理', menuPath: '/admin/user', key: 'AdminUserEdit', name: '编辑账号', type: 'edit', typeText: '编辑', describe: '编辑账号与启用/禁用', enabled: true, apiIds: [4, 5] },
  { id: 4, menuId: 2, menuName: '账号管理', menuPath: '/admin/user', key: 'AdminUserResetPwd', name: '重置密码', type: 'edit', typeText: '编辑', describe: '重置账号登录密码', enabled: true, apiIds: [43] },
  { id: 5, menuId: 2, menuName: '账号管理', menuPath: '/admin/user', key: 'AdminUserBindRoles', name: '绑定角色', type: 'edit', typeText: '编辑', describe: '为账号绑定/解绑角色', enabled: true, apiIds: [7, 40] },
  { id: 6, menuId: 2, menuName: '账号管理', menuPath: '/admin/user', key: 'AdminUserDelete', name: '删除账号', type: 'delete', typeText: '删除', describe: '删除账号', enabled: true, apiIds: [6] },
  { id: 7, menuId: 3, menuName: '角色管理', menuPath: '/admin/role', key: 'AdminRoleView', name: '查看角色', type: 'view', typeText: '查看', describe: '查看角色列表与详情', enabled: true, apiIds: [8, 10, 15, 41] },
  { id: 8, menuId: 3, menuName: '角色管理', menuPath: '/admin/role', key: 'AdminRoleAdd', name: '新建角色', type: 'edit', typeText: '编辑', describe: '创建角色', enabled: true, apiIds: [9] },
  { id: 9, menuId: 3, menuName: '角色管理', menuPath: '/admin/role', key: 'AdminRoleEdit', name: '编辑角色', type: 'edit', typeText: '编辑', describe: '编辑角色与启用/禁用', enabled: true, apiIds: [11, 12] },
  { id: 10, menuId: 3, menuName: '角色管理', menuPath: '/admin/role', key: 'AdminRoleBindPermissions', name: '绑定权限', type: 'edit', typeText: '编辑', describe: '为角色绑定/解绑权限', enabled: true, apiIds: [14, 15, 41] },
  { id: 11, menuId: 3, menuName: '角色管理', menuPath: '/admin/role', key: 'AdminRoleDelete', name: '删除角色', type: 'delete', typeText: '删除', describe: '删除角色', enabled: true, apiIds: [13] },
  { id: 12, menuId: 4, menuName: '菜单管理', menuPath: '/admin/menu', key: 'AdminMenuView', name: '查看菜单', type: 'view', typeText: '查看', describe: '查看菜单列表与详情', enabled: true, apiIds: [16, 17, 19] },
  { id: 13, menuId: 4, menuName: '菜单管理', menuPath: '/admin/menu', key: 'AdminMenuAdd', name: '新建菜单', type: 'edit', typeText: '编辑', describe: '创建菜单', enabled: true, apiIds: [18] },
  { id: 14, menuId: 4, menuName: '菜单管理', menuPath: '/admin/menu', key: 'AdminMenuEdit', name: '编辑菜单', type: 'edit', typeText: '编辑', describe: '编辑/启停/显示隐藏菜单与权限配置', enabled: true, apiIds: [20, 21, 23, 31, 44] },
  { id: 15, menuId: 4, menuName: '菜单管理', menuPath: '/admin/menu', key: 'AdminMenuDelete', name: '删除菜单', type: 'delete', typeText: '删除', describe: '删除菜单', enabled: true, apiIds: [22] },
  { id: 16, menuId: 5, menuName: '权限管理', menuPath: '/admin/permission', key: 'AdminPermissionView', name: '查看权限', type: 'view', typeText: '查看', describe: '查看权限列表与详情', enabled: true, apiIds: [25, 27, 16] },
  { id: 17, menuId: 5, menuName: '权限管理', menuPath: '/admin/permission', key: 'AdminPermissionAdd', name: '新建权限', type: 'edit', typeText: '编辑', describe: '创建权限', enabled: true, apiIds: [26, 24] },
  { id: 18, menuId: 5, menuName: '权限管理', menuPath: '/admin/permission', key: 'AdminPermissionEdit', name: '编辑权限', type: 'edit', typeText: '编辑', describe: '编辑/启停权限与绑定接口', enabled: true, apiIds: [28, 29, 32, 39] },
  { id: 19, menuId: 5, menuName: '权限管理', menuPath: '/admin/permission', key: 'AdminPermissionDelete', name: '删除权限', type: 'delete', typeText: '删除', describe: '删除权限', enabled: true, apiIds: [30] },
  { id: 20, menuId: 6, menuName: '接口管理', menuPath: '/admin/api', key: 'AdminApiView', name: '查看接口', type: 'view', typeText: '查看', describe: '查看接口列表与详情', enabled: true, apiIds: [33, 35] },
  { id: 21, menuId: 6, menuName: '接口管理', menuPath: '/admin/api', key: 'AdminApiAdd', name: '新建接口', type: 'edit', typeText: '编辑', describe: '创建接口', enabled: true, apiIds: [34] },
  { id: 22, menuId: 6, menuName: '接口管理', menuPath: '/admin/api', key: 'AdminApiEdit', name: '编辑接口', type: 'edit', typeText: '编辑', describe: '编辑接口与启用/禁用', enabled: true, apiIds: [36, 37] },
  { id: 23, menuId: 6, menuName: '接口管理', menuPath: '/admin/api', key: 'AdminApiDelete', name: '删除接口', type: 'delete', typeText: '删除', describe: '删除接口', enabled: true, apiIds: [38] },
];

/** 由菜单树构建登录返回的菜单 map */
export function buildMenusMap(tree: MenuTreeItem[]): Record<string, any> {
  const map: Record<string, any> = {};
  const walk = (list: MenuTreeItem[]) => {
    list.forEach((item) => {
      map[item.key!] = {
        id: item.id,
        key: item.key,
        name: item.name,
        parentId: item.parentId,
        path: item.path,
        redirect: item.redirect,
        icon: item.icon,
        hideChildrenInMenu: item.hideChildrenInMenu,
        hideInMenu: item.hideInMenu,
      };
      if (item.children?.length) {
        walk(item.children);
      }
    });
  };
  walk(tree);
  return map;
}

/** 构建登录返回的权限 map */
export function buildPermissionsMap(permissions: MockPermission[]): Record<string, string> {
  const map: Record<string, string> = {};
  permissions.forEach((item) => {
    map[item.key] = item.name;
  });
  return map;
}

/** 当前登录用户 */
export const currentAdminUserDetail: AdminInfo = {
  adminId: 1,
  username: 'admin',
  nickname: '骑着八戒游天河',
  avatar: '',
  email: 'ddd@q1.com',
  loginTotal: 39,
  lastLoginIp: '127.0.0.1, 127.0.0.1',
  lastLoginTime: '2024-06-23 11:31:49',
  isEnabled: true,
  createdAt: '2024-06-23 03:31:49',
  updatedAt: '2024-06-23 11:31:49',
  token:
    'eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJhZG1pbiIsImNyZWF0ZWQiOjE2NjAzOTY5MDc3NDYsImV4cCI6MTY2MTAwMTcwNywidmVyc2lvbiI6IlsxMCwgLTEyLCAtMTI3LCAtNDYsIC03LCAxMCwgMTA0LCAtMTksIDQ1LCAxMTYsIC03OCwgLTQ4LCAtNTAsIC00NCwgLTU2LCAtMTI2XSJ9.2xyB-TRM9NARwABmmADbJTZqlgALEXlgzFxV99prwvD-Ng_OV1JaeQK9c_oRjbU3vDQOHuy-3YRbjNpC6AJ2nA',
  /** 过期时间戳（秒），7 天后过期 */
  expire: Math.floor(Date.now() / 1000) + 604800,
  menus: buildMenusMap(menuTreeData),
  permissions: buildPermissionsMap(permissionData),
  roles: [{ roleId: 1, roleName: '管理员' }],
};
