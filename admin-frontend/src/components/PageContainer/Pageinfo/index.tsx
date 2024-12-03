import React from 'react';
import { Pagination } from 'antd';
import { DEFAULT_PAGE_NO, DEFAULT_PAGE_SIZE } from '@/pages/Admin/User/common';

export type PageinfoType = {
  total: number | undefined;
  pageNum: number | undefined;
  pageSize: number | undefined;
  children?: any;
  onChange?: (pageNum: number, pageSize?: number) => void;
  onShowSizeChange?: (current: number, size: number) => void;
};

export type PageDefaultType = {
  total: number;
  current: number;
  pageSize: number;
};

type StateType = { current: number; pageSize?: number };

// 列表页容器
class Pageinfo extends React.Component<PageinfoType, StateType> {
  current: number;
  pageSize: number;
  total: number;

  constructor(props: PageinfoType) {
    super(props);

    this.current = props?.pageSize || DEFAULT_PAGE_NO;
    this.pageSize = props?.pageSize || DEFAULT_PAGE_SIZE;
    this.total = props?.total || 0;

    this.state = { current: this.current, pageSize: this.pageSize };
    this.onChange = this.onChange.bind(this);
    this.onShowSizeChange = this.onShowSizeChange.bind(this);
  }

  onChange(pageNum: number, pageSize?: number) {}
  onShowSizeChange(current: number, size: number) {}

  render(): React.ReactNode {
    return (
      <Pagination
        showQuickJumper
        defaultCurrent={DEFAULT_PAGE_NO}
        style={{ marginTop: '1rem', textAlign: 'right' }}
        total={this.total}
        current={this.current}
        pageSize={this.pageSize}
        showTotal={(total) => `共 ${total} 条数据`}
        onChange={(page, pageSize) => this.onChange(page, pageSize)}
        onShowSizeChange={(current, size) => this.onShowSizeChange(current, size)}
        {...this.props.children}
      />
    );
  }
}

export default Pageinfo;
