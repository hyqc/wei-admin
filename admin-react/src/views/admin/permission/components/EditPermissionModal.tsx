import { useEffect, useState } from 'react';
import { Form, Input, Modal, Radio, Switch, message } from 'antd';
import { DEFAULT_PERMISSION_TYPES, handleKey } from './common';
import { editAdminPermission, getAdminPermissionInfo } from '@/api/admin/permission';
import { AdminPerssionKey } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { PermissionListItem } from '@/types/admin_permission';

interface EditPermissionModalProps {
  open: boolean;
  detailData?: PermissionListItem;
  onClose: () => void;
  onNotice: () => void;
}

/** 编辑权限（切换类型时唯一键跟随菜单路径+类型重新生成） */
export default function EditPermissionModal({ open, detailData, onClose, onNotice }: EditPermissionModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);
  /** 服务端返回的菜单名称与路径 */
  const [detailMenuName, setDetailMenuName] = useState('');
  const [detailPath, setDetailPath] = useState('');

  // 打开时实时拉取详情回填，避免编辑列表中的过期数据
  useEffect(() => {
    if (open && detailData?.id) {
      getAdminPermissionInfo({ id: detailData.id }).then((res) => {
        setDetailMenuName(res.data.menuName || '');
        setDetailPath(res.data.menuPath || '');
        form.setFieldsValue({
          name: res.data.name || '',
          type: res.data.type || 'view',
          key: res.data.key || '',
          describe: res.data.describe || '',
          enabled: res.data.isEnabled ?? true,
        });
      });
    }
  }, [open, detailData, form]);

  const onTypeChange = () => {
    const menuPath = detailPath || detailData?.menuPath || '';
    if (menuPath) {
      const type = form.getFieldValue('type');
      form.setFieldValue('key', handleKey(menuPath, type));
    }
  };

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await editAdminPermission({
        id: detailData?.id,
        menuId: detailData?.menuId,
        name: values.name,
        type: values.type,
        key: values.key,
        describe: values.describe,
        enabled: values.enabled,
      });
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal open={open} title="编辑权限" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 14 }}>
        <Form.Item label="菜单名称">
          <Input value={detailMenuName} disabled />
        </Form.Item>
        <Form.Item label="权限名称" name="name" rules={[{ required: true, message: '请添加权限名称' }, { max: 50, message: '名称长度不能超过50个字符' }]}>
          <Input placeholder="请输入权限名称" allowClear />
        </Form.Item>
        <Form.Item label="权限类型" name="type">
          <Radio.Group onChange={onTypeChange} optionType="button" options={DEFAULT_PERMISSION_TYPES.map((item) => ({ key: item.key, value: item.key, label: item.name }))} />
        </Form.Item>
        <Form.Item label="唯一键" name="key" rules={[{ required: true, pattern: AdminPerssionKey, message: '请按照驼峰法命名' }]}>
          <Input disabled />
        </Form.Item>
        <Form.Item label="权限描述" name="describe">
          <Input.TextArea placeholder="请输入权限描述" rows={3} />
        </Form.Item>
        <Form.Item label="是否启用" name="enabled" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="禁用" />
        </Form.Item>
      </Form>
    </Modal>
  );
}
