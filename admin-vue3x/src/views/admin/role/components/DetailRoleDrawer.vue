<template>
  <a-drawer
    :open="open"
    title="角色详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-descriptions :column="1" bordered>
        <a-descriptions-item label="角色名称">{{ detail?.name }}</a-descriptions-item>
        <a-descriptions-item label="描述">{{ detail?.describe }}</a-descriptions-item>
        <a-descriptions-item label="创建人">{{ detail?.createAdminName }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-badge :status="detail?.isEnabled ? 'success' : 'error'" :text="detail?.isEnabled ? '启用' : '禁用'" />
        </a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ detail?.createdAt }}</a-descriptions-item>
        <a-descriptions-item label="更新时间">{{ detail?.updatedAt }}</a-descriptions-item>
      </a-descriptions>
      <a-divider orientation="left">绑定权限</a-divider>
      <a-empty v-if="!loading && treeData.length === 0" description="该角色未绑定任何权限" />
      <a-tree
        v-if="treeData.length > 0"
        :key="treeKey"
        :tree-data="treeData"
        checkable
        :checked-keys="checkedKeys"
        :default-expanded-keys="expandedKeys"
        selectable
      />
    </a-spin>
  </a-drawer>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminMenuMode } from '@/api/admin/menu';
import { getAdminRoleInfo, getAdminRolePermissions } from '@/api/admin/role';
import type { RoleItem } from '@/types/admin_role';
import type { MenuModeItem } from '@/types/admin_menu';

const props = defineProps<{
  open: boolean;
  detailData?: RoleItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const loading = ref(false);
/** 角色详情（实时拉取，避免展示列表中的过期数据） */
const detail = ref<RoleItem>();
const treeKey = ref(0);
const treeData = ref<any[]>([]);
const checkedKeys = ref<number[]>([]);
const expandedKeys = ref<(string | number)[]>([]);

/** 默认只展开第一级（模型），第二级页面默认收起，超过 2 级的内容（权限）不显示 */
function collectExpandKeys(nodes: any[]): (string | number)[] {
  return nodes.filter((node) => node.children?.length).map((node) => node.key);
}

/** 仅保留该角色已绑定的权限分支，构建 模型→页面→权限 树 */
function buildBoundTree(modes: MenuModeItem[], boundIds: Set<number>) {
  const result: any[] = [];
  for (const mode of modes) {
    const pages = (mode.pages || [])
      .map((page) => {
        const perms = (page.permissions || []).filter((perm) => boundIds.has(perm.permissionId as number));
        if (perms.length === 0) return null;
        return {
          key: `page-${page.pageId}`,
          title: page.pageName,
          children: perms.map((perm) => ({
            key: perm.permissionId,
            title: `${perm.permissionName}（${perm.permissionTypeName}）`,
          })),
        };
      })
      .filter((page): page is NonNullable<typeof page> => !!page);
    if (pages.length > 0) {
      result.push({ key: `model-${mode.modelId}`, title: mode.modelName, children: pages });
    }
  }
  return result;
}

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      loading.value = true;
      treeData.value = [];
      checkedKeys.value = [];
      Promise.all([
        getAdminRoleInfo({ id: props.detailData.id }),
        getAdminMenuMode({}),
        getAdminRolePermissions({ id: props.detailData.id }),
      ])
        .then(([infoRes, modeRes, permRes]) => {
          detail.value = infoRes.data;
          const rolePerms = permRes.data.list || [];
          const boundIds = new Set(rolePerms.map((item) => item.permissionId).filter((id): id is number => !!id));
          checkedKeys.value = Array.from(boundIds);
          treeData.value = buildBoundTree(modeRes.data.modes || [], boundIds);
          treeKey.value += 1;
          expandedKeys.value = collectExpandKeys(treeData.value);
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);
</script>
