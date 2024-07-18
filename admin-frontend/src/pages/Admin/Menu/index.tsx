import React, { useEffect, useState } from 'react';
import { Container, Content } from '@/components/PageListContainer';
import { Form, Button, Space, Table, message, Popconfirm, Switch } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { ResponseBodyType } from '@/services/apis/types';
import {
  adminMenuDelete,
  adminMenuDetail,
  adminMenuEdit,
  adminMenuEnable,
  adminMenuPermissions,
  adminMenuTree,
  RequestAdminMenuEditParamsType,
  RequestAdminMenuTreeParamsType,
  ResponseAdminMenuListItemType,
  ResponseAdminMenuPermissionsType,
} from '@/services/apis/admin/menu';
import { ColumnsType } from 'antd/lib/table';
import Authorization from '@/components/Autuorization';
import AdminMenuDetailModal, { NoticeModalPropsType } from './detail';
import { history } from 'umi';
import SavePermissionsModal from './components/PermissionsSave';
import FetchButton from '@/components/FetchButton';

const Admin: React.FC = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [detailData, setDetailData] = useState<any>();
  const [rowsData, setRowsData] = useState<ResponseAdminMenuListItemType[]>([]);
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [savePermissionsModalStatus, setSaveMenuPermissionsModalStatus] = useState<boolean>(false);
  const [menuPermissionsDetail, setMenuPermissionsDetail] = useState<
    ResponseAdminMenuPermissionsType | undefined
  >(undefined);

  const columns: ColumnsType<any> = [
    {
      title: '名称',
      align: 'left',
      dataIndex: 'name',
    },
    {
      title: '路由',
      align: 'left',
      dataIndex: 'path',
    },
    {
      title: '菜单',
      align: 'left',
      dataIndex: 'hideInMenu',
      render(enabled: boolean, record: ResponseAdminMenuListItemType) {
        if (record.path == '/') {
          return <></>;
        }
        return (
          <Authorization
            name="AdminMenuEdit"
            forbidden={
              <>
                <Switch
                  disabled
                  checkedChildren={'隐藏'}
                  unCheckedChildren={'显示'}
                  checked={enabled}
                />
              </>
            }
          >
            <Popconfirm
              title={`确定在菜单中${record.hideInMenu ? '显示' : '隐藏'}该项吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateMenuStatus(record, 'hideInMenu')}
            >
              <Switch checkedChildren={'隐藏'} unCheckedChildren={'显示'} checked={enabled} />
            </Popconfirm>
          </Authorization>
        );
      },
    },
    {
      title: '子菜单',
      align: 'left',
      dataIndex: 'hideChildrenInMenu',
      render(enabled: boolean, record: ResponseAdminMenuListItemType) {
        if (record.path == '/') {
          return <></>;
        }
        return (
          <Authorization
            name="AdminMenuEdit"
            forbidden={
              <>
                <Switch
                  disabled
                  checkedChildren={'隐藏'}
                  unCheckedChildren={'显示'}
                  checked={enabled}
                />
              </>
            }
          >
            <Popconfirm
              title={`确定在菜单中${record.hideInMenu ? '显示' : '隐藏'}其子菜单项吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateMenuStatus(record, 'hideChildrenInMenu')}
            >
              <Switch checkedChildren={'隐藏'} unCheckedChildren={'显示'} checked={enabled} />
            </Popconfirm>
          </Authorization>
        );
      },
    },
    {
      title: '状态',
      width: '6rem',
      align: 'center',
      dataIndex: 'enabled',
      render(enabled: boolean, record: ResponseAdminMenuListItemType) {
        if (record.path == '/') {
          return <></>;
        }
        return (
          <Authorization
            name="AdminMenuEdit"
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
              title={`确定要${record.enabled ? '禁用' : '启用'}该菜单吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateMenuEnabled(record)}
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
      width: '10rem',
      render(text, record: ResponseAdminMenuListItemType) {
        if (record.path == '/') {
          return <></>;
        }
        return (
          <Space>
            <Authorization name="AdminMenuView">
              <FetchButton onClick={() => openDetailModal(record)}>详情</FetchButton>
            </Authorization>
            <Authorization name="AdminMenuEdit">
              <FetchButton onClick={() => openAddModal(record)}>添加子菜单</FetchButton>
            </Authorization>
            <Authorization name="AdminMenuEdit">
              {record.hideInMenu ||
              (record.children !== undefined &&
                record?.children.length > 0 &&
                !record.hideChildrenInMenu) ? (
                ''
              ) : (
                <FetchButton onClick={() => openSavePermissionsModal(record)}>创建权限</FetchButton>
              )}
            </Authorization>
            <Authorization name="AdminMenuEdit">
              <FetchButton onClick={() => openEditModal(record)}>编辑</FetchButton>
            </Authorization>

            {/* 禁用的才能删除 */}
            <Authorization name="AdminMenuDelete">
              {!record.enabled ? (
                <Popconfirm
                  title="确定要删除该菜单吗？"
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

  // 获取菜单列表
  function getRows(data?: RequestAdminMenuTreeParamsType) {
    setLoading(true);
    adminMenuTree(data)
      .then((res: ResponseBodyType) => {
        const resData: ResponseAdminMenuListItemType[] = res.data;
        setRowsData(resData);
      })
      .catch((err) => {
        console.log('error', err);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  function tableChange(pagination: any, filters: any, sorter: any) {
    getRows({
      ...form.getFieldsValue(),
    });
  }

  // 菜单状态更新
  function updateMenuStatus(record: ResponseAdminMenuListItemType, field: string) {
    const updateData: RequestAdminMenuEditParamsType = {
      id: record.id,
    };
    updateData[field] = !record[field];
    adminMenuEdit(updateData).then((res) => {
      message.success(res.message, MessageDuritain, () => {
        getRows({ ...form.getFieldsValue() });
      });
    });
  }

  function updateMenuEnabled(record: ResponseAdminMenuListItemType) {
    adminMenuEnable({
      menuId: record.id,
      enabled: !record.enabled,
    }).then((res) => {
      message.success(res.message, MessageDuritain, () => {
        getRows({ ...form.getFieldsValue() });
      });
    });
  }

  // 菜单详情
  function openDetailModal(record: ResponseAdminMenuListItemType) {
    adminMenuDetail({ menuId: record.id }).then((res) => {
      setDetailData(res.data);
      setDetailModalStatus(true);
    });
  }

  // 菜单编辑
  function openEditModal(record: ResponseAdminMenuListItemType) {
    adminMenuDetail({ menuId: record.id }).then((res) => {
      history.push(`/admin/menu/edit?menuId=${record.id}`);
    });
  }

  // 菜单添加
  function openAddModal(record?: ResponseAdminMenuListItemType) {
    if (record?.id) {
      history.push(`/admin/menu/add?menuId=${record.id}`);
      return;
    }
    history.push('/admin/menu/add');
  }

  /**
   * 保存菜单权限
   * @param record
   */
  function openSavePermissionsModal(record: ResponseAdminMenuListItemType) {
    adminMenuPermissions({ menuId: record.id }).then((res) => {
      setMenuPermissionsDetail(res.data || { menuId: record.id });
      setSaveMenuPermissionsModalStatus(true);
    });
  }

  function noticeDetailModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setDetailModalStatus(false);
    if (data.reload) {
      getRows({ ...form.getFieldsValue() });
    }
  }
  function noticeAddPermissionModal(data: NoticeModalPropsType) {
    setDetailData(undefined);
    setMenuPermissionsDetail(undefined);
    setSaveMenuPermissionsModalStatus(false);
    if (data.reload) {
      getRows({ ...form.getFieldsValue() });
    }
  }

  // 删除菜单
  function onDelete(record: ResponseAdminMenuListItemType) {
    adminMenuDelete({ menuId: record.id, enabled: record.enabled }).then((res) => {
      message.success(res.message, MessageDuritain);
      getRows({ ...form.getFieldsValue() });
    });
  }

  useEffect(() => {
    getRows();
  }, []);

  return (
    <Container>
      <Content>
        {/* button */}
        <Space style={{ marginBottom: '1rem' }}>
          <Authorization name="AdminMenuEdit">
            <Button type="primary" onClick={() => openAddModal()}>
              <PlusOutlined />
              新建菜单
            </Button>
          </Authorization>
        </Space>

        {/* table */}
        {rowsData.length ? (
          <Table
            sticky
            key={'id'}
            rowKey="id"
            scroll={{ x: 'auto' }}
            expandable={{
              defaultExpandAllRows: true,
              fixed: 'left',
              indentSize: 24,
            }}
            loading={loading}
            columns={columns}
            pagination={false}
            dataSource={rowsData}
            onChange={tableChange}
          />
        ) : (
          <></>
        )}
      </Content>

      {/* modal */}

      <AdminMenuDetailModal
        modalStatus={detailModalStatus}
        detailData={detailData}
        noticeModal={noticeDetailModal}
      />

      {menuPermissionsDetail && menuPermissionsDetail.menu && menuPermissionsDetail.menu?.id > 0 ? (
        <SavePermissionsModal
          modalStatus={savePermissionsModalStatus}
          detailData={menuPermissionsDetail}
          noticeModal={noticeAddPermissionModal}
        />
      ) : (
        ''
      )}
    </Container>
  );
};

export default Admin;
