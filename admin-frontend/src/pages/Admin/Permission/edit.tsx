import {
  adminPermissionEdit,
} from '@/services/apis/admin/permission';
import { App, Form, Input, message, Modal, Select, Switch } from 'antd';
import { ChangeEvent, useEffect, useState } from 'react';
import { DEFAULT_RULES, path2UpperCamelCase } from './components/common';
import PageMenus from './components/PageMenus';
import { first2Upcase } from '@/utils/common';
import { AdminPermissionInfo } from '@/proto/admin_ts/admin_permission';
import { MenuTreeItem } from '@/proto/admin_ts/admin_menu';

export type NoticeModalPropsType = {
  reload?: boolean;
};

export type EditModalPropsType = {
  pageMenusData: MenuTreeItem[];
  modalStatus: boolean;
  detailData: AdminPermissionInfo;
  noticeModal: (data: NoticeModalPropsType) => void;
};

const EditModal: React.FC<EditModalPropsType> = (props) => {
  const {message} = App.useApp();
  const [form] = Form.useForm();
  const { pageMenusData, modalStatus, detailData, noticeModal } = props;
  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);
  const [menuCheckedItem, setMenuCheckedItem] = useState<MenuTreeItem | undefined>();
  const [pageMenusDataMap, setPageMenusDataMap] = useState<any>();
  const rules: any = DEFAULT_RULES;

  function handleOk() {
    setConfirmLoading(true);
    form
      .validateFields()
      .then((values) => {
        adminPermissionEdit(values).then((res) => {
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

  function onChangePath(e: ChangeEvent) {
    makePermissionKey(form.getFieldValue('path'));
  }

  function onChangeType(e: ChangeEvent) {
    makePermissionKey();
  }

  function menuIdChange(index: number) {
    if (pageMenusDataMap[index] !== undefined) {
      const menu = pageMenusDataMap[index];
      setMenuCheckedItem(menu);
      makePermissionKey(menu.path);
    }
  }

  function makePermissionKey(path?: string) {
    const type = form.getFieldValue('type');
    path = path ??  (menuCheckedItem?.path ?? '');
    const key = path2UpperCamelCase(path) + first2Upcase(type);
    form.setFieldsValue({ key });
  }

  useEffect(() => {
    let data: any = {};
    pageMenusData?.forEach((item) => {
      data[item.id ?? 0] = item;
    });
    setPageMenusDataMap(data);
  }, [pageMenusData]);

  useEffect(() => {
    form.setFieldsValue(detailData);
    if (detailData?.menuId) {
      menuIdChange(detailData.menuId);
    }
  }, [detailData, pageMenusData]);

  return (
    (<Modal
      forceRender
      title="编辑权限"
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
        <Form.Item label="权限ID" hidden name="id">
          <Input disabled />
        </Form.Item>
        <Form.Item label="菜单名称" name="menuId">
          <PageMenus data={pageMenusData} onChange={menuIdChange} />
        </Form.Item>
        <Form.Item label="菜单路由">
          <Input disabled value={menuCheckedItem?.path} />
        </Form.Item>
        <Form.Item label="名称" name="name" rules={rules.name}>
          <Input />
        </Form.Item>
        <Form.Item label="类型" name="type">
          <Select style={{ offset: 0, width: '160' }} onChange={onChangeType}>
            <Select.Option value="view">查看</Select.Option>
            <Select.Option value="edit">编辑</Select.Option>
            <Select.Option value="delete">删除</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item label="键名" name="key" rules={rules.key}>
          <Input disabled onChange={onChangePath} />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input.TextArea />
        </Form.Item>
        <Form.Item label="状态" name="isEnabled" valuePropName="checked">
          <Switch checkedChildren={'启用'} unCheckedChildren={'禁用'} />
        </Form.Item>
      </Form>
    </Modal>)
  );
};

export default EditModal;
