import { useState } from 'react';
import { Card, Tabs } from 'antd';
import CurrentAccount from './components/CurrentAccount';
import CurrentPassword from './components/CurrentPassword';

/** 个人中心 */
export default function Account() {
  const [activeKey, setActiveKey] = useState('account');
  return (
    <div style={{ maxWidth: 800 }}>
      <Card>
        <Tabs activeKey={activeKey} onChange={setActiveKey}>
          <Tabs.TabPane key="account" tab="个人中心">
            <CurrentAccount />
          </Tabs.TabPane>
          <Tabs.TabPane key="password" tab="修改密码">
            <CurrentPassword />
          </Tabs.TabPane>
        </Tabs>
      </Card>
    </div>
  );
}
