import { Drawer, Form, Input, Switch } from 'antd';
import { useEffect } from 'react';
import { INPUT_STYLE } from '@/services/apis/config';
import { ResponseAdminMenuModeTypeData } from '@/services/apis/admin/menu';
import BindPermissions from './components/BindPermissions';
import { RespAdminRoleInfoData } from '@/proto/admin_ts/admin_role';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type DetailModalPropsType = {
  modalStatus: boolean;
  detailData: RespAdminRoleInfoData;
  menuPageData: ResponseAdminMenuModeTypeData[];
  noticeModal: (data: NoticeModalPropsType) => void;
};
const inputStyle = INPUT_STYLE;

const DetailModal: React.FC<DetailModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { modalStatus, detailData, menuPageData, noticeModal } = props;

  function onClose() {
    noticeModal({ reload: false });
  }

  useEffect(() => {
    form.setFieldsValue(detailData);
  });

  return (
    (<Drawer
      forceRender
      title="角色详情"
      footer={null}
      width={760}
      destroyOnClose={true}
      getContainer={false}
      open={modalStatus}
      onClose={onClose}
    >
      <Form form={form} labelCol={{ span: 4 }} wrapperCol={{ span: 20 }}>
        <Form.Item label="名称" name="name">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input.TextArea disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="创建时间" name="createTime">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="最后更新时间" name="modifyTime">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="状态" name="enabled" valuePropName="checked">
          <Switch disabled checkedChildren={'启用'} unCheckedChildren={'禁用'} />
        </Form.Item>
        <Form.Item label="权限" name="permissionIds">
          <BindPermissions
            disabled
            datasource={menuPageData}
            permissionIds={detailData?.permissionIds}
          />
        </Form.Item>
      </Form>
    </Drawer>)
  );
};

export default DetailModal;
