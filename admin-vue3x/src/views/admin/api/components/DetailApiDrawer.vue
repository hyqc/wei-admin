<template>
  <a-drawer
    :open="open"
    title="接口详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-descriptions :column="1" bordered>
        <a-descriptions-item label="接口ID">{{ detail?.id }}</a-descriptions-item>
        <a-descriptions-item label="接口名称">{{ detail?.name }}</a-descriptions-item>
        <a-descriptions-item label="唯一键">{{ detail?.key }}</a-descriptions-item>
        <a-descriptions-item label="接口路径">{{ detail?.path }}</a-descriptions-item>
        <a-descriptions-item label="描述">{{ detail?.describe }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-badge :status="detail?.isEnabled ? 'success' : 'error'" :text="detail?.isEnabled ? '启用' : '禁用'" />
        </a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ detail?.createdAt }}</a-descriptions-item>
        <a-descriptions-item label="更新时间">{{ detail?.updatedAt }}</a-descriptions-item>
      </a-descriptions>
    </a-spin>
  </a-drawer>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminApiInfo } from '@/api/admin/api';
import type { AdminApiItem } from '@/types/common';

const props = defineProps<{
  open: boolean;
  detailData?: AdminApiItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const loading = ref(false);
/** 接口详情（实时拉取，避免展示列表中的过期数据） */
const detail = ref<AdminApiItem>();

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData?.id) {
      loading.value = true;
      detail.value = undefined;
      getAdminApiInfo({ id: props.detailData.id })
        .then((res) => {
          detail.value = res.data;
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);
</script>
