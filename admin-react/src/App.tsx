import { useEffect } from 'react';
import { Navigate, Route, Routes, useLocation } from 'react-router-dom';
import { HomePath, LoginPath } from '@/api/config';
import { getRouteTitle } from '@/router/menu';
import ProtectedLayout from '@/layout/ProtectedLayout';
import Login from '@/views/login';
import Home from '@/views/home';
import Account from '@/views/account';
import AdminUser from '@/views/admin/user';
import AdminRole from '@/views/admin/role';
import AdminMenu from '@/views/admin/menu';
import AdminPermission from '@/views/admin/permission';
import AdminApi from '@/views/admin/api';
import Forbidden from '@/views/error/403';
import NotFound from '@/views/error/404';

export default function App() {
  const location = useLocation();

  // 路由变化后同步浏览器标题（与页签一致，取菜单名称）
  useEffect(() => {
    const title = getRouteTitle(location.pathname);
    document.title = `${title} - Admin React`;
  }, [location.pathname]);

  return (
    <Routes>
      <Route path={LoginPath} element={<Login />} />
      <Route element={<ProtectedLayout />}>
        <Route index element={<Navigate to={HomePath} replace />} />
        <Route path="home" element={<Home />} />
        <Route path="account" element={<Account />} />
        <Route path="admin/user" element={<AdminUser />} />
        <Route path="admin/role" element={<AdminRole />} />
        <Route path="admin/menu" element={<AdminMenu />} />
        <Route path="admin/permission" element={<AdminPermission />} />
        <Route path="admin/api" element={<AdminApi />} />
      </Route>
      <Route path="/403" element={<Forbidden />} />
      <Route path="/404" element={<NotFound />} />
      <Route path="*" element={<Navigate to="/404" replace />} />
    </Routes>
  );
}
