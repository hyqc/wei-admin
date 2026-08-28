import type { PageInfoType } from './types';

/**
 * 处理分页参数
 * @param params 请求参数
 */
export function handlePagination(params: Record<string, unknown>) {
  const { pageSize, pageNum } = params;
  return {
    pageNum: pageNum ?? 1,
    pageSize: pageSize ?? 10,
  };
}

/**
 * 搜索时重置分页
 * @param pageInfo 当前分页信息
 * @param setPageInfo 设置分页
 * @param searchFn 搜索函数
 */
export function searchResetPageInfo(
  pageInfo: PageInfoType,
  setPageInfo: (pageInfo: PageInfoType) => void,
  searchFn: () => void,
) {
  if (pageInfo.pageNum !== 1) {
    setPageInfo({ ...pageInfo, pageNum: 1 });
  } else {
    searchFn();
  }
}
