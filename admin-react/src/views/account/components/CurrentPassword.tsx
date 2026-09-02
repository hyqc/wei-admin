import { useState } from 'react';
import { Button, Form, Input, message } from 'antd';
import { currentAdminEditPassword } from '@/api/admin/account';
import { AdminUserPassword } from '@/api/pattern';

/** 修改密码 */
export default function CurrentPassword() {
  const [form] = Form.useForm();
  const [saving, setSaving] = useState(false);

  const onSave = async () => {
    const values = await form.validateFields();
    setSaving(true);
    try {
      const res = await currentAdminEditPassword({
        oldPassword: values.oldPassword,
        newPassword: values.newPassword,
        confirmPassword: values.confirmPassword,
      });
      message.success(res.msg, 2);
      form.resetFields();
    } finally {
      setSaving(false);
    }
  };

  return (
    <div>
      {/* 禁用浏览器自动填充，避免带入已保存的登录密码 */}
      <Form form={form} labelCol={{ span: 4 }} style={{ maxWidth: 480 }} autoComplete="off">
        <Form.Item label="原密码" name="oldPassword" rules={[{ required: true, message: '请输入原密码' }]}>
          <Input.Password placeholder="请输入原密码" autoComplete="current-password" />
        </Form.Item>
        <Form.Item
          label="新密码"
          name="newPassword"
          rules={[
            { required: true, message: '请输入新密码' },
            { pattern: AdminUserPassword, message: '密码格式不正确' },
          ]}
        >
          <Input.Password placeholder="请输入新密码" autoComplete="new-password" />
        </Form.Item>
        <Form.Item
          label="确认密码"
          name="confirmPassword"
          dependencies={['newPassword']}
          rules={[
            { required: true, message: '请再次输入新密码' },
            ({ getFieldValue }) => ({
              validator(_rule, value: string) {
                if (!value || value === getFieldValue('newPassword')) {
                  return Promise.resolve();
                }
                return Promise.reject(new Error('两次输入的密码不一致'));
              },
            }),
          ]}
        >
          <Input.Password placeholder="请再次输入新密码" autoComplete="new-password" />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 4 }}>
          <Button type="primary" loading={saving} onClick={onSave}>
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}
