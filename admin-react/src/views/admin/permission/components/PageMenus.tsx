import { useEffect, useMemo, useState } from 'react';
import { Tree } from 'antd';
import { getAdminMenuPages } from '@/api/admin/menu';
import type { MenuTreeItem } from '@/types/admin_menu';

interface PageMenusProps {
  value?: number;
  onChange?: (node?: MenuTreeItem) => void;
}

/** 菜单页面树（仅显示未隐藏菜单，选中后回传完整菜单信息） */
export default function PageMenus({ value, onChange }: PageMenusProps) {
  const [menuPages, setMenuPages] = useState<MenuTreeItem[]>([]);
  const [selectedKeys, setSelectedKeys] = useState<number[]>(value ? [value] : []);
  const [expandedKeys, setExpandedKeys] = useState<number[]>([]);
  const [menuMap] = useState(() => new Map<number, MenuTreeItem>());

  // 仅显示未隐藏节点
  const treeData = useMemo(() => {
    const walk = (list: MenuTreeItem[]): MenuTreeItem[] =>
      list
        .filter((item) => item.hideInMenu !== true)
        .map((item) => ({ ...item, children: item.children?.length ? walk(item.children) : undefined }));
    return walk(menuPages);
  }, [menuPages]);

  useEffect(() => {
    setSelectedKeys(value ? [value] : []);
  }, [value]);

  useEffect(() => {
    getAdminMenuPages({ all: true }).then((res) => {
      setMenuPages(res.data || []);
      const collectMap = (list: MenuTreeItem[]) => {
        list.forEach((item) => {
          menuMap.set(item.id as number, item);
          if (item.children?.length) collectMap(item.children);
        });
      };
      collectMap(res.data || []);
      // 递归收集所有含子节点的菜单 id，全部展开
      const keys: number[] = [];
      const collect = (list: MenuTreeItem[]) => {
        for (const item of list) {
          if (item.children?.length) {
            keys.push(item.id as number);
            collect(item.children);
          }
        }
      };
      collect(res.data || []);
      setExpandedKeys(keys);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const onSelect = (keys: React.Key[]) => {
    const id = keys.length ? Number(keys[0]) : undefined;
    setSelectedKeys(id ? [id] : []);
    onChange?.(id ? menuMap.get(id) : undefined);
  };

  return (
    <Tree
      treeData={treeData as never}
      height={320}
      expandedKeys={expandedKeys}
      selectable
      selectedKeys={selectedKeys}
      fieldNames={{ title: 'name', key: 'id', children: 'children' }}
      onSelect={onSelect}
      onExpand={(keys) => setExpandedKeys(keys as number[])}
      titleRender={(node) => (
        <span>
          {(node as MenuTreeItem).name}
          {(node as MenuTreeItem).path && (
            <span style={{ color: 'rgba(0,0,0,0.45)', fontSize: 12 }}>（{(node as MenuTreeItem).path}）</span>
          )}
        </span>
      )}
    />
  );
}
