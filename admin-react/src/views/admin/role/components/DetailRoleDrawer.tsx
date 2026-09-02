import { useEffect, useState } from 'react';
import { Badge, Descriptions, Divider, Drawer, Empty, Spin, Tree } from 'antd';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminMenuMode } from '@/api/admin/menu';
import { getAdminRoleInfo, getAdminRolePermissions } from '@/api/admin/role';
import type { RoleItem } from '@/types/admin_role';
import type { MenuModeItem } from '@/types/admin_menu';

interface DetailRoleDrawerProps {
  open: boolean;
  detailData?: RoleItem;
  onClose: () => void;
}

interface TreeNode {
  key: string | number;
  title: string;
  children?: TreeNode[];
}

/** 仅保留该角色已绑定的权限分支，构建 模型→页面→权限 树 */
function buildBoundTree(modes: MenuModeItem[], boundIds: Set<number>): TreeNode[] {
  const result: TreeNode[] = [];
  for (const mode of modes) {
    const pages: TreeNode[] = [];
    for (const page of mode.pages || []) {
      const perms = (page.permissions || []).filter((perm) => boundIds.has(perm.permissionId as number));
      if (perms.length === 0) continue;
      pages.push({
        key: `page-${page.pageId}`,
        title: page.pageName ?? '',
        children: perms.map((perm) => ({
          key: perm.permissionId as number,
          title: `${perm.permissionName}（${perm.permissionTypeName}）`,
        })),
      });
    }
    if (pages.length > 0) {
      result.push({ key: `model-${mode.modelId}`, title: mode.modelName ?? '', children: pages });
    }
  }
  return result;
}

/** 默认只展开第一级（模型），第二级页面默认收起 */
function collectExpandKeys(nodes: TreeNode[]): (string | number)[] {
  return nodes.filter((node) => node.children?.length).map((node) => node.key);
}

/** 角色详情：基础信息 + 绑定权限树 */
export default function DetailRoleDrawer({ open, detailData, onClose }: DetailRoleDrawerProps) {
  const [loading, setLoading] = useState(false);
  /** 角色详情（实时拉取，避免展示列表中的过期数据） */
  const [detail, setDetail] = useState<RoleItem>();
  const [treeData, setTreeData] = useState<TreeNode[]>([]);
  const [checkedKeys, setCheckedKeys] = useState<number[]>([]);
  const [expandedKeys, setExpandedKeys] = useState<(string | number)[]>([]);
  const [treeKey, setTreeKey] = useState(0);

  useEffect(() => {
    if (open && detailData) {
      setLoading(true);
      setTreeData([]);
      setCheckedKeys([]);
      Promise.all([
        getAdminRoleInfo({ id: detailData.id }),
        getAdminMenuMode({}),
        getAdminRolePermissions({ id: detailData.id }),
      ])
        .then(([infoRes, modeRes, permRes]) => {
          setDetail(infoRes.data);
          const rolePerms = permRes.data.list || [];
          const boundIds = new Set(rolePerms.map((item) => item.permissionId).filter((id): id is number => !!id));
          setCheckedKeys(Array.from(boundIds));
          const tree = buildBoundTree(modeRes.data.modes || [], boundIds);
          setTreeData(tree);
          setExpandedKeys(collectExpandKeys(tree));
          setTreeKey((k) => k + 1);
        })
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  return (
    <Drawer open={open} title="角色详情" width={DefaultDrawerWidth} onClose={onClose}>
      <Spin spinning={loading}>
        <Descriptions column={1} bordered>
          <Descriptions.Item label="角色名称">{detail?.name}</Descriptions.Item>
          <Descriptions.Item label="描述">{detail?.describe}</Descriptions.Item>
          <Descriptions.Item label="创建人">{detail?.createAdminName}</Descriptions.Item>
          <Descriptions.Item label="状态">
            <Badge status={detail?.isEnabled ? 'success' : 'error'} text={detail?.isEnabled ? '启用' : '禁用'} />
          </Descriptions.Item>
          <Descriptions.Item label="创建时间">{detail?.createdAt}</Descriptions.Item>
          <Descriptions.Item label="更新时间">{detail?.updatedAt}</Descriptions.Item>
        </Descriptions>
        <Divider orientation="left">绑定权限</Divider>
        {!loading && treeData.length === 0 && <Empty description="该角色未绑定任何权限" />}
        {treeData.length > 0 && (
          <Tree
            key={treeKey}
            treeData={treeData}
            checkable
            checkedKeys={checkedKeys}
            defaultExpandedKeys={expandedKeys}
            selectable
          />
        )}
      </Spin>
    </Drawer>
  );
}
