<template>
  <a-drawer
    :open="open"
    title="权限详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-descriptions :column="1" bordered>
      <a-descriptions-item label="菜单名称">{{ detailData?.menuName }}</a-descriptions-item>
      <a-descriptions-item label="菜单路径">{{ detailData?.menuPath }}</a-descriptions-item>
      <a-descriptions-item label="权限名称">{{ detailData?.name }}</a-descriptions-item>
      <a-descriptions-item label="唯一键">{{ detailData?.key }}</a-descriptions-item>
      <a-descriptions-item label="权限类型">{{ detailData?.typeText }}</a-descriptions-item>
      <a-descriptions-item label="描述">{{ detailData?.describe }}</a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-badge :status="detailData?.enabled ? 'success' : 'error'" :text="detailData?.enabled ? '启用' : '禁用'" />
      </a-descriptions-item>
      <a-descriptions-item label="接口数量">{{ detailData?.apis?.length || 0 }}</a-descriptions-item>
      <a-descriptions-item label="接口列表">
        <div v-if="detailData?.apis?.length">
          <a-tag v-for="api in detailData?.apis" :key="api.id" color="green">
            {{ api.name }}
          </a-tag>
        </div>
        <span v-else>暂无接口</span>
      </a-descriptions-item>
    </a-descriptions>
    <div class="drawer-footer">
      <Authorization permission="AdminPermissionEdit">
        <a-button type="primary" @click="bindApisStatus = true">
          绑定接口
        </a-button>
      </Authorization>
    </div>
    <BindApisModal v-model:open="bindApisStatus" :detail-data="detailData" @notice="onNotice" />
  </a-drawer>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { DefaultDrawerWidth } from '@/api/config';
import Authorization from '@/components/Authorization.vue';
import BindApisModal from './BindApisModal.vue';
import type { PermissionListItem } from '@/types/admin_permission';

defineProps<{
  open: boolean;
  detailData?: PermissionListItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const bindApisStatus = ref(false);

function onNotice() {
  emit('notice');
}
</script>

<style scoped lang="less">
.drawer-footer {
  margin-top: 16px;
  text-align: right;
}
</style>
