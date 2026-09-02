import { useEffect, useState } from 'react';
import { Alert, Empty, Modal, Spin, Tree, message } from 'antd';
import { getAdminPermissionList, getAdminPermissionInfo, bindAdminPermissionApis } from '@/api/admin/permission';
import { getAdminApiAll } from '@/api/admin/api';
import { DefaultModalWidth } from '@/api/config';
import type { AdminApiItem } from '@/types/common';
import type { PermissionListItem } from '@/types/admin_permission';

interface BindApisModalProps {
  open: boolean;
  detailData?: PermissionListItem;
  onClose: () => void;
  onNotice: () => void;
}

interface TreeNode {
  key: string;
  title: string;
  children?: TreeNode[];
}

/** 节点 key 前缀：menu-{menuId} / api-{apiId} */
const API_PREFIX = 'api-';

/** 递归收集所有接口节点 key，用于回显已绑定接口 */
function collectApiKeys(nodes: TreeNode[], checkedSet: Set<number>, out: string[]) {
  for (const node of nodes) {
    if (node.key.startsWith(API_PREFIX)) {
      const apiId = Number(node.key.slice(API_PREFIX.length));
      if (checkedSet.has(apiId)) out.push(node.key);
    }
    if (node.children?.length) collectApiKeys(node.children, checkedSet, out);
  }
}

/** 菜单 → 接口 两级树：接口按所属菜单分组（经权限绑定关系推导，接口不重复） */
function buildTree(permList: PermissionListItem[], apiList: AdminApiItem[]): TreeNode[] {
  const apiMenuMap = new Map<number, { menuId: number; menuName: string }>();
  const sortedPerms = [...permList].sort((x, y) => (x.menuId ?? 0) - (y.menuId ?? 0));
  for (const perm of sortedPerms) {
    if (!perm.menuId) continue;
    for (const api of perm.apis || []) {
      if (api.id != null && !apiMenuMap.has(api.id)) {
        apiMenuMap.set(api.id, { menuId: perm.menuId, menuName: perm.menuName || '未分组' });
      }
    }
  }
  const menuNodes = new Map<number, TreeNode>();
  for (const api of apiList) {
    const owner = api.id != null ? apiMenuMap.get(api.id) : undefined;
    const menuId = owner?.menuId ?? 0;
    const menuName = owner?.menuName || '未分组';
    if (!menuNodes.has(menuId)) {
      menuNodes.set(menuId, { key: `menu-${menuId}`, title: menuName, children: [] });
    }
    menuNodes.get(menuId)!.children!.push({
      key: `api-${api.id}`,
      title: `${api.name}（${api.key}）`,
    });
  }
  return [...menuNodes.values()].sort((a, b) => {
    if (a.key === 'menu-0') return 1;
    if (b.key === 'menu-0') return -1;
    return Number(a.key.slice(5)) - Number(b.key.slice(5));
  });
}

/** 权限绑定接口（打开时默认收起所有菜单分组） */
export default function BindApisModal({ open, detailData, onClose, onNotice }: BindApisModalProps) {
  const [loading, setLoading] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [treeData, setTreeData] = useState<TreeNode[]>([]);
  const [checkedKeys, setCheckedKeys] = useState<string[]>([]);
  const [permissionName, setPermissionName] = useState('');

  useEffect(() => {
    if (open && detailData) {
      setTreeData([]);
      setCheckedKeys([]);
      setPermissionName(detailData.name || '');
      setLoading(true);
      Promise.all([
        getAdminApiAll(),
        // 一次取回全部权限（含各自绑定的接口与所属菜单），用于构建树
        getAdminPermissionList({ pageNum: 1, pageSize: 10000 }),
        getAdminPermissionInfo({ id: detailData.id }),
      ])
        .then(([apiRes, permRes, permInfoRes]) => {
          const tree = buildTree(permRes.data?.list || [], apiRes.data || []);
          setTreeData(tree);
          // 默认收起所有节点，接口列表需用户点击菜单展开
          const checkedSet = new Set(permInfoRes.data.apiIds || []);
          const checked: string[] = [];
          collectApiKeys(tree, checkedSet, checked);
          setCheckedKeys(checked);
        })
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  const handleOk = () => {
    setConfirmLoading(true);
    const apiIds = [
      ...new Set(
        checkedKeys
          .filter((k) => k.startsWith(API_PREFIX))
          .map((k) => Number(k.slice(API_PREFIX.length)))
          .filter((id) => !Number.isNaN(id)),
      ),
    ];
    bindAdminPermissionApis({ permissionId: detailData?.id, apiIds })
      .then((res) => {
        message.success(res.msg, 2);
        onNotice();
        onClose();
      })
      .finally(() => setConfirmLoading(false));
  };

  return (
    <Modal open={open} title="绑定接口" width={DefaultModalWidth} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Spin spinning={loading}>
        {permissionName && <Alert message={`权限：${permissionName}`} type="info" showIcon style={{ marginBottom: 12 }} />}
        {treeData.length === 0 ? (
          <Empty description="暂无接口" />
        ) : (
          <Tree
            treeData={treeData}
            checkable
            checkedKeys={checkedKeys}
            height={480}
            onCheck={(keys) => setCheckedKeys(keys as string[])}
          />
        )}
      </Spin>
    </Modal>
  );
}
