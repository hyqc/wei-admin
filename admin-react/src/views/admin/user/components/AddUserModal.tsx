import { useEffect, useState } from 'react';
import { Form, Input, Modal, Select, message } from 'antd';
import { addAdminUser } from '@/api/admin/user';
import { getAdminRoleAll } from '@/api/admin/role';
import { AdminUsername, AdminUserPassword, AdminEmail } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { RoleItem } from '@/types/admin_role';

interface AddUserModalProps {
  open: boolean;
  onClose: () => void;
  onNotice: () => void;
}

/** 新建账号 */
export default function AddUserModal({ open, onClose, onNotice }: AddUserModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [roleOptions, setRoleOptions] = useState<RoleItem[]>([]);

  useEffect(() => {
    if (open) {
      form.resetFields();
      getAdminRoleAll().then((res) => setRoleOptions(res.data || []));
    }
  }, [open, form]);

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await addAdminUser(values);
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
      title="新建账号"
      width={DefaultModalWidth}
      confirmLoading={confirmLoading}
      maskClosable={false}
      okText="保存"
      cancelText="取消"
      onOk={handleOk}
      onCancel={onClose}
    >
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 12 }} autoComplete="off">
        {/* 禁用浏览器自动填充，避免带入已保存的登录账号密码 */}
        <Form.Item
          label="账号"
          name="username"
          rules={[
            { required: true, message: '请输入账号' },
            { pattern: AdminUsername, message: '账号格式不正确' },
          ]}
        >
          <Input placeholder="请输入账号" allowClear autoComplete="off" />
        </Form.Item>
        <Form.Item
          label="密码"
          name="password"
          rules={[
            { required: true, message: '请输入密码' },
            { pattern: AdminUserPassword, message: '密码格式不正确' },
          ]}
        >
          <Input.Password placeholder="请输入密码" autoComplete="new-password" />
        </Form.Item>
        <Form.Item label="昵称" name="nickname" rules={[{ max: 50, message: '昵称长度不能超过50个字符' }]}>
          <Input placeholder="请输入昵称" allowClear />
        </Form.Item>
        <Form.Item label="邮箱" name="email" rules={[{ pattern: AdminEmail, message: '邮箱格式不正确' }]}>
          <Input placeholder="请输入邮箱" allowClear />
        </Form.Item>
        <Form.Item label="角色" name="roleIds">
          <Select mode="multiple" placeholder="请选择角色" allowClear options={roleOptions.map((r) => ({ value: r.id, label: r.name }))} />
        </Form.Item>
      </Form>
    </Modal>
  );
}
