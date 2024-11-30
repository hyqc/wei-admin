import {
  adminUserBindRoles,
} from '@/services/apis/admin/user';
import { App, Form, Input, Modal, Select } from 'antd';
import { useEffect, useState } from 'react';
import { adminRoleAll } from '@/services/apis/admin/role';
import { AdminUserListItem, AdminUserRoleItem } from '@/proto/admin_ts/common';
import { RoleItem } from '@/proto/admin_ts/admin_role';
import { ReqAdminUserBindRoles } from '@/proto/admin_ts/admin_user';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type BindModalPropsType = {
  modalStatus: boolean;
  detailData: AdminUserListItem;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const BindModal: React.FC<BindModalPropsType> = (props) => {
  const { message } = App.useApp();
  const [form] = Form.useForm();
  const { modalStatus, detailData, noticeModal } = props;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);
  const [roleOptions, setRoleOptions] = useState<RoleItem[]>([]);

  const roleIdsValue: number[] = detailData?.roles?.map((item: AdminUserRoleItem) => {
    return item.roleId ?? 0;
  }).filter(roleId => roleId > 0) ?? [];

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        const data: ReqAdminUserBindRoles = {
          adminId: detailData.adminId,
          roleIds: values.roleIds || [],
        };
        adminUserBindRoles(data)
          .then((res) => {
            message.destroy();
            message.success(res.msg, MessageDuritain, () => {
              noticeModal({ reload: true });
            });
          })
          .catch((err) => {
            console.log(err);
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

  function fetchAdminRoles(name?: string) {
    adminRoleAll({ name }).then((res) => {
      setRoleOptions(res.data || []);
    });
  }

  useEffect(() => {
    form.resetFields();
    if (detailData && detailData.username) {
      fetchAdminRoles();
      form.setFieldsValue({ username: detailData?.username, roleIds: roleIdsValue });
    }
    return () => { };
  }, [detailData]);

  return (
    (<Modal
      forceRender
      title="分配角色"
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
        <Form.Item label="ID" name="adminId" hidden>
          <Input disabled value={detailData.adminId} />
        </Form.Item>
        <Form.Item label="账号" name="username">
          <Input disabled />
        </Form.Item>
        <Form.Item label="角色" name="roleIds">
          <Select
            showSearch
            mode={'multiple'}
            filterOption={(input, option) => {
              return (option!.children as unknown as string)
                .toLowerCase()
                .includes(input.toLowerCase());
            }}
          >
            {roleOptions?.map((item) => {
              return (
                <Select.Option key={item.id} value={item.id}>
                  {item.name}
                </Select.Option>
              );
            })}
          </Select>
        </Form.Item>
      </Form>
    </Modal>)
  );
};

export default BindModal;
