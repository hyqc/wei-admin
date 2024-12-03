import { Col, Layout, Menu, Row } from 'antd';
import React, { useState } from 'react';
import MenuDoc from './menu';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import 'github-markdown-css';
import './index.less';
import { Link } from 'umi';

const { Header, Content, Sider } = Layout;

const Document: React.FC = () => {
  const [mdContent, setMdContent] = useState<string>('');
  async function parseMdFileContent(url: string) {
    const res = await fetch(url);
    const content = await res.text();
    setMdContent(content);
  }

  function openMd(path?: string) {
    path && path.length > 0 && parseMdFileContent(path);
  }

  return (
    <Layout className="layout">
      <Header className="header">
        <Row className="row">
          <Col span={6}>
            <h1>
              <Link to={'/'}>
                <img className="logo" src="/favicon.ico" />
                WEI
              </Link>
            </h1>
          </Col>
          <Col span={14}>
            <Menu mode="horizontal">
              <Menu.Item key="mail">更多</Menu.Item>
            </Menu>
          </Col>
          <Col span={4}></Col>
        </Row>
      </Header>
      <Layout className="layout-content">
        <Sider
          width={310}
          className="slider-left"
          theme="light"
          breakpoint="lg"
          collapsedWidth="0"
          onBreakpoint={(broken) => {}}
          onCollapse={(collapsed, type) => {}}
        >
          <MenuDoc changeDoc={openMd} />
        </Sider>
        <Content className="content">
          <div className="site-layout-background" style={{ padding: 24, minHeight: 360 }}>
            <ReactMarkdown
              className="markdown-body"
              remarkPlugins={[remarkGfm]}
              children={mdContent}
            />
          </div>
        </Content>
      </Layout>
    </Layout>
  );
};

export default Document;
