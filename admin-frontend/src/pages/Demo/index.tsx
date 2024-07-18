import React from 'react';
import { Container, Content, Search } from '@/components/PageListContainer';
import { Form, Input, Button, Select, Row, Col } from 'antd';
import { SearchOutlined, ReloadOutlined } from '@ant-design/icons';
import type { Gutter } from 'antd/lib/grid/row';

const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];
const FormSearchRowColSpan = 6;

const Demo: React.FC = () => {
  const [form] = Form.useForm();

  return (
    <Container>
      <Search>
        <Form form={form}>
          <Row gutter={FormSearchRowGutter}>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="名称">
                <Input  />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="昵称">
                <Input  />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="角色名">
                <Input  />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="邮箱">
                <Input  />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="状态">
                <Select defaultValue={'0'} style={{ offset: 0, width: '120' }}>
                  <Select.Option value="0">全部</Select.Option>
                  <Select.Option value="1">启用</Select.Option>
                  <Select.Option value="2">禁用</Select.Option>
                </Select>
              </Form.Item>
            </Col>
          </Row>
          <Row>
            <Col span={24} style={{ textAlign: 'right' }}>
              <Button type="primary" htmlType="submit" style={{ marginRight: '2rem' }}>
                <SearchOutlined />
                查询
              </Button>
              <Button type="primary" htmlType="submit">
                <ReloadOutlined />
                清除
              </Button>
            </Col>
          </Row>
        </Form>
      </Search>
      <Content>这里是内容</Content>
    </Container>
  );
};

export default Demo;
