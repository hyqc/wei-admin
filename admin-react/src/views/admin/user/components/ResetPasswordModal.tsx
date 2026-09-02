import { useEffect, useState } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { resetAdminUserPassword } from '@/api/admin/user';
import { AdminUserPassword } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';

interface ResetPasswordModalProps {
  open: boolean;
  detailData?: AdminUserListItem;
  onClose: () => void;
}

/** 重置密码 */
export default function ResetPasswordModal({ open, detailData, onClose }: ResetPasswordModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);

  useEffect(() => {
    if (open) {
      form.resetFields();
    }
  }, [open, form]);

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await resetAdminUserPassword({ adminId: detailData?.adminId, password: values.password });
      message.success(res.msg, 2);
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal
      open={open}
      title="重置密码"
      width={DefaultModalWidth}
      confirmLoading={confirmLoading}
      maskClosable={false}
      okText="保存"
      cancelText="取消"
      onOk={handleOk}
      onCancel={onClose}
    >
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 12 }} autoComplete="off">
        <Form.Item label="账号">
          <Input value={detailData?.username} disabled autoComplete="off" />
        </Form.Item>
        <Form.Item
          label="新密码"
          name="password"
          rules={[
            { required: true, message: '请输入新密码' },
            { pattern: AdminUserPassword, message: '密码格式不正确' },
          ]}
        >
          {/* 禁用浏览器自动填充，避免带入已保存的登录密码 */}
          <Input.Password placeholder="请输入新密码" autoComplete="new-password" />
        </Form.Item>
        <Form.Item
          label="确认密码"
          name="confirmPassword"
          dependencies={['password']}
          rules={[
            { required: true, message: '请再次输入新密码' },
            ({ getFieldValue }) => ({
              validator(_rule, value: string) {
                if (!value || value === getFieldValue('password')) {
                  return Promise.resolve();
                }
                return Promise.reject(new Error('两次输入的密码不一致'));
              },
            }),
          ]}
        >
          <Input.Password placeholder="请再次输入新密码" autoComplete="new-password" />
        </Form.Item>
      </Form>
    </Modal>
  );
}
