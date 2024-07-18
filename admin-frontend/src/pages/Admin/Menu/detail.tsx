import { Drawer, Form, Input, Switch } from 'antd';
import { useEffect } from 'react';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import { INPUT_STYLE } from '@/services/apis/config';
import { ResponseAdminMenuDetailType } from '@/services/apis/admin/menu';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type DetailModalPropsType = {
  modalStatus: boolean;
  detailData: ResponseAdminMenuDetailType;
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
    form.setFieldsValue(detailData);
  });

  return (
    (<Drawer
      forceRender
      title="菜单详情"
      footer={null}
      width={DefaultDrawerWidth}
      destroyOnClose={true}
      getContainer={false}
      open={modalStatus}
      onClose={onClose}
    >
      <Form form={form} labelAlign="left" labelCol={{ span: 6 }} wrapperCol={{ span: 12 }}>
        <Form.Item label="菜单ID" name="id">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="名称" name="name">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="路由" name="path">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="唯一键名" name="key">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="图标" name="icon">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="排序值" name="sort">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="重定向路由" name="redirect">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="菜单中隐藏" name="hideInMenu" valuePropName="checked">
          <Switch disabled checkedChildren={'隐藏'} unCheckedChildren={'显示'} />
        </Form.Item>
        <Form.Item label="隐藏子菜单" name="hideChildrenInMenu" valuePropName="checked">
          <Switch disabled checkedChildren={'隐藏'} unCheckedChildren={'显示'} />
        </Form.Item>
        <Form.Item label="状态" name="enabled" valuePropName="checked">
          <Switch disabled checkedChildren={'启用'} unCheckedChildren={'禁用'} />
        </Form.Item>
        <Form.Item label="创建时间" name="createTime">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="最后更新时间" name="modifyTime">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
      </Form>
    </Drawer>)
  );
};

export default DetailModal;
