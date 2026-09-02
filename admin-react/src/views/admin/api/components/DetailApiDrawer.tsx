import { Badge, Descriptions, Drawer } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import type { AdminApiItem } from '@/types/common';

interface DetailApiDrawerProps {
  open: boolean;
  detailData?: AdminApiItem;
  onClose: () => void;
}

/** 接口详情 */
export default function DetailApiDrawer({ open, detailData, onClose }: DetailApiDrawerProps) {
  return (
    <Drawer open={open} title="接口详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Descriptions column={1} bordered>
        <Descriptions.Item label="接口ID">{detailData?.id}</Descriptions.Item>
        <Descriptions.Item label="接口名称">{detailData?.name}</Descriptions.Item>
        <Descriptions.Item label="唯一键">{detailData?.key}</Descriptions.Item>
        <Descriptions.Item label="接口路径">{detailData?.path}</Descriptions.Item>
        <Descriptions.Item label="描述">{detailData?.describe}</Descriptions.Item>
        <Descriptions.Item label="状态">
          <Badge status={detailData?.isEnabled ? 'success' : 'error'} text={detailData?.isEnabled ? '启用' : '禁用'} />
        </Descriptions.Item>
        <Descriptions.Item label="创建时间">{detailData?.createdAt}</Descriptions.Item>
        <Descriptions.Item label="更新时间">{detailData?.updatedAt}</Descriptions.Item>
      </Descriptions>
    </Drawer>
  );
}
