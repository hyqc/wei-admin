import { useEffect, useState } from 'react';
import { Button, Form, Input, Popconfirm, Space, Table, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import {
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
import AddApiModal from './components/AddApiModal';
import EditApiModal from './components/EditApiModal';
import DetailApiDrawer from './components/DetailApiDrawer';
import { getAdminApiList, deleteAdminApi, enableAdminApi } from '@/api/admin/api';
import type { AdminApiItem } from '@/types/common';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

/** 接口管理 */
export default function AdminApi() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<AdminApiItem[]>([]);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [searchForm] = Form.useForm<{ key?: string; name?: string; path?: string }>();
  const [addOpen, setAddOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [editData, setEditData] = useState<AdminApiItem>();
  const [detailData, setDetailData] = useState<AdminApiItem>();

  function getRows(page = pageInfo, search = searchForm.getFieldsValue()) {
    setLoading(true);
    getAdminApiList({
      pageNum: page.pageNum,
      pageSize: page.pageSize,
      key: search.key,
      name: search.name,
      path: search.path,
    })
      .then((res) => {
        setRows(res.data.list || []);
        setPageInfo((prev) => ({ ...prev, total: res.data.pageInfo?.total || 0 }));
      })
      .finally(() => setLoading(false));
  }

  useEffect(() => {
    getRows();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function onSearch() {
    const next = { ...pageInfo, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  }

  function onReset() {
    searchForm.setFieldsValue({ key: undefined, name: undefined, path: undefined });
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

  function updateEnabled(record: AdminApiItem) {
    enableAdminApi({ id: record.id, enabled: !record.isEnabled }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function onDelete(record: AdminApiItem) {
    deleteAdminApi({ id: record.id }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function onModalNotice() {
    getRows();
  }

  const columns: ColumnsType<AdminApiItem> = [
    { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
    { title: '接口名称', dataIndex: 'name', key: 'name', width: 160 },
    { title: '唯一键', dataIndex: 'key', key: 'key', width: 200 },
    { title: '接口路径', dataIndex: 'path', key: 'path', width: 240 },
    {
      title: '状态',
      dataIndex: 'isEnabled',
      key: 'isEnabled',
      width: 100,
      render: (_v, record) => (
        <Popconfirm
          title={`确定要${record.isEnabled ? '禁用' : '启用'}该接口吗？`}
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
      width: 260,
      render: (_v, record) => (
        <Space>
          <Authorization permission="AdminApiView">
            <Button type="link" size="small" icon={<EyeOutlined />} onClick={() => { setDetailData(record); setDetailOpen(true); }}>
              详情
            </Button>
          </Authorization>
          <Authorization permission="AdminApiEdit">
            <Button type="link" size="small" icon={<EditOutlined />} onClick={() => { setEditData(record); setEditOpen(true); }}>
              编辑
            </Button>
          </Authorization>
          <Authorization permission="AdminApiDelete">
            {!record.isEnabled && (
              <Popconfirm title="确定要删除该接口吗？" okText="确定" cancelText="取消" onConfirm={() => onDelete(record)}>
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
          <Form.Item label="唯一键" name="key">
            <Input placeholder="请输入唯一键" allowClear />
          </Form.Item>
          <Form.Item label="名称" name="name">
            <Input placeholder="请输入接口名称" allowClear />
          </Form.Item>
          <Form.Item label="路径" name="path">
            <Input placeholder="请输入接口路径" allowClear />
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
        <Button type="primary" icon={<PlusOutlined />} onClick={() => setAddOpen(true)}>新建接口</Button>
      }
    >
      <Table columns={columns} dataSource={rows} loading={loading} pagination={false} rowKey="id" scroll={{ x: 1060 }} />
      <AddApiModal open={addOpen} onClose={() => setAddOpen(false)} onNotice={onModalNotice} />
      <EditApiModal open={editOpen} detailData={editData} onClose={() => setEditOpen(false)} onNotice={onModalNotice} />
      <DetailApiDrawer open={detailOpen} detailData={detailData} onClose={() => setDetailOpen(false)} />
    </PageContainer>
  );
}
