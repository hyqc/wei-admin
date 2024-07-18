import { adminAPIEdit, ResponseAdminAPIDetailType } from '@/services/apis/admin/resource';
import { Form, Input, message, Modal, Switch } from 'antd';
import { ChangeEvent, useEffect, useState } from 'react';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import { DEFAULT_RULES, path2UpperCamelCase } from './components/common';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type EditModalPropsType = {
  modalStatus: boolean;
  detailData: ResponseAdminAPIDetailType;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const EditModal: React.FC<EditModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { modalStatus, detailData, noticeModal } = props;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);

  const rules: any = DEFAULT_RULES;

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        adminAPIEdit(values).then((res) => {
          message.success(res.message, MessageDuritain, () => {
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

  function onChangePath(e: ChangeEvent) {
    form.setFieldsValue({ key: path2UpperCamelCase(form.getFieldValue('path')) });
  }

  useEffect(() => {
    form.setFieldsValue(detailData);
  }, [detailData]);

  return (
    (<Modal
      forceRender
      title="编辑接口"
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
        <Form.Item label="ID" name="id" hidden>
          <Input disabled />
        </Form.Item>
        <Form.Item label="名称" name="name" rules={rules.name}>
          <Input />
        </Form.Item>
        <Form.Item label="路由" name="path" rules={rules.path}>
          <Input onChange={onChangePath} />
        </Form.Item>
        <Form.Item label="键名" name="key" rules={rules.key}>
          <Input />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input.TextArea />
        </Form.Item>
        <Form.Item label="状态" name="enabled" valuePropName="checked">
          <Switch checkedChildren={'启用'} unCheckedChildren={'禁用'} defaultChecked />
        </Form.Item>
      </Form>
    </Modal>)
  );
};

export default EditModal;
