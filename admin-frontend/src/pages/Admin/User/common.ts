import { AdminUsername, AdminUserPassword } from '@/services/common/pattern';
import { FormInstance } from 'antd';

export const DEFAULT_PAGE_NO = 1;
export const DEFAULT_PAGE_SIZE = 10;

export const AdminUserFormRules = (form: FormInstance<any>): { [key: string]: any[] } => {
  return {
    username: [
      { required: true, type: 'string', message: '账号不能为空!' },
      {
        required: true,
        pattern: AdminUsername,
        message: '账号由任意大小写字母，数字组成，且长度为1-50个字符',
      },
    ],
    nickname: [
      { required: true, type: 'string', message: '昵称不能为空!' },
      { max: 50, type: 'string', message: '昵称不能超过50个字符!' },
      { min: 1, type: 'string', message: '昵称最少一个字符!' },
    ],
    email: [{ type: 'email', message: '邮箱格式错误!' }],
    oldPassword: [
      { required: true, type: 'string', message: '密码不能为空' },
      {
        required: true,
        pattern: AdminUserPassword,
        message: '密码长度6-30位',
      },
    ],
    password: [
      { required: true, type: 'string', message: '密码不能为空' },
      {
        required: true,
        pattern: AdminUserPassword,
        message: '密码长度6-30位',
      },
      {
        required: true,
        validator: (rule: any, value: string) => {
          const oldPassword: string = form.getFieldValue('oldPassword');
          if (oldPassword !== undefined && oldPassword === value) {
            return Promise.reject(new Error('新旧密码不能相同'));
          }
          return Promise.resolve();
        },
      },
    ],
    confirmPassword: [
      { required: true, type: 'string', message: '密码不能为空' },
      {
        required: true,
        pattern: AdminUserPassword,
        message: '密码长度6-30位',
      },
      {
        required: true,
        validator: (rule: any, value: string) => {
          if (form.getFieldValue('password') !== value) {
            return Promise.reject(new Error('两次输入的密码不一致'));
          }
          return Promise.resolve();
        },
      },
    ],
  };
};
