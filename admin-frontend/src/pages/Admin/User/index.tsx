import React, { useState, useEffect } from 'react';
import { Content, Search, Container } from '@/components/PageContainer';
import { RowEnabledButton } from '@/components/CustomButton/enabled'
import {
  Form,
  Input,
  Button,
  Select,
  Row,
  Col,
  Table,
  Avatar,
  Tag,
  Popconfirm,
  Space,
  App,
} from 'antd';
import { SearchOutlined, ReloadOutlined, PlusOutlined } from '@ant-design/icons';
import type { Gutter } from 'antd/lib/grid/row';
import { ColumnsType } from 'antd/lib/table';
import {
  adminUserList,
  adminUserDetail,
  adminUserDelete,
  adminUserEnable,
} from '@/services/apis/admin/user';
import type {
} from '@/services/apis/admin/user';
import type {
  PageInfoType,
  ResponseBodyType,
} from '@/services/apis/types';
import { DEFAULT_PAGE_INFO } from '@/services/apis/config';
import AdminUserAddModal, { NoticeModalPropsType } from './add';
import AdminUserEditModal from './edit';
import AdminUserDetailModal from './detail';
import AdminUserBindRolesModal from './bind';
import AdminUserEditPasswordModal from './password';
import { adminRoleAll } from '@/services/apis/admin/role';
import Authorization from '@/components/Autuorization';
import FetchButton from '@/components/FetchButton';
import { AdminUserListItem, AdminUserRoleItem } from '@/proto/admin_ts/common';
import { ReqAdminUserEnabled, ReqAdminUserList, RespAdminUserInfoData, RespAdminUserListData } from '@/proto/admin_ts/admin_user';
import { handlePagination, searchResetPageInfo } from '@/services/common/utils';
import { DefaultPagination } from '@/components/PageContainer/Pagination';
import { RoleItem } from '@/proto/admin_ts/admin_role';


const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];
const FormSearchRowColSpan = 5.2;

const Admin: React.FC = () => {
  const { message } = App.useApp();
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [rowsData, setRowsData] = useState<AdminUserListItem[]>([]);
  const [detailData, setDetailData] = useState<any>();
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [editModalStatus, setEditModalStatus] = useState<boolean>(false);
  const [addModalStatus, setAddModalStatus] = useState<boolean>(false);
  const [bindModalStatus, setBindRolesModalStatus] = useState<boolean>(false);
  const [editPasswordModalStatus, setEditPasswordModalStatus] = useState<boolean>(false);
  const [roleOptions, setRoleOptions] = useState<RoleItem[]>([]);

  // 获取账号列表
  function getRows(data?: ReqAdminUserList) {
    setLoading(true);
    adminUserList(data)
      .then((res: ResponseBodyType) => {
        const data: RespAdminUserListData = res.data;
        const rows = data?.list || [];
        const total = data?.total ?? 0
        setRowsData(rows);
        setPageInfo((pageInfo)=>({...pageInfo,total}))
      })
      .catch((err) => {
        console.log('error', err);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  // 账号状态更新
  function updateEnabled(record: AdminUserListItem) {
    const updateData: ReqAdminUserEnabled = {
      adminId: record.adminId,
      enabled: !record.isEnabled,
    };
    adminUserEnable(updateData).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        getRows({ base: { ...pageInfo }, ...form.getFieldsValue() });
      });
    });
  }

  function fetchAdminRoles(name?: string) {
    adminRoleAll({ name }).then((res) => {
      res.data.unshift({ roleId: 0, roleName: '全部' });
      setRoleOptions(res.data || []);
    });
  }

  // 删除账号
  function onDelete(record: AdminUserListItem) {
    adminUserDelete({ adminId: record.adminId }).then(
      (res) => {
        message.success(res.msg, MessageDuritain, () => {
          const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
          getRows({ base, ...form.getFieldsValue() });
        });
      },
    );
  }

  // 账号详情
  function openDetailModal(record: AdminUserListItem) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      const data: RespAdminUserInfoData = res.data;
      setDetailData(data.data);
      setDetailModalStatus(true);
    });
  }

  // 分配角色
  function openBindRolesModal(record: AdminUserListItem) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      const data: RespAdminUserInfoData = res.data;
      setDetailData(data.data);
      setBindRolesModalStatus(true);
    });
  }

  // 账号编辑
  function openEditModal(record: AdminUserListItem) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      const data: RespAdminUserInfoData = res.data;
      setDetailData(data.data);
      setEditModalStatus(true);
    });
  }

  // 账号添加
  function openAddModal() {
    setAddModalStatus(true);
  }

  function openEditPasswordModal(record: AdminUserListItem) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      const data: RespAdminUserInfoData = res.data;
      setDetailData(data.data);
      setEditPasswordModalStatus(true);
    });
  }

  function noticeAddModal(data: NoticeModalPropsType) {
    setAddModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  function noticeEditModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  function noticeDetailModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setDetailModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  function noticeBindRolesModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setBindRolesModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  function noticeEditPasswordModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditPasswordModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  // 列表搜索
  function onSearchFinish(values: ReqAdminUserList) {
    const base = handlePagination(1, pageInfo.pageSize)
    getRows({ ...values, base });
  }

  // 搜索重置
  function onSearchReset() {
    form.resetFields();
    setPageInfo(searchResetPageInfo(pageInfo.pageSize))
    const base = handlePagination(1, pageInfo.pageSize)
    getRows({ base });
  }

  function onShowSizeChange(current: number, size: number) {
    const page = {...pageInfo, pageSize: size, pageNum: current }
    setPageInfo(()=>({...page}));
  }

  function tableChange(pagination: any, filters: any, sorter: any) {
    const page = handlePagination(pagination.current, pagination.pageSize)
    setPageInfo({ ...pageInfo, ...page });
    const base = { ...page, sortField: sorter.field, sortType: sorter.order, }
    getRows({
      ...form.getFieldsValue(),
      base,
    });
  }

  useEffect(() => {
    onSearchReset();
  }, []);

  useEffect(() => {
    fetchAdminRoles();
  }, []);

  useEffect(() => {
    console.log('pageInfo: ', pageInfo)
  }, [pageInfo]);


  const columns: ColumnsType<any> = [
    {
      title: '账号',
      align: 'left',
      dataIndex: 'username',
      width: '8rem',
      sorter: true,
      fixed: 'left',
    },
    {
      title: '昵称',
      align: 'left',
      dataIndex: 'nickname',
      width: '8rem',
    },
    {
      title: '头像',
      align: 'center',
      dataIndex: 'avatar',
      width: '3rem',
      render: (avatar) => {
        return <Avatar src={avatar} />;
      },
    },
    {
      title: '邮箱',
      align: 'center',
      width: '12rem',
      dataIndex: 'email',
    },
    {
      title: '角色',
      width: '8rem',
      dataIndex: 'roles',
      render: (roles, record: AdminUserListItem) => {
        if (record.adminId === AdminId) {
          return (
            <Tag color="geekblue" style={{ cursor: 'default' }}>
              超管
            </Tag>
          );
        }
        return roles?.map((item: AdminUserRoleItem) => {
          return (
            <Tag color="geekblue" style={{ cursor: 'default' }} key={item.roleId}>
              {item.roleName}
            </Tag>
          );
        });
      },
    },
    {
      title: '登录次数',
      align: 'center',
      width: '7rem',
      dataIndex: 'loginTotal',
      sorter: true,
      render(loginTotal: number) {
        return <>{loginTotal ?? 0}</>
      },
    },
    {
      title: '更新时间',
      align: 'center',
      width: '12rem',
      dataIndex: 'updatedAt',
      sorter: true,
    },
    {
      title: '状态',
      width: '6rem',
      align: 'center',
      dataIndex: 'isEnabled',
      render(isEnabled: boolean, record: AdminUserListItem) {
        if (record.adminId === AdminId) {
          return (
            <RowEnabledButton isEnabled={isEnabled} disabled={true} />
          );
        }
        return (
          <Authorization
            name="AdminUserEdit"
            forbidden={
              <RowEnabledButton isEnabled={isEnabled} disabled={false} />
            }
          >
            <Popconfirm
              title={`确定要${record.isEnabled ? '禁用' : '启用'}该账号吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateEnabled(record)}
            >
              <RowEnabledButton isEnabled={isEnabled} disabled={false} />
            </Popconfirm>
          </Authorization>
        );
      },
    },
    {
      title: '操作',
      align: 'left',
      width: '1rem',
      fixed: 'right',
      render(text, record: AdminUserListItem) {
        return (
          <Space>
            <Authorization name="AdminUserView">
              <FetchButton onClick={() => openDetailModal(record)} >详情</FetchButton>
            </Authorization>
            {record.adminId === AdminId ? (
              <></>
            ) : (
              <>
                <Authorization name="AdminUserEdit">
                  <FetchButton onClick={() => openBindRolesModal(record)}>分配角色</FetchButton>
                </Authorization>

                <Authorization name="AdminUserEdit">
                  <FetchButton onClick={() => openEditModal(record)}>编辑</FetchButton>
                </Authorization>

                <Authorization name="AdminUserEdit">
                  {/* 非超管修改密码 */}
                  <FetchButton onClick={() => openEditPasswordModal(record)}>修改密码</FetchButton>
                </Authorization>

                {/* 禁用的才能删除 */}
                <Authorization name="AdminUserDelete">
                  {!record.isEnabled ? (
                    <Popconfirm
                      title="确定要删除该账号吗？"
                      okText="确定"
                      cancelText="取消"
                      onConfirm={() => onDelete(record)}
                    >
                      <FetchButton danger>删除</FetchButton>
                    </Popconfirm>
                  ) : (
                    ''
                  )}
                </Authorization>
              </>
            )}
          </Space>
        );
      },
    },
  ];


  return (
    <Container>
      <Search>
        <Form form={form} onFinish={onSearchFinish}>
          <Row gutter={FormSearchRowGutter}>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="账号" name="username" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="昵称" name="nickname" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="邮箱" name="email" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="角色名" name="roleId" initialValue={0}>
                <Select
                  style={{ offset: 0, width: '120' }}
                  showSearch
                  filterOption={(input, option) => {
                    return (option!.children as unknown as string)
                      .toLowerCase()
                      .includes(input.toLowerCase());
                  }}
                >
                  {roleOptions?.map((item) => {
                    return (
                      <Select.Option key={item.id} value={item.id}>
                        {item.name}
                      </Select.Option>
                    );
                  })}
                </Select>
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="状态" name="enabled" initialValue={0}>
                <Select style={{ offset: 0, width: '120' }}>
                  <Select.Option value={0}>全部</Select.Option>
                  <Select.Option value={1}>启用</Select.Option>
                  <Select.Option value={2}>禁用</Select.Option>
                </Select>
              </Form.Item>
            </Col>
          </Row>
          <Row>
            <Col span={24} style={{ textAlign: 'right' }}>
              <Space>
                <Button type="primary" htmlType="submit">
                  <SearchOutlined />
                  查询
                </Button>
                <Button type="primary" htmlType="button" onClick={onSearchReset}>
                  <ReloadOutlined />
                  清除
                </Button>
              </Space>
            </Col>
          </Row>
        </Form>
      </Search>

      <Content>
        {/* button */}
        <Space style={{ marginBottom: '1rem' }}>
          <Authorization name="AdminUserEdit">
            <Button type="primary" onClick={openAddModal}>
              <PlusOutlined />
              新建账号
            </Button>
          </Authorization>
        </Space>

        {/* table */}
        <Table
          sticky
          rowKey="adminId"
          scroll={{ x: 'auto' }}
          loading={loading}
          columns={columns}
          dataSource={rowsData}
          onChange={tableChange}
          pagination={{
            current: pageInfo.pageNum,
            pageSize: pageInfo.pageSize,
            total: pageInfo.total,
            onShowSizeChange,
            ...DefaultPagination,
          }}
        />
      </Content>

      {/* modal */}

      <AdminUserAddModal modalStatus={addModalStatus} noticeModal={noticeAddModal} />

      {detailData ? (
        <>
          <AdminUserDetailModal
            modalStatus={detailModalStatus}
            detailData={detailData}
            noticeModal={noticeDetailModal}
          />

          <AdminUserEditModal
            modalStatus={editModalStatus}
            detailData={detailData}
            noticeModal={noticeEditModal}
          />

          <AdminUserBindRolesModal
            modalStatus={bindModalStatus}
            detailData={detailData}
            noticeModal={noticeBindRolesModal}
          />

          <AdminUserEditPasswordModal
            modalStatus={editPasswordModalStatus}
            detailData={detailData}
            noticeModal={noticeEditPasswordModal}
          />
        </>
      ) : (
        <></>
      )}
    </Container>
  );
};

export default Admin;
