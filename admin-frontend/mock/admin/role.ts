import {
  ResponseAdminRoleAllItemType,
  ResponseAdminRoleDetailType,
  ResponseAdminRoleListItemType,
} from '@/services/apis/admin/role';
import { ResponseDetailType, ResponseListType } from '@/services/apis/types';
import { success } from '../common';
import { Request, Response } from 'express';

const allRoles: ResponseAdminRoleAllItemType[] = [
  {
    roleId: 1,
    roleName: '管理员',
  },
  {
    roleId: 2,
    roleName: '运维',
  },
  {
    roleId: 3,
    roleName: '运维1',
  },
  {
    roleId: 4,
    roleName: '运维2',
  },
  {
    roleId: 5,
    roleName: '运维3',
  },
];

const rolesAll = {
  code: 0,
  message: '成功',
  type: 1,
  data: allRoles,
};

const rolesListRows: ResponseAdminRoleListItemType[] = [
  {
    roleId: 1,
    roleName: '管理员',
    createAdminId: 1,
    createAdminName: 'admin',
    enabled: false,
    createTime: '2022-08-11 02:09:58',
    modifyTime: '2022-08-11 02:09:59',
  },
];

const list: ResponseListType = {
  code: 0,
  message: '成功',
  type: 'SUCCESS',
  data: {
    total: 22,
    pageNum: 1,
    pageSize: 10,
    list: rolesListRows,
  },
};

const roleDetailData: ResponseAdminRoleDetailType = {
  id: 1,
  name: '管理员',
  permissionIds: [1, 2, 3, 4, 5, 6, 7, 8],
  describe: '拥有全部权限',
  createAdminId: 1,
  createAdminName: 'admin',
  modifyAdminId: 9,
  modifyAdminName: 'test00004',
  enabled: true,
  createTime: '2022-08-11 02:09:58',
  modifyTime: '2022-08-11 02:09:59',
};

const roleDetail: ResponseDetailType = {
  code: 0,
  message: '成功',
  type: 'SUCCESS',
  data: roleDetailData,
};

export default {
  'POST /api/admin/role/all': rolesAll,
  'POST /api/admin/role/list': (req: Request, res: Response) => {
    list.data.pageNum = req.body.pageNum;
    list.data.pageSize = req.body.pageSize;
    res.json(list);
  },
  'POST /api/admin/role/add': success,
  'POST /api/admin/role/detail': roleDetail,
  'POST /api/admin/role/edit': success,
  'POST /api/admin/role/enable': success,
  'POST /api/admin/role/delete': success,
};
