import {
  ResponseAdminUserDetailType,
  ResponseAdminUserListItemType,
} from '@/services/apis/admin/user';
import { ResponseDetailType, ResponseListType } from '@/services/apis/types';
import { success } from '../common';
import { Request, Response } from 'express';

const userListDataRows: ResponseAdminUserListItemType[] = [
  {
    adminId: 5,
    username: 'test00001',
    nickname: '测试00001',
    email: '',
    avatar: '',
    roles: [],
    enabled: false,
    loginTotal: 0,
    lastLoginIp: '',
    createTime: '2022-08-07 21:09:23',
    modifyTime: '2022-08-07 21:09:23',
  },
  {
    adminId: 7,
    username: 'test00002',
    nickname: '测试00002',
    email: '',
    avatar: '',
    roles: [],
    enabled: true,
    loginTotal: 0,
    lastLoginIp: '',
    createTime: '2022-08-07 21:11:05',
    modifyTime: '2022-08-07 21:11:05',
  },
  {
    adminId: 8,
    username: 'test00003',
    nickname: '测试00003',
    email: '',
    avatar: '',
    roles: [],
    enabled: true,
    loginTotal: 0,
    lastLoginIp: '',
    createTime: '2022-08-07 21:13:59',
    modifyTime: '2022-08-07 21:13:59',
  },
  {
    adminId: 9,
    username: 'test00004',
    nickname: '测试00004',
    email: 'test0004@qq.com',
    avatar: '',
    roleIds: '1',
    roles: [
      {
        adminId: 9,
        roleId: 1,
        roleName: '管理员',
      },
    ],
    enabled: true,
    loginTotal: 19,
    lastLoginIp: '127.0.0.1,127.0.0.1',
    createTime: '2022-08-11 02:17:28',
    modifyTime: '2022-08-11 02:17:28',
  },
  {
    adminId: 1,
    username: 'admin',
    nickname: '超管',
    email: 'ddd@q1.com',
    avatar: '',
    roles: [],
    enabled: true,
    loginTotal: 23,
    lastLoginIp: '171.113.167.161,43.132.157.77',
    createTime: '2022-08-11 23:29:20',
    modifyTime: '2022-08-11 23:29:20',
  },
];

const list: ResponseListType = {
  code: 0,
  type: 'success',
  message: '成功',
  data: {
    total: 105,
    pageNum: 1,
    pageSize: 10,
    list: userListDataRows,
  },
};

const userDetail: ResponseAdminUserDetailType = {
  adminId: 9,
  username: 'test00004',
  nickname: '测试00004',
  email: 'test0004@qq.com',
  avatar: '',
  roles: [
    {
      adminId: 9,
      roleId: 1,
      roleName: '管理员',
    },
  ],
  enabled: true,
  lastLoginIp: '',
  lastLoginTime: '2022-08-11 02:17:29',
  createTime: '2022-08-11 02:17:28',
  modifyTime: '2022-08-11 02:17:28',
};

const userDetailResponse: ResponseDetailType = {
  code: 0,
  message: '成功',
  type: 'SUCCESS',
  data: userDetail,
};

export default {
  'POST /api/admin/user/list': (req: Request, res: Response) => {
    list.data.pageNum = req.body.pageNum;
    list.data.pageSize = req.body.pageSize;
    res.json(list);
  },
  'POST /api/admin/user/detail': userDetailResponse,
  'POST /api/admin/user/add': success,
  'POST /api/admin/user/edit': success,
  'POST /api/admin/user/delete': success,
  'POST /api/admin/user/enable': success,
  'POST /api/admin/user/bind': success,
};
