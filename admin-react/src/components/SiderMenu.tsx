import { useMemo, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { Menu } from 'antd';
import type { MenuProps } from 'antd';
import { useUserStore } from '@/store/user';
import { getAntIcon } from '@/utils/icon';
import { HandleRemoteMenuIntoLocal, localMenuData, type MenuConfigItem } from '@/router/menu';
import { HomePath } from '@/api/config';

type MenuItem = Required<MenuProps>['items'][number];

/** 渲染菜单标题：有图标显示图标，无图标占位保持对齐 */
function renderTitle(name: string, icon?: string) {
  const Icon = getAntIcon(icon);
  return (
    <>
      {Icon ? <Icon /> : <span className="menu-icon-placeholder" aria-hidden="true" />}
      <span>{name}</span>
    </>
  );
}

function toMenuItems(list: MenuConfigItem[]): MenuItem[] {
  return list.map((item) => {
    if (!item.children || item.children.length === 0) {
      return { key: item.key, icon: <></>, label: renderTitle(item.name, item.icon) };
    }
    return {
      key: item.key,
      label: renderTitle(item.name, item.icon),
      children: toMenuItems(item.children),
    };
  });
}

/** 侧边菜单：本地菜单配置 + 后端返回的菜单过滤 */
export default function SiderMenu({ collapsed, onNavigate }: { collapsed: boolean; onNavigate?: () => void }) {
  const navigate = useNavigate();
  const location = useLocation();
  const menus = useUserStore((s) => s.menus);

  const menuItems = useMemo<MenuConfigItem[]>(() => {
    const remoteMenus = menus || {};
    const result = HandleRemoteMenuIntoLocal([], localMenuData, remoteMenus, 'children');
    // 顶部固定“首页”
    const homeItem: MenuConfigItem = { key: 'Home', path: HomePath, name: '首页', icon: 'HomeOutlined' };
    return [homeItem, ...result];
  }, [menus]);

  const items = useMemo(() => toMenuItems(menuItems), [menuItems]);

  // 当前选中项：按路径匹配
  const selectedKeys = useMemo(() => {
    const hit = menuItems.find((item) => item.path === location.pathname);
    if (hit) return [hit.key];
    const child = menuItems
      .flatMap((item) => item.children ?? [])
      .find((item) => item.path === location.pathname);
    return child ? [child.key] : [];
  }, [menuItems, location.pathname]);

  // 默认展开当前菜单所在分组
  const [openKeys, setOpenKeys] = useState<string[]>(() => {
    const parent = menuItems.find((item) => item.children?.some((child) => child.path === location.pathname));
    return parent ? [parent.key] : [];
  });

  const onClick: MenuProps['onClick'] = ({ key }) => {
    const hit =
      menuItems.find((item) => item.key === key) ??
      menuItems.flatMap((item) => item.children ?? []).find((item) => item.key === key);
    if (hit) {
      navigate(hit.path);
      onNavigate?.();
    }
  };

  return (
    <div className="sider-menu">
      <div className="logo" onClick={() => navigate(HomePath)}>
        <img src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg" alt="logo" />
        {!collapsed && <span>Admin React</span>}
      </div>
      <Menu
        theme="dark"
        mode="inline"
        items={items}
        selectedKeys={selectedKeys}
        openKeys={openKeys}
        onOpenChange={(keys) => setOpenKeys(keys as string[])}
        onClick={onClick}
      />
    </div>
  );
}
