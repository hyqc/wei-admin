import type { RequestLoginParamsType } from '@/services/apis/admin/account';
import { login } from '@/services/apis/admin/account';
import { SetLoginToken } from '@/utils/common';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Card, Checkbox, Form, Input } from 'antd';
import { parse } from 'query-string';
import React, { useEffect, useState } from 'react';
import { history, SelectLang, useModel } from 'umi';
import { AdminUserFormRules } from '../Admin/User/common';
import styles from './index.less';

const Login: React.FC = () => {
  const [form] = Form.useForm();
  const { setInitialState } = useModel('@@initialState');
  const [loginBtnLoading, setLoginBtnLoading] = useState<boolean>(false);
  const rules: any = AdminUserFormRules(form);

  const handleSubmit = async (values: RequestLoginParamsType) => {
    try {
      setLoginBtnLoading(true);
      const res = await login(values);
      console.log('登录返回：', res, history);
      // 设置token
      SetLoginToken(res.data.token, res.data.expire, values.remember || false);
      // 设置菜单
      await setInitialState((s: any) => ({
        ...s,
        currentUser: res.data,
      }));
      /** 此方法会跳转到 redirect 参数所在的位置 */
      if (!history) return;
      const query = parse(history.location.search);
      const { redirect } = query as { redirect: string };
      const redirectURL =  redirect || '/home';
      console.log('===', redirectURL)
      window.location.href = redirectURL
      return;
    } catch (error) {
      console.log(error);
    } finally {
      setLoginBtnLoading(false);
    }
  };

  useEffect(() => {
    form.resetFields();
  }, []);

  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.login}>
        <Card>
          <Form form={form} name="login" className={styles.loginForm} onFinish={handleSubmit}>
            <Form.Item>
              <p className={styles.title}>
                <img className={styles.logo} src="/favicon.ico" />
                <span>管理后台</span>
              </p>
            </Form.Item>
            <Form.Item name="username" rules={rules.username} validateFirst>
              <Input
                type="text"
                className={styles.loginInput}
                prefix={<UserOutlined />}
                placeholder="账号"
              />
            </Form.Item>
            <Form.Item name="password" rules={rules.password} validateFirst>
              <Input.Password
                className={styles.loginInput}
                prefix={<LockOutlined />}
                placeholder="密码"
              />
            </Form.Item>
            <Form.Item>
              <Form.Item name="remember" valuePropName="checked" noStyle>
                <Checkbox>自动登录</Checkbox>
              </Form.Item>
            </Form.Item>

            <Form.Item>
              <Button
                type="primary"
                loading={loginBtnLoading}
                htmlType="submit"
                className={styles.loginButton}
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </Card>
      </div>
    </div>
  );
};

export default Login;
