import React, { useEffect, useState } from 'react';
import { Container, Content, Search } from '@/components/PageContainer';
import {
  Form,
  Input,
  Button,
  Select,
  Row,
  Col,
  Space,
  Table,
  message,
  Popconfirm,
} from 'antd';
import { SearchOutlined, ReloadOutlined, PlusOutlined } from '@ant-design/icons';
import type { Gutter } from 'antd/lib/grid/row';
import { PageInfoType } from '@/services/apis/types';
import {
  adminRoleDelete,
  adminRoleDetail,
  adminRoleEnable,
  adminRoleList,
} from '@/services/apis/admin/role';
import { DEFAULT_PAGE_INFO } from '@/services/apis/config';
import { ColumnsType } from 'antd/lib/table';
import Authorization from '@/components/Autuorization';
import AdminRoleAddModal, { NoticeModalPropsType } from './add';
import AdminRoleEditModal from './edit';
import AdminRoleDetailModal from './detail';
import AdminBindModal from './bind';
import { adminMenuMode, ResponseAdminMenuModeTypeData } from '@/services/apis/admin/menu';
import FetchButton from '@/components/FetchButton';
import { RowEnabledButton } from '@/components';
import { handlePagination } from '@/services/common/utils';
import { ReqAdminRoleEnable, ReqAdminRoleList, RespAdminRoleListData, RoleItem } from '@/proto/admin_ts/admin_role';
import { DefaultPagination } from '@/components/PageContainer/Pagination';

const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];
const FormSearchRowColSpan = 5.2;

const Admin: React.FC = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [detailData, setDetailData] = useState<any>(undefined);
  const [rowsData, setRowsData] = useState<RoleItem[]>([]);
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [editModalStatus, setEditModalStatus] = useState<boolean>(false);
  const [addModalStatus, setAddModalStatus] = useState<boolean>(false);
  const [bindPermissionsModalStatus, setBindPermissionsModalStatus] = useState<boolean>(false);
  const [modelPageData, setModelPageData] = useState<ResponseAdminMenuModeTypeData[]>([]);
  const columns: ColumnsType<any> = [
    {
      title: 'ID',
      dataIndex: 'id',
      width: '6rem',
      align: 'center',
      sorter: true,
    },
    {
      title: '名称',
      align: 'center',
      width: '12rem',
      dataIndex: 'name',
    },
    {
      title: '名称',
      align: 'center',
      width: '12rem',
      dataIndex: 'describe',
    },
    {
      title: '创建时间',
      align: 'center',
      dataIndex: 'createdAt',
      width: '10rem',
      sorter: true,
    },
    {
      title: '更新时间',
      align: 'center',
      dataIndex: 'updatedAt',
      width: '10rem',
      sorter: true,
    },
    {
      title: '状态',
      width: '6rem',
      align: 'center',
      dataIndex: 'isEnabled',
      render(isEnabled: boolean, record: RoleItem) {
        return (
          <Authorization
            name="AdminRoleEdit"
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
      width: '2rem',
      render(text, record: RoleItem) {
        return (
          <Space>
            <Authorization name="AdminRoleView">
              <FetchButton onClick={() => openDetailModal(record)}>详情</FetchButton>
            </Authorization>
            <Authorization name="AdminRoleEdit">
              <FetchButton onClick={() => openEditModal(record)}>编辑</FetchButton>
            </Authorization>

            <Authorization name="AdminRoleEdit">
              <FetchButton onClick={() => openBindPermissionsModal(record)}>分配权限</FetchButton>
            </Authorization>

            {/* 禁用的才能删除 */}
            <Authorization name="AdminRoleDelete">
              {!record.isEnabled ? (
                <Popconfirm
                  title="确定要删除该角色吗？"
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
          </Space>
        );
      },
    },
  ];

  // 获取角色列表
  function getRows(data?: ReqAdminRoleList) {
    setLoading(true);
    adminRoleList(data)
      .then((res) => {
        const data: RespAdminRoleListData = res.data;
        const rows = data?.list || [];
        const total = data.total ?? 0
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

  // 角色状态更新
  function updateEnabled(record: RoleItem) {
    const updateData: ReqAdminRoleEnable = {
      id: record.id,
      enabled: !record.isEnabled,
    };
    adminRoleEnable(updateData).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        getRows({ base: { ...pageInfo }, ...form.getFieldsValue() });
      });
    });
  }

  // 角色详情
  function openDetailModal(record: RoleItem) {
    adminRoleDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setDetailModalStatus(true);
    });
  }

  // 角色绑定权限
  function openBindPermissionsModal(record: RoleItem) {
    adminRoleDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setBindPermissionsModalStatus(true);
    });
  }

  // 角色编辑
  function openEditModal(record: RoleItem) {
    adminRoleDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setEditModalStatus(true);
    });
  }

  // 角色添加
  function openAddModal() {
    setAddModalStatus(true);
  }

  function noticeAddModal(data: NoticeModalPropsType) {
    setAddModalStatus(false);
    if (data.reload) {
      getRows({ base: { ...pageInfo }, ...form.getFieldsValue() });
    }
  }

  function noticeDetailModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setDetailModalStatus(false);
    if (data.reload) {
      getRows({ base: { ...pageInfo }, ...form.getFieldsValue() });
    }
  }

  function noticeEditModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditModalStatus(false);
    if (data.reload) {
      getRows({ base: { ...pageInfo }, ...form.getFieldsValue() });
    }
  }

  function noticeBindModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setBindPermissionsModalStatus(false);
  }

  // 删除角色
  function onDelete(record: RoleItem) {
    adminRoleDelete({ id: record.roleId }).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
          getRows({ base, ...form.getFieldsValue() });
      });
    });
  }

  // 列表搜索
  function onSearchFinish(values: ReqAdminRoleList) {
    const base = handlePagination(1, pageInfo.pageSize)
    getRows({ ...values, base });
  }

  // 搜索重置
  function onSearchReset() {
    form.resetFields();
    const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
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

  function getAdminMenuModeData() {
    adminMenuMode().then((res) => {
      setModelPageData(res.data || []);
    });
  }

  useEffect(() => {
    onSearchReset();
    getAdminMenuModeData();
  }, []);

  return (
    <Container>
      <Search>
        <Form form={form} onFinish={onSearchFinish}>
          <Row gutter={FormSearchRowGutter}>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="名称" name="roleName" initialValue={''}>
                <Input />
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
          <Authorization name="AdminRoleEdit">
            <Button type="primary" onClick={openAddModal}>
              <PlusOutlined />
              新建角色
            </Button>
          </Authorization>
        </Space>

        {/* table */}
        <Table
          sticky
          rowKey="roleId"
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

      <AdminRoleAddModal modalStatus={addModalStatus} noticeModal={noticeAddModal} />
      {detailData ? (
        <>
          <AdminRoleDetailModal
            modalStatus={detailModalStatus}
            detailData={detailData}
            menuPageData={modelPageData}
            noticeModal={noticeDetailModal}
          />

          <AdminRoleEditModal
            modalStatus={editModalStatus}
            detailData={detailData}
            noticeModal={noticeEditModal}
          />

          <AdminBindModal
            modalStatus={bindPermissionsModalStatus}
            detailData={detailData}
            menuPageData={modelPageData}
            noticeModal={noticeBindModal}
          />
        </>
      ) : (
        <></>
      )}
    </Container>
  );
};

export default Admin;
