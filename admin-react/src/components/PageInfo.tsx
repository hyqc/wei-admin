import type { PageInfoType } from '@/api/types';

/** 分页信息文案：共 X 条 第 Y/Z 页 */
export default function PageInfo({ pageInfo }: { pageInfo: PageInfoType }) {
  const totalPages = !pageInfo.total || !pageInfo.pageSize ? 1 : Math.ceil(pageInfo.total / pageInfo.pageSize);
  return (
    <div className="page-info">
      共 {pageInfo.total} 条 第 {pageInfo.pageNum}/{totalPages} 页
    </div>
  );
}
