import { useEffect, useState } from 'react';
import { Button, Form, Input, Popconfirm, Select, Space, Table, Tag, Tooltip, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import {
  DeleteOutlined,
  EditOutlined,
  EyeOutlined,
  KeyOutlined,
  PlusOutlined,
  ReloadOutlined,
  SearchOutlined,
  UserSwitchOutlined,
} from '@ant-design/icons';
import PageContainer from '@/components/PageContainer';
import Authorization from '@/components/Authorization';
import RowEnabledButton from '@/components/RowEnabledButton';
import AddUserModal from './components/AddUserModal';
import EditUserModal from './components/EditUserModal';
import DetailUserDrawer from './components/DetailUserDrawer';
import BindRolesModal from './components/BindRolesModal';
import ResetPasswordModal from './components/ResetPasswordModal';
import { deleteAdminUser, enableAdminUser, getAdminUserList } from '@/api/admin/user';
import type { AdminUserListItem } from '@/types/common';
import { AdminId, DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

/** 拆分登录IP：last 取倒数第 2 个（上次登录），current 取最后一个（当前登录） */
function getLoginIp(ipStr?: string, type: 'last' | 'current' = 'last'): string {
  const s = ipStr || '';
  let list: string[] = [];
  try {
    const parsed = JSON.parse(s);
    if (Array.isArray(parsed)) list = parsed.map(String);
  } catch {
    list = [];
  }
  if (list.length === 0) {
    list = s
      .split(',')
      .map((x) => x.trim())
      .filter(Boolean);
  }
  const ip = type === 'current' ? list[list.length - 1] : list[list.length - 2];
  return ip || '-';
}

/** 昵称超 10 字符截断显示 */
function displayNickname(nickname?: string): string {
  if (!nickname) return '';
  return nickname.length > 10 ? `${nickname.slice(0, 10)}...` : nickname;
}

/** 账号管理 */
export default function AdminUser() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<AdminUserListItem[]>([]);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [searchForm] = Form.useForm<{ username?: string; nickname?: string; email?: string; enabled?: number }>();

  const [addOpen, setAddOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [detailOpen, setDetailOpen] = useState(false);
  const [bindRolesOpen, setBindRolesOpen] = useState(false);
  const [resetPwdOpen, setResetPwdOpen] = useState(false);
  const [editData, setEditData] = useState<AdminUserListItem>();
  const [detailData, setDetailData] = useState<AdminUserListItem>();

  function getRows(page = pageInfo, search = searchForm.getFieldsValue()) {
    setLoading(true);
    getAdminUserList({
      pageNum: page.pageNum,
      pageSize: page.pageSize,
      username: search.username,
      nickname: search.nickname,
      email: search.email,
      enabled: search.enabled ?? 0,
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
    searchForm.setFieldsValue({ username: undefined, nickname: undefined, email: undefined, enabled: 0 });
    const next = { ...pageInfo, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
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

  function updateEnabled(record: AdminUserListItem) {
    enableAdminUser({ adminId: record.adminId, enabled: !record.isEnabled }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function onDelete(record: AdminUserListItem) {
    deleteAdminUser({ adminId: record.adminId }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  }

  function openEditModal(record: AdminUserListItem) {
    setEditData(record);
    setEditOpen(true);
  }

  function openDetailModal(record: AdminUserListItem) {
    setDetailData(record);
    setDetailOpen(true);
  }

  function openBindRolesModal(record: AdminUserListItem) {
    setDetailData(record);
    setBindRolesOpen(true);
  }

  function openResetPwdModal(record: AdminUserListItem) {
    setDetailData(record);
    setResetPwdOpen(true);
  }

  function onModalNotice() {
    getRows();
  }

  const columns: ColumnsType<AdminUserListItem> = [
    { title: 'ID', dataIndex: 'adminId', key: 'adminId', fixed: 'left', width: 80 },
    { title: '账号', dataIndex: 'username', key: 'username', fixed: 'left', width: 120 },
    {
      title: '昵称',
      dataIndex: 'nickname',
      key: 'nickname',
      width: 160,
      render: (v: string) =>
        v && v.length > 10 ? (
          <Tooltip title={v}>
            <span
              style={{
                display: 'inline-block',
                maxWidth: '100%',
                overflow: 'hidden',
                textOverflow: 'ellipsis',
                whiteSpace: 'nowrap',
                verticalAlign: 'middle',
              }}
            >
              {displayNickname(v)}
            </span>
          </Tooltip>
        ) : (
          displayNickname(v)
        ),
    },
    { title: '邮箱', dataIndex: 'email', key: 'email', width: 180 },
    {
      title: '角色',
      dataIndex: 'roles',
      key: 'roles',
      width: 180,
      // 超管无需角色，自动拥有系统全部权限
      render: (_v, record) =>
        record.isSuperAdmin ? (
          <Tag color="gold">超级管理员</Tag>
        ) : (
          (record.roles ?? []).map((role) => (
            <Tag key={role.roleId} color="blue">
              {role.roleName}
            </Tag>
          ))
        ),
    },
    {
      title: '状态',
      dataIndex: 'isEnabled',
      key: 'isEnabled',
      width: 100,
      render: (_v, record) => (
        <Popconfirm
          title={`确定要${record.isEnabled ? '禁用' : '启用'}该账号吗？`}
          okText="确定"
          cancelText="取消"
          onConfirm={() => updateEnabled(record)}
        >
          <RowEnabledButton isEnabled={record.isEnabled} isEnabledButtonDisabled={record.adminId === AdminId} />
        </Popconfirm>
      ),
    },
    { title: '登录次数', dataIndex: 'loginTotal', key: 'loginTotal', width: 100 },
    {
      title: '上次登录IP',
      dataIndex: 'lastLoginIp',
      key: 'lastLoginIp',
      width: 140,
      render: (v: string) => v || '-',
    },
    {
      title: '当前登录IP',
      dataIndex: 'currentLoginIp',
      key: 'currentLoginIp',
      width: 140,
      render: (v: string, record) => v || getLoginIp(record.lastLoginIp, 'current'),
    },
    { title: '上次登录时间', dataIndex: 'lastLoginTime', key: 'lastLoginTime', width: 180 },
    { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
      fixed: 'right',
      width: 420,
      render: (_v, record) => (
        <Space>
          <Authorization permission="AdminUserView">
            <Button type="link" size="small" icon={<EyeOutlined />} onClick={() => openDetailModal(record)}>
              详情
            </Button>
          </Authorization>
          <Authorization permission="AdminUserEdit">
            <Button type="link" size="small" icon={<EditOutlined />} onClick={() => openEditModal(record)}>
              编辑
            </Button>
          </Authorization>
          <Authorization permission="AdminUserEdit">
            <Button type="link" size="small" icon={<KeyOutlined />} onClick={() => openResetPwdModal(record)}>
              重置密码
            </Button>
          </Authorization>
          <Authorization permission="AdminUserEdit">
            {/* 超管账号自动拥有全部权限，不允许绑定角色 */}
            {!record.isSuperAdmin && (
              <Button type="link" size="small" icon={<UserSwitchOutlined />} onClick={() => openBindRolesModal(record)}>
                绑定角色
              </Button>
            )}
          </Authorization>
          <Authorization permission="AdminUserDelete">
            {/* 超管不可删除，仅被禁用的账号可删除 */}
            {record.username !== 'admin' && !record.isEnabled && (
              <Popconfirm title="确定要删除该账号吗？" okText="确定" cancelText="取消" onConfirm={() => onDelete(record)}>
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
          <Form.Item label="账号" name="username">
            <Input placeholder="请输入账号" allowClear />
          </Form.Item>
          <Form.Item label="昵称" name="nickname">
            <Input placeholder="请输入昵称" allowClear />
          </Form.Item>
          <Form.Item label="邮箱" name="email">
            <Input placeholder="请输入邮箱" allowClear />
          </Form.Item>
          <Form.Item label="状态" name="enabled" initialValue={0}>
            <Select style={{ width: 120 }} placeholder="全部" options={[{ value: 0, label: '全部' }, { value: 1, label: '启用' }, { value: 2, label: '禁用' }]} />
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
        <Button type="primary" icon={<PlusOutlined />} onClick={() => setAddOpen(true)}>
          新建账号
        </Button>
      }
    >
      <Table columns={columns} dataSource={rows} loading={loading} pagination={false} rowKey="adminId" scroll={{ x: 1940 }} />

      <AddUserModal open={addOpen} onClose={() => setAddOpen(false)} onNotice={onModalNotice} />
      <EditUserModal open={editOpen} detailData={editData} onClose={() => setEditOpen(false)} onNotice={onModalNotice} />
      <DetailUserDrawer open={detailOpen} detailData={detailData} onClose={() => setDetailOpen(false)} />
      <BindRolesModal open={bindRolesOpen} detailData={detailData} onClose={() => setBindRolesOpen(false)} onNotice={onModalNotice} />
      <ResetPasswordModal open={resetPwdOpen} detailData={detailData} onClose={() => setResetPwdOpen(false)} />
    </PageContainer>
  );
}
