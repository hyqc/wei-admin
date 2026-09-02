import { useState } from 'react';
import { Outlet } from 'react-router-dom';
import { Layout } from 'antd';
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import SiderMenu from '@/components/SiderMenu';
import HeaderRight from '@/components/HeaderRight';
import FooterBar from '@/components/FooterBar';
import TabsView from '@/components/TabsView';

const { Sider, Header, Content, Footer } = Layout;

/** 基础布局：侧边菜单 + 头部 + 页签 + 内容 + 页脚 */
export default function BasicLayout() {
  const [collapsed, setCollapsed] = useState(false);

  return (
    <Layout className="basic-layout">
      <Sider trigger={null} collapsible collapsed={collapsed} width={220} className="layout-sider">
        <SiderMenu collapsed={collapsed} />
      </Sider>
      <Layout className="basic-layout-inner">
        <Header className="layout-header">
          <div className="header-left">
            {collapsed ? (
              <MenuUnfoldOutlined className="trigger" onClick={() => setCollapsed(!collapsed)} />
            ) : (
              <MenuFoldOutlined className="trigger" onClick={() => setCollapsed(!collapsed)} />
            )}
          </div>
          <HeaderRight />
        </Header>
        <Content className="layout-content">
          <TabsView />
          <Outlet />
        </Content>
        <Footer className="layout-footer">
          <FooterBar />
        </Footer>
      </Layout>
    </Layout>
  );
}
