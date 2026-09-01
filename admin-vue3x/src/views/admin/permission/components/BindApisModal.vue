<template>
  <a-modal
    :open="open"
    title="绑定接口"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-alert
        v-if="permissionName"
        :message="`权限：${permissionName}`"
        type="info"
        show-icon
        style="margin-bottom: 12px"
      />
      <a-empty v-if="treeData.length === 0" description="暂无接口" />
      <a-tree
        v-else
        ref="treeRef"
        :key="treeKey"
        :tree-data="treeData"
        checkable
        :checked-keys="checkedKeys"
        :default-expanded-keys="expandedKeys"
        :height="480"
        @check="onCheck"
      />
    </a-spin>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { getAdminPermissionList, getAdminPermissionInfo, bindAdminPermissionApis } from '@/api/admin/permission';
import { getAdminApiAll } from '@/api/admin/api';
import { DefaultModalWidth } from '@/api/config';
import type { AdminApiItem } from '@/types/common';
import type { PermissionListItem } from '@/types/admin_permission';

const props = defineProps<{
  open: boolean;
  detailData?: PermissionListItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

/** 节点 key 前缀：menu-{menuId} / api-{apiId} */
const API_PREFIX = 'api-';

const treeRef = ref();
const loading = ref(false);
const confirmLoading = ref(false);
const treeKey = ref(0);
const treeData = ref<any[]>([]);
const checkedKeys = ref<(string | number)[]>([]);
const expandedKeys = ref<(string | number)[]>([]);
const permissionName = ref('');

/** 递归收集所有接口节点 key，用于回显已绑定接口 */
function collectApiKeys(nodes: any[], checkedSet: Set<number>, out: (string | number)[]) {
  for (const node of nodes) {
    if (typeof node.key === 'string' && node.key.startsWith(API_PREFIX)) {
      const apiId = Number(node.key.slice(API_PREFIX.length));
      if (checkedSet.has(apiId)) out.push(node.key);
    }
    if (node.children?.length) collectApiKeys(node.children, checkedSet, out);
  }
}

/** 菜单 → 接口 两级树：接口按所属菜单分组（经权限绑定关系推导，接口不重复） */
function buildTree(permList: PermissionListItem[], apiList: AdminApiItem[]) {
  // 接口 → 所属菜单（同一接口被多个权限绑定时取 menuId 最小者）
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
  // 按菜单分组
  const menuNodes = new Map<number, { key: string; title: string; children: any[] }>();
  for (const api of apiList) {
    const owner = api.id != null ? apiMenuMap.get(api.id) : undefined;
    const menuId = owner?.menuId ?? 0;
    const menuName = owner?.menuName || '未分组';
    if (!menuNodes.has(menuId)) {
      menuNodes.set(menuId, { key: `menu-${menuId}`, title: menuName, children: [] });
    }
    menuNodes.get(menuId)!.children.push({
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

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      treeData.value = [];
      checkedKeys.value = [];
      permissionName.value = props.detailData.name || '';
      loading.value = true;
      Promise.all([
        getAdminApiAll(),
        // 一次取回全部权限（含各自绑定的接口与所属菜单），用于构建树
        getAdminPermissionList({ pageNum: 1, pageSize: 10000 }),
        getAdminPermissionInfo({ id: props.detailData.id }),
      ])
        .then(([apiRes, permRes, permResData]) => {
          treeData.value = buildTree(permRes.data?.list || [], apiRes.data || []);
          treeKey.value += 1;
          // 默认收起所有节点，接口列表需用户点击菜单展开
          expandedKeys.value = [];
          const checkedSet = new Set(permResData.data.apiIds || []);
          collectApiKeys(treeData.value, checkedSet, checkedKeys.value);
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

function onCheck(keys: (string | number)[]) {
  checkedKeys.value = keys;
}

function handleOk() {
  confirmLoading.value = true;
  const apiIds = [...new Set(
    checkedKeys.value
      .filter((k) => typeof k === 'string' && k.startsWith(API_PREFIX))
      .map((k) => Number((k as string).slice(API_PREFIX.length)))
      .filter((id) => !Number.isNaN(id)),
  )];
  bindAdminPermissionApis({
    permissionId: props.detailData?.id,
    apiIds,
  })
    .then((res) => {
      message.success(res.msg, 2);
      emit('notice');
      emit('update:open', false);
    })
    .finally(() => {
      confirmLoading.value = false;
    });
}
</script>
