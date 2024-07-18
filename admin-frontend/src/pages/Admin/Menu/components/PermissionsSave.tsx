import { Form, Input, message, Modal, Space, Switch } from 'antd';
import { useEffect, useState } from 'react';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import {
  adminAddMenuPermission,
  RequestAdminPermissionAddForMenuParamsType,
  ResponseAdminMenuPermissionsItemType,
} from '@/services/apis/admin/permission';
import { PERMIDDION_RULES } from '../../Permission/components/common';
import { ResponseAdminMenuPermissionsType } from '@/services/apis/admin/menu';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type AddModalPropsType = {
  detailData: ResponseAdminMenuPermissionsType;
  modalStatus: boolean;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const AddPermissionsModal: React.FC<AddModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { detailData, modalStatus, noticeModal } = props;
  const { menu, permissions } = detailData;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);
  const [savePermissions] = useState<ResponseAdminMenuPermissionsItemType[]>(permissions);
  const rules = PERMIDDION_RULES;

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        const data: RequestAdminPermissionAddForMenuParamsType = {
          menuId: detailData.menu.id,
          permissions: handleFormValues(values),
        };
        adminAddMenuPermission(data).then((res) => {
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
        setConfirmLoading(false);
      });
  }

  function handleFormValues(values: any) {
    return permissions.map((item: ResponseAdminMenuPermissionsItemType, index: number) => {
      let tmp: ResponseAdminMenuPermissionsItemType = { ...item };
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

  useEffect(() => {
    console.log('detailData', detailData);
  }, [detailData]);

  return (
    (<Modal
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
          <Input disabled value={menu?.name} style={{ width: 260 }} />
        </Form.Item>
        <Form.Item label="权限">
          {savePermissions.map((item, index) => {
            return (
              <Form.Item label={item.typeText} name={index} key={index}>
                <Space>
                  <Form.Item
                    label="名称"
                    name={`name[${index}]`}
                    rules={rules.view.name}
                    initialValue={item.name}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="唯一键"
                    name={`key[${index}]`}
                    rules={rules.view.key}
                    initialValue={item.key}
                  >
                    <Input />
                  </Form.Item>
                  <Form.Item
                    label="状态"
                    name={`enabled[${index}]`}
                    initialValue={item.enabled}
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
    </Modal>)
  );
};

export default AddPermissionsModal;
