import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { Dropdown, Tabs } from 'antd';
import type { MenuProps } from 'antd';
import { CloseOutlined, ColumnWidthOutlined, MinusOutlined } from '@ant-design/icons';
import { HomePath } from '@/api/config';
import { getRouteTitle } from '@/router/menu';
import { useTabsStore } from '@/store/tabs';

/** 多页签：路由变化时同步，支持右键关闭当前/其他/全部 */
export default function TabsView() {
  const navigate = useNavigate();
  const location = useLocation();
  const tabs = useTabsStore((s) => s.tabs);
  const addTab = useTabsStore((s) => s.addTab);
  const removeTab = useTabsStore((s) => s.removeTab);
  const removeOther = useTabsStore((s) => s.removeOther);
  const removeAll = useTabsStore((s) => s.removeAll);
  const getNeighborPath = useTabsStore((s) => s.getNeighborPath);

  const [activeKey, setActiveKey] = useState(location.pathname);

  // 路由变化时同步页签（不存在则新增，标题取对应菜单名称）
  useEffect(() => {
    setActiveKey(location.pathname);
    addTab({
      path: location.pathname,
      title: getRouteTitle(location.pathname),
      closable: location.pathname !== HomePath,
    });
  }, [location.pathname, addTab]);

  const closeTab = (path: string) => {
    const nextPath = getNeighborPath(path, location.pathname);
    removeTab(path);
    if (nextPath) navigate(nextPath);
  };

  const onMenuClick = (key: string, path: string) => {
    if (key === 'current') {
      closeTab(path);
      return;
    }
    if (key === 'other') {
      removeOther(path);
      if (location.pathname !== path) navigate(path);
      return;
    }
    if (key === 'all') {
      removeAll();
      if (location.pathname !== HomePath) navigate(HomePath);
    }
  };

  const items: MenuProps['items'] = [
    { key: 'current', icon: <CloseOutlined />, label: '关闭当前' },
    { key: 'other', icon: <ColumnWidthOutlined />, label: '关闭其他' },
    { key: 'all', icon: <MinusOutlined />, label: '关闭全部' },
  ];

  return (
    <div className="tabs-view">
      <Tabs
        activeKey={activeKey}
        type="editable-card"
        hideAdd
        size="small"
        className="tabs"
        onChange={(key) => {
          if (key !== location.pathname) navigate(key);
        }}
        onEdit={(targetKey, action) => {
          if (action === 'remove') closeTab(String(targetKey));
        }}
        items={tabs.map((tab) => ({
          key: tab.path,
          closable: tab.closable,
          label: (
            <Dropdown menu={{ items, onClick: ({ key }) => onMenuClick(String(key), tab.path) }} trigger={['contextMenu']}>
              <span className="tab-title">{getRouteTitle(tab.path)}</span>
            </Dropdown>
          ),
        }))}
      />
    </div>
  );
}
