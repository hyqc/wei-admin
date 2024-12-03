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
  Popconfirm,
  Switch,
  Tag,
  Tooltip,
  App,
} from 'antd';
import { SearchOutlined, ReloadOutlined, PlusOutlined } from '@ant-design/icons';
import type { Gutter } from 'antd/lib/grid/row';
import {
  PageInfoType,
  ResponseBodyType,
} from '@/services/apis/types';
import {
  adminPermissionDelete,
  adminPermissionDetail,
  adminPermissionEnable,
  adminPermissionList,
  adminPermissionBindApis,
  adminPermissionUnBindApi,
} from '@/services/apis/admin/permission';
import { DEFAULT_PAGE_INFO } from '@/services/apis/config';
import { ColumnsType } from 'antd/lib/table';
import Authorization from '@/components/Autuorization';
import AdminPermissionAddModal, { NoticeModalPropsType } from './add';
import AdminPermissionEditModal from './edit';
import AdminPermissionDetailModal from './detail';
import AdminPermissionBindApiModal from './bind';
import { adminMenuPages } from '@/services/apis/admin/menu';
import PageMenus from './components/PageMenus';
import FetchButton from '@/components/FetchButton';
import { DefaultPagination } from '@/components/PageContainer/Pagination';
import { PermissionListItem, ReqAdminPermissionEnable, ReqAdminPermissionList, RespAdminPermissionListData } from '@/proto/admin_ts/admin_permission';
import { AdminApiItem } from '@/proto/admin_ts/common';
import { handlePagination, searchResetPageInfo } from '@/services/common/utils';
import { RowEnabledButton } from '@/components';
import { MenuTreeItem } from '@/proto/admin_ts/admin_menu';

const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];
const FormSearchRowColSpan = 5.2;

const Admin: React.FC = () => {
  const { message } = App.useApp()
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [detailData, setDetailData] = useState<any>();
  const [rowsData, setRowsData] = useState<PermissionListItem[]>([]);
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [editModalStatus, setEditModalStatus] = useState<boolean>(false);
  const [addModalStatus, setAddModalStatus] = useState<boolean>(false);
  const [bindAPIModalStatus, setBindAPIModalStatus] = useState<boolean>(false);
  const [pageMenusData, setPageMenusData] = useState<MenuTreeItem[]>([]);

  const columns: ColumnsType<any> = [
    {
      title: 'ID',
      dataIndex: 'id',
      width: '4rem',
      align: 'center',
      sorter: true,
    },
    {
      title: '菜单名称',
      align: 'left',
      dataIndex: 'menuName',
    },
    {
      title: '名称',
      align: 'left',
      dataIndex: 'name',
      render(name: string, record: PermissionListItem) {
        return <Tooltip title={record.key}>{name}</Tooltip>;
      },
    },
    {
      title: '类型',
      align: 'left',
      dataIndex: 'typeText',
      width: '3rem',
      render(type: string, record: PermissionListItem) {
        return record.type === 'view' ? (
          <Tag color="#87d068">{record.typeText}</Tag>
        ) : record.type === 'edit' ? (
          <Tag color="#108ee9">{record.typeText}</Tag>
        ) : (
          <Tag color="#f50">{record.typeText}</Tag>
        );
      },
    },
    {
      title: '关联接口',
      align: 'left',
      dataIndex: 'apis',
      render(apis: PermissionListItem[], record: PermissionListItem) {
        return showApisTag(apis, record);
      },
    },
    {
      title: '状态',
      width: '6rem',
      align: 'center',
      dataIndex: 'isEnabled',
      render(isEnabled: boolean, record: PermissionListItem) {
        return (
          <Authorization
            name="AdminPermissionEdit"
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
      width: '15rem',
      render(text, record: PermissionListItem) {
        return (
          <Space>
            <Authorization name="AdminPermissionView">
              <FetchButton onClick={() => openDetailModal(record)}>详情</FetchButton>
            </Authorization>
            <Authorization name="AdminPermissionEdit">
              <FetchButton onClick={() => openEditModal(record)}>编辑</FetchButton>
            </Authorization>
            <Authorization name="AdminPermissionEdit">
              <FetchButton onClick={() => openBindAPIModal(record)}>绑定接口</FetchButton>
            </Authorization>
            {/* 禁用的才能删除 */}
            <Authorization name="AdminPermissionDelete">
              {!record.isEnabled ? (
                <Popconfirm
                  title="确定要删除该权限吗？"
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

  // 获取权限列表
  function getRows(data?: ReqAdminPermissionList) {
    setLoading(true);
    adminPermissionList(data)
      .then((res) => {
        const data: RespAdminPermissionListData = res.data;
        const rows = data?.list || [];
        const total = data?.total ?? 0
        setRowsData(rows);
        setPageInfo((pageInfo) => ({ ...pageInfo, total }))
      })
      .catch((err) => {
        console.log('error', err);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  // 权限状态更新
  function updateEnabled(record: PermissionListItem) {
    const updateData: ReqAdminPermissionEnable = {
      id: record.id,
      isEnabled: !record.isEnabled,
    };
    adminPermissionEnable(updateData).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
        getRows({ base, ...form.getFieldsValue() });
      });
    });
  }

  // 权限添加
  function openAddModal() {
    setAddModalStatus(true);
  }

  // 权限详情
  function openDetailModal(record: PermissionListItem) {
    adminPermissionDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setDetailModalStatus(true);
    });
  }

  // 权限编辑
  function openEditModal(record: PermissionListItem) {
    adminPermissionDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setEditModalStatus(true);
    });
  }

  // 删除绑定的接口
  function deleteBindApi(permissionId: number, apiId: number) {
    adminPermissionUnBindApi({ permissionId, apiId }).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
        getRows({ base, ...form.getFieldsValue() });
      });
    });
  }

  function showApisTag(
    apis: AdminApiItem[],
    record: PermissionListItem,
  ) {
    return apis?.map((item) => {
      return (
        <Authorization
          key={item.id}
          name="AdminPermissionEdit"
          forbidden={
            <>
              <Tag style={{ cursor: 'no-drop' }}>{item.name}</Tag>
            </>
          }
        >
          <Popconfirm
            title={`确定要解绑${item.path}接口吗？`}
            okText="确定"
            cancelText="取消"
            onConfirm={() => deleteBindApi(record.id ?? 0, item.id ?? 0)}
          >
            <Tag style={{ cursor: 'pointer', margin: '4px' }}>{item.name}</Tag>
          </Popconfirm>
        </Authorization>
      );
    });
  }

  function openBindAPIModal(record: PermissionListItem) {
    adminPermissionDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setBindAPIModalStatus(true);
    });
  }

  function noticeAddModal(data: NoticeModalPropsType) {
    setAddModalStatus(false);
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

  function noticeEditModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  function noticeBindModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setBindAPIModalStatus(false);
    if (data.reload) {
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    }
  }

  // 删除权限
  function onDelete(record: PermissionListItem) {
    adminPermissionDelete({ id: record.id ?? 0 }).then((res) => {
      message.success(res.msg, MessageDuritain);
      const base = handlePagination(pageInfo.pageNum, pageInfo.pageSize)
      getRows({ base, ...form.getFieldsValue() });
    });
  }

  // 管理员列表搜索
  function onSearchFinish(values: ReqAdminPermissionList) {
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
    const page = { ...pageInfo, pageSize: size, pageNum: current }
    setPageInfo(() => ({ ...page }));
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

  function getPageMenus() {
    adminMenuPages().then((res: ResponseBodyType) => {
      const data: MenuTreeItem[] = res.data.list ?? []
      setPageMenusData(data);
    });
  }

  useEffect(() => {
    getPageMenus();
  }, []);

  useEffect(() => {
    onSearchReset();
  }, []);

  return (
    <Container>
      <Search>
        <Form form={form} onFinish={onSearchFinish}>
          <Row gutter={FormSearchRowGutter}>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="菜单名称" name="menuId" initialValue={0}>
                <PageMenus data={pageMenusData}>
                  <Select.Option value={0}>全部</Select.Option>
                </PageMenus>
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="权限名称" name="name" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="权限键名" name="key" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="权限类型" name="type" initialValue={''}>
                <Select>
                  <Select.Option value="">全部</Select.Option>
                  <Select.Option value="view">查看</Select.Option>
                  <Select.Option value="edit">编辑</Select.Option>
                  <Select.Option value="delete">删除</Select.Option>
                </Select>
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="权限状态" name="enabled" initialValue={0}>
                <Select>
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
                  重置
                </Button>
              </Space>
            </Col>
          </Row>
        </Form>
      </Search>

      <Content>
        {/* button */}
        <Space style={{ marginBottom: '1rem' }}>
          <Authorization name="AdminPermissionEdit">
            <Button type="primary" onClick={openAddModal}>
              <PlusOutlined />
              新建权限
            </Button>
          </Authorization>
        </Space>

        {/* table */}
        <Table
          sticky
          rowKey="id"
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

      <AdminPermissionAddModal
        pageMenusData={pageMenusData}
        modalStatus={addModalStatus}
        noticeModal={noticeAddModal}
      />

      <AdminPermissionDetailModal
        detailData={detailData}
        modalStatus={detailModalStatus}
        noticeModal={noticeDetailModal}
      />

      <AdminPermissionEditModal
        pageMenusData={pageMenusData}
        modalStatus={editModalStatus}
        detailData={detailData}
        noticeModal={noticeEditModal}
      />

      <AdminPermissionBindApiModal
        modalStatus={bindAPIModalStatus}
        detailData={detailData}
        noticeModal={noticeBindModal}
      />
    </Container>
  );
};

export default Admin;
