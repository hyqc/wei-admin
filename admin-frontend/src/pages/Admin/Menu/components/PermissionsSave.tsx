import {
  adminAddMenuPermission,
} from '@/services/apis/admin/permission';
import { App, Form, Input, Modal, Space, Switch } from 'antd';
import { useEffect, useState } from 'react';
import { PERMIDDION_RULES } from '../../Permission/components/common';
import { MenuPermissionItem, RespAdminMenuPermissionsData } from '@/proto/admin_ts/admin_menu';
import { ReqPermissionBindMenu } from '@/proto/admin_ts/admin_permission';
import { AdminMenuModel } from '@/proto/admin_ts/admin_model';
import { set } from 'lodash';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type ModalPropsType = {
  menuInfo: AdminMenuModel;
  permissions: MenuPermissionItem[];
  modalStatus: boolean;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const AddPermissionsModal: React.FC<ModalPropsType> = (props) => {
  const { message } = App.useApp()
  const [form] = Form.useForm();
  const { menuInfo, permissions, modalStatus, noticeModal } = props;

  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);
  const [savePermissions, setSavePermissions] = useState<MenuPermissionItem[]>([]);
  const rules = PERMIDDION_RULES;

  useEffect(() => {
    setSavePermissions([...permissions])

    const formData: { [key: string]: any } = {}
    permissions.map((item, index) => {
      formData[`name[${index}]`] = item.name ?? ''
      formData[`key[${index}]`] = item.key ?? ''
      formData[`enabled[${index}]`] = item.enabled ?? false
    })
    form.setFieldsValue(formData)

  }, [menuInfo, permissions])


  function handleFormValues(values: any) {
    return savePermissions?.map((item: MenuPermissionItem, index: number) => {
      let tmp: MenuPermissionItem = { ...item };
      tmp.name = values[`name[${index}]`];
      tmp.key = values[`key[${index}]`];
      tmp.enabled = values[`enabled[${index}]`];
      return tmp;
    });
  }

  function handleCancel() {
    form.resetFields();
    noticeModal({ reload: false });
  }

  function handleOk() {
    setConfirmLoading(true);
    form.validateFields()
      .then((values) => {
        const data: ReqPermissionBindMenu = {
          menuId: menuInfo?.id,
          permissions: handleFormValues(values),
        };
        adminAddMenuPermission(data).then((res) => {
          message.success(res.msg, MessageDuritain, () => {
            noticeModal({ reload: true });
            form.resetFields();
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

  return (
    <Modal
      forceRender
      title="添加菜单操作权限"
      width={900}
      destroyOnClose={true}
      getContainer={false}
      maskClosable={false}
      open={modalStatus}
      confirmLoading={confirmLoading}
      onOk={handleOk}
      onCancel={handleCancel}
      okText="保存"
      cancelText="取消"
    >
      <Form form={form} labelCol={{ span: 3 }} wrapperCol={{ span: 20 }}>
        <Form.Item label="菜单">
          <Input disabled value={menuInfo?.name} style={{ width: 260 }} />
        </Form.Item>
        <Form.Item label="权限">
          {[...permissions].map((nv, index) => {
            const item = { ...nv }
            return (
              <Form.Item label={item.typeName} key={index}>
                <Space>
                  <Form.Item
                    label="名称"
                    name={`name[${index}]`}
                    rules={rules.view.name}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="唯一键"
                    name={`key[${index}]`}
                    rules={rules.view.key}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="状态"
                    name={`enabled[${index}]`}
                    valuePropName="checked"
                  >
                    <Switch
                      checkedChildren={'启用'}
                      unCheckedChildren={'禁用'}
                      defaultChecked={item.enabled}
                    />
                  </Form.Item>
                </Space>
              </Form.Item>
            );
          })}
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default AddPermissionsModal;
