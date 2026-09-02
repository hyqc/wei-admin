import { useEffect, useState } from 'react';
import { Col, Form, Input, Modal, Radio, Row, Switch, message } from 'antd';
import PageMenus from './PageMenus';
import { DEFAULT_PERMISSION_TYPES, handleKey } from './common';
import { addAdminPermission } from '@/api/admin/permission';
import type { MenuTreeItem } from '@/types/admin_menu';

interface AddPermissionModalProps {
  open: boolean;
  onClose: () => void;
  onNotice: () => void;
}

/** 新建权限（唯一键由菜单路径+类型自动生成） */
export default function AddPermissionModal({ open, onClose, onNotice }: AddPermissionModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [menuId, setMenuId] = useState<number>();
  const [menuName, setMenuName] = useState('');
  const [menuPath, setMenuPath] = useState('');

  useEffect(() => {
    if (open) {
      form.resetFields();
      setMenuId(undefined);
      setMenuName('');
      setMenuPath('');
    }
  }, [open, form]);

  // 选择菜单后：记录名称与路径，并生成唯一键
  const onMenuChange = (node?: MenuTreeItem) => {
    setMenuId(node?.id);
    if (node) {
      const path = node.path || '';
      setMenuName(node.name || '');
      setMenuPath(path);
      const type = form.getFieldValue('type') || 'view';
      form.setFieldValue('key', handleKey(path, type));
    } else {
      setMenuName('');
      setMenuPath('');
      form.setFieldValue('key', '');
    }
  };

  // 切换权限类型时，唯一键跟随菜单路径+类型重新生成
  const onTypeChange = () => {
    if (menuId) {
      const type = form.getFieldValue('type');
      form.setFieldValue('key', handleKey(menuPath, type));
    }
  };

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await addAdminPermission({
        menuId,
        menuName,
        menuPath,
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
    <Modal open={open} title="新建权限" width={900} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Row gutter={16}>
        <Col span={10}>
          <div style={{ marginBottom: 8, fontWeight: 600 }}>选择菜单页面</div>
          <div style={{ maxHeight: 360, overflow: 'auto', border: '1px solid #f0f0f0', borderRadius: 4, padding: 8 }}>
            <PageMenus key={open ? 'open' : 'closed'} value={menuId} onChange={onMenuChange} />
          </div>
        </Col>
        <Col span={14}>
          <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 16 }} initialValues={{ type: 'view', enabled: true, describe: '' }}>
            <Form.Item label="权限名称" name="name" rules={[{ required: true, message: '请添加权限名称' }, { max: 50, message: '名称长度不能超过50个字符' }]}>
              <Input placeholder="请输入权限名称" allowClear />
            </Form.Item>
            <Form.Item label="权限类型" name="type">
              <Radio.Group onChange={onTypeChange} optionType="button" options={DEFAULT_PERMISSION_TYPES.map((item) => ({ key: item.key, value: item.key, label: item.name }))} />
            </Form.Item>
            <Form.Item label="唯一键" name="key" rules={[{ required: true, pattern: /^([A-Z][a-zA-Z0-9]*)+$/g, message: '请按照驼峰法命名' }]}>
              <Input disabled />
            </Form.Item>
            <Form.Item label="权限描述" name="describe">
              <Input.TextArea placeholder="请输入权限描述" rows={3} />
            </Form.Item>
            <Form.Item label="是否启用" name="enabled" valuePropName="checked">
              <Switch checkedChildren="启用" unCheckedChildren="禁用" />
            </Form.Item>
          </Form>
        </Col>
      </Row>
    </Modal>
  );
}
