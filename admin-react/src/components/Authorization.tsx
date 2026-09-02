import type { ReactNode } from 'react';
import { useUserStore } from '@/store/user';

/** 按钮级权限控制：permission 为空时始终渲染 */
export default function Authorization({ permission, children }: { permission?: string; children?: ReactNode }) {
  const hasPermission = useUserStore((s) => s.hasPermission);
  if (!hasPermission(permission)) return null;
  return <>{children}</>;
}
