import React, { useEffect, useState } from 'react';
import { Container, Content } from '@/components/PageContainer';
import { Form, Button, Space, Table, message, Popconfirm, Switch } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { ResponseBodyType } from '@/services/apis/types';
import {
  adminMenuDelete,
  adminMenuDetail,
  adminMenuEnable,
  adminMenuPermissions,
  adminMenuShow,
  adminMenuTree,
} from '@/services/apis/admin/menu';
import { ColumnsType } from 'antd/lib/table';
import Authorization from '@/components/Autuorization';
import AdminMenuDetailModal, { NoticeModalPropsType } from './detail';
import AdminMenuSaveModal from './save';
import SavePermissionsModal from './components/PermissionsSave';
import FetchButton from '@/components/FetchButton';
import { RowEnabledButton } from '@/components';
import { MenuTreeItem, ReqAdminMenuShow, RespAdminMenuPermissionsData } from '@/proto/admin_ts/admin_menu';
import { AdminMenuModel } from '@/proto/admin_ts/admin_model';

const Admin: React.FC = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState<boolean>(false);
  const [detailData, setDetailData] = useState<AdminMenuModel>();
  const [rowsData, setRowsData] = useState<MenuTreeItem[]>([]);
  const [detailModalStatus, setDetailModalStatus] = useState<boolean>(false);
  const [saveModalStatus, setSaveModalStatus] = useState<boolean>(false);
  const [savePermissionsModalStatus, setSaveMenuPermissionsModalStatus] = useState<boolean>(false);
  const [menuPermissionsDetail, setMenuPermissionsDetail] = useState<RespAdminMenuPermissionsData>({menuInfo: {}, permissions: []});


  // 获取菜单列表
  function getRows(data?: object) {
    console.log('data:', data)
    setLoading(true);
    adminMenuTree()
      .then((res: ResponseBodyType) => {
        const resData: MenuTreeItem[] = res.data;
        setRowsData(resData);
      })
      .catch((err) => {
        console.log('error', err);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  function tableChange() {
    getRows();
  }

  // 菜单状态更新
  function updateMenuStatus(record: MenuTreeItem, field: string) {
    const updateData: ReqAdminMenuShow = {
      menuId: record.id,
      field: field,
      show: !record[field]
    };

    adminMenuShow(updateData).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        getRows();
      });
    });
  }

  function updateMenuEnabled(record: MenuTreeItem) {
    adminMenuEnable({
      menuId: record.id,
      enabled: !record.enabled,
    }).then((res) => {
      message.success(res.msg, MessageDuritain, () => {
        getRows();
      });
    });
  }

  // 菜单详情
  function openDetailModal(record: MenuTreeItem) {
    adminMenuDetail({ menuId: record.id }).then((res) => {
      setDetailData(res.data.data);
      setDetailModalStatus(true);
    });
  }

  // 菜单编辑
  function openEditModal(record: MenuTreeItem) {
    adminMenuDetail({ menuId: record.id }).then((res) => {
      setDetailData(res.data.data)
      setSaveModalStatus(true)
    });
  }

  // 菜单添加
  function openAddModal(record?: MenuTreeItem) {
    setDetailData({
      parentId: record?.id ?? 0
    })
    setSaveModalStatus(true)
  }

  /**
   * 保存菜单权限
   * @param record
   */
  function openSavePermissionsModal(record: MenuTreeItem) {
    adminMenuPermissions({ menuId: record.id }).then((res) => {
      const data = res.data || {menuInfo: { }, permissions: []};
      setMenuPermissionsDetail((old)=>{
        const nv = {menuInfo: {...data.menuInfo}, permissions: [...data.permissions]}
        return nv
      });
      setSaveMenuPermissionsModalStatus(true);
    });
  }

  function noticeDetailModal(data: NoticeModalPropsType) {
    setDetailData({});
    setDetailModalStatus(false);
    if (data.reload) {
      getRows();
    }
  }

  function noticeSaveModal(data: NoticeModalPropsType) {
    setDetailData({});
    setSaveModalStatus(false);
    if (data.reload) {
      getRows();
    }
  }

  function noticeAddPermissionModal(data: NoticeModalPropsType) {
    setDetailData({});
    setMenuPermissionsDetail({menuInfo: {}, permissions: []});
    setSaveMenuPermissionsModalStatus(false);
    if (data.reload) {
      getRows();
    }
  }

  // 删除菜单
  function onDelete(record: MenuTreeItem) {
    adminMenuDelete({ menuId: record.id }).then((res) => {
      message.success(res.msg, MessageDuritain);
      getRows();
    });
  }

  useEffect(() => {
    getRows();
  }, []);

  useEffect(()=>{},[menuPermissionsDetail])


  const showTrueText = '隐藏'
  const showFalseText = '显示'

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
      render(hideInMenu: boolean, record: MenuTreeItem) {
        if (record.path === '/') {
          return <></>;
        }
        return (
          <Authorization
            name="AdminUserEdit"
            forbidden={
              <RowEnabledButton danger={!!hideInMenu} isEnabled={hideInMenu} trueText={showTrueText} falseText={showFalseText} />
            }
          >
            <Popconfirm
              title={`确定在菜单中${record.hideInMenu ? '隐藏' : '显示'}该项吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateMenuStatus(record, 'hideInMenu')}
            >
              <RowEnabledButton danger={!!hideInMenu} isEnabled={hideInMenu} trueText={showTrueText} falseText={showFalseText} />
            </Popconfirm>
          </Authorization>
        );
      },
    },
    {
      title: '子菜单',
      align: 'left',
      dataIndex: 'hideChildrenInMenu',
      render(hideChildrenInMenu: boolean, record: MenuTreeItem) {
        if (record.path === '/') {
          return <></>;
        }
        return (
          <Authorization
            name="AdminUserEdit"
            forbidden={
              <RowEnabledButton danger={!!hideChildrenInMenu} isEnabled={hideChildrenInMenu} trueText={showTrueText} falseText={showFalseText} />
            }
          >
            <Popconfirm
              title={`确定在菜单中${record.hideChildrenInMenu ? '隐藏' : '显示'}该项吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateMenuStatus(record, 'hideChildrenInMenu')}
            >
              <RowEnabledButton danger={!!hideChildrenInMenu} isEnabled={hideChildrenInMenu} trueText={showTrueText} falseText={showFalseText} />
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
      render(enabled: boolean, record: MenuTreeItem) {
        if (record.path === '/') {
          return <></>;
        }
        return (
          <Authorization
            name="AdminUserEdit"
            forbidden={
              <RowEnabledButton isEnabled={enabled} />
            }
          >
            <Popconfirm
              title={`确定${record.enabled ? '禁用' : '启用'}该项吗？`}
              okText="确定"
              cancelText="取消"
              onConfirm={() => updateMenuEnabled(record)}
            >
              <RowEnabledButton isEnabled={enabled} />
            </Popconfirm>
          </Authorization>
        );
      },
    },
    {
      title: '操作',
      align: 'left',
      width: '10rem',
      render(text, record: MenuTreeItem) {
        if (record.path === '/') {
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
                <FetchButton onClick={() => openSavePermissionsModal(record)}>权限</FetchButton>
              )}
            </Authorization>
            <Authorization name="AdminMenuEdit">
              <FetchButton onClick={() => openEditModal(record)}>编辑</FetchButton>
            </Authorization>

            {/* 禁用的才能删除 */}
            <Authorization name="AdminMenuDelete">
              {!record.enabled && (record.children === undefined || record.children.length === 0) ?  (
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

      <AdminMenuSaveModal
        modalStatus={saveModalStatus}
        detailData={detailData}
        noticeModal={noticeSaveModal}
      />

      <SavePermissionsModal
        modalStatus={savePermissionsModalStatus}
        menuInfo={menuPermissionsDetail.menuInfo || {}}
        permissions={menuPermissionsDetail.permissions || []}
        noticeModal={noticeAddPermissionModal}
      />
    </Container>
  );
};

export default Admin;
