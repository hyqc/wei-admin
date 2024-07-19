﻿import { layout } from "@/app";

export default [
  {
    key: 'Login',
    path: '/login',
    component: './Login',
    title: 'login',
    layout: false,
    access: true
  },
  {
    path: '/',
    redirect: '/home',
    key: 'Home',
    access: true
  },
  {
    path: '/home',
    name: 'home',
    icon: 'HomeOutlined',
    component: './Home',
    key: 'Home',
    access: true
  },
  {
    path: '/account',
    name: 'account',
    icon: 'UserOutlined',
    component: './Account',
    key: 'Account',
    access: true
  },
  {
    path: '/admin',
    key: 'Admin',
    name: 'admin',
    icon: 'SettingOutlined',
    routes: [
      {
        key: 'AdminUser',
        path: '/admin/user',
        name: 'user',
        icon: 'TeamOutlined',
        component: './Admin/User',
        hideInMenu: true,
      },
      {
        key: 'AdminRole',
        name: 'role',
        icon: 'UserSwitchOutlined',
        hideInMenu: true,
        path: '/admin/role',
        component: './Admin/Role',
      },
      {
        key: 'AdminMenu',
        name: 'menu',
        icon: 'MenuUnfoldOutlined',
        hideInMenu: true,
        path: '/admin/menu',
        parentKeys: ["Admin"],
        routes: [
          {
            key: 'AdminMenu',
            icon: 'MenuUnfoldOutlined',
            path: '/admin/menu',
            component: './Admin/Menu',
            hideInMenu: true,
          },
          {
            key: 'AdminMenuAdd',
            path: '/admin/menu/add',
            name: 'add',
            component: './Admin/Menu/add',
            hideInMenu: true,
          },
          {
            key: 'AdminMenuEdit',
            path: '/admin/menu/edit',
            name: 'edit',
            component: './Admin/Menu/edit',
            hideInMenu: true,
          },
        ],
      },
      {
        key: 'AdminPermission',
        path: '/admin/permission',
        name: 'permission',
        icon: 'UnlockOutlined',
        component: './Admin/Permission',
        hideInMenu: true,
      },
      {
        key: 'AdminApi',
        path: '/admin/api',
        name: 'api',
        icon: 'ApiOutlined',
        component: './Admin/Api',
        hideInMenu: true,
      },
      {
        component: './404',
        access: true,
        key: '404',
      },
    ],
  },
  {
    name: 'doc',
    path: '/doc',
    icon: 'BookOutlined',
    component: './Document',
    key: 'Doc',
    access: true,
    layout: false,
  },
  {
    name: 'demo',
    path: '/demo',
    icon: 'BookOutlined',
    component: './Demo',
    key: 'Demo',
    access: true,
  },
  {
    component: './404',
    key: '404',
    access: true,
  },
];