<template>
  <a-drawer
    :open="open"
    title="权限详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-descriptions :column="1" bordered>
        <a-descriptions-item label="菜单名称">{{ detail?.menuName || '-' }}</a-descriptions-item>
        <a-descriptions-item label="菜单路径">{{ detail?.menuPath || '-' }}</a-descriptions-item>
        <a-descriptions-item label="权限名称">{{ detail?.name }}</a-descriptions-item>
        <a-descriptions-item label="唯一键">{{ detail?.key }}</a-descriptions-item>
        <a-descriptions-item label="权限类型">{{ detail?.typeText || detail?.type }}</a-descriptions-item>
        <a-descriptions-item label="描述">{{ detail?.describe || '-' }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-badge :status="detail?.isEnabled ? 'success' : 'error'" :text="detail?.isEnabled ? '启用' : '禁用'" />
        </a-descriptions-item>
        <a-descriptions-item label="接口数量">{{ detail?.apis?.length || 0 }}</a-descriptions-item>
        <a-descriptions-item label="接口列表">
          <div v-if="detail?.apis?.length">
            <a-tag v-for="api in detail?.apis" :key="api.id" color="green">
              {{ api.name }}
            </a-tag>
          </div>
          <span v-else>暂无接口</span>
        </a-descriptions-item>
      </a-descriptions>
    </a-spin>
    <div class="drawer-footer">
      <Authorization permission="AdminPermissionEdit">
        <a-button type="primary" @click="bindApisStatus = true">
          绑定接口
        </a-button>
      </Authorization>
    </div>
    <BindApisModal v-model:open="bindApisStatus" :detail-data="detail" @notice="onNotice" />
  </a-drawer>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { DefaultDrawerWidth } from '@/api/config';
import Authorization from '@/components/Authorization.vue';
import BindApisModal from './BindApisModal.vue';
import { getAdminPermissionInfo } from '@/api/admin/permission';
import type { ResponseAdminPermissionInfoType } from '@/api/admin/permission';
import type { PermissionListItem } from '@/types/admin_permission';

const props = defineProps<{
  open: boolean;
  detailData?: PermissionListItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const bindApisStatus = ref(false);
const loading = ref(false);
const detail = ref<ResponseAdminPermissionInfoType>();

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData?.id) {
      loading.value = true;
      detail.value = undefined;
      getAdminPermissionInfo({ id: props.detailData.id })
        .then((res) => {
          detail.value = res.data;
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

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
