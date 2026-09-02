import { Spin } from 'antd';

/** 请求中的加载指示器（对应 Vue 版 FetchButton） */
export default function FetchButton({ loading, size = 'default' }: { loading?: boolean; size?: 'small' | 'default' | 'large' }) {
  if (!loading) return null;
  return <Spin size={size} />;
}
