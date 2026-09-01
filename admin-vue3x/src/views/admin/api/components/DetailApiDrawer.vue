<template>
  <a-drawer
    :open="open"
    title="接口详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-descriptions :column="1" bordered>
      <a-descriptions-item label="接口ID">{{ detailData?.id }}</a-descriptions-item>
      <a-descriptions-item label="接口名称">{{ detailData?.name }}</a-descriptions-item>
      <a-descriptions-item label="唯一键">{{ detailData?.key }}</a-descriptions-item>
      <a-descriptions-item label="接口路径">{{ detailData?.path }}</a-descriptions-item>
      <a-descriptions-item label="描述">{{ detailData?.describe }}</a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-badge :status="detailData?.isEnabled ? 'success' : 'error'" :text="detailData?.isEnabled ? '启用' : '禁用'" />
      </a-descriptions-item>
      <a-descriptions-item label="创建时间">{{ detailData?.createdAt }}</a-descriptions-item>
      <a-descriptions-item label="更新时间">{{ detailData?.updatedAt }}</a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script setup lang="ts">
import { DefaultDrawerWidth } from '@/api/config';
import type { AdminApiItem } from '@/types/common';

defineProps<{
  open: boolean;
  detailData?: AdminApiItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();
</script>
