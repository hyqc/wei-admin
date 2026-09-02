import { useEffect, useState } from 'react';
import { Form, Input, Modal, Select, message } from 'antd';
import { bindAdminUserRoles, getAdminUserInfo } from '@/api/admin/user';
import { getAdminRoleAll } from '@/api/admin/role';
import { DefaultModalWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';
import type { RoleItem } from '@/types/admin_role';

interface BindRolesModalProps {
  open: boolean;
  detailData?: AdminUserListItem;
  onClose: () => void;
  onNotice: () => void;
}

/** 账号绑定角色 */
export default function BindRolesModal({ open, detailData, onClose, onNotice }: BindRolesModalProps) {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [roleOptions, setRoleOptions] = useState<RoleItem[]>([]);

  useEffect(() => {
    if (open && detailData) {
      setLoading(true);
      getAdminRoleAll()
        .then((res) => setRoleOptions(res.data || []))
        .finally(() => setLoading(false));
      getAdminUserInfo({ adminId: detailData.adminId }).then((res) => {
        form.setFieldsValue({ roleIds: res.data.roleIds || [] });
      });
    }
  }, [open, detailData, form]);

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await bindAdminUserRoles({ adminId: detailData?.adminId, roleIds: values.roleIds ?? [] });
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal
      open={open}
      title="绑定角色"
      width={DefaultModalWidth}
      confirmLoading={confirmLoading}
      maskClosable={false}
      okText="保存"
      cancelText="取消"
      onOk={handleOk}
      onCancel={onClose}
    >
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 12 }}>
        <Form.Item label="账号">
          <Input value={detailData?.username} disabled />
        </Form.Item>
        <Form.Item label="角色" name="roleIds">
          <Select mode="multiple" placeholder="请选择角色" allowClear loading={loading} options={roleOptions.map((r) => ({ value: r.id, label: r.name }))} />
        </Form.Item>
      </Form>
    </Modal>
  );
}
