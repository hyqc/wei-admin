import { useState } from 'react';
import { Avatar, Button, Form, Input, Spin, Upload, message } from 'antd';
import type { UploadProps } from 'antd';
import { useUserStore } from '@/store/user';
import { currentAdminEdit } from '@/api/admin/account';
import { uploadAdminFile } from '@/api/admin/upload';
import { AdminEmail } from '@/api/pattern';

/** 头像上传分组：与个人中心菜单路径前缀保持一致 */
const AVATAR_UPLOAD_GROUP = '/account/';
/** 允许的图片扩展名（与后端 upload.allowed_exts 保持一致） */
const IMAGE_EXTS = ['jpg', 'jpeg', 'png', 'gif', 'webp', 'ico'];
/** 图片大小上限（与后端 upload.max_size 保持一致） */
const IMAGE_MAX_SIZE = 10 * 1024 * 1024;

/** 个人中心：编辑当前账号资料 */
export default function CurrentAccount() {
  const userInfo = useUserStore((s) => s.userInfo);
  const setCurrentUser = useUserStore((s) => s.setCurrentUser);
  const [form] = Form.useForm();
  const [loading] = useState(false);
  const [saving, setSaving] = useState(false);
  const [uploading, setUploading] = useState(false);
  // 左侧大头像跟随表单实时预览
  const avatarPreview = Form.useWatch('avatar', form) || userInfo?.avatar || '';

  const displayName = userInfo?.nickname || userInfo?.username || '?';

  /** 选择图片后立即上传（同步返回 false，仅拦截选择，不触发 antd 自动上传） */
  const onPickAvatar: UploadProps['beforeUpload'] = (file) => {
    const ext = (file.name.split('.').pop() || '').toLowerCase();
    if (!IMAGE_EXTS.includes(ext) || (file.type && !file.type.startsWith('image/'))) {
      message.error(`仅支持上传 ${IMAGE_EXTS.join('/')} 图片`);
      return false;
    }
    if (file.size > IMAGE_MAX_SIZE) {
      message.error('图片大小不能超过 10MB');
      return false;
    }
    void (async () => {
      setUploading(true);
      try {
        const res = await uploadAdminFile({ file, uploadGroup: AVATAR_UPLOAD_GROUP });
        form.setFieldValue('avatar', res.data?.url || '');
        message.success('头像上传成功');
      } finally {
        setUploading(false);
      }
    })();
    return false;
  };

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
          <Avatar size={96} src={avatarPreview} style={{ background: '#1677ff', fontSize: 36 }}>
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
                {
                  // 邮箱非必填：未填写时跳过校验，填写后才校验格式
                  validator(_rule, value: string) {
                    if (!value || AdminEmail.test(value)) {
                      return Promise.resolve();
                    }
                    return Promise.reject(new Error('邮箱格式不正确'));
                  },
                },
              ]}
            >
              <Input placeholder="请输入邮箱" allowClear />
            </Form.Item>
            <Form.Item
              label="头像"
              name="avatar"
              extra="支持 jpg/jpeg/png/gif/webp/ico，大小不超过 10MB；上传后自动填入地址并预览"
            >
              <Input
                placeholder="请输入头像图片地址"
                allowClear
                addonAfter={
                  <Upload accept="image/*" showUploadList={false} beforeUpload={onPickAvatar}>
                    <Button type="link" size="small" loading={uploading}>
                      点击上传
                    </Button>
                  </Upload>
                }
              />
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
