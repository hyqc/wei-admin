<template>
  <div class="page-info">
    共 {{ pageInfo.total }} 条 第 {{ pageInfo.pageNum }}/{{ totalPages }} 页
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { PageInfoType } from '@/api/types';

const props = defineProps<{ pageInfo: PageInfoType }>();

const totalPages = computed(() => {
  const { total, pageSize } = props.pageInfo;
  if (!total || !pageSize) return 1;
  return Math.ceil(total / pageSize);
});
</script>

<style scoped lang="less">
.page-info {
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
}
</style>
