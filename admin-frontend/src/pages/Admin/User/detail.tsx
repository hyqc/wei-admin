import { ResponseAdminUserDetailType } from '@/services/apis/admin/user';
import { INPUT_STYLE } from '@/services/apis/config';
import { Drawer, Form, Input, Switch, Tag } from 'antd';

// import 'antd/es/modal/style';


// import 'antd/es/slider/style';


import Avatar from 'antd/lib/avatar/avatar';
import { useEffect } from 'react';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type DetailModalPropsType = {
  modalStatus: boolean;
  detailData: ResponseAdminUserDetailType;
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
      title="管理员详情"
      footer={null}
      width={DefaultDrawerWidth}
      destroyOnClose={true}
      getContainer={false}
      open={modalStatus}
      onClose={onClose}
    >
      <Form form={form} labelAlign="left" labelCol={{ span: 5 }} wrapperCol={{ span: 12 }}>
        <Form.Item label="ID" name="adminId">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="账号" name="username">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="头像" name="avatar" initialValue={true}>
          <Avatar size={64} src={detailData?.avatar} />
        </Form.Item>
        <Form.Item label="昵称" name="nickname">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="邮箱" name="email">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="角色" name="roles">
          <span>
            {detailData?.adminId === AdminId ? (
              <Tag color="blue">超管</Tag>
            ) : (
              detailData?.roles?.map((item) => {
                return (
                  <Tag key={item.roleId} color="blue">
                    {item.roleName}
                  </Tag>
                );
              })
            )}
          </span>
        </Form.Item>
        <Form.Item label="状态" name="enabled" valuePropName="checked">
          <Switch disabled checkedChildren={'启用'} unCheckedChildren={'禁用'} />
        </Form.Item>
        <Form.Item label="登录次数" name="totalLogin">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="最后登录IP" name="lastLoginIp">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="最后登录时间" name="lastLoginTime">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="创建时间" name="createTime">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
        <Form.Item label="更新时间" name="modifyTime">
          <Input disabled rootStyle={inputStyle} />
        </Form.Item>
      </Form>
    </Drawer>)
  );
};

export default DetailModal;
