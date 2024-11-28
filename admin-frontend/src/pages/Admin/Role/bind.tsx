import {
  adminRoleBind,
} from '@/services/apis/admin/role';
import { Button, Drawer, Form, Input, message, Popconfirm } from 'antd';
import { useEffect, useState } from 'react';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import { INPUT_STYLE } from '@/services/apis/config';
import BindPermissions from './components/BindPermissions';
import { ResponseAdminMenuModeTypeData } from '@/services/apis/admin/menu';
import { RespAdminRoleInfoData } from '@/proto/admin_ts/admin_role';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type BindModalPropsType = {
  modalStatus: boolean;
  detailData: RespAdminRoleInfoData;
  menuPageData: ResponseAdminMenuModeTypeData[];
  noticeModal: (data: NoticeModalPropsType) => void;
};
const inputStyle = INPUT_STYLE;
const ButtonStyles = { marginRight: '2rem' };
const BindModal: React.FC<BindModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { modalStatus, detailData, menuPageData, noticeModal } = props;
  const [saveBtnLoading, setSaveBtnLoading] = useState<boolean>(false);

  function onClose() {
    noticeModal({ reload: false });
    setSaveBtnLoading(false);
    form.resetFields();
  }

  function handleOk() {
    setSaveBtnLoading(true);
    form
      .validateFields()
      .then((values) => {
        const data: RequestAdminRoleBindPermissionsParamsType = {
          permissionIds: form.getFieldValue('permissionIds'),
          id: detailData.id,
        };
        console.log('value', values);
        adminRoleBind(data).then((res) => {
          message.success(res.message, MessageDuritain, () => {
            noticeModal({ reload: true });
            form.resetFields();
          });
        });
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setSaveBtnLoading(false);
      });
  }

  useEffect(() => {
    console.log(detailData);
    form.setFieldsValue(detailData);
  }, [form]);

  return (
    (<Drawer
      forceRender
      mask={false}
      maskClosable={false}
      title="分配权限"
      width={700}
      destroyOnClose={true}
      getContainer={false}
      open={modalStatus}
      onClose={onClose}
    >
      <Form form={form} labelCol={{ span: 2 }} wrapperCol={{ span: 22 }}>
        <Form.Item label="ID" name="id" hidden>
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="名称">{detailData.name}</Form.Item>
        <Form.Item label="权限" name="permissionIds">
          <BindPermissions datasource={menuPageData} permissionIds={detailData?.permissionIds} />
        </Form.Item>
        <Form.Item rootStyle={{ marginTop: '8rem', marginLeft: '4rem' }}>
          <Button onClick={onClose} rootStyle={ButtonStyles}>
            取消
          </Button>
          <Popconfirm
            title="确定要删除该角色吗？"
            okText="确定"
            cancelText="取消"
            onConfirm={handleOk}
          >
            <Button loading={saveBtnLoading} type="primary" rootStyle={ButtonStyles}>
              保存
            </Button>
          </Popconfirm>
        </Form.Item>
      </Form>
    </Drawer>)
  );
};

export default BindModal;
