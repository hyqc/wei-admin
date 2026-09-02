import { useEffect, useState } from 'react';
import { Empty, Form, Input, Modal, Spin, Switch, message } from 'antd';
import { getAdminMenuPermissions } from '@/api/admin/menu';
import { addAdminMenuPermissions } from '@/api/admin/permission';
import { DEFAULT_PERMISSION_TYPES, handleKey } from '@/views/admin/permission/components/common';
import type { MenuTreeItem, MenuPermissionItem } from '@/types/admin_menu';

interface Props {
  open: boolean;
  detailData?: MenuTreeItem;
  onClose: () => void;
}

/** 添加菜单操作权限（权限名称与唯一键由系统生成，不可修改） */
export default function PermissionsSaveModal({ open, detailData, onClose }: Props) {
  const [loading, setLoading] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [permissions, setPermissions] = useState<MenuPermissionItem[]>([]);

  useEffect(() => {
    if (open && detailData) {
      setPermissions([]);
      setLoading(true);
      const menuId = detailData.id;
      const menuName = detailData.name || '';
      const menuPath = detailData.path || '';
      getAdminMenuPermissions({ menuId })
        .then((res) => {
          const existing = res.data.permissions || [];
          if (existing.length > 0) {
            // 后端未返回类型名称，按类型补充展示
            setPermissions(
              existing.map((item) => ({
                ...item,
                typeName: DEFAULT_PERMISSION_TYPES.find((t) => t.key === item.type)?.name || item.type || '',
              })),
            );
            return;
          }
          // 菜单尚未配置权限时，按菜单信息自动生成查看/编辑/删除三类
          setPermissions(
            DEFAULT_PERMISSION_TYPES.map((type) => ({
              type: type.key,
              typeName: type.name,
              name: `${menuName}${type.name}`,
              key: handleKey(menuPath, type.key),
              enabled: true,
              describe: '',
            })),
          );
        })
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  const handleOk = () => {
    if (!detailData) return;
    setConfirmLoading(true);
    const list = permissions.map((item) => ({
      menuId: detailData.id,
      menuName: detailData.name,
      menuPath: detailData.path,
      id: item.id,
      name: item.name,
      key: item.key,
      type: item.type,
      describe: item.describe,
      enabled: item.enabled,
    }));
    addAdminMenuPermissions(list)
      .then((res) => {
        message.success(res.msg, 2);
        onClose();
      })
      .finally(() => setConfirmLoading(false));
  };

  return (
    <Modal open={open} title="添加菜单操作权限" width={900} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Spin spinning={loading}>
        <Form labelCol={{ span: 3 }} wrapperCol={{ span: 20 }}>
          <Form.Item label="菜单">
            <Input value={detailData?.name} disabled style={{ width: 260 }} />
          </Form.Item>
          <Form.Item label="权限">
            {permissions.map((item, index) => (
              <div key={index} style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 12 }}>
                <span style={{ flex: '0 0 48px', fontWeight: 600 }}>{item.typeName}</span>
                <Input value={item.name} style={{ width: 160 }} disabled />
                <Input value={item.key} style={{ width: 200 }} disabled />
                <Switch
                  checked={item.enabled}
                  checkedChildren="启用"
                  unCheckedChildren="禁用"
                  onChange={(checked) =>
                    setPermissions((prev) => prev.map((p, i) => (i === index ? { ...p, enabled: checked } : p)))
                  }
                />
              </div>
            ))}
            <div className="form-tip">权限名称与唯一键由系统根据菜单自动生成，不允许手动修改</div>
            {permissions.length === 0 && <Empty description="暂无权限配置" />}
          </Form.Item>
        </Form>
      </Spin>
    </Modal>
  );
}
