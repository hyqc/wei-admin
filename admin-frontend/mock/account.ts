import { currentAdminUserDetail } from './login';

const CurrentAdminUserDetail = {
  code: 0,
  message: '成功',
  type: 'SUCCESS',
  data: currentAdminUserDetail,
};
export default {
  'POST /api/admin/account/detail': CurrentAdminUserDetail,
};
