import { useEffect, useState } from 'react';
import { Button, Form, Input, Popconfirm, Select, Space, Table, Tag, TreeSelect, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import {
  ApiOutlined,
  DeleteOutlined,
  EditOutlined,
  EyeOutlined,
  PlusOutlined,
  ReloadOutlined,
  SearchOutlined,
} from '@ant-design/icons';
import PageContainer from '@/components/PageContainer';
import Authorization from '@/components/Authorization';
import RowEnabledButton from '@/components/RowEnabledButton';
import AddPermissionModal from './components/AddPermissionModal';
import EditPermissionModal from './components/EditPermissionModal';
import DetailPermissionDrawer from './components/DetailPermissionDrawer';
import BindApisModal from './components/BindApisModal';
import { DEFAULT_PERMISSION_TYPES } from './components/common';
import { getAdminPermissionList, deleteAdminPermission, enableAdminPermission } from '@/api/admin/permission';
import { getAdminMenuTree } from '@/api/admin/menu';
import type { PermissionListItem } from '@/types/admin_permission';
import type { MenuTreeItem } from '@/types/admin_menu';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

/** 权限管理 */
export default function AdminPermission() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<PermissionListItem[]>([]);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [menuTreeOptions, setMenuTreeOptions] = useState<unknown[]>([]);  const [searchForm] = Form.useForm<{ menuId?: number; key?: string; name?: string; type?: string }>();
  const [addOpen, setAddOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [bindApisOpen, setBindApisOpen] = useState(false);
  const [editData, setEditData] = useState<PermissionListItem>();
  const [detailData, setDetailData] = useState<PermissionListItem>();

  function loadMenuTree() {
    getAdminMenuTree().then((res) => {
      const build = (list: MenuTreeItem[]): unknown[] =>
        list.map((item) => ({
          key: item.id,
          value: item.id,
          title: item.name,
          children: item.children?.length ? build(item.children) : undefined,
        }));
      setMenuTreeOptions(build(res.data.list || []));
    });
  }

  function getRows(page = pageInfo, search = searchForm.getFieldsValue()) {
    setLoading(true);
    getAdminPermissionList({
      pageNum: page.pageNum,
      pageSize: page.pageSize,
      menuId: search.menuId,
      key: search.key,
      name: search.name,
      type: search.type,
    })
      .then((res) => {
        setRows(res.data.list || []);
        setPageInfo((prev) => ({ ...prev, total: res.data.pageInfo?.total || (res.data.list || []).length }));
      })
      .finally(() => setLoading(false));
  }

  useEffect(() => {
    loadMenuTree();
    getRows();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function onSearch() {
    const next = { ...pageInfo, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  }

  function onReset() {
    searchForm.setFieldsValue({ menuId: undefined, key: undefined, name: undefined, type: undefined });
    onSearch();
  }

  function onPageChange(pageNum: number) {
    const next = { ...pageInfo, pageNum };
    setPageInfo(next);
    getRows(next);
  }

  function onPageSizeChange(pageSize: number) {
    const next = { ...pageInfo, pageSize, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  }

  function updateEnabled(record: PermissionListItem) {
    enableAdminPermission({ id: record.id, enabled: !record.isEnabled }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function onDelete(record: PermissionListItem) {
    deleteAdminPermission({ id: record.id }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function onModalNotice() {
    getRows();
  }

  const columns: ColumnsType<PermissionListItem> = [
    { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
    { title: '权限名称', dataIndex: 'name', key: 'name', width: 140 },
    { title: '菜单ID', dataIndex: 'menuId', key: 'menuId', width: 80 },
    { title: '所属菜单', dataIndex: 'menuName', key: 'menuName', width: 140, render: (v: string) => v || '-' },
    { title: '唯一键', dataIndex: 'key', key: 'key', width: 180 },
    {
      title: '类型',
      dataIndex: 'type',
      key: 'type',
      width: 100,
      render: (_v, record) => <Tag color="blue">{record.typeText}</Tag>,
    },
    {
      title: '接口数',
      dataIndex: 'apis',
      key: 'apis',
      width: 100,
      render: (_v, record) =>
        record.apis?.length ? (
          <Tag color="green">{record.apis.length} 个</Tag>
        ) : (
          <span style={{ color: 'rgba(0,0,0,0.45)' }}>未绑定</span>
        ),
    },
    {
      title: '状态',
      dataIndex: 'enabled',
      key: 'enabled',
      width: 100,
      render: (_v, record) => (
        <Popconfirm
          title={`确定要${record.isEnabled ? '禁用' : '启用'}该权限吗？`}
          okText="确定"
          cancelText="取消"
          onConfirm={() => updateEnabled(record)}
        >
          <RowEnabledButton isEnabled={record.isEnabled} />
        </Popconfirm>
      ),
    },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
      fixed: 'right',
      width: 320,
      render: (_v, record) => (
        <Space>
          <Authorization permission="AdminPermissionView">
            <Button type="link" size="small" icon={<EyeOutlined />} onClick={() => { setDetailData(record); setDetailOpen(true); }}>
              详情
            </Button>
          </Authorization>
          <Authorization permission="AdminPermissionEdit">
            <Button type="link" size="small" icon={<EditOutlined />} onClick={() => { setEditData(record); setEditOpen(true); }}>
              编辑
            </Button>
          </Authorization>
          <Authorization permission="AdminPermissionEdit">
            <Button type="link" size="small" icon={<ApiOutlined />} onClick={() => { setDetailData(record); setBindApisOpen(true); }}>
              绑定接口
            </Button>
          </Authorization>
          <Authorization permission="AdminPermissionDelete">
            {!record.isEnabled && (
              <Popconfirm title="确定要删除该权限吗？" okText="确定" cancelText="取消" onConfirm={() => onDelete(record)}>
                <Button type="link" size="small" danger icon={<DeleteOutlined />}>
                  删除
                </Button>
              </Popconfirm>
            )}
          </Authorization>
        </Space>
      ),
    },
  ];

  return (
    <PageContainer
      pageInfo={pageInfo}
      onPageChange={onPageChange}
      onPageSizeChange={onPageSizeChange}
      searchArea={
        <Form layout="inline" form={searchForm} onFinish={onSearch}>
          <Form.Item label="菜单" name="menuId">
            <TreeSelect treeData={menuTreeOptions as never} treeDefaultExpandAll placeholder="请选择菜单" allowClear style={{ width: 180 }} />
          </Form.Item>
          <Form.Item label="唯一键" name="key">
            <Input placeholder="请输入唯一键" allowClear />
          </Form.Item>
          <Form.Item label="名称" name="name">
            <Input placeholder="请输入权限名称" allowClear />
          </Form.Item>
          <Form.Item label="类型" name="type">
            <Select
              placeholder="全部"
              allowClear
              style={{ width: 120 }}
              options={DEFAULT_PERMISSION_TYPES.map((item) => ({ value: item.key, label: item.name }))}
            />
          </Form.Item>
          <Form.Item>
            <Space>
              <Button type="primary" htmlType="submit" icon={<SearchOutlined />}>查询</Button>
              <Button icon={<ReloadOutlined />} onClick={onReset}>重置</Button>
            </Space>
          </Form.Item>
        </Form>
      }
      extra={
        <Button type="primary" icon={<PlusOutlined />} onClick={() => setAddOpen(true)}>新建权限</Button>
      }
    >
      <Table columns={columns} dataSource={rows} loading={loading} pagination={false} rowKey="id" scroll={{ x: 1250 }} />
      <AddPermissionModal open={addOpen} onClose={() => setAddOpen(false)} onNotice={onModalNotice} />
      <EditPermissionModal open={editOpen} detailData={editData} onClose={() => setEditOpen(false)} onNotice={onModalNotice} />
      <DetailPermissionDrawer open={detailOpen} detailData={detailData} onClose={() => setDetailOpen(false)} onNotice={onModalNotice} />
      <BindApisModal open={bindApisOpen} detailData={detailData} onClose={() => setBindApisOpen(false)} onNotice={onModalNotice} />
    </PageContainer>
  );
}
