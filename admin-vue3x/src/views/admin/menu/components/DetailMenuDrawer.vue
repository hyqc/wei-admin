<template>
  <a-drawer
    :open="open"
    title="菜单详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-descriptions :column="1" bordered>
      <a-descriptions-item label="菜单名称">{{ detailData?.name }}</a-descriptions-item>
      <a-descriptions-item label="键名">{{ detailData?.key }}</a-descriptions-item>
      <a-descriptions-item label="菜单路径">{{ detailData?.path }}</a-descriptions-item>
      <a-descriptions-item label="重定向地址">{{ detailData?.redirect }}</a-descriptions-item>
      <a-descriptions-item label="排序">{{ detailData?.sort }}</a-descriptions-item>
      <a-descriptions-item label="图标">
        <span v-if="menuIcon(detailData?.icon)" class="icon-cell">
          <component :is="menuIcon(detailData?.icon)" />
          <span>{{ detailData?.icon }}</span>
        </span>
        <span v-else>{{ detailData?.icon || '-' }}</span>
      </a-descriptions-item>
      <a-descriptions-item label="描述">{{ detailData?.describe }}</a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-badge :status="detailData?.enabled ? 'success' : 'error'" :text="detailData?.enabled ? '启用' : '禁用'" />
      </a-descriptions-item>
      <a-descriptions-item label="是否显示">
        {{ detailData?.hideInMenu ? '隐藏' : '显示' }}
      </a-descriptions-item>
      <a-descriptions-item label="是否隐藏子菜单">
        {{ detailData?.hideChildrenInMenu ? '隐藏' : '显示' }}
      </a-descriptions-item>
      <a-descriptions-item label="创建时间">{{ formatTime(detailData?.createTime) }}</a-descriptions-item>
      <a-descriptions-item label="更新时间">{{ formatTime(detailData?.modifyTime) }}</a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script setup lang="ts">
import { DefaultDrawerWidth } from '@/api/config';
import { getAntIcon } from '@/utils/icon';
import type { MenuTreeItem } from '@/types/admin_menu';

defineProps<{
  open: boolean;
  detailData?: MenuTreeItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

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
