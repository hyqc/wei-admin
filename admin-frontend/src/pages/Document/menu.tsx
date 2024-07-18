import { MenuDataItem } from '@umijs/route-utils';
import { Menu } from 'antd';
import React from 'react';

export const DocMenus: MenuDataItem[] = [
  {
    name: '接口文档',
    icon: 'BookOutlined',
    disabled: true,
  },
  {
    name: '后台文档',
    icon: 'BookOutlined',
    disabled: true,
  },
  {
    name: '账号管理',
    path: '/doc/admin/user',
    icon: 'BookOutlined',
    component: './Document',
    uri: '/docs/admin/账号管理.md',
    paddingLeft: '1rem',
  },
  {
    name: '角色管理',
    path: '/doc/admin/role',
    icon: 'BookOutlined',
    component: './Document',
    uri: '/docs/admin/角色管理.md',
    paddingLeft: '1rem',
  },
  {
    name: '菜单管理',
    path: '/doc/admin/menu',
    icon: 'BookOutlined',
    component: './Document',
    uri: '/docs/admin/菜单管理.md',
    paddingLeft: '1rem',
  },
  {
    name: '权限管理',
    path: '/doc/admin/menu',
    icon: 'BookOutlined',
    component: './Document',
    uri: '/docs/admin/权限管理.md',
    paddingLeft: '1rem',
  },
  {
    name: '接口管理',
    path: '/doc/admin/api',
    icon: 'BookOutlined',
    component: './Document',
    uri: '/docs/admin/接口管理.md',
    paddingLeft: '1rem',
  },
  {
    name: '前台文档',
    icon: 'BookOutlined',
    disabled: true,
  },
  {
    name: '测试',
    icon: 'BookOutlined',
  },
];

export type MenuDocType = {
  changeDoc?: (path: string) => void;
};

const MenuDoc: React.FC<MenuDocType> = (props: any) => {
  const { changeDoc } = props;
  function openMd(path?: string) {
    changeDoc && changeDoc(path);
  }

  return (
    <Menu className="menu-item" theme="light" mode="inline" defaultSelectedKeys={['0']}>
      {DocMenus.map((item, index) => {
        return (
          <Menu.Item
            disabled={item.disabled}
            style={{ cursor: item.disabled ? 'text' : 'pointer' }}
            key={index}
            onClick={() => openMd(item?.uri)}
          >
            {item.name}
          </Menu.Item>
        );
      })}
    </Menu>
  );
};

export default MenuDoc;
