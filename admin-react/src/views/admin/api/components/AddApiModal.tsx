import { useEffect, useState } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { addAdminApi } from '@/api/admin/api';
import { AdminAPIKey } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import { generateApiKeyByPath } from '@/utils/apiKey';

interface AddApiModalProps {
  open: boolean;
  onClose: () => void;
  onNotice: () => void;
}

/** 新建接口（唯一键由接口路径自动生成，不可编辑） */
export default function AddApiModal({ open, onClose, onNotice }: AddApiModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);

  useEffect(() => {
    if (open) form.resetFields();
  }, [open, form]);

  // 路径变化后自动生成唯一键
  const onPathChange = (path: string) => {
    form.setFieldValue('key', generateApiKeyByPath(path));
  };

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await addAdminApi(values);
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal open={open} title="新建接口" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 14 }}>
        <Form.Item label="接口名称" name="name" rules={[{ required: true, message: '请输入接口名称' }, { max: 50, message: '名称长度不能超过50个字符' }]}>
          <Input placeholder="请输入接口名称" allowClear />
        </Form.Item>
        <Form.Item label="唯一键" name="key" rules={[{ required: true, message: '唯一键由接口路径自动生成，请先输入接口路径' }, { pattern: AdminAPIKey, message: '唯一键格式不正确（示例：adminUser::list）' }]}>
          <Input placeholder="根据接口路径自动生成" disabled />
        </Form.Item>
        <Form.Item label="接口路径" name="path" rules={[{ required: true, message: '请输入接口路径' }]}>
          <Input placeholder="示例：/admin/user/list" allowClear onChange={(e) => onPathChange(e.target.value)} />
        </Form.Item>
        <Form.Item label="描述" name="describe" rules={[{ max: 200, message: '描述长度不能超过200个字符' }]}>
          <Input.TextArea placeholder="请输入接口描述" rows={3} />
        </Form.Item>
      </Form>
    </Modal>
  );
}
