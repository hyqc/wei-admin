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
  Switch,
  App,
} from 'antd';
import { SearchOutlined, ReloadOutlined, PlusOutlined } from '@ant-design/icons';
import type { Gutter } from 'antd/lib/grid/row';
import {
  adminAPIDelete,
  adminAPIDetail,
  adminAPIEnable,
  adminAPIList,
} from '@/services/apis/admin/resource';
import { DEFAULT_PAGE_INFO } from '@/services/apis/config';
import { ColumnsType } from 'antd/lib/table';
import Authorization from '@/components/Autuorization';
import AdminAPIAddModal, { NoticeModalPropsType } from './add';
import AdminAPIEditModal from './edit';
import AdminAPIDetailModal from './detail';
import FetchButton from '@/components/FetchButton';
import { DefaultPagination } from '@/components/PageContainer/Pagination';
import { ReqAdminApiEnable, ReqAdminApiList, RespAdminApiListData } from '@/proto/admin_ts/admin_api';
import { AdminApiItem } from '@/proto/admin_ts/common';
import { PageInfoType } from '@/services/apis/types';
import { handlePagination, searchResetPageInfo } from '@/services/common/utils';
import { RowEnabledButton } from '@/components';

const FormSearchRowGutter: [Gutter, Gutter] = [12, 0];
const FormSearchRowColSpan = 5.2;

const Admin: React.FC = () => {
  const { message } = App.useApp()
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [detailData, setDetailData] = useState<AdminApiItem>();
  const [rowsData, setRowsData] = useState<AdminApiItem[]>([]);
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [editModalStatus, setEditModalStatus] = useState<boolean>(false);
  const [addModalStatus, setAddModalStatus] = useState<boolean>(false);

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
      align: 'left',
      width: '12rem',
      dataIndex: 'name',
    },
    {
      title: '路由',
      align: 'left',
      dataIndex: 'path',
    },
    {
      title: '唯一键',
      align: 'left',
      dataIndex: 'key',
      sorter: true,
    },
    {
      title: '创建时间',
      align: 'center',
      width: '15rem',
      dataIndex: 'createdAt',
      sorter: true,
    },
    {
      title: '更新时间',
      align: 'center',
      width: '15rem',
      dataIndex: 'updatedAt',
      sorter: true,
    },
    {
      title: '状态',
      width: '6rem',
      align: 'center',
      dataIndex: 'isEnabled',
      render(isEnabled: boolean, record: AdminApiItem) {
        return (
          <Authorization
            name="AdminApiEdit"
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
      width: '10rem',
      fixed: 'right',
      render(text, record: AdminApiItem) {
        return (
          <Space>
            <Authorization name="AdminApiView">
              <FetchButton onClick={() => openDetailModal(record)}>详情</FetchButton>
            </Authorization>
            <Authorization name="AdminApiEdit">
              <FetchButton onClick={() => openEditModal(record)}>编辑</FetchButton>
            </Authorization>

            {/* 禁用的才能删除 */}
            <Authorization name="AdminApiDelete">
              {!record.isEnabled ? (
                <Popconfirm
                  title="确定要删除该接口资源吗？"
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

  // 获取接口资源列表
  function getRows(data?: ReqAdminApiList) {
    setLoading(true);
    adminAPIList(data)
      .then((res) => {
        const data: RespAdminApiListData = res.data;
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

  // 接口资源状态更新
  function updateEnabled(record: AdminApiItem) {
    const updateData: ReqAdminApiEnable = {
      id: record.id,
      enabled: !record.isEnabled,
    };
    adminAPIEnable(updateData).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        getRows({ ...form.getFieldsValue() });
      });
    });
  }

  // 接口资源添加
  function openAddModal() {
    setAddModalStatus(true);
  }

  // 接口资源详情
  function openDetailModal(record: AdminApiItem) {
    adminAPIDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setDetailModalStatus(true);
    });
  }

  // 接口资源编辑
  function openEditModal(record: AdminApiItem) {
    adminAPIDetail({ id: record.id }).then((res) => {
      setDetailData(res.data);
      setEditModalStatus(true);
    });
  }

  function noticeAddModal(data: NoticeModalPropsType) {
    setAddModalStatus(false);
    if (data.reload) {
      getRows({ ...form.getFieldsValue() });
    }
  }

  function noticeDetailModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setDetailModalStatus(false);
    if (data.reload) {
      getRows({ ...form.getFieldsValue() });
    }
  }

  function noticeEditModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setEditModalStatus(false);
    if (data.reload) {
      getRows({ ...form.getFieldsValue() });
    }
  }

  // 删除接口资源
  function onDelete(record: AdminApiItem) {
    adminAPIDelete({ id: record.id }).then((res) => {
      message.success(res.msg, MessageDuritain);
      getRows({ ...form.getFieldsValue() });
    });
  }
  // 列表搜索
  function onSearchFinish(values: ReqAdminApiList) {
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

  useEffect(() => {
    onSearchReset();
  }, []);

  return (
    <Container>
      <Search>
        <Form form={form} onFinish={onSearchFinish}>
          <Row gutter={FormSearchRowGutter}>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="名称" name="name" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="路由" name="path" initialValue={''}>
                <Input />
              </Form.Item>
            </Col>
            <Col span={FormSearchRowColSpan}>
              <Form.Item label="键名" name="key" initialValue={''}>
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
          <Authorization name="AdminApiEdit">
            <Button type="primary" onClick={openAddModal}>
              <PlusOutlined />
              新建接口
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

      <AdminAPIAddModal modalStatus={addModalStatus} noticeModal={noticeAddModal} />

      {detailData ? (
        <>
          <AdminAPIDetailModal
            modalStatus={detailModalStatus}
            detailData={detailData}
            noticeModal={noticeDetailModal}
          />
          <AdminAPIEditModal
            modalStatus={editModalStatus}
            detailData={detailData}
            noticeModal={noticeEditModal}
          />
        </>
      ) : (
        <></>
      )}
    </Container>
  );
};

export default Admin;
