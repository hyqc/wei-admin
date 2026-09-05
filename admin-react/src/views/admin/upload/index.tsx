import { useEffect, useState } from 'react';
import { Button, Form, Image, Input, Popconfirm, Select, Space, Table, Tag, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { DeleteOutlined, LinkOutlined, ReloadOutlined, SearchOutlined, UploadOutlined } from '@ant-design/icons';
import PageContainer from '@/components/PageContainer';
import Authorization from '@/components/Authorization';
import UploadFileModal from './components/UploadFileModal';
import { deleteAdminUpload, getAdminUploadList } from '@/api/admin/upload';
import { UPLOAD_DRIVER_OPTIONS } from '@/types/admin_upload';
import type { UploadItem } from '@/types/admin_upload';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const IMAGE_EXTS = ['png', 'jpg', 'jpeg', 'gif', 'webp', 'bmp', 'svg'];
const isImage = (record: UploadItem) => !!record.ext && IMAGE_EXTS.includes(record.ext.toLowerCase());

/** 上传管理：所有上传记录列表，支持上传与删除 */
export default function AdminUpload() {
  const [loading, setLoading] = useState(false);
  const [rows, setRows] = useState<UploadItem[]>([]);
  const [pageInfo, setPageInfo] = useState<PageInfoType>({ ...DEFAULT_PAGE_INFO });
  const [uploadOpen, setUploadOpen] = useState(false);
  const [searchForm] = Form.useForm<{ uploadGroup?: string; originalName?: string; ext?: string; driver?: string }>();

  const getRows = (page = pageInfo, search = searchForm.getFieldsValue()) => {
    setLoading(true);
    getAdminUploadList({
      pageNum: page.pageNum,
      pageSize: page.pageSize,
      uploadGroup: search.uploadGroup,
      originalName: search.originalName,
      ext: search.ext,
      driver: search.driver,
    })
      .then((res) => {
        setRows(res.data.list || []);
        setPageInfo((prev) => ({ ...prev, total: res.data.pageInfo?.total || (res.data.list || []).length }));
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    getRows();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const onSearch = () => {
    const next = { ...pageInfo, pageNum: 1 };
    setPageInfo(next);
    getRows(next);
  };

  const onReset = () => {
    searchForm.setFieldsValue({ uploadGroup: undefined, originalName: undefined, ext: undefined, driver: undefined });
    onSearch();
  };

  const onDelete = (record: UploadItem) => {
    deleteAdminUpload({ id: record.id }).then((res) => {
      message.success(res.msg, 2);
      getRows();
    });
  };

  const copyUrl = async (record: UploadItem) => {
    if (!record.url) return;
    try {
      await navigator.clipboard.writeText(record.url);
      message.success('链接已复制', 2);
    } catch {
      message.warning('复制失败，请手动复制');
    }
  };

  const columns: ColumnsType<UploadItem> = [
    { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
    {
      title: '文件',
      dataIndex: 'file',
      key: 'file',
      width: 280,
      render: (_v, record) => (
        <Space>
          {isImage(record) && <Image src={record.url} width={36} height={36} style={{ objectFit: 'cover', borderRadius: 4 }} preview={false} fallback={record.url} />}
          <div style={{ lineHeight: 1.4 }}>
            <div style={{ maxWidth: 200, overflow: 'hidden', whiteSpace: 'nowrap', textOverflow: 'ellipsis' }} title={record.originalName}>
              {record.originalName}
            </div>
            <div style={{ color: 'rgba(0,0,0,0.45)', fontSize: 12 }}>{record.newName}</div>
          </div>
        </Space>
      ),
    },
    {
      title: '分组',
      dataIndex: 'uploadGroup',
      key: 'uploadGroup',
      width: 180,
      render: (_v, record) => <Tag color="blue">/{record.uploadGroup}/</Tag>,
    },
    { title: '大小', dataIndex: 'sizeText', key: 'sizeText', width: 100 },
    {
      title: '存储',
      dataIndex: 'driver',
      key: 'driver',
      width: 120,
      render: (_v, record) => <Tag>{record.driverText || record.driver}</Tag>,
    },
    {
      title: '访问地址',
      dataIndex: 'url',
      key: 'url',
      width: 90,
      render: (_v, record) => (
        <a href={record.url} target="_blank" rel="noopener noreferrer">
          打开
        </a>
      ),
    },
    { title: '上传者', dataIndex: 'adminName', key: 'adminName', width: 120 },
    { title: '上传日期', dataIndex: 'uploadDate', key: 'uploadDate', width: 120 },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
      fixed: 'right',
      width: 180,
      render: (_v, record) => (
        <Space>
          <Button type="link" size="small" icon={<LinkOutlined />} onClick={() => copyUrl(record)}>
            复制链接
          </Button>
          <Authorization permission="AdminUploadDelete">
            <Popconfirm title="删除后存储上的文件也会一并删除，确定要删除吗？" okText="确定" cancelText="取消" onConfirm={() => onDelete(record)}>
              <Button type="link" size="small" danger icon={<DeleteOutlined />}>
                删除
              </Button>
            </Popconfirm>
          </Authorization>
        </Space>
      ),
    },
  ];

  return (
    <PageContainer
      pageInfo={pageInfo}
      onPageChange={(pageNum) => {
        const next = { ...pageInfo, pageNum };
        setPageInfo(next);
        getRows(next);
      }}
      onPageSizeChange={(pageSize) => {
        const next = { ...pageInfo, pageSize, pageNum: 1 };
        setPageInfo(next);
        getRows(next);
      }}
      searchArea={
        <Form layout="inline" form={searchForm} onFinish={onSearch}>
          <Form.Item label="分组" name="uploadGroup">
            <Input placeholder="如 /admin/user" allowClear />
          </Form.Item>
          <Form.Item label="文件名" name="originalName">
            <Input placeholder="原始文件名" allowClear />
          </Form.Item>
          <Form.Item label="类型" name="ext">
            <Input placeholder="如 png" allowClear style={{ width: 100 }} />
          </Form.Item>
          <Form.Item label="存储" name="driver">
            <Select placeholder="全部" allowClear style={{ width: 140 }} options={UPLOAD_DRIVER_OPTIONS} />
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
        <Authorization permission="AdminUploadEdit">
          <Button type="primary" icon={<UploadOutlined />} onClick={() => setUploadOpen(true)}>
            上传文件
          </Button>
        </Authorization>
      }
    >
      <Table columns={columns} dataSource={rows} loading={loading} pagination={false} rowKey="id" scroll={{ x: 1400 }} />
      <UploadFileModal open={uploadOpen} onClose={() => setUploadOpen(false)} onNotice={() => getRows()} />
    </PageContainer>
  );
}
