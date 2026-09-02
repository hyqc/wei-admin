<template>
  <a-drawer
    :open="open"
    title="菜单详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-descriptions :column="1" bordered>
        <a-descriptions-item label="菜单名称">{{ detail?.name }}</a-descriptions-item>
        <a-descriptions-item label="键名">{{ detail?.key }}</a-descriptions-item>
        <a-descriptions-item label="菜单路径">{{ detail?.path }}</a-descriptions-item>
        <a-descriptions-item label="重定向地址">{{ detail?.redirect }}</a-descriptions-item>
        <a-descriptions-item label="排序">{{ detail?.sort }}</a-descriptions-item>
        <a-descriptions-item label="图标">
          <span v-if="menuIcon(detail?.icon)" class="icon-cell">
            <component :is="menuIcon(detail?.icon)" />
            <span>{{ detail?.icon }}</span>
          </span>
          <span v-else>{{ detail?.icon || '-' }}</span>
        </a-descriptions-item>
        <a-descriptions-item label="描述">{{ detail?.describe }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-badge :status="detail?.enabled ? 'success' : 'error'" :text="detail?.enabled ? '启用' : '禁用'" />
        </a-descriptions-item>
        <a-descriptions-item label="是否显示">
          {{ detail?.hideInMenu ? '隐藏' : '显示' }}
        </a-descriptions-item>
        <a-descriptions-item label="是否隐藏子菜单">
          {{ detail?.hideChildrenInMenu ? '隐藏' : '显示' }}
        </a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ formatTime(detail?.createTime) }}</a-descriptions-item>
        <a-descriptions-item label="更新时间">{{ formatTime(detail?.modifyTime) }}</a-descriptions-item>
      </a-descriptions>
    </a-spin>
  </a-drawer>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminMenuInfo } from '@/api/admin/menu';
import { getAntIcon } from '@/utils/icon';
import type { MenuTreeItem } from '@/types/admin_menu';

const props = defineProps<{
  open: boolean;
  detailData?: MenuTreeItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const loading = ref(false);
/** 菜单详情（实时拉取，避免展示列表中的过期数据） */
const detail = ref<MenuTreeItem>();

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData?.id) {
      loading.value = true;
      detail.value = undefined;
      getAdminMenuInfo({ menuId: props.detailData.id })
        .then((res) => {
          detail.value = { ...(res.data as MenuTreeItem) };
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

/** 按名称取图标组件，名称为空或不存在时不渲染 */
function menuIcon(name?: string) {
  return getAntIcon(name);
}

function formatTime(time?: number) {
  if (!time) return '';
  const d = new Date(time * 1000);
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
}
</script>

<style scoped lang="less">
.icon-cell {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
</style>
