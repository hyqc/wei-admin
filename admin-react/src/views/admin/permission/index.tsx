import { useEffect, useState } from 'react';
import { Button, Form, Input, Select, Space, Table, Tag, TreeSelect } from 'antd';
import { EyeOutlined, ReloadOutlined, SearchOutlined, SettingOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import type { ColumnsType } from 'antd/es/table';
import PageContainer from '@/components/PageContainer';
import Authorization from '@/components/Authorization';
import RowEnabledButton from '@/components/RowEnabledButton';
import DetailPermissionDrawer from './components/DetailPermissionDrawer';
import { PERMISSION_TYPE_OPTIONS } from './components/common';
import { getAdminPermissionList } from '@/api/admin/permission';
import { getAdminMenuTree } from '@/api/admin/menu';
import { getAdminApiAll } from '@/api/admin/api';
import type { PermissionListItem } from '@/types/admin_permission';
import type { MenuTreeItem } from '@/types/admin_menu';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

/**
 * 权限管理（只读检索/审计）
 * 权限点及其接口绑定统一在「菜单管理 → 权限配置」中维护，此页只提供查询、详情与按接口反查
 */
export default function AdminPermission() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<PermissionListItem[]>([]);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [menuTreeOptions, setMenuTreeOptions] = useState<unknown[]>([]);
  const [apiOptions, setApiOptions] = useState<{ label: string; value: number }[]>([]);
  const [detailOpen, setDetailOpen] = useState(false);
  const [detailData, setDetailData] = useState<PermissionListItem>();
  const [searchForm] = Form.useForm<{ menuId?: number; key?: string; name?: string; type?: string; apiId?: number }>();
  const navigate = useNavigate();

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

  /** 接口下拉：用于反查“绑定了某接口的权限点” */
  function loadApiOptions() {
    getAdminApiAll().then((res) => {
      setApiOptions(
        (res.data || [])
          .slice()
          .sort((a, b) => (a.path || '').localeCompare(b.path || ''))
          .map((item) => ({ label: `${item.name} (${item.path})`, value: item.id as number })),
      );
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
      apiId: search.apiId,
    })
      .then((res) => {
        setRows(res.data.list || []);
        setPageInfo((prev) => ({ ...prev, total: res.data.pageInfo?.total || (res.data.list || []).length }));
      })
      .finally(() => setLoading(false));
  }

  useEffect(() => {
    loadMenuTree();
    loadApiOptions();
    getRows();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function onSearch() {
    const next = { ...pageInfo, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  }

  function onReset() {
    searchForm.setFieldsValue({ menuId: undefined, key: undefined, name: undefined, type: undefined, apiId: undefined });
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
      render: (_v, record) => <RowEnabledButton isEnabled={record.isEnabled} />,
    },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
      fixed: 'right',
      width: 180,
      render: (_v, record) => (
        <Space>
          <Authorization permission="AdminPermissionView">
            <Button
              type="link"
              size="small"
              icon={<EyeOutlined />}
              onClick={() => {
                setDetailData(record);
                setDetailOpen(true);
              }}
            >
              详情
            </Button>
          </Authorization>
          <Authorization permission="AdminMenuEdit">
            <Button type="link" size="small" icon={<SettingOutlined />} onClick={() => navigate('/admin/menu')}>
              去配置
            </Button>
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
            <Select placeholder="全部" allowClear style={{ width: 120 }} options={PERMISSION_TYPE_OPTIONS} />
          </Form.Item>
          <Form.Item label="接口" name="apiId">
            <Select
              placeholder="全部"
              allowClear
              showSearch
              optionFilterProp="label"
              options={apiOptions}
              style={{ width: 240 }}
            />
          </Form.Item>
          <Form.Item>
            <Space>
              <Button type="primary" htmlType="submit" icon={<SearchOutlined />}>
                查询
              </Button>
              <Button icon={<ReloadOutlined />} onClick={onReset}>
                重置
              </Button>
            </Space>
          </Form.Item>
        </Form>
      }
      extra={
        <span style={{ color: 'rgba(0,0,0,0.45)' }}>
          权限点及其接口绑定由各菜单的“权限配置”统一维护，此页仅用于查询与审计
        </span>
      }
    >
      <Table columns={columns} dataSource={rows} loading={loading} pagination={false} rowKey="id" scroll={{ x: 1250 }} />
      <DetailPermissionDrawer open={detailOpen} detailData={detailData} onClose={() => setDetailOpen(false)} />
    </PageContainer>
  );
}
