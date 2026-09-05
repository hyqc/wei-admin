import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { Spin } from 'antd';
import BasicLayout from './BasicLayout';
import { useUserStore } from '@/store/user';
import { LoginPath } from '@/api/config';
import { IsAuthForbiddenCode } from '@/api/code';
import { IsLogin, Logout } from '@/utils/common';

/**
 * 登录守卫：未登录跳转登录页；
 * 已登录但内存无用户信息（刷新页面）时，用 token 恢复菜单/权限，失败则退出。
 */
export default function ProtectedLayout() {
  const navigate = useNavigate();
  const location = useLocation();
  const menus = useUserStore((s) => s.menus);
  const fetchUserInfo = useUserStore((s) => s.fetchUserInfo);
  const [ready, setReady] = useState(false);

  useEffect(() => {
    let canceled = false;
    async function check() {
      if (!IsLogin()) {
        navigate(`${LoginPath}?redirect=${encodeURIComponent(location.pathname + location.search)}`, {
          replace: true,
        });
        return;
      }
      if (Object.keys(menus).length === 0) {
        try {
          await fetchUserInfo();
        } catch (e) {
          // 没有访问权限时仅提示并跳 403，不退出登录
          const code = (e as { code?: number })?.code;
          if (IsAuthForbiddenCode(code)) {
            navigate('/403', { replace: true });
            return;
          }
          Logout();
          return;
        }
      }
      if (!canceled) setReady(true);
    }
    check();
    return () => {
      canceled = true;
    };
  }, [location.pathname, location.search, menus, fetchUserInfo, navigate]);

  if (!ready) {
    return (
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
        <Spin size="large" />
      </div>
    );
  }

  // 子路由由 BasicLayout 内部的 Outlet 渲染
  return <BasicLayout />;
}
