import { useEffect, useState } from 'react';
import { Badge, Descriptions, Drawer, Spin, Tag } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminUserInfo } from '@/api/admin/user';
import type { AdminUserListItem } from '@/types/common';

interface DetailUserDrawerProps {
  open: boolean;
  detailData?: AdminUserListItem;
  onClose: () => void;
}

/** 账号详情（打开时实时拉取，避免展示列表中的过期数据） */
export default function DetailUserDrawer({ open, detailData, onClose }: DetailUserDrawerProps) {
  const [loading, setLoading] = useState(false);
  const [detail, setDetail] = useState<AdminUserListItem>();

  useEffect(() => {
    if (open && detailData?.adminId) {
      setLoading(true);
      setDetail(undefined);
      getAdminUserInfo({ adminId: detailData.adminId })
        .then((res) => setDetail(res.data))
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  return (
    <Drawer open={open} title="账号详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Spin spinning={loading}>
      <Descriptions column={1} bordered>
        <Descriptions.Item label="账号">{detail?.username}</Descriptions.Item>
        <Descriptions.Item label="昵称">{detail?.nickname}</Descriptions.Item>
        <Descriptions.Item label="邮箱">{detail?.email}</Descriptions.Item>
        <Descriptions.Item label="角色">
          {/* 超管无需绑定角色，自动拥有系统全部权限 */}
          {detail?.isSuperAdmin ? (
            <Tag color="gold">超级管理员</Tag>
          ) : detail?.roles?.length ? (
            detail.roles.map((role) => (
              <Tag key={role.roleId} color="blue">
                {role.roleName}
              </Tag>
            ))
          ) : (
            '-'
          )}
        </Descriptions.Item>
        <Descriptions.Item label="状态">
          <Badge status={detail?.isEnabled ? 'success' : 'error'} text={detail?.isEnabled ? '启用' : '禁用'} />
        </Descriptions.Item>
        <Descriptions.Item label="登录次数">{detail?.loginTotal}</Descriptions.Item>
        <Descriptions.Item label="上次登录IP">{detail?.lastLoginIp || '-'}</Descriptions.Item>
        <Descriptions.Item label="本次登录IP">{detail?.currentLoginIp || '-'}</Descriptions.Item>
        <Descriptions.Item label="本次登录时间">{detail?.currentLoginTime || '-'}</Descriptions.Item>
        <Descriptions.Item label="创建时间">{detail?.createdAt}</Descriptions.Item>
        <Descriptions.Item label="更新时间">{detail?.updatedAt}</Descriptions.Item>
      </Descriptions>
      </Spin>
    </Drawer>
  );
}
