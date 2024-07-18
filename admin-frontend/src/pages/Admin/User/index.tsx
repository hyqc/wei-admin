import React, { useState, useEffect } from 'react';
import { Content, Search, Container } from '@/components/PageListContainer';
import {
  Form,
  Input,
  Button,
  Select,
  Row,
  Col,
  Table,
  Avatar,
  Switch,
  message,
  Tag,
  Popconfirm,
  Space,
} from 'antd';
import { SearchOutlined, ReloadOutlined, PlusOutlined } from '@ant-design/icons';
import type { Gutter } from 'antd/lib/grid/row';
import { ColumnsType } from 'antd/lib/table';
import {
  adminUserList,
  ResponseAdminUserListItemRolesItemType,
  adminUserDetail,
  adminUserDelete,
  adminUserEnable,
  RequestAdminUserEnableParamsType,
} from '@/services/apis/admin/user';
import type {
  ResponseAdminUserListItemType,
  RequestAdminUserListParamsType,
} from '@/services/apis/admin/user';
import type {
  PageInfoType,
  ResponseBodyType,
  ResponseListDataType,
  ResponseListType,
} from '@/services/apis/types';
import { DEFAULT_PAGE_INFO } from '@/services/apis/config';
import AdminUserAddModal, { NoticeModalPropsType } from './add';
import AdminUserEditModal from './edit';
import AdminUserDetailModal from './detail';
import AdminUserBindRolesModal from './bind';
import AdminUserEditPasswordModal from './password';
import { adminRoleAll, ResponseAdminRoleAllItemType } from '@/services/apis/admin/role';
import Authorization from '@/components/Autuorization';
import FetchButton from '@/components/FetchButton';

const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];
const FormSearchRowColSpan = 5.2;

const Admin: React.FC = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [rowsData, setRowsData] = useState<ResponseAdminUserListItemType[]>([]);
  const [detailData, setDetailData] = useState<any>();
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [editModalStatus, setEditModalStatus] = useState<boolean>(false);
  const [addModalStatus, setAddModalStatus] = useState<boolean>(false);
  const [bindModalStatus, setBindRolesModalStatus] = useState<boolean>(false);
  const [editPasswordModalStatus, setEditPasswordModalStatus] = useState<boolean>(false);
  const [roleOptions, setRoleOptions] = useState<ResponseAdminRoleAllItemType[]>([]);

  const columns: ColumnsType<any> = [
    {
      title: '账号',
      align: 'left',
      dataIndex: 'username',
      width: '8rem',
      sorter: true,
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
      render: (avatar, record) => {
        return <Avatar src={avatar} />;
      },
    },
    {
      title: '角色',
      width: '8rem',
      dataIndex: 'roles',
      render: (roles, record: ResponseAdminUserListItemType) => {
        if (record.adminId === AdminId) {
          return (
            <Tag color="geekblue" style={{ cursor: 'default' }}>
              超管
            </Tag>
          );
        }
        return roles?.map((item: ResponseAdminUserListItemRolesItemType) => {
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
    },
    {
      title: '更新时间',
      align: 'center',
      width: '12rem',
      dataIndex: 'modifyTime',
      sorter: true,
    },
    {
      title: '状态',
      width: '6rem',
      align: 'center',
      dataIndex: 'enabled',
      render(enabled: boolean, record: ResponseAdminUserListItemType) {
        if (record.adminId === AdminId) {
          return (
            <Switch
              checkedChildren={'启用'}
              disabled
              unCheckedChildren={'禁用'}
              checked={enabled}
            />
          );
        }
        return (
          <Authorization
            name="AdminUserEdit"
            forbidden={
              <>
                <Switch
                  disabled
                  checkedChildren={'启用'}
                  unCheckedChildren={'禁用'}
                  checked={enabled}
                />
              </>
            }
          >
            <Popconfirm
              title={`确定要${record.enabled ? '禁用' : '启用'}该账号吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateEnabled(record)}
            >
              <Switch checkedChildren={'启用'} unCheckedChildren={'禁用'} checked={enabled} />
            </Popconfirm>
          </Authorization>
        );
      },
    },
    {
      title: '操作',
      align: 'left',
      width: '1rem',
      render(text, record: ResponseAdminUserListItemType) {
        return (
          <Space>
            <Authorization name="AdminUserView">
              <FetchButton onClick={() => openDetailModal(record)}>详情</FetchButton>
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
                  {!record.enabled ? (
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

  // 获取账号列表
  function getRows(data?: RequestAdminUserListParamsType) {
    setLoading(true);
    adminUserList(data)
      .then((res: ResponseListType) => {
        const data: ResponseListDataType = res.data;
        const rows = data?.list || [];
        const page = { total: data.total, pageSize: data.pageSize, pageNum: data.pageNum };
        setPageInfo(page);
        setRowsData(rows);
      })
      .catch((err) => {
        console.log('error', err);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  // 账号状态更新
  function updateEnabled(record: ResponseAdminUserListItemType) {
    const updateData: RequestAdminUserEnableParamsType = {
      adminId: record.adminId,
      enabled: !record.enabled,
    };
    adminUserEnable(updateData).then((res) => {
      message.success(res.message, MessageDuritain, () => {
        getRows({ ...pageInfo, ...form.getFieldsValue() });
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
  function onDelete(record: ResponseAdminUserListItemType) {
    adminUserDelete({ adminId: record.adminId, enabled: record.enabled }).then(
      (res: ResponseBodyType) => {
        message.success(res.message, MessageDuritain, () => {
          getRows({ ...pageInfo, ...form.getFieldsValue() });
        });
      },
    );
  }

  // 账号详情
  function openDetailModal(record: ResponseAdminUserListItemType) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      setDetailData(res.data);
      setDetailModalStatus(true);
    });
  }

  // 分配角色
  function openBindRolesModal(record: ResponseAdminUserListItemType) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      setDetailData(res.data);
      setBindRolesModalStatus(true);
    });
  }

  // 账号编辑
  function openEditModal(record: ResponseAdminUserListItemType) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      setDetailData(res.data);
      setEditModalStatus(true);
    });
  }

  // 账号添加
  function openAddModal() {
    setAddModalStatus(true);
  }

  function openEditPasswordModal(record: ResponseAdminUserListItemType) {
    adminUserDetail({ adminId: record.adminId }).then((res) => {
      setDetailData(res.data);
      setEditPasswordModalStatus(true);
    });
  }

  function noticeAddModal(data: NoticeModalPropsType) {
    setAddModalStatus(false);
    if (data.reload) {
      getRows({ ...pageInfo, ...form.getFieldsValue() });
    }
  }

  function noticeEditModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditModalStatus(false);
    if (data.reload) {
      getRows({ ...pageInfo, ...form.getFieldsValue() });
    }
  }

  function noticeDetailModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setDetailModalStatus(false);
    if (data.reload) {
      getRows({ ...pageInfo, ...form.getFieldsValue() });
    }
  }

  function noticeBindRolesModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setBindRolesModalStatus(false);
    if (data.reload) {
      getRows({ ...pageInfo, ...form.getFieldsValue() });
    }
  }

  function noticeEditPasswordModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditPasswordModalStatus(false);
    if (data.reload) {
      getRows({ ...pageInfo, ...form.getFieldsValue() });
    }
  }

  // 列表搜索
  function onSearchFinish(values: RequestAdminUserListParamsType) {
    const page = { ...pageInfo, pageNum: 1 };
    getRows({ ...values, ...page });
  }

  // 搜索重置
  function onSearchReset() {
    form.resetFields();
    getRows({ pageNum: 1, pageSize: pageInfo.pageSize });
  }

  function onShowSizeChange(current: number, size: number) {
    const page = { pageSize: size, pageNum: current };
    setPageInfo({ ...pageInfo, ...page });
  }

  function tableChange(pagination: any, filters: any, sorter: any) {
    const page = { pageSize: pagination.pageSize, pageNum: pagination.current };
    setPageInfo({ ...pageInfo, ...page });
    getRows({
      ...form.getFieldsValue(),
      ...page,
      sortField: sorter.field,
      sortType: sorter.order,
    });
  }

  useEffect(() => {
    onSearchReset();
  }, []);

  useEffect(() => {
    fetchAdminRoles();
  }, []);

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
                      <Select.Option key={item.roleId} value={item.roleId}>
                        {item.roleName}
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
            showQuickJumper: true,
            position: ['bottomRight'],
            showTotal: (total) => `共 ${total} 条数据`,
            onShowSizeChange,
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
