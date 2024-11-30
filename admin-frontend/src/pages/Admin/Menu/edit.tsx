import { MenuTreeItem, ReqAdminMenuEdit } from '@/proto/admin_ts/admin_menu';
import { AdminMenuModel } from '@/proto/admin_ts/admin_model';
import { adminMenuAdd, adminMenuEdit, adminMenuTree } from '@/services/apis/admin/menu';
import { App, Button, Col, Form, Input, Modal, Row, Switch } from 'antd';
import { Gutter } from 'antd/lib/grid/row';
import { useEffect, useState } from 'react';
import MenuTreeSelect from './components/MenuTreeSelect';
import { FORM_RULES, makeMenuSpreadTreeData, path2UpperCamelCase } from './components/common';

const ButtonStyles = { marginRight: '2rem' };
const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];

type NoticeSaveModalPropsType = {
  reload?: boolean;
};

type SaveModalPropsType = {
  modalStatus: boolean;
  noticeModal: (data: NoticeSaveModalPropsType) => void;
  detailData?: AdminMenuModel;
};

const EditModal: React.FC<SaveModalPropsType> = (props) => {
  const { message } = App.useApp();
  const [form] = Form.useForm();
  const { modalStatus, noticeModal, detailData } = props;

  const [confirmLoading, setConfirmLoading] = useState<boolean>(false);
  const [menuData, setMenuData] = useState<MenuTreeItem[]>([]);

  const modalTitle = detailData?.id ? '编辑菜单' : '新建菜单'
  const rules: any = FORM_RULES;

  useEffect(() => {
    adminMenuTree().then((res) => {
      const spreadTreeData = makeMenuSpreadTreeData(res.data, detailData?.id);
      setMenuData(spreadTreeData);
    });
  }, []);

  useEffect(() => {
    form.setFieldsValue(detailData);
  }, [detailData]);

  function handleOk() {
    if(form.getFieldValue('parentId')===undefined){
      form.setFieldValue('parentId', 0)
    }
    form.validateFields()
      .then((values) => {
        console.log(values);
        const data: ReqAdminMenuEdit = {
          ...values
        };
        if(values.id){
          adminMenuEdit(data).then((res) => {
            message.success(res.msg, MessageDuritain, () => {
              noticeModal({ reload: true });
              form.resetFields();
            });
          });
        }else{
          adminMenuAdd(data).then((res) => {
            message.success(res.msg, MessageDuritain, () => {
              noticeModal({ reload: true });
              form.resetFields();
            });
          });
        }
        
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setConfirmLoading(false);
      });
  }

  function handleCancel() {
    form.resetFields();
    noticeModal({ reload: false });
  }

  function onChangePath() {
    form.setFieldsValue({ key: path2UpperCamelCase(form.getFieldValue('path')) });
  }

  return (
    <Modal
      forceRender
      title={modalTitle}
      width={DefaultModalWidth}
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
      <Form
        form={form}
        labelCol={{ span: 6 }}
        wrapperCol={{ span: 16 }}
        labelAlign="left"
        labelWrap
      >
        <Row gutter={FormSearchRowGutter}>
          <Col span={12}>
            <Form.Item label="ID" name="id" hidden>
              <Input disabled />
            </Form.Item>
            <Form.Item label="父级菜单" name="parentId" rules={rules.parentId}>
              <MenuTreeSelect data={menuData} />
            </Form.Item>
            <Form.Item label="名称" name="name" rules={rules.name}>
              <Input />
            </Form.Item>
            <Form.Item label="路由" name="path" rules={rules.path}>
              <Input onChange={onChangePath} />
            </Form.Item>
            <Form.Item label="唯一键名" name="key" rules={rules.key} >
              <Input placeholder="示例：AdminMenu"/>
            </Form.Item>
            <Form.Item label="排序值" name="sort" initialValue={0} rules={rules.sort}>
              <Input min={0} />
            </Form.Item>
            <Form.Item label="图标" name="icon">
              <Input />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item label="重定向地址" name="redirect">
              <Input />
            </Form.Item>
            <Form.Item label="描述" name="describe" initialValue={''}>
              <Input.TextArea />
            </Form.Item>
            <Form.Item label="菜单中隐藏" name="isHideInMenu" valuePropName="checked">
              <Switch checkedChildren={'隐藏'} unCheckedChildren={'显示'} />
            </Form.Item>
            <Form.Item label="隐藏子菜单" name="isHideChildrenInMenu" valuePropName="checked">
              <Switch checkedChildren={'隐藏'} unCheckedChildren={'显示'} />
            </Form.Item>
            <Form.Item label="状态" name="isEnabled" valuePropName="checked">
              <Switch checkedChildren={'启用'} unCheckedChildren={'禁用'} />
            </Form.Item>
          </Col>
        </Row>
      </Form>
    </Modal>
  );
};

export default EditModal;
