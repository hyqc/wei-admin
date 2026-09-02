import type { ReactNode } from 'react';
import { Pagination } from 'antd';
import type { PageInfoType } from '@/api/types';
import PageInfo from './PageInfo';

interface PageContainerProps {
  pageInfo?: PageInfoType;
  pageSizeOptions?: string[];
  searchArea?: ReactNode;
  extra?: ReactNode;
  children?: ReactNode;
  onPageChange?: (pageNum: number) => void;
  onPageSizeChange?: (pageSize: number) => void;
}

/** 页面容器：搜索区 + 操作区 + 内容 + 分页 */
export default function PageContainer({
  pageInfo,
  pageSizeOptions,
  searchArea,
  extra,
  children,
  onPageChange,
  onPageSizeChange,
}: PageContainerProps) {
  return (
    <div className="page-container">
      <div className="page-header">
        <div className="header-search">{searchArea}</div>
        <div className="header-extra">{extra}</div>
      </div>
      <div className="page-content">{children}</div>
      {pageInfo && (
        <div className="page-footer">
          <div className="footer-pageinfo">
            <PageInfo pageInfo={pageInfo} />
          </div>
          <div className="footer-pagination">
            <Pagination
              current={pageInfo.pageNum}
              pageSize={pageInfo.pageSize}
              total={pageInfo.total}
              showSizeChanger
              showQuickJumper
              showLessItems
              pageSizeOptions={pageSizeOptions || ['10', '20', '50', '100']}
              onChange={(pageNum, pageSize) => {
                // 页大小变化时 antd 会同时触发 onChange，交由父级按需处理
                if (pageSize !== pageInfo.pageSize) {
                  onPageSizeChange?.(pageSize);
                  return;
                }
                onPageChange?.(pageNum);
              }}
            />
          </div>
        </div>
      )}
    </div>
  );
}
