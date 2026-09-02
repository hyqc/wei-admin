import { useNavigate } from 'react-router-dom';
import { Button, Result } from 'antd';
import { HomePath } from '@/api/config';

/** 404 页面不存在 */
export default function NotFound() {
  const navigate = useNavigate();
  return (
    <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
      <Result
        status="404"
        title="404"
        subTitle="抱歉，您访问的页面不存在"
        extra={
          <Button type="primary" onClick={() => navigate(HomePath, { replace: true })}>
            返回首页
          </Button>
        }
      />
    </div>
  );
}
