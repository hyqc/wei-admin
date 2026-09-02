import { useEffect, useState } from 'react';
import { Badge, Descriptions, Drawer, Spin } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminMenuInfo } from '@/api/admin/menu';
import { getAntIcon } from '@/utils/icon';
import type { MenuTreeItem } from '@/types/admin_menu';

interface DetailMenuDrawerProps {
  open: boolean;
  detailData?: MenuTreeItem;
  onClose: () => void;
}

function formatTime(time?: number) {
  if (!time) return '';
  const d = new Date(time * 1000);
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
}

/** 菜单详情（打开时实时拉取，避免展示列表中的过期数据） */
export default function DetailMenuDrawer({ open, detailData, onClose }: DetailMenuDrawerProps) {
  const [loading, setLoading] = useState(false);
  const [detail, setDetail] = useState<MenuTreeItem>();

  useEffect(() => {
    if (open && detailData?.id) {
      setLoading(true);
      setDetail(undefined);
      getAdminMenuInfo({ menuId: detailData.id })
        .then((res) => setDetail({ ...(res.data as MenuTreeItem) }))
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  const Icon = getAntIcon(detail?.icon);
  return (
    <Drawer open={open} title="菜单详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Spin spinning={loading}>
      <Descriptions column={1} bordered>
        <Descriptions.Item label="菜单名称">{detail?.name}</Descriptions.Item>
        <Descriptions.Item label="键名">{detail?.key}</Descriptions.Item>
        <Descriptions.Item label="菜单路径">{detail?.path}</Descriptions.Item>
        <Descriptions.Item label="重定向地址">{detail?.redirect}</Descriptions.Item>
        <Descriptions.Item label="排序">{detail?.sort}</Descriptions.Item>
        <Descriptions.Item label="图标">
          {Icon ? (
            <span className="icon-cell">
              <Icon />
              <span>{detail?.icon}</span>
            </span>
          ) : (
            detail?.icon || '-'
          )}
        </Descriptions.Item>
        <Descriptions.Item label="描述">{detail?.describe}</Descriptions.Item>
        <Descriptions.Item label="状态">
          <Badge status={detail?.enabled ? 'success' : 'error'} text={detail?.enabled ? '启用' : '禁用'} />
        </Descriptions.Item>
        <Descriptions.Item label="是否显示">{detail?.hideInMenu ? '隐藏' : '显示'}</Descriptions.Item>
        <Descriptions.Item label="是否隐藏子菜单">{detail?.hideChildrenInMenu ? '隐藏' : '显示'}</Descriptions.Item>
        <Descriptions.Item label="创建时间">{formatTime(detail?.createTime)}</Descriptions.Item>
        <Descriptions.Item label="更新时间">{formatTime(detail?.modifyTime)}</Descriptions.Item>
      </Descriptions>
      </Spin>
    </Drawer>
  );
}

