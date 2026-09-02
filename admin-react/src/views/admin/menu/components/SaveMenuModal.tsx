import { useEffect, useMemo, useState } from 'react';
import { Form, Input, InputNumber, Modal, Switch, TreeSelect, message } from 'antd';
import { addAdminMenu, editAdminMenu, getAdminMenuInfo } from '@/api/admin/menu';
import { AdminMenuKey, AdminRouterPath } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import IconSelect from '@/components/IconSelect';
import type { MenuTreeItem, ReqAdminMenuAdd } from '@/types/admin_menu';

/** 把路径转为键名（大驼峰） */
function path2UpperCamelCase(path: string) {
  return path
    ?.split('/')
    .filter((name) => name.length > 0)
    .map((name) => name[0].toUpperCase() + name.substring(1))
    .join('');
}

interface SaveMenuModalProps {
  open: boolean;
  tree: MenuTreeItem[];
  detailData?: MenuTreeItem;
  onClose: () => void;
  onNotice: () => void;
}

/** 新建/编辑菜单（键名由路径自动生成，不可编辑） */
export default function SaveMenuModal(props: SaveMenuModalProps) {
  const { open, tree, detailData, onClose, onNotice } = props;
  const [form] = Form.useForm();
  const [confirmLoading, setConfirmLoading] = useState(false);

  const treeSelectData = useMemo(() => {
    interface SelectNode {
      key: number;
      value: number;
      title: string;
      children?: SelectNode[];
    }
    const build = (list: MenuTreeItem[]): SelectNode[] =>
      list.map((item) => ({
        key: item.id as number,
        value: item.id as number,
        title: item.name ?? '',
        children: item.children?.length ? build(item.children) : undefined,
      }));
    return [{ key: 0, value: 0, title: '顶级菜单' }, ...build(tree)];
  }, [tree]);  useEffect(() => {
    if (open) {
      if (detailData?.id) {
        // 编辑：实时拉取详情回填，避免编辑列表中的过期数据
        getAdminMenuInfo({ menuId: detailData.id }).then((res) => {
          const info = res.data as MenuTreeItem;
          form.setFieldsValue({
            parentId: info.parentId || 0,
            name: info.name || '',
            path: info.path || '',
            key: info.key || '',
            sort: info.sort ?? 0,
            icon: info.icon || '',
            redirect: info.redirect || '/',
            describe: info.describe || '',
            isHideInMenu: !!info.hideInMenu,
            isHideChildrenInMenu: !!info.hideChildrenInMenu,
            isEnabled: !!info.enabled,
          });
        });
      } else {
        form.resetFields();
        form.setFieldsValue({
          parentId: detailData?.parentId ?? 0,
          sort: 0,
          icon: '',
          redirect: '/',
          isHideInMenu: false,
          isHideChildrenInMenu: false,
          isEnabled: true,
        });
      }
    }
  }, [open, detailData, form]);

  // 键名由路径自动生成
  const onPathChange = (path: string) => {
    form.setFieldValue('key', path2UpperCamelCase(path));
  };

  const handleOk = async () => {
    const values = await form.validateFields();
    const data: ReqAdminMenuAdd & { id?: number } = { ...values, parentId: values.parentId ?? 0 };
    setConfirmLoading(true);
    try {
      const res = detailData?.id ? await editAdminMenu({ ...data, id: detailData.id }) : await addAdminMenu(data);
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal open={open} title={detailData?.id ? '编辑菜单' : '新建菜单'} width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Form form={form} labelCol={{ span: 8 }} wrapperCol={{ span: 14 }}>
        <Form.Item label="父级菜单" name="parentId">
          <TreeSelect treeData={treeSelectData} treeDefaultExpandAll placeholder="请选择父级菜单" allowClear />
        </Form.Item>
        <Form.Item label="菜单名称" name="name" rules={[{ required: true, message: '请输入菜单名称' }, { max: 50, message: '菜单名称长度不能超过50个字符' }]}>
          <Input placeholder="请输入菜单名称" allowClear />
        </Form.Item>
        <Form.Item label="菜单路径" name="path" rules={[{ required: true, message: '请输入菜单路径' }, { pattern: AdminRouterPath, message: '路径格式不正确' }]}>
          <Input placeholder="请输入菜单路径" allowClear onChange={(e) => onPathChange(e.target.value)} />
        </Form.Item>
        <Form.Item label="键名" name="key" rules={[{ required: true, pattern: AdminMenuKey, message: '请按照驼峰法命名' }]}>
          <Input disabled placeholder="根据菜单路径自动生成" />
        </Form.Item>
        <Form.Item label="排序" name="sort">
          <InputNumber min={0} style={{ width: '100%' }} />
        </Form.Item>
        <Form.Item label="图标" name="icon">
          <IconSelect placeholder="请选择图标" />
        </Form.Item>
        <Form.Item label="重定向地址" name="redirect">
          <Input placeholder="请输入重定向地址" allowClear />
        </Form.Item>
        <Form.Item label="描述" name="describe">
          <Input placeholder="请输入菜单描述" allowClear />
        </Form.Item>
        <Form.Item label="是否隐藏菜单" name="isHideInMenu" valuePropName="checked">
          <Switch checkedChildren="隐藏" unCheckedChildren="显示" />
        </Form.Item>
        <Form.Item label="是否隐藏子菜单" name="isHideChildrenInMenu" valuePropName="checked">
          <Switch checkedChildren="隐藏" unCheckedChildren="显示" />
        </Form.Item>
        <Form.Item label="是否启用" name="isEnabled" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="禁用" />
        </Form.Item>
      </Form>
    </Modal>
  );
}
