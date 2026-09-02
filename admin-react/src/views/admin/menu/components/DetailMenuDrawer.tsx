import { Badge, Descriptions, Drawer } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
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

/** 菜单详情 */
export default function DetailMenuDrawer({ open, detailData, onClose }: DetailMenuDrawerProps) {
  const Icon = getAntIcon(detailData?.icon);
  return (
    <Drawer open={open} title="菜单详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Descriptions column={1} bordered>
        <Descriptions.Item label="菜单名称">{detailData?.name}</Descriptions.Item>
        <Descriptions.Item label="键名">{detailData?.key}</Descriptions.Item>
        <Descriptions.Item label="菜单路径">{detailData?.path}</Descriptions.Item>
        <Descriptions.Item label="重定向地址">{detailData?.redirect}</Descriptions.Item>
        <Descriptions.Item label="排序">{detailData?.sort}</Descriptions.Item>
        <Descriptions.Item label="图标">
          {Icon ? (
            <span className="icon-cell">
              <Icon />
              <span>{detailData?.icon}</span>
            </span>
          ) : (
            detailData?.icon || '-'
          )}
        </Descriptions.Item>
        <Descriptions.Item label="描述">{detailData?.describe}</Descriptions.Item>
        <Descriptions.Item label="状态">
          <Badge status={detailData?.enabled ? 'success' : 'error'} text={detailData?.enabled ? '启用' : '禁用'} />
        </Descriptions.Item>
        <Descriptions.Item label="是否显示">{detailData?.hideInMenu ? '隐藏' : '显示'}</Descriptions.Item>
        <Descriptions.Item label="是否隐藏子菜单">{detailData?.hideChildrenInMenu ? '隐藏' : '显示'}</Descriptions.Item>
        <Descriptions.Item label="创建时间">{formatTime(detailData?.createTime)}</Descriptions.Item>
        <Descriptions.Item label="更新时间">{formatTime(detailData?.modifyTime)}</Descriptions.Item>
      </Descriptions>
    </Drawer>
  );
}

