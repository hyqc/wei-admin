import React from 'react';
import { Card } from 'antd';

export type SearchType = {
  children?: React.ReactNode;
};

// 列表页容器
const Search: React.FC<SearchType> = (props: any) => {
  return (
    <Card style={{ backgroundColor: '#FFF', ...props?.style }} bordered={false}>
      {props?.children}
    </Card>
  );
};

export default Search;
