import { Drawer, Form, Input, Select, Switch } from 'antd';
import { useEffect } from 'react';
import { INPUT_STYLE } from '@/services/apis/config';
import { AdminPermissionInfo } from '@/proto/admin_ts/admin_permission';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type DetailModalPropsType = {
  modalStatus: boolean;
  detailData: AdminPermissionInfo;
  noticeModal: (data: NoticeModalPropsType) => void;
};
const inputStyle = INPUT_STYLE;

const DetailModal: React.FC<DetailModalPropsType> = (props) => {
  const [form] = Form.useForm();
  const { modalStatus, detailData, noticeModal } = props;

  function onClose() {
    noticeModal({ reload: false });
  }

  useEffect(() => {
    form.setFieldsValue({ ...detailData, isEnabled: detailData?.isEnabled ?? false });
  });

  return (
    (<Drawer
      forceRender
      title="接口详情"
      footer={null}
      width={DefaultDrawerWidth}
      destroyOnClose={true}
      getContainer={false}
      open={modalStatus}
      onClose={onClose}
    >
      <Form form={form} labelAlign="left" labelCol={{ span: 4 }} wrapperCol={{ span: 12 }}>
        <Form.Item label="菜单ID" name="menuId">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="菜单名称" name="menuName">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="菜单路由" name="menuPath">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="权限ID" name="id">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="名称" name="name">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="类型" name="type">
          <Select disabled style={{ offset: 0, width: '160' }}>
            <Select.Option value="view">查看</Select.Option>
            <Select.Option value="edit">编辑</Select.Option>
            <Select.Option value="delete">删除</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item label="键名" name="key">
          <Input disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input.TextArea disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="创建时间" name="createdAt">
          <Input.TextArea disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="更新时间" name="updatedAt">
          <Input.TextArea disabled style={inputStyle} />
        </Form.Item>
        <Form.Item label="状态" name="isEnabled" valuePropName="checked">
          <Switch disabled style={inputStyle} checkedChildren={'启用'} unCheckedChildren={'禁用'} />
        </Form.Item>
      </Form>
    </Drawer>)
  );
};

export default DetailModal;
