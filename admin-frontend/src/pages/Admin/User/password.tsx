import {
  adminUserEditPwd,
} from '@/services/apis/admin/user';
import { Form, Input, message, Modal } from 'antd';
import { useEffect, useState } from 'react';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import { AdminUserFormRules } from './common';
import { ReqAdminUserEditPassword } from '@/proto/admin_ts/admin_user';
import { AdminUserListItem } from '@/proto/admin_ts/common';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type AdminUserEditPasswordModalPropsType = {
  modalStatus: boolean;
  detailData: AdminUserListItem;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const Password: React.FC<AdminUserEditPasswordModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { modalStatus, detailData, noticeModal } = props;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);

  const rules = AdminUserFormRules(form);

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        const data: ReqAdminUserEditPassword = {
          adminId: detailData.adminId,
          password: values.password,
          confirmPassword: values.confirmPassword,
        };
        adminUserEditPwd(data).then((res) => {
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
  }, []);

  return (
    (<Modal
      forceRender
      title="修改密码"
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
        <Form.Item label="账号" name="username">
          <Input disabled />
        </Form.Item>
        <Form.Item label="密码" name="password" initialValue={''} rules={rules.password}>
          <Input.Password />
        </Form.Item>
        <Form.Item
          label="确认密码"
          name="confirmPassword"
          dependencies={['password']}
          initialValue={''}
          rules={rules.confirmPassword}
        >
          <Input.Password />
        </Form.Item>
      </Form>
    </Modal>)
  );
};

export default Password;
