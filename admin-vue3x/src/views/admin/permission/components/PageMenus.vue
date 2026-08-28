<template>
  <a-tree
    :tree-data="treeData"
    default-expand-all
    selectable
    :selected-keys="selectedKeys"
    :field-names="{ title: 'name', key: 'id', children: 'children' }"
    @select="onSelect"
  >
    <template #title="{ name, path }">
      <span>{{ name }}</span>
      <span v-if="path" class="menu-path">（{{ path }}）</span>
    </template>
  </a-tree>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { getAdminMenuPages } from '@/api/admin/menu';
import type { MenuTreeItem } from '@/types/admin_menu';

const props = defineProps<{
  value?: number;
}>();
const emit = defineEmits<{
  (e: 'update:value', value?: number): void;
  (e: 'change', value?: MenuTreeItem): void;
}>();

const menuPages = ref<MenuTreeItem[]>([]);
const selectedKeys = ref<number[]>([]);

/** 仅显示叶子节点（页面） */
const treeData = computed(() => {
  const walk = (list: MenuTreeItem[]): MenuTreeItem[] =>
    list
      .filter((item) => item.hideInMenu !== true)
      .map((item) => ({
        ...item,
        children: item.children?.length ? walk(item.children) : undefined,
      }));
  return walk(menuPages.value);
});

watch(
  () => props.value,
  (val) => {
    selectedKeys.value = val ? [val] : [];
  },
);

function onSelect(keys: number[], info: { node: any }) {
  const id = keys?.[0];
  const node = info.node;
  emit('update:value', id);
  emit('change', node?.data ? { ...node.data } : undefined);
}

function load() {
  getAdminMenuPages({ all: true }).then((res) => {
    menuPages.value = res.data || [];
  });
}

defineExpose({ load });
</script>

<style scoped lang="less">
.menu-path {
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
}
</style>
