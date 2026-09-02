import { useState } from 'react';
import { Avatar, Button, Form, Input, Spin, message } from 'antd';
import { useUserStore } from '@/store/user';
import { currentAdminEdit } from '@/api/admin/account';
import { AdminEmail } from '@/api/pattern';

/** 个人中心：编辑当前账号资料 */
export default function CurrentAccount() {
  const userInfo = useUserStore((s) => s.userInfo);
  const setCurrentUser = useUserStore((s) => s.setCurrentUser);
  const [form] = Form.useForm();
  const [loading] = useState(false);
  const [saving, setSaving] = useState(false);

  const displayName = userInfo?.nickname || userInfo?.username || '?';

  const onSave = async () => {
    const values = await form.validateFields();
    setSaving(true);
    try {
      const res = await currentAdminEdit({
        adminId: userInfo?.adminId,
        nickname: values.nickname,
        email: values.email,
        avatar: values.avatar,
      });
      // 更新本地用户信息
      if (userInfo) {
        setCurrentUser({
          ...userInfo,
          nickname: values.nickname,
          email: values.email,
          avatar: values.avatar,
        });
      }
      message.success(res.msg, 2);
    } finally {
      setSaving(false);
    }
  };

  return (
    <Spin spinning={loading}>
      <div style={{ display: 'flex', padding: 24 }}>
        <div style={{ flex: '0 0 160px', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
          <Avatar size={96} src={userInfo?.avatar} style={{ background: '#1677ff', fontSize: 36 }}>
            {displayName[0]}
          </Avatar>
          <div style={{ marginTop: 12, fontSize: 16, fontWeight: 600 }}>{userInfo?.nickname}</div>
        </div>
        <div style={{ flex: 1 }}>
          <Form
            form={form}
            labelCol={{ span: 4 }}
            initialValues={{
              username: userInfo?.username || '',
              nickname: userInfo?.nickname || '',
              email: userInfo?.email || '',
              avatar: userInfo?.avatar || '',
            }}
          >
            <Form.Item label="账号" name="username">
              <Input disabled />
            </Form.Item>
            <Form.Item
              label="昵称"
              name="nickname"
              rules={[
                { required: true, message: '请输入昵称' },
                { max: 50, message: '昵称长度不能超过50个字符' },
              ]}
            >
              <Input placeholder="请输入昵称" allowClear />
            </Form.Item>
            <Form.Item
              label="邮箱"
              name="email"
              rules={[
                { required: true, message: '请输入邮箱' },
                { pattern: AdminEmail, message: '邮箱格式不正确' },
              ]}
            >
              <Input placeholder="请输入邮箱" allowClear />
            </Form.Item>
            <Form.Item label="头像" name="avatar">
              <Input placeholder="请输入头像地址" allowClear />
            </Form.Item>
            <Form.Item wrapperCol={{ offset: 4 }}>
              <Button type="primary" loading={saving} onClick={onSave}>
                保存
              </Button>
            </Form.Item>
          </Form>
        </div>
      </div>
    </Spin>
  );
}
