import { useEffect, useState } from 'react';
import { Badge, Descriptions, Drawer, Spin, Tag } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminPermissionInfo } from '@/api/admin/permission';
import type { ResponseAdminPermissionInfoType } from '@/api/admin/permission';
import type { PermissionListItem } from '@/types/admin_permission';

interface DetailPermissionDrawerProps {
  open: boolean;
  detailData?: PermissionListItem;
  onClose: () => void;
}

/** 权限详情（只读；权限点与接口的绑定在菜单权限配置中维护） */
export default function DetailPermissionDrawer({ open, detailData, onClose }: DetailPermissionDrawerProps) {
  const [loading, setLoading] = useState(false);
  const [detail, setDetail] = useState<ResponseAdminPermissionInfoType>();

  useEffect(() => {
    if (open && detailData?.id) {
      setLoading(true);
      setDetail(undefined);
      getAdminPermissionInfo({ id: detailData.id })
        .then((res) => setDetail(res.data))
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  return (
    <Drawer open={open} title="权限详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Spin spinning={loading}>
        <Descriptions column={1} bordered>
          <Descriptions.Item label="菜单名称">{detail?.menuName || '-'}</Descriptions.Item>
          <Descriptions.Item label="菜单路径">{detail?.menuPath || '-'}</Descriptions.Item>
          <Descriptions.Item label="权限名称">{detail?.name}</Descriptions.Item>
          <Descriptions.Item label="唯一键">{detail?.key}</Descriptions.Item>
          <Descriptions.Item label="权限类型">{detail?.typeText || detail?.type}</Descriptions.Item>
          <Descriptions.Item label="描述">{detail?.describe || '-'}</Descriptions.Item>
          <Descriptions.Item label="状态">
            <Badge status={detail?.isEnabled ? 'success' : 'error'} text={detail?.isEnabled ? '启用' : '禁用'} />
          </Descriptions.Item>
          <Descriptions.Item label="接口数量">{detail?.apis?.length || 0}</Descriptions.Item>
          <Descriptions.Item label="接口列表">
            {detail?.apis?.length ? (
              <div>
                {detail.apis.map((api) => (
                  <Tag key={api.id} color="green" title={api.path}>
                    {api.name}
                  </Tag>
                ))}
              </div>
            ) : (
              <span style={{ color: 'rgba(0,0,0,0.45)' }}>暂无接口</span>
            )}
          </Descriptions.Item>
        </Descriptions>
        <div style={{ marginTop: 16, fontSize: 12, lineHeight: 1.5, color: 'rgba(0,0,0,0.45)' }}>
          权限点与接口的绑定在“菜单管理 → 对应菜单 → 权限配置”中维护
        </div>
      </Spin>
    </Drawer>
  );
}
