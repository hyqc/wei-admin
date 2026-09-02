import { useEffect, useState } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { addAdminRole } from '@/api/admin/role';
import { DefaultModalWidth } from '@/api/config';

interface AddRoleModalProps {
  open: boolean;
  onClose: () => void;
  onNotice: () => void;
}

/** 新建角色 */
export default function AddRoleModal({ open, onClose, onNotice }: AddRoleModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);

  useEffect(() => {
    if (open) form.resetFields();
  }, [open, form]);

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await addAdminRole(values);
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal open={open} title="新建角色" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 12 }}>
        <Form.Item
          label="角色名称"
          name="name"
          rules={[
            { required: true, message: '请输入角色名称' },
            { max: 50, message: '角色名称长度不能超过50个字符' },
          ]}
        >
          <Input placeholder="请输入角色名称" allowClear />
        </Form.Item>
        <Form.Item label="描述" name="describe" rules={[{ max: 200, message: '描述长度不能超过200个字符' }]}>
          <Input.TextArea placeholder="请输入角色描述" rows={4} />
        </Form.Item>
      </Form>
    </Modal>
  );
}
