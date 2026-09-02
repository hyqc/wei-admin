import { useEffect, useRef, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { Button, Card, Checkbox, Form, Input, Row, Col } from 'antd';
import { LockOutlined, SafetyOutlined, UserOutlined } from '@ant-design/icons';
import { useUserStore } from '@/store/user';
import { getCaptcha } from '@/api/admin/account';
import { AdminUsername, AdminUserPassword } from '@/api/pattern';
import { HomePath } from '@/api/config';
import type { ReqLogin } from '@/types/admin_account';

/** 登录页 */
export default function Login() {
  const navigate = useNavigate();
  const location = useLocation();
  const loginAsync = useUserStore((s) => s.loginAsync);
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [captcha, setCaptcha] = useState<{ captchaId: string; image: string }>({ captchaId: '', image: '' });
  const [remember, setRemember] = useState(true);
  const captchaIdRef = useRef('');

  /** 刷新验证码 */
  async function refreshCaptcha() {
    try {
      const res = await getCaptcha();
      captchaIdRef.current = res.data.captchaId;
      setCaptcha({ captchaId: res.data.captchaId, image: res.data.image });
      form?.setFieldValue('captchaCode', '');
    } catch {
      /* 拦截器已提示 */
    }
  }

  useEffect(() => {
    refreshCaptcha();
  }, []);

  const onFinish = async (values: ReqLogin) => {
    setLoading(true);
    try {
      await loginAsync({
        username: values.username,
        password: values.password,
        remember,
        captchaId: captchaIdRef.current,
        captchaCode: values.captchaCode,
      });
      const redirect = new URLSearchParams(location.search).get('redirect') || HomePath;
      navigate(redirect, { replace: true });
    } catch {
      // 登录失败（验证码失效/账号密码错误）后刷新验证码
      await refreshCaptcha();
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-page">
      <Card className="login-card">
        <div className="login-title">
          <img
            src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg"
            alt="logo"
            className="login-logo"
          />
          <span>Admin React</span>
        </div>
        <div className="login-subtitle">ant design 后台管理系统</div>
        <Form form={form} size="large" initialValues={{ username: 'admin', password: '123456' }} onFinish={onFinish}>
          <Form.Item
            name="username"
            rules={[
              { required: true, message: '请输入用户名' },
              { pattern: AdminUsername, message: '用户名格式不正确' },
            ]}
          >
            <Input prefix={<UserOutlined />} placeholder="用户名：admin" allowClear />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[
              { required: true, message: '请输入密码' },
              { pattern: AdminUserPassword, message: '密码格式不正确' },
            ]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="密码：123456" />
          </Form.Item>
          <Form.Item name="captchaCode" rules={[{ required: true, message: '请输入验证码' }]}>
            <Row gutter={8}>
              <Col span={15}>
                <Input prefix={<SafetyOutlined />} placeholder="请输入验证码" allowClear maxLength={8} />
              </Col>
              <Col span={9}>
                {captcha.image ? (
                  <img
                    src={captcha.image}
                    className="captcha-img"
                    alt="验证码"
                    title="看不清？点击刷新"
                    onClick={refreshCaptcha}
                  />
                ) : (
                  <Button block onClick={refreshCaptcha}>
                    获取验证码
                  </Button>
                )}
              </Col>
            </Row>
          </Form.Item>
          <Form.Item>
            <Checkbox checked={remember} onChange={(e) => setRemember(e.target.checked)}>
              自动登录
            </Checkbox>
          </Form.Item>
          <Form.Item style={{ marginBottom: 0 }}>
            <Button type="primary" htmlType="submit" block loading={loading}>
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
}
