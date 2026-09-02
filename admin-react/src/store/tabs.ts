import { create } from 'zustand';
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

/** 持久化到本地，刷新页面后不丢失 */
function persist(list: TabItem[]) {
  localStorage.setItem(LocalStorageTabsKey, JSON.stringify(list));
}

/** 页签状态 */
interface TabsState {
  tabs: TabItem[];
  /** 新增页签，已存在则不重复添加 */
  addTab: (tab: TabItem) => void;
  /** 关闭页签 */
  removeTab: (path: string) => void;
  /** 关闭其他页签（保留首页与指定页签） */
  removeOther: (path: string) => void;
  /** 关闭全部页签（保留首页） */
  removeAll: () => void;
  /** 关闭当前页签后需要跳转的路径；关闭的不是当前页时返回 null */
  getNeighborPath: (path: string, currentPath: string) => string | null;
}

export const useTabsStore = create<TabsState>((set, get) => ({
  tabs: loadTabs(),

  addTab: (tab) => {
    const { tabs } = get();
    if (tabs.some((item) => item.path === tab.path)) return;
    const next = [...tabs, { ...tab, closable: tab.path !== HomePath }];
    persist(next);
    set({ tabs: next });
  },

  removeTab: (path) => {
    if (path === HomePath) return;
    const next = get().tabs.filter((item) => item.path !== path);
    persist(next);
    set({ tabs: next });
  },

  removeOther: (path) => {
    const next = get().tabs.filter((item) => item.path === path || item.path === HomePath);
    persist(next);
    set({ tabs: next });
  },

  removeAll: () => {
    const next = [{ ...HOME_TAB }];
    persist(next);
    set({ tabs: next });
  },

  getNeighborPath: (path, currentPath) => {
    if (path !== currentPath) return null;
    const { tabs } = get();
    const index = tabs.findIndex((item) => item.path === path);
    if (index === -1) return null;
    const next = tabs[index + 1] ?? tabs[index - 1];
    return next?.path ?? HomePath;
  },
}));
