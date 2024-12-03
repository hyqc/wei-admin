import { Button, Result } from 'antd';
import React from 'react';
import { history } from 'umi';

const ForbiddenPage: React.FC = () => (
  <Result
    status="403"
    title="403"
    subTitle="抱歉, 您没有访问该页面的权限，请联系管理员~."
    extra={
      <Button type="primary" onClick={() => history.push('/')}>
        返回首页
      </Button>
    }
  />
);

export default ForbiddenPage;
