<template>
  <a-modal
    :open="open"
    title="绑定接口"
    :width="720"
    :mask-closable="false"
    ok-text="确定"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-alert :message="tip" type="info" show-icon style="margin-bottom: 12px" />
      <a-empty v-if="treeData.length === 0" description="暂无接口" />
      <a-tree
        v-else
        :tree-data="treeData"
        checkable
        default-expand-all
        :checked-keys="checkedKeys"
        :height="420"
        @check="onCheck"
      />
    </a-spin>
  </a-modal>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { getAdminApiAll } from '@/api/admin/api';
import { getAdminMenuAll } from '@/api/admin/menu';
import {
  API_NODE_PREFIX,
  buildApiTree,
  buildMenuNameByPath,
  parseCheckedApiIds,
} from '@/views/admin/permission/components/common';

const props = defineProps<{
  open: boolean;
  /** 已绑定的接口ID */
  apiIds?: number[];
  /** 权限名称，用于弹窗提示 */
  permissionName?: string;
}>();
const emit = defineEmits<{
  (e: 'update:open', value: boolean): void;
  (e: 'ok', apiIds: number[]): void;
}>();

const loading = ref(false);
const treeData = ref<{ key: string; title: string; children?: { key: string; title: string }[] }[]>([]);
const checkedKeys = ref<(string | number)[]>([]);

const tip = computed(() => `权限：${props.permissionName || '-'}，勾选该操作需要访问的接口`);

watch(
  () => props.open,
  (val) => {
    if (!val) return;
    checkedKeys.value = (props.apiIds || []).map((id) => `${API_NODE_PREFIX}${id}`);
    loading.value = true;
    // 菜单名用于把分组标题渲染为「接口管理 (/admin/api)」
    Promise.all([getAdminApiAll(), getAdminMenuAll()])
      .then(([apiRes, menuRes]) => {
        treeData.value = buildApiTree(apiRes.data || [], buildMenuNameByPath(menuRes.data || []));
      })
      .finally(() => {
        loading.value = false;
      });
  },
);

function onCheck(keys: (string | number)[] | { checked: (string | number)[] }) {
  checkedKeys.value = Array.isArray(keys) ? keys : keys.checked || [];
}

function handleOk() {
  emit('ok', parseCheckedApiIds(checkedKeys.value));
  emit('update:open', false);
}
</script>
