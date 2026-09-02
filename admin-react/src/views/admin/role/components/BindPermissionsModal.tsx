import { useEffect, useState } from 'react';
import { Alert, Modal, Spin, Tree, message } from 'antd';
import { getAdminRolePermissions, bindAdminRolePermissions } from '@/api/admin/role';
import { getAdminMenuMode } from '@/api/admin/menu';
import { DefaultModalWidth } from '@/api/config';
import type { RoleItem } from '@/types/admin_role';
import type { MenuModeItem } from '@/types/admin_menu';

interface BindPermissionsModalProps {
  open: boolean;
  detailData?: RoleItem;
  onClose: () => void;
  onNotice: () => void;
}

interface TreeNode {
  key: string | number;
  title: string;
  children?: TreeNode[];
}

/** 默认只展开第一级（模型），第二级页面默认收起 */
function collectExpandKeys(nodes: TreeNode[]): (string | number)[] {
  return nodes.filter((node) => node.children?.length).map((node) => node.key);
}

/** 构建 模型→页面→权限 全量树 */
function buildTree(modes: MenuModeItem[]): TreeNode[] {
  return modes.map((mode) => ({
    key: `model-${mode.modelId}`,
    title: mode.modelName ?? '',
    children: (mode.pages || []).map((page) => ({
      key: `page-${page.pageId}`,
      title: page.pageName ?? '',
      children: (page.permissions || []).map((perm) => ({
        key: perm.permissionId as number,
        title: `${perm.permissionName}（${perm.permissionTypeName}）`,
      })),
    })),
  }));
}

/** 角色绑定权限 */
export default function BindPermissionsModal({ open, detailData, onClose, onNotice }: BindPermissionsModalProps) {
  const [loading, setLoading] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [treeData, setTreeData] = useState<TreeNode[]>([]);
  const [checkedKeys, setCheckedKeys] = useState<number[]>([]);
  const [halfCheckedKeys, setHalfCheckedKeys] = useState<number[]>([]);
  const [expandedKeys, setExpandedKeys] = useState<(string | number)[]>([]);
  const [treeKey, setTreeKey] = useState(0);

  useEffect(() => {
    if (open && detailData) {
      setLoading(true);
      setCheckedKeys([]);
      Promise.all([getAdminMenuMode({}), getAdminRolePermissions({ id: detailData.id })])
        .then(([modeRes, permRes]) => {
          const tree = buildTree(modeRes.data.modes || []);
          setTreeData(tree);
          setExpandedKeys(collectExpandKeys(tree));
          const rolePerms = permRes.data.list || [];
          setCheckedKeys(rolePerms.map((item) => item.permissionId).filter((id): id is number => !!id));
          setTreeKey((k) => k + 1);
        })
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  const onCheck = (checked: { checked: (string | number)[]; halfChecked: (string | number)[] }) => {
    setCheckedKeys(checked.checked.filter((k) => typeof k === 'number') as number[]);
    setHalfCheckedKeys(checked.halfChecked.filter((k) => typeof k === 'number') as number[]);
  };

  const handleOk = () => {
    setConfirmLoading(true);
    bindAdminRolePermissions({
      id: detailData?.id,
      permissionIds: [...checkedKeys, ...halfCheckedKeys],
    })
      .then((res) => {
        message.success(res.msg, 2);
        onNotice();
        onClose();
      })
      .finally(() => setConfirmLoading(false));
  };

  return (
    <Modal open={open} title="绑定权限" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Spin spinning={loading}>
        {detailData && <Alert message={`角色：${detailData.name}`} type="info" showIcon style={{ marginBottom: 12 }} />}
        {treeData.length > 0 && (
          <Tree
            key={treeKey}
            treeData={treeData}
            checkable
            checkedKeys={checkedKeys}
            defaultExpandedKeys={expandedKeys}
            onCheck={onCheck as never}
          />
        )}
      </Spin>
    </Modal>
  );
}
