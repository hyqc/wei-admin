import React, { ReactNode } from 'react';
import { Card } from 'antd';

// 列表页容器
export type ContentType = {
  style?: any;
  children?: React.ReactNode;
};
const Content: React.FC<ContentType> = (props: any) => {
  const defaultMinHeight = {
    minHeight: '61vh',
  }
  return (
    <Card style={{ ...defaultMinHeight,backgroundColor: '#FFF', ...props?.style, marginTop: '1rem' }} bordered={false}>
      {props?.children}
    </Card>
  );
};

export default Content;
