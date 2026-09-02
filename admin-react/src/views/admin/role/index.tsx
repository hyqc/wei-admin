import { useEffect, useState } from 'react';
import { Button, Form, Input, Popconfirm, Select, Space, Table, Tag, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import {
  DeleteOutlined,
  EditOutlined,
  EyeOutlined,
  PlusOutlined,
  ReloadOutlined,
  SafetyCertificateOutlined,
  SearchOutlined,
} from '@ant-design/icons';
import PageContainer from '@/components/PageContainer';
import Authorization from '@/components/Authorization';
import RowEnabledButton from '@/components/RowEnabledButton';
import AddRoleModal from './components/AddRoleModal';
import EditRoleModal from './components/EditRoleModal';
import DetailRoleDrawer from './components/DetailRoleDrawer';
import BindPermissionsModal from './components/BindPermissionsModal';
import { deleteAdminRole, enableAdminRole, getAdminRoleList } from '@/api/admin/role';
import type { RoleItem } from '@/types/admin_role';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

/** 角色管理 */
export default function AdminRole() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<RoleItem[]>([]);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [searchForm] = Form.useForm<{ name?: string; enabled?: number }>();
  const [addOpen, setAddOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [bindOpen, setBindOpen] = useState(false);
  const [editData, setEditData] = useState<RoleItem>();
  const [detailData, setDetailData] = useState<RoleItem>();

  const getRows = (page = pageInfo, search = searchForm.getFieldsValue()) => {
    setLoading(true);
    getAdminRoleList({
      pageNum: page.pageNum,
      pageSize: page.pageSize,
      name: search.name,
      enabled: search.enabled ?? 0,
    })
      .then((res) => {
        setRows(res.data.list || []);
        setPageInfo((prev) => ({ ...prev, total: res.data.pageInfo?.total || 0 }));
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    getRows();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const onSearch = () => {
    const next = { ...pageInfo, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  };

  const onReset = () => {
    searchForm.setFieldsValue({ name: undefined, enabled: 0 });
    onSearch();
  };

  const onPageChange = (pageNum: number) => {
    const next = { ...pageInfo, pageNum };
    setPageInfo(next);
    getRows(next);
  };

  const onPageSizeChange = (pageSize: number) => {
    const next = { ...pageInfo, pageSize, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  };

  const updateEnabled = (record: RoleItem) => {
    enableAdminRole({ id: record.id, enabled: !record.isEnabled }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  };

  const onDelete = (record: RoleItem) => {
    deleteAdminRole({ id: record.id }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  };

  const columns: ColumnsType<RoleItem> = [
    { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
    {
      title: '角色名称',
      dataIndex: 'name',
      key: 'name',
      width: 160,
      render: (_v, record) => (
        <Space>
          <span>{record.name}</span>
          {record.isSuperAdmin && <Tag color="gold">超级管理员</Tag>}
        </Space>
      ),
    },
    { title: '描述', dataIndex: 'describe', key: 'describe', width: 180 },
    { title: '创建人', dataIndex: 'createAdminName', key: 'createAdminName', width: 120 },
    {
      title: '状态',
      dataIndex: 'isEnabled',
      key: 'isEnabled',
      width: 100,
      render: (_v, record) => (
        <Popconfirm
          title={`确定要${record.isEnabled ? '禁用' : '启用'}该角色吗？`}
          okText="确定"
          cancelText="取消"
          onConfirm={() => updateEnabled(record)}
        >
          <RowEnabledButton isEnabled={record.isEnabled} />
        </Popconfirm>
      ),
    },
    { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
    { title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 180 },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
      fixed: 'right',
      width: 320,
      render: (_v, record) => (
        <Space>
          <Authorization permission="AdminRoleView">
            <Button type="link" size="small" icon={<EyeOutlined />} onClick={() => { setDetailData(record); setDetailOpen(true); }}>
              详情
            </Button>
          </Authorization>
          <Authorization permission="AdminRoleEdit">
            {/* 超管角色允许编辑，但仅可修改描述 */}
            <Button type="link" size="small" icon={<EditOutlined />} onClick={() => { setEditData(record); setEditOpen(true); }}>
              编辑
            </Button>
          </Authorization>
          <Authorization permission="AdminRoleEdit">
            {/* 超级管理员角色不允许绑定权限 */}
            {record.id !== 1 && (
              <Button type="link" size="small" icon={<SafetyCertificateOutlined />} onClick={() => { setDetailData(record); setBindOpen(true); }}>
                绑定权限
              </Button>
            )}
          </Authorization>
          <Authorization permission="AdminRoleDelete">
            {!record.isEnabled && (
              <Popconfirm title="确定要删除该角色吗？" okText="确定" cancelText="取消" onConfirm={() => onDelete(record)}>
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
          <Form.Item label="角色名称" name="name">
            <Input placeholder="请输入角色名称" allowClear />
          </Form.Item>
          <Form.Item label="状态" name="enabled" initialValue={0}>
            <Select style={{ width: 120 }} options={[{ value: 0, label: '全部' }, { value: 1, label: '启用' }, { value: 2, label: '禁用' }]} />
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
        <Button type="primary" icon={<PlusOutlined />} onClick={() => setAddOpen(true)}>新建角色</Button>
      }
    >
      <Table columns={columns} dataSource={rows} loading={loading} pagination={false} rowKey="id" scroll={{ x: 1320 }} />
      <AddRoleModal open={addOpen} onClose={() => setAddOpen(false)} onNotice={onModalNotice} />
      <EditRoleModal open={editOpen} detailData={editData} onClose={() => setEditOpen(false)} onNotice={onModalNotice} />
      <DetailRoleDrawer open={detailOpen} detailData={detailData} onClose={() => setDetailOpen(false)} />
      <BindPermissionsModal open={bindOpen} detailData={detailData} onClose={() => setBindOpen(false)} onNotice={onModalNotice} />
    </PageContainer>
  );

  function onModalNotice() {
    getRows();
  }
}
