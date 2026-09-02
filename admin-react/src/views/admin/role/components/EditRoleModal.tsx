import { useEffect, useState } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { editAdminRole, getAdminRoleInfo } from '@/api/admin/role';
import { DefaultModalWidth } from '@/api/config';
import type { RoleItem } from '@/types/admin_role';

interface EditRoleModalProps {
  open: boolean;
  detailData?: RoleItem;
  onClose: () => void;
  onNotice: () => void;
}

/** 编辑角色：超管角色仅允许修改描述 */
export default function EditRoleModal({ open, detailData, onClose, onNotice }: EditRoleModalProps) {
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);
  /** 超管角色仅允许修改描述（以服务端最新数据为准） */
  const [isSuperAdmin, setIsSuperAdmin] = useState(false);

  // 打开时实时拉取详情回填，避免编辑列表中的过期数据
  useEffect(() => {
    if (open && detailData?.id) {
      getAdminRoleInfo({ id: detailData.id }).then((res) => {
        form.setFieldsValue({ name: res.data.name || '', describe: res.data.describe || '' });
        setIsSuperAdmin(!!res.data.isSuperAdmin);
      });
    }
  }, [open, detailData, form]);

  const handleOk = async () => {
    const values = await form.validateFields();
    setConfirmLoading(true);
    try {
      const res = await editAdminRole({
        id: detailData?.id as number,
        name: values.name,
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
    <Modal open={open} title="编辑角色" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 12 }}>
        <Form.Item
          label="角色名称"
          name="name"
          rules={[
            { required: true, message: '请输入角色名称' },
            { max: 50, message: '角色名称长度不能超过50个字符' },
          ]}
        >
          <Input placeholder="请输入角色名称" allowClear disabled={isSuperAdmin} />
        </Form.Item>
        {isSuperAdmin && <div className="form-tip" style={{ marginTop: -18, marginBottom: 16 }}>超级管理员角色名称不可修改，仅可修改描述</div>}
        <Form.Item label="描述" name="describe" rules={[{ max: 200, message: '描述长度不能超过200个字符' }]}>
          <Input.TextArea placeholder="请输入角色描述" rows={4} />
        </Form.Item>
      </Form>
    </Modal>
  );
}
