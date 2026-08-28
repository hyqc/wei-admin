<template>
  <a-modal
    :open="open"
    title="绑定权限"
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
        v-if="detailData"
        :message="`角色：${detailData.name}`"
        type="info"
        show-icon
        style="margin-bottom: 12px"
      />
      <a-tree
        v-if="treeData.length > 0"
        ref="treeRef"
        :key="treeKey"
        :tree-data="treeData"
        checkable
        :checked-keys="checkedKeys"
        :default-expanded-keys="expandedKeys"
        @check="onCheck"
      />
    </a-spin>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { getAdminRolePermissions, bindAdminRolePermissions } from '@/api/admin/role';
import { getAdminMenuMode } from '@/api/admin/menu';
import { DefaultModalWidth } from '@/api/config';
import type { RoleItem } from '@/types/admin_role';
import type { MenuModeItem } from '@/types/admin_menu';

const props = defineProps<{
  open: boolean;
  detailData?: RoleItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const treeRef = ref();
const loading = ref(false);
const confirmLoading = ref(false);
const treeKey = ref(0);
const treeData = ref<any[]>([]);
const checkedKeys = ref<number[]>([]);
const halfCheckedKeys = ref<number[]>([]);
const expandedKeys = ref<(string | number)[]>([]);

/** 默认只展开第一级（模型），第二级页面默认收起，超过 2 级的内容（权限）不显示 */
function collectExpandKeys(nodes: any[]): (string | number)[] {
  return nodes.filter((node) => node.children?.length).map((node) => node.key);
}

function buildTree(modes: MenuModeItem[]) {
  return modes.map((mode) => ({
    key: `model-${mode.modelId}`,
    title: mode.modelName,
    children: (mode.pages || []).map((page) => ({
      key: `page-${page.pageId}`,
      title: page.pageName,
      children: (page.permissions || []).map((perm) => ({
        key: perm.permissionId as number,
        title: `${perm.permissionName}（${perm.permissionTypeName}）`,
      })),
    })),
  }));
}

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      loading.value = true;
      checkedKeys.value = [];
      Promise.all([
        getAdminMenuMode({}),
        getAdminRolePermissions({ id: props.detailData.id }),
      ])
        .then(([modeRes, permRes]) => {
          treeData.value = buildTree(modeRes.data.modes || []);
          treeKey.value += 1;
          expandedKeys.value = collectExpandKeys(treeData.value);
          const rolePerms = permRes.data.list || [];
          checkedKeys.value = rolePerms.map((item) => item.permissionId).filter((id): id is number => !!id);
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

function onCheck(keys: (string | number)[], info: { halfCheckedKeys: (string | number)[] }) {
  checkedKeys.value = keys.filter((k) => typeof k === 'number') as number[];
  halfCheckedKeys.value = info.halfCheckedKeys.filter((k) => typeof k === 'number') as number[];
}

function handleOk() {
  confirmLoading.value = true;
  bindAdminRolePermissions({
    id: props.detailData?.id,
    permissionIds: [...checkedKeys.value, ...halfCheckedKeys.value],
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
