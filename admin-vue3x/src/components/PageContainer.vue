<template>
  <div class="page-container">
    <div class="page-header">
      <div class="header-title">
        <span v-if="title">{{ title }}</span>
        <slot name="title" />
      </div>
      <div class="header-search">
        <slot name="searchArea" />
      </div>
      <div class="header-extra">
        <slot name="extra" />
      </div>
    </div>
    <div class="page-content">
      <slot />
    </div>
    <div v-if="pageInfo" class="page-footer">
      <div class="footer-pageinfo">
        <PageInfo :page-info="pageInfo" />
      </div>
      <div class="footer-pagination">
        <a-pagination
          v-model:current="pageInfo.pageNum"
          v-model:page-size="pageInfo.pageSize"
          :total="pageInfo.total"
          :show-size-changer="true"
          :show-quick-jumper="true"
          :page-size-options="pageSizeOptions"
          show-less-items
          @change="onPageChange"
          @showSizeChange="onSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { PageInfoType } from '@/api/types';
import PageInfo from './PageInfo.vue';

const props = defineProps<{
  title?: string;
  pageInfo?: PageInfoType;
  pageSizeOptions?: string[];
}>();

const emit = defineEmits<{
  (e: 'pageChange', pageNum: number): void;
  (e: 'pageSizeChange', pageSize: number): void;
}>();

const pageSizeOptions = computed(() => props.pageSizeOptions || ['10', '20', '50', '100']);

const onPageChange = (pageNum: number) => {
  emit('pageChange', pageNum);
};

const onSizeChange = (_current: number, size: number) => {
  emit('pageSizeChange', size);
};
</script>

<style scoped lang="less">
.page-container {
  padding: 16px;
  background: #fff;
  border-radius: 2px;

  .page-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;

    .header-title {
      font-size: 16px;
      font-weight: 600;
      color: rgba(0, 0, 0, 0.85);
      white-space: nowrap;
    }
  }

  .page-content {
    margin-bottom: 16px;
  }

  .page-footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
}
</style>
