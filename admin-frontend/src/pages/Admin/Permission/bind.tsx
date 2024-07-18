import {
  adminPermissionBindApi,
  RequestAdminPermissionBindApiParamsType,
  ResponseAdminPermissionDetailType,
} from '@/services/apis/admin/permission';
import { Form, Input, message, Modal, Select } from 'antd';
import { useEffect, useState } from 'react';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import { adminAPIAll, ResponseAdminAPIAllItemType } from '@/services/apis/admin/resource';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type BindModalPropsType = {
  modalStatus: boolean;
  detailData: ResponseAdminPermissionDetailType;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const BindModal: React.FC<BindModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { modalStatus, detailData, noticeModal } = props;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);
  const [apiOptions, setAPIOptions] = useState<ResponseAdminAPIAllItemType[]>([]);

  const apiIdsValue: number[] =
    detailData?.apis?.map((item: ResponseAdminAPIAllItemType) => {
      return item.id;
    }) || [];

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        const data: RequestAdminPermissionBindApiParamsType = {
          permissionId: detailData.id,
          apiIds: values.apiIds || [],
        };
        adminPermissionBindApi(data)
          .then((res) => {
            message.destroy();
            message.success(res.message, MessageDuritain, () => {
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

  function fetchAdminAPIs() {
    adminAPIAll().then((res) => {
      setAPIOptions(res.data || []);
    });
  }

  useEffect(() => {
    form.resetFields();
    if (detailData) {
      fetchAdminAPIs();
      form.setFieldsValue({ name: detailData?.name, apiIds: apiIdsValue });
    }
  }, [detailData]);

  return (
    (<Modal
      forceRender
      title="绑定接口"
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
        <Form.Item label="权限名称" name="name">
          <Input disabled />
        </Form.Item>
        <Form.Item label="接口" name="apiIds">
          <Select
            showSearch
            mode={'multiple'}
            filterOption={(input, option) => {
              return (option!.children as unknown as string)
                .toLowerCase()
                .includes(input.toLowerCase());
            }}
          >
            {apiOptions?.map((item) => {
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
