import { defineStore } from 'pinia';
import { ref, watch } from 'vue';
import { HomePath, LocalStorageTabsKey } from '@/api/config';

/** 页签项 */
export interface TabItem {
  /** 路由路径，作为页签唯一键 */
  path: string;
  /** 页签标题（菜单名称） */
  title: string;
  /** 是否可关闭（首页不可关闭） */
  closable: boolean;
}

/** 首页页签，固定且不可关闭 */
const HOME_TAB: TabItem = { path: HomePath, title: '首页', closable: false };

/** 从本地存储恢复页签 */
function loadTabs(): TabItem[] {
  try {
    const raw = localStorage.getItem(LocalStorageTabsKey);
    if (!raw) return [{ ...HOME_TAB }];
    const list = JSON.parse(raw) as TabItem[];
    if (!Array.isArray(list)) return [{ ...HOME_TAB }];
    // 过滤无效数据；首页固定在最前且不重复
    const rest = list.filter((item) => item?.path && item.path !== HomePath);
    return [{ ...HOME_TAB }, ...rest];
  } catch {
    return [{ ...HOME_TAB }];
  }
}

/** 页签状态 */
export const useTabsStore = defineStore('tabs', () => {
  const tabs = ref<TabItem[]>(loadTabs());

  // 页签变化持久化到本地，刷新页面后不丢失
  watch(
    tabs,
    (list) => {
      localStorage.setItem(LocalStorageTabsKey, JSON.stringify(list));
    },
    { deep: true },
  );

  /** 新增页签，已存在则不重复添加 */
  function addTab(tab: TabItem) {
    if (tabs.value.some((item) => item.path === tab.path)) return;
    tabs.value.push({ ...tab, closable: tab.path !== HomePath });
  }

  /**
   * 关闭页签后需要跳转的路径
   * @param path 被关闭的页签路径
   * @param currentPath 当前激活路径
   * @returns 关闭的是当前页签时返回相邻页签路径，否则返回 null
   */
  function getNeighborPath(path: string, currentPath: string): string | null {
    if (path !== currentPath) return null;
    const index = tabs.value.findIndex((item) => item.path === path);
    if (index === -1) return null;
    const next = tabs.value[index + 1] ?? tabs.value[index - 1];
    return next?.path ?? HomePath;
  }

  /** 关闭当前页签 */
  function removeTab(path: string) {
    if (path === HomePath) return;
    tabs.value = tabs.value.filter((item) => item.path !== path);
  }

  /** 关闭其他页签（保留首页与指定页签） */
  function removeOther(path: string) {
    tabs.value = tabs.value.filter((item) => item.path === path || item.path === HomePath);
  }

  /** 关闭全部页签（保留首页） */
  function removeAll() {
    tabs.value = [{ ...HOME_TAB }];
  }

  return { tabs, addTab, removeTab, removeOther, removeAll, getNeighborPath };
});
