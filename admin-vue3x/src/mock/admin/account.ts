import { ok, fail } from '../common';
import type { MockEntry } from '../plugin';
import { currentAdminUserDetail } from '../data';

let userDetail = { ...currentAdminUserDetail };

const entries: MockEntry[] = [
  {
    url: '/api/admin/account/login',
    method: 'POST',
    response: ({ body }) => {
      const { username, password } = body;
      if (!username || !password) {
        return fail('请输入账号或密码');
      }
      if (username !== 'admin') {
        return fail('账号不存在');
      }
      if (password !== '123456') {
        return fail('密码错误');
      }
      return ok(userDetail, '登录成功');
    },
  },
  {
    url: '/api/admin/account/captcha',
    method: 'POST',
    response: () => {
      const code = '8888';
      const svg =
        `<svg xmlns="http://www.w3.org/2000/svg" width="120" height="44">` +
        `<rect width="120" height="44" fill="#eef2f7"/>` +
        `<text x="60" y="31" font-size="24" font-family="monospace" text-anchor="middle" fill="#3b5bdb">${code}</text>` +
        `</svg>`;
      return ok({ captchaId: 'mock-captcha-id', image: `data:image/svg+xml;base64,${btoa(svg)}` });
    },
  },
  {
    url: '/api/admin/account/logout',
    method: 'POST',
    response: () => ok(null, '退出成功'),
  },
  {
    url: '/api/admin/account/info',
    method: 'POST',
    response: () => ok(userDetail),
  },
  {
    url: '/api/admin/account/edit',
    method: 'POST',
    response: ({ body }) => {
      userDetail = { ...userDetail, ...body };
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/account/password',
    method: 'POST',
    response: ({ body }) => {
      if (body.oldPassword !== '123456') {
        return fail('原密码错误');
      }
      if (body.newPassword !== body.confirmPassword) {
        return fail('两次密码不一致');
      }
      return ok(null, '修改成功');
    },
  },
  {
    url: '/api/admin/common/upload',
    method: 'POST',
    response: () => ok({ url: userDetail.avatar }, '上传成功'),
  },
];

export default entries;
