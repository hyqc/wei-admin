<template>
  <a-tree
    :tree-data="treeData"
    :height="320"
    :expanded-keys="expandedKeys"
    selectable
    :selected-keys="selectedKeys"
    :field-names="{ title: 'name', key: 'id', children: 'children' }"
    @select="onSelect"
    @expand="(keys: number[]) => (expandedKeys = keys)"
  >
    <template #title="{ name, path }">
      <span>{{ name }}</span>
      <span v-if="path" class="menu-path">（{{ path }}）</span>
    </template>
  </a-tree>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
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
/** 展开的节点 key，数据异步加载后全部展开 */
const expandedKeys = ref<number[]>([]);
/** 菜单 id 与原始菜单数据映射，用于选中时回填完整菜单信息 */
let menuMap = new Map<number, MenuTreeItem>();

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

function onSelect(keys: number[]) {
  const id = keys?.[0];
  emit('update:value', id);
  emit('change', id ? menuMap.get(id) : undefined);
}

/** 递归收集所有含子节点的菜单 id */
function collectParentKeys(list: MenuTreeItem[]): number[] {
  const keys: number[] = [];
  for (const item of list) {
    if (item.children?.length) {
      keys.push(item.id as number);
      keys.push(...collectParentKeys(item.children));
    }
  }
  return keys;
}

function load() {
  getAdminMenuPages({ all: true }).then((res) => {
    menuPages.value = res.data || [];
    const map = new Map<number, MenuTreeItem>();
    const collectMap = (list: MenuTreeItem[]) => {
      list.forEach((item) => {
        map.set(item.id as number, item);
        if (item.children?.length) collectMap(item.children);
      });
    };
    collectMap(menuPages.value);
    menuMap = map;
    expandedKeys.value = collectParentKeys(treeData.value);
  });
}

// 弹窗首次打开时子组件才挂载，父组件调用 load 拿不到实例，这里挂载自身即加载
onMounted(() => {
  load();
});

defineExpose({ load });
</script>

<style scoped lang="less">
.menu-path {
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
}
</style>
