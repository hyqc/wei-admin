import { POST } from '../config';

export type APIItemType = {
  url: string;
  method?: string;
};

export type APIType = {
  [key: string]: APIItemType;
};

export const APICommon = {
  upload: {
    url: '/admin/common/upload',
    method: POST,
  },
};

export const APIAccount = {
  login: {
    url: '/admin/account/login',
    method: POST,
  },
  logout: {
    url: '/admin/account/logout',
    method: POST,
  },
  detail: {
    url: '/admin/account/detail',
    method: POST,
  },
  edit: {
    url: '/admin/account/edit',
    method: POST,
  },
  password: {
    url: '/admin/account/password',
    method: POST,
  },
};

export const APIAdminUsers = {
  list: {
    url: '/admin/user/list',
    method: POST,
  },
  detail: {
    url: '/admin/user/detail',
    method: POST,
  },
  add: {
    url: '/admin/user/add',
    method: POST,
  },
  delete: {
    url: '/admin/user/delete',
    method: POST,
  },
  edit: {
    url: '/admin/user/edit',
    method: POST,
  },
  enable: {
    url: '/admin/user/enable',
    method: POST,
  },
  bindRoles: {
    url: '/admin/user/bindRoles',
    method: POST,
  },
};

export const APIAdminRoles = {
  list: {
    url: '/admin/role/list',
    method: POST,
  },
  detail: {
    url: '/admin/role/detail',
    method: POST,
  },
  add: {
    url: '/admin/role/add',
    method: POST,
  },
  delete: {
    url: '/admin/role/delete',
    method: POST,
  },
  edit: {
    url: '/admin/role/edit',
    method: POST,
  },
  enable: {
    url: '/admin/role/enable',
    method: POST,
  },
  all: {
    url: '/admin/role/all',
    method: POST,
  },
  bindPermissions: {
    url: '/admin/role/bindPermissions',
    method: POST,
  },
};

export const APIAdminMenus = {
  tree: {
    url: '/admin/menu/tree',
    method: POST,
  },
  list: {
    url: '/admin/menu/list',
    method: POST,
  },
  detail: {
    url: '/admin/menu/detail',
    method: POST,
  },
  add: {
    url: '/admin/menu/add',
    method: POST,
  },
  delete: {
    url: '/admin/menu/delete',
    method: POST,
  },
  edit: {
    url: '/admin/menu/edit',
    method: POST,
  },
  enable: {
    url: '/admin/menu/enable',
    method: POST,
  },
  all: {
    url: '/admin/menu/all',
    method: POST,
  },
  permissions: {
    url: '/admin/menu/permissions',
    method: POST,
  },
  pages: {
    url: '/admin/menu/pages',
    method: POST,
  },
  mode: {
    url: '/admin/menu/mode',
    method: POST,
  },
};

export const APIAdminPermissions = {
  list: {
    url: '/admin/permission/list',
    method: POST,
  },
  detail: {
    url: '/admin/permission/detail',
    method: POST,
  },
  add: {
    url: '/admin/permission/add',
    method: POST,
  },
  addMenuPermissions: {
    url: '/admin/permission/addMenuPermissions',
    method: POST,
  },
  delete: {
    url: '/admin/permission/delete',
    method: POST,
  },
  edit: {
    url: '/admin/permission/edit',
    method: POST,
  },
  enable: {
    url: '/admin/permission/enable',
    method: POST,
  },
  all: {
    url: '/admin/permission/all',
    method: POST,
  },
  bindApis: {
    url: '/admin/permission/bindApis',
    method: POST,
  },
  unbind: {
    url: '/admin/permission/unbind',
    method: POST,
  },
};

export const APIAdminAPIResources = {
  list: {
    url: '/admin/api/list',
    method: POST,
  },
  detail: {
    url: '/admin/api/detail',
    method: POST,
  },
  add: {
    url: '/admin/api/add',
    method: POST,
  },
  delete: {
    url: '/admin/api/delete',
    method: POST,
  },
  edit: {
    url: '/admin/api/edit',
    method: POST,
  },
  enable: {
    url: '/admin/api/enable',
    method: POST,
  },
  all: {
    url: '/admin/api/all',
    method: POST,
  },
};
