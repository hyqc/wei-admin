import { AdminInfo } from '@/proto/admin_ts/admin_account';
import type { ReponseCurrentAdminUserDetailType } from '@/services/apis/admin/account';

const AvatarImage =
  'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png';


const currentAdminUserDetail: AdminInfo = {
  adminId: 1,
  username: 'admin',
  nickname: '超管',
  avatar: AvatarImage,
  email: 'ddd@q1.com',
  createTime: '2022-08-13 21:21:46',
  modifyTime: '2022-08-13 21:21:46',
  lastLoginTime: '2022-08-13 21:21:47',
  lastLoginIp: '127.0.0.1',
  loginTotal: 23,
  enabled: true,
  token:
    'eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJhZG1pbiIsImNyZWF0ZWQiOjE2NjAzOTY5MDc3NDYsImV4cCI6MTY2MTAwMTcwNywidmVyc2lvbiI6IlsxMCwgLTEyLCAtMTI3LCAtNDYsIC03LCAxMCwgMTA0LCAtMTksIDQ1LCAxMTYsIC03OCwgLTQ4LCAtNTAsIC00NCwgLTU2LCAtMTI2XSJ9.2xyB-TRM9NARwABmmADbJTZqlgALEXlgzFxV99prwvD-Ng_OV1JaeQK9c_oRjbU3vDQOHuy-3YRbjNpC6AJ2nA',
  expire: 604800,
  menus: {
    AdminMenu: {
      id: 4,
      key: 'AdminMenu',
      name: '菜单管理',
      parentId: 1,
      path: '/admin/menu',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: true,
      hideInMenu: false,
    },
    AdminMenuAdd: {
      id: 7,
      key: 'AdminMenuAdd',
      name: '添加菜单',
      parentId: 4,
      path: '/admin/menu/add',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: true,
      hideInMenu: true,
    },
    AdminPermission: {
      id: 5,
      key: 'AdminPermission',
      name: '权限管理',
      parentId: 1,
      path: '/admin/permission',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: true,
      hideInMenu: false,
    },
    // AdminUser: {
    //   id: 2,
    //   key: 'AdminUser',
    //   name: '账号管理',
    //   parentId: 1,
    //   path: '/admin/user',
    //   redirect: '/',
    //   icon: '',
    //   hideChildrenInMenu: true,
    //   hideInMenu: false,
    // },
    AdminRole: {
      id: 3,
      key: 'AdminRole',
      name: '角色管理',
      parentId: 1,
      path: '/admin/role',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: true,
      hideInMenu: false,
    },
    Admin: {
      id: 1,
      key: 'Admin',
      name: '系统设置',
      parentId: 0,
      path: '/admin',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: false,
      hideInMenu: false,
    },
    AdminApi: {
      id: 6,
      key: 'AdminApi',
      name: '接口管理',
      parentId: 1,
      path: '/admin/api',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: false,
      hideInMenu: false,
    },
    AdminMenuEdit: {
      id: 8,
      key: 'AdminMenuEdit',
      name: '编辑菜单',
      parentId: 4,
      path: '/admin/menu/edit',
      redirect: '/',
      icon: '',
      hideChildrenInMenu: true,
      hideInMenu: true,
    },
  },
  permissions: {
    AdminUserEdit: '账号编辑',
    AdminRoleDelete: '角色删除',
    AdminUserDelete: '账号删除',
    AdminApiDelete: '接口删除',
    AdminPermissionDelete: '权限删除',
    AdminRoleView: '角色查看',
    AdminRoleEdit: '角色编辑',
    AdminApiEdit: '接口编辑',
    AdminUserView: '账号查看',
    AdminMenuDelete: '菜单删除',
    AdminApiView: '接口查看',
    AdminPermissionView: '权限查看',
    AdminMenuView: '菜单查看',
    AdminPermissionEdit: '权限编辑',
    AdminMenuEdit: '菜单编辑',
  },
};

const LoginData = {
  code: 0,
  type: 'SUCCESS',
  message: '登录成功',
  data: currentAdminUserDetail,
};

export { currentAdminUserDetail };

// 代码中会兼容本地 service mock 以及部署站点的静态数据
export default {
  'POST /api/admin/account/login': LoginData,
};
