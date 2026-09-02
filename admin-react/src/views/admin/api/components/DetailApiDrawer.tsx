import { useEffect, useState } from 'react';
import { Badge, Descriptions, Drawer, Spin } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminApiInfo } from '@/api/admin/api';
import type { AdminApiItem } from '@/types/common';

interface DetailApiDrawerProps {
  open: boolean;
  detailData?: AdminApiItem;
  onClose: () => void;
}

/** 接口详情（打开时实时拉取，避免展示列表中的过期数据） */
export default function DetailApiDrawer({ open, detailData, onClose }: DetailApiDrawerProps) {
  const [loading, setLoading] = useState(false);
  const [detail, setDetail] = useState<AdminApiItem>();

  useEffect(() => {
    if (open && detailData?.id) {
      setLoading(true);
      setDetail(undefined);
      getAdminApiInfo({ id: detailData.id })
        .then((res) => setDetail(res.data))
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  return (
    <Drawer open={open} title="接口详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Spin spinning={loading}>
      <Descriptions column={1} bordered>
        <Descriptions.Item label="接口ID">{detail?.id}</Descriptions.Item>
        <Descriptions.Item label="接口名称">{detail?.name}</Descriptions.Item>
        <Descriptions.Item label="唯一键">{detail?.key}</Descriptions.Item>
        <Descriptions.Item label="接口路径">{detail?.path}</Descriptions.Item>
        <Descriptions.Item label="描述">{detail?.describe}</Descriptions.Item>
        <Descriptions.Item label="状态">
          <Badge status={detail?.isEnabled ? 'success' : 'error'} text={detail?.isEnabled ? '启用' : '禁用'} />
        </Descriptions.Item>
        <Descriptions.Item label="创建时间">{detail?.createdAt}</Descriptions.Item>
        <Descriptions.Item label="更新时间">{detail?.updatedAt}</Descriptions.Item>
      </Descriptions>
      </Spin>
    </Drawer>
  );
}
