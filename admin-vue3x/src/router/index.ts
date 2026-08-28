import { createRouter, createWebHistory } from 'vue-router';
import { IsLogin, Logout } from '@/utils/common';
import { HomePath, LoginPath } from '@/api/config';
import { useUserStore } from '@/store/user';

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { title: '登录' },
  },
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/layout/BasicLayout.vue'),
    redirect: HomePath,
    children: [
      {
        path: '/home',
        name: 'Home',
        component: () => import('@/views/home/index.vue'),
        meta: { title: '首页' },
      },
      {
        path: '/account',
        name: 'Account',
        component: () => import('@/views/account/index.vue'),
        meta: { title: '个人中心' },
      },
      {
        path: '/admin/user',
        name: 'AdminUser',
        component: () => import('@/views/admin/user/index.vue'),
        meta: { title: '账号管理' },
      },
      {
        path: '/admin/role',
        name: 'AdminRole',
        component: () => import('@/views/admin/role/index.vue'),
        meta: { title: '角色管理' },
      },
      {
        path: '/admin/menu',
        name: 'AdminMenu',
        component: () => import('@/views/admin/menu/index.vue'),
        meta: { title: '菜单管理' },
      },
      {
        path: '/admin/permission',
        name: 'AdminPermission',
        component: () => import('@/views/admin/permission/index.vue'),
        meta: { title: '权限管理' },
      },
      {
        path: '/admin/api',
        name: 'AdminApi',
        component: () => import('@/views/admin/api/index.vue'),
        meta: { title: '接口管理' },
      },
    ],
  },
  {
    path: '/403',
    name: 'Forbidden',
    component: () => import('@/views/error/403.vue'),
    meta: { title: '403' },
  },
  {
    path: '/404',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue'),
    meta: { title: '404' },
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _from, next) => {
  if (IsLogin()) {
    if (to.path === LoginPath) {
      next({ path: HomePath });
      return;
    }
    // 刷新页面后 Pinia 内存为空，用 token 恢复用户信息（菜单/权限）
    const store = useUserStore();
    if (Object.keys(store.menus).length === 0) {
      try {
        await store.fetchUserInfo();
      } catch {
        Logout();
        next({ path: LoginPath, query: { redirect: to.fullPath } });
        return;
      }
    }
    next();
    return;
  }
  if (to.path === LoginPath) {
    next();
    return;
  }
  next({ path: LoginPath, query: { redirect: to.fullPath } });
});

router.afterEach((to) => {
  const title = to.meta?.title;
  document.title = title ? `${title} - Admin Vue3x` : 'Admin Vue3x';
});

export default router;
