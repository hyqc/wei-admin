import { Badge, Descriptions, Drawer, Tag } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';

interface DetailUserDrawerProps {
  open: boolean;
  detailData?: AdminUserListItem;
  onClose: () => void;
}

/** 账号详情 */
export default function DetailUserDrawer({ open, detailData, onClose }: DetailUserDrawerProps) {
  return (
    <Drawer open={open} title="账号详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Descriptions column={1} bordered>
        <Descriptions.Item label="账号">{detailData?.username}</Descriptions.Item>
        <Descriptions.Item label="昵称">{detailData?.nickname}</Descriptions.Item>
        <Descriptions.Item label="邮箱">{detailData?.email}</Descriptions.Item>
        <Descriptions.Item label="角色">
          {/* 超管无需绑定角色，自动拥有系统全部权限 */}
          {detailData?.isSuperAdmin ? (
            <Tag color="gold">超级管理员</Tag>
          ) : detailData?.roles?.length ? (
            detailData.roles.map((role) => (
              <Tag key={role.roleId} color="blue">
                {role.roleName}
              </Tag>
            ))
          ) : (
            '-'
          )}
        </Descriptions.Item>
        <Descriptions.Item label="状态">
          <Badge status={detailData?.isEnabled ? 'success' : 'error'} text={detailData?.isEnabled ? '启用' : '禁用'} />
        </Descriptions.Item>
        <Descriptions.Item label="登录次数">{detailData?.loginTotal}</Descriptions.Item>
        <Descriptions.Item label="上次登录IP">{detailData?.lastLoginIp || '-'}</Descriptions.Item>
        <Descriptions.Item label="本次登录IP">{detailData?.currentLoginIp || '-'}</Descriptions.Item>
        <Descriptions.Item label="本次登录时间">{detailData?.currentLoginTime || '-'}</Descriptions.Item>
        <Descriptions.Item label="创建时间">{detailData?.createdAt}</Descriptions.Item>
        <Descriptions.Item label="更新时间">{detailData?.updatedAt}</Descriptions.Item>
      </Descriptions>
    </Drawer>
  );
}
