import { useEffect, useState } from 'react';
import { Alert, Empty, Modal, Spin, Tree } from 'antd';
import { getAdminApiAll } from '@/api/admin/api';
import { getAdminMenuAll } from '@/api/admin/menu';
import {
  API_NODE_PREFIX,
  buildApiTree,
  buildMenuNameByPath,
  parseCheckedApiIds,
} from '@/views/admin/permission/components/common';

interface Props {
  open: boolean;
  /** 已绑定的接口ID */
  apiIds?: number[];
  /** 权限名称，用于弹窗提示 */
  permissionName?: string;
  onOk: (apiIds: number[]) => void;
  onClose: () => void;
}

/** 权限点绑定接口（菜单权限配置内联使用，勾选结果由父级保存时统一提交） */
export default function PermissionApiBindModal({ open, apiIds, permissionName, onOk, onClose }: Props) {
  const [loading, setLoading] = useState(false);
  const [treeData, setTreeData] = useState<{ key: string; title: string; children?: unknown[] }[]>([]);
  const [checkedKeys, setCheckedKeys] = useState<React.Key[]>([]);

  useEffect(() => {
    if (!open) return;
    setCheckedKeys((apiIds || []).map((id) => `${API_NODE_PREFIX}${id}`));
    setLoading(true);
    // 菜单名用于把分组标题渲染为「接口管理 (/admin/api)」
    Promise.all([getAdminApiAll(), getAdminMenuAll()])
      .then(([apiRes, menuRes]) => setTreeData(buildApiTree(apiRes.data || [], buildMenuNameByPath(menuRes.data || []))))
      .finally(() => setLoading(false));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [open]);

  return (
    <Modal
      open={open}
      title="绑定接口"
      width={720}
      maskClosable={false}
      okText="确定"
      cancelText="取消"
      onOk={() => {
        onOk(parseCheckedApiIds(checkedKeys as (string | number)[]));
        onClose();
      }}
      onCancel={onClose}
    >
      <Spin spinning={loading}>
        <Alert type="info" showIcon message={`权限：${permissionName || '-'}，勾选该操作需要访问的接口`} style={{ marginBottom: 12 }} />
        {treeData.length === 0 ? (
          <Empty description="暂无接口" />
        ) : (
          <Tree
            checkable
            treeData={treeData as never}
            checkedKeys={checkedKeys}
            defaultExpandAll
            height={420}
            onCheck={(keys) => setCheckedKeys(Array.isArray(keys) ? keys : keys.checked)}
          />
        )}
      </Spin>
    </Modal>
  );
}
