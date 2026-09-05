import { useEffect, useState } from 'react';
import { Button, Popconfirm, Space, Switch, Table, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import {
  DeleteOutlined,
  EditOutlined,
  EyeOutlined,
  PlusOutlined,
  SafetyCertificateOutlined,
} from '@ant-design/icons';
import PageContainer from '@/components/PageContainer';
import Authorization from '@/components/Authorization';
import SaveMenuModal from './components/SaveMenuModal';
import DetailMenuDrawer from './components/DetailMenuDrawer';
import PermissionsSaveModal from './components/PermissionsSaveModal';
import { deleteAdminMenu, enableAdminMenu, getAdminMenuTree, showAdminMenu } from '@/api/admin/menu';
import { getAntIcon } from '@/utils/icon';
import type { MenuTreeItem } from '@/types/admin_menu';

const menuIcon = (name?: string) => getAntIcon(name);

function formatTime(time?: number) {
  if (!time) return '';
  const d = new Date(time * 1000);
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
}

/** 菜单管理 */
export default function AdminMenu() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<MenuTreeItem[]>([]);
  const [expandedRowKeys, setExpandedRowKeys] = useState<number[]>([]);
  const [saveOpen, setSaveOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [permissionsOpen, setPermissionsOpen] = useState(false);
  const [saveData, setSaveData] = useState<MenuTreeItem>();
  const [detailData, setDetailData] = useState<MenuTreeItem>();

  function getRows() {
    setLoading(true);
    getAdminMenuTree()
      .then((res) => {
        const list = res.data.list || [];
        setRows(list);
        // 加载后默认全部展开（递归收集含子节点的菜单 id）
        const keys: number[] = [];
        const collect = (items: MenuTreeItem[]) => {
          for (const item of items) {
            if (item.children?.length) {
              keys.push(item.id as number);
              collect(item.children);
            }
          }
        };
        collect(list);
        setExpandedRowKeys(keys);
      })
      .finally(() => setLoading(false));
  }

  useEffect(() => {
    getRows();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function updateEnabled(record: MenuTreeItem) {
    enableAdminMenu({ menuId: record.id, enabled: !record.enabled }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function updateShow(record: MenuTreeItem) {
    showAdminMenu({ menuId: record.id, show: !!record.hideInMenu }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function onDelete(record: MenuTreeItem) {
    deleteAdminMenu({ menuId: record.id }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function openAddModal(parent?: MenuTreeItem) {
    // 新增子菜单时只传父级ID：若直接传父菜单对象，会被弹窗识别为“编辑父菜单”
    setSaveData(parent ? ({ parentId: parent.id } as MenuTreeItem) : undefined);
    setSaveOpen(true);
  }

  function openEditModal(record: MenuTreeItem) {
    setSaveData({ ...record });
    setSaveOpen(true);
  }

  const columns: ColumnsType<MenuTreeItem> = [
    { title: '菜单名称', dataIndex: 'name', key: 'name', width: 220 },
    {
      title: '图标',
      dataIndex: 'icon',
      key: 'icon',
      width: 160,
      render: (_v, record) => {
        const Icon = menuIcon(record.icon);
        return Icon ? (
          <span className="icon-cell">
            <Icon />
            <span>{record.icon}</span>
          </span>
        ) : (
          '-'
        );
      },
    },
    { title: '键名', dataIndex: 'key', key: 'key', width: 160 },
    { title: '菜单路径', dataIndex: 'path', key: 'path', width: 180 },
    { title: '排序', dataIndex: 'sort', key: 'sort', width: 80 },
    {
      title: '状态',
      dataIndex: 'enabled',
      key: 'isEnabled',
      width: 100,
      render: (_v, record) => (
        <Authorization permission="AdminMenuEdit">
          <Popconfirm
            title={`确定要${record.enabled ? '禁用' : '启用'}该菜单吗？`}
            okText="确定"
            cancelText="取消"
            onConfirm={() => updateEnabled(record)}
          >
            <Switch checked={record.enabled} checkedChildren="启用" unCheckedChildren="禁用" />
          </Popconfirm>
        </Authorization>
      ),
    },
    {
      title: '是否显示',
      dataIndex: 'hideInMenu',
      key: 'hideInMenu',
      width: 100,
      render: (_v, record) => (
        <Authorization permission="AdminMenuEdit">
          <Switch
            checked={!record.hideInMenu}
            checkedChildren="显示"
            unCheckedChildren="隐藏"
            onChange={() => updateShow(record)}
          />
        </Authorization>
      ),
    },
    { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 180, render: (v: number) => formatTime(v) },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
      fixed: 'right',
      width: 380,
      render: (_v, record) => (
        <Space>
          <Authorization permission="AdminMenuView">
            <Button type="link" size="small" icon={<EyeOutlined />} onClick={() => { setDetailData(record); setDetailOpen(true); }}>
              详情
            </Button>
          </Authorization>
          <Authorization permission="AdminMenuAdd">
            <Button type="link" size="small" icon={<PlusOutlined />} onClick={() => openAddModal(record)}>
              新增子菜单
            </Button>
          </Authorization>
          <Authorization permission="AdminMenuEdit">
            <Button type="link" size="small" icon={<EditOutlined />} onClick={() => openEditModal(record)}>
              编辑
            </Button>
          </Authorization>
          <Authorization permission="AdminMenuEdit">
            {/* 目录型菜单（含子菜单）不是页面，没有可授权的操作，不提供权限配置 */}
            {!record.children?.length && (
              <Button type="link" size="small" icon={<SafetyCertificateOutlined />} onClick={() => { setDetailData(record); setPermissionsOpen(true); }}>
                权限配置
              </Button>
            )}
          </Authorization>
          <Authorization permission="AdminMenuDelete">
            {!record.enabled && (
              <Popconfirm title="确定要删除该菜单吗？" okText="确定" cancelText="取消" onConfirm={() => onDelete(record)}>
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
      extra={
        <Authorization permission="AdminMenuAdd">
          <Button type="primary" icon={<PlusOutlined />} onClick={() => openAddModal()}>
            新建菜单
          </Button>
        </Authorization>
      }
    >
      <Table
        columns={columns}
        dataSource={rows}
        loading={loading}
        pagination={false}
        rowKey="id"
        scroll={{ x: 1320 }}
        expandable={{ expandedRowKeys, onExpandedRowsChange: (keys) => setExpandedRowKeys(keys as number[]) }}
      />
      <SaveMenuModal open={saveOpen} tree={rows} detailData={saveData} onClose={() => setSaveOpen(false)} onNotice={getRows} />
      <DetailMenuDrawer open={detailOpen} detailData={detailData} onClose={() => setDetailOpen(false)} />
      <PermissionsSaveModal open={permissionsOpen} detailData={detailData} onClose={() => setPermissionsOpen(false)} />
    </PageContainer>
  );
}
