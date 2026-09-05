import { useEffect, useState } from 'react';
import { Button, Form, Input, Modal, Upload, message } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import { uploadAdminFile } from '@/api/admin/upload';

interface Props {
  open: boolean;
  onClose: () => void;
  onNotice: () => void;
}

/** 上传文件弹窗：分组（上传路径前缀）必填，选择文件后统一提交 */
export default function UploadFileModal({ open, onClose, onNotice }: Props) {
  const [group, setGroup] = useState('');
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (open) {
      setGroup('');
      setFile(null);
    }
  }, [open]);

  const beforeUpload = (uploadFile: File) => {
    // 不自动上传：先校验分组，再由弹窗统一提交
    setFile(uploadFile);
    return false;
  };

  const onOk = async () => {
    const groupVal = group.trim();
    if (!groupVal) {
      message.warning('请填写分组（上传路径前缀）');
      return;
    }
    if (!file) {
      message.warning('请选择要上传的文件');
      return;
    }
    setLoading(true);
    try {
      const res = await uploadAdminFile({ file, uploadGroup: groupVal });
      message.success(res.msg, 2);
      onNotice();
      onClose();
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal open={open} title="上传文件" width={560} confirmLoading={loading} maskClosable={false} okText="开始上传" cancelText="取消" onOk={onOk} onCancel={onClose}>
      <Form labelCol={{ span: 5 }} wrapperCol={{ span: 18 }}>
        <Form.Item label="分组" required>
          <Input value={group} placeholder="上传路径前缀，如 /admin/user/" allowClear onChange={(e) => setGroup(e.target.value)} />
          <div style={{ color: 'rgba(0,0,0,0.45)', fontSize: 12, lineHeight: 1.5 }}>
            分组即上传路径前缀，通常与所属菜单路径保持一致；存储目录结构为「分组/年/月/文件」
          </div>
        </Form.Item>
        <Form.Item label="文件" required>
          <Upload maxCount={1} beforeUpload={beforeUpload} fileList={file ? [{ uid: '1', name: file.name, status: 'done' }] : []} onRemove={() => setFile(null)}>
            <Button icon={<UploadOutlined />}>选择文件</Button>
          </Upload>
        </Form.Item>
      </Form>
    </Modal>
  );
}
