import { Container, Content } from '@/components/PageListContainer';
import { Tabs } from 'antd';
import React from 'react';

import CurrentAccountInfo from './components/CurrentAccount';
import CurrentAccountPassword from './components/CurrentPassword';

const CurrentAccount: React.FC = () => {
  const items = [
    {
      key: 'center',
      label: '个人信息',
      children: <CurrentAccountInfo />,
    },
    {
      key: 'password',
      label: '修改密码',
      children: <CurrentAccountPassword />,
    },
  ];

  return (
    <Container>
      <Content>
        <Tabs tabPosition="left" items={items} />
      </Content>
    </Container>
  );
};

export default CurrentAccount;
