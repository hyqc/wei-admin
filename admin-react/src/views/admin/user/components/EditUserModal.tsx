import { useEffect, useState } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { editAdminUser, getAdminUserInfo } from '@/api/admin/user';
import { AdminEmail } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';

interface EditUserModalProps {
  open: boolean;
  detailData?: AdminUserListItem;
  onClose: () => void;
  onNotice: () => void;
}

/** 编辑账号 */
export default function EditUserModal({ open, detailData, onClose, onNotice }: EditUserModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);

  // 打开时实时拉取详情回填，避免编辑列表中的过期数据
  useEffect(() => {
    if (open && detailData?.adminId) {
      getAdminUserInfo({ adminId: detailData.adminId }).then((res) => {
        form.setFieldsValue({ nickname: res.data.nickname || '', email: res.data.email || '' });
        // 清空上一次打开残留的校验状态
        form.setFields([
          { name: 'nickname', errors: [] },
          { name: 'email', errors: [] },
        ]);
      });
    }
  }, [open, detailData, form]);

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await editAdminUser({
        adminId: detailData?.adminId as number,
        nickname: values.nickname,
        email: values.email,
      });
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
      title="编辑账号"
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
      </Form>
    </Modal>
  );
}
