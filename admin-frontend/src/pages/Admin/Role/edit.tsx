import { RespAdminRoleInfoData } from '@/proto/admin_ts/admin_role';
import { adminRoleEdit } from '@/services/apis/admin/role';
import { App, Form, Input, Modal, Switch } from 'antd';
import { useEffect, useState } from 'react';


export type NoticeModalPropsType = {
  reload?: boolean;
};

export type EditModalPropsType = {
  modalStatus: boolean;
  detailData: RespAdminRoleInfoData;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const EditModal: React.FC<EditModalPropsType> = (props) => {
  const { message } = App.useApp();
  const [form] = Form.useForm();
  const { modalStatus, detailData, noticeModal } = props;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);

  const rules: any = {
    roleName: [{ required: true, type: 'string', message: '请输入角色名称!' }],
  };

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        adminRoleEdit(values).then((res) => {
          message.success(res.msg, MessageDuritain, () => {
            noticeModal({ reload: true });
          });
        });
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setConfirmLoading(false);
      });
  }

  function handleCancel() {
    noticeModal({ reload: false });
  }

  useEffect(() => {
    form.setFieldsValue(detailData);
  }, [detailData]);

  return (
    (<Modal
      forceRender
      title="编辑角色"
      width={DefaultModalWidth}
      destroyOnClose={true}
      maskClosable={false}
      getContainer={false}
      open={modalStatus}
      confirmLoading={confirmLoading}
      onOk={handleOk}
      onCancel={handleCancel}
      okText="保存"
      cancelText="取消"
    >
      <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 12 }}>
        <Form.Item label="ID" name="id">
          <Input disabled />
        </Form.Item>
        <Form.Item label="名称" name="name" rules={rules.name}>
          <Input />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input.TextArea />
        </Form.Item>
        <Form.Item label="状态" name="isEnabled" valuePropName="checked">
          <Switch checkedChildren={'启用'} unCheckedChildren={'禁用'} />
        </Form.Item>
      </Form>
    </Modal>)
  );
};

export default EditModal;
