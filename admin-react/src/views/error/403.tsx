import { useNavigate } from 'react-router-dom';
import { Button, Result } from 'antd';
import { HomePath } from '@/api/config';

/** 403 无权限页 */
export default function Forbidden() {
  const navigate = useNavigate();
  return (
    <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
      <Result
        status="403"
        title="403"
        subTitle="抱歉，您没有权限访问该页面"
        extra={
          <Button type="primary" onClick={() => navigate(HomePath, { replace: true })}>
            返回首页
          </Button>
        }
      />
    </div>
  );
}
