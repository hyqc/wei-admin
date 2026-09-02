import { useEffect, useState } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { editAdminApi, getAdminApiInfo } from '@/api/admin/api';
import { AdminAPIKey } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import { generateApiKeyByPath } from '@/utils/apiKey';
import type { AdminApiItem } from '@/types/common';

interface EditApiModalProps {
  open: boolean;
  detailData?: AdminApiItem;
  onClose: () => void;
  onNotice: () => void;
}

/** 编辑接口（唯一键只读，随路径自动生成） */
export default function EditApiModal({ open, detailData, onClose, onNotice }: EditApiModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);

  // 打开时实时拉取详情回填，避免编辑列表中的过期数据
  useEffect(() => {
    if (open && detailData?.id) {
      getAdminApiInfo({ id: detailData.id }).then((res) => {
        form.setFieldsValue({
          name: res.data.name || '',
          key: res.data.key || '',
          path: res.data.path || '',
          describe: res.data.describe || '',
        });
      });
    }
  }, [open, detailData, form]);

  // 路径变化后按路径规则自动生成唯一键（回填不触发该回调）
  const onPathChange = (path: string) => {
    form.setFieldValue('key', generateApiKeyByPath(path));
  };

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await editAdminApi({
        id: detailData?.id as number,
        name: values.name,
        key: values.key,
        path: values.path,
        describe: values.describe,
      });
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal open={open} title="编辑接口" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
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
