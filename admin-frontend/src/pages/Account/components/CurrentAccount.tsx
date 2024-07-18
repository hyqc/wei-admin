import { AdminUserFormRules } from '@/pages/Admin/User/common';
import {
  currentAdminEdit,
  RequestCurrentAdminEditParamsType,
  upload,
} from '@/services/apis/admin/account';
import { APICommon } from '@/services/apis/admin/api';
import { GetLoginToken } from '@/utils/common';
import { CloudUploadOutlined } from '@ant-design/icons';
import { Button, Form, Input, message, Upload, UploadFile } from 'antd';
import { RcFile } from 'antd/lib/upload';
import { useEffect, useState } from 'react';
import { useModel } from 'umi';

const CurrentAccountInfo: React.FC = () => {
  const [form] = Form.useForm();
  const { initialState } = useModel('@@initialState');
  const detailData = { ...initialState?.currentUser };

  const [fileList, setFielList] = useState<any[]>();
  const [avatar, setAvatar] = useState<string>(detailData.avatar || '');
  const tokenInfo = GetLoginToken();
  const token = tokenInfo !== undefined ? tokenInfo.token : '';
  const headers = {
    Authorization: 'Bearer ' + token,
  };
  const styleButton = {
    marginLeft: '5.6rem',
  };

  const rules: any = AdminUserFormRules(form);

  function handleOk() {
    form
      .validateFields()
      .then((values) => {
        let data: RequestCurrentAdminEditParamsType;
        data = {
          nickname: values.nickname,
          email: values.email,
          avatar,
        };
        console.log(data);
        currentAdminEdit(data).then((res) => {
          message.success(res.message, MessageDuritain, () => {});
        });
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {});
  }

  function uploadOnChange(data: any) {
    if (data && data.fileList && data.fileList[0]?.response) {
      setAvatar(data.fileList[0]?.response?.data?.url || '');
    }
    setFielList(data.fileList || []);
  }

  async function onPreview(file: UploadFile) {
    let src = file.url as string;
    if (!src) {
      src = await new Promise((resolve) => {
        const reader = new FileReader();
        reader.readAsDataURL(file.originFileObj as RcFile);
        reader.onload = () => resolve(reader.result as string);
      });
    }
    const image = new Image();
    image.src = src;
    const imgWindow = window.open(src);
    imgWindow?.document.write(image.outerHTML);
  }

  function uploadFile(e: any) {
    upload({
      file: e.file,
      fileType: 1,
    })
      .then(() => {
        e.onSuccess();
      })
      .catch(() => {
        e.onError();
      });
  }

  useEffect(() => {
    form.setFieldsValue(detailData);
    if (detailData && detailData.avatar) {
      setFielList([
        {
          adminId: -1,
          nickname: detailData.nickname,
          url: detailData.avatar,
          enabled: detailData.enabled,
        },
      ]);
    }
  }, []);

  return (
    <div>
      <Form form={form} key="center" labelCol={{ span: 2 }} wrapperCol={{ span: 10 }}>
        <Form.Item label="名称" name="username" rules={rules.username}>
          <Input disabled />
        </Form.Item>
        <Form.Item label="昵称" name="nickname" rules={rules.nickname}>
          <Input />
        </Form.Item>
        <Form.Item label="邮箱" name="email" rules={rules.email}>
          <Input />
        </Form.Item>
        <Form.Item label="头像" name="avatar">
          <Upload
            maxCount={1}
            // customRequest={uploadFile}
            headers={headers}
            data={{
              fileType: 1,
            }}
            accept={UploadImageAccept}
            action={`${BaseAPI}${APICommon.upload.url}`}
            fileList={fileList}
            listType="picture-card"
            onPreview={onPreview}
            onChange={uploadOnChange}
          >
            <CloudUploadOutlined />
            上传头像
          </Upload>
        </Form.Item>
        <Form.Item>
          <Button style={styleButton} htmlType="submit" type="primary" onClick={handleOk}>
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default CurrentAccountInfo;
