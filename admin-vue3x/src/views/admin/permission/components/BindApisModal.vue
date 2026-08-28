<template>
  <a-modal
    :open="open"
    title="绑定接口"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-alert
        v-if="permissionName"
        :message="`权限：${permissionName}`"
        type="info"
        show-icon
        style="margin-bottom: 12px"
      />
      <a-checkbox-group v-model:value="apiIds" style="width: 100%">
        <a-checkbox v-for="api in apiOptions" :key="api.id" :value="api.id" style="display: block; margin-bottom: 8px">
          <span>{{ api.name }}</span>
          <span class="api-key">（{{ api.key }}）</span>
          <span class="api-path">{{ api.path }}</span>
        </a-checkbox>
      </a-checkbox-group>
    </a-spin>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { getAdminPermissionInfo, bindAdminPermissionApis } from '@/api/admin/permission';
import { getAdminApiAll } from '@/api/admin/api';
import { DefaultModalWidth } from '@/api/config';
import type { AdminApiItem } from '@/types/common';
import type { PermissionListItem } from '@/types/admin_permission';

const props = defineProps<{
  open: boolean;
  detailData?: PermissionListItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const loading = ref(false);
const confirmLoading = ref(false);
const apiOptions = ref<AdminApiItem[]>([]);
const apiIds = ref<number[]>([]);
const permissionName = ref('');

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      apiIds.value = [];
      permissionName.value = props.detailData.name || '';
      loading.value = true;
      Promise.all([
        getAdminApiAll(),
        getAdminPermissionInfo({ id: props.detailData.id }),
      ])
        .then(([apiRes, permRes]) => {
          apiOptions.value = apiRes.data || [];
          apiIds.value = permRes.data.apiIds || [];
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

function handleOk() {
  confirmLoading.value = true;
  bindAdminPermissionApis({
    permissionId: props.detailData?.id,
    apiIds: apiIds.value,
  })
    .then((res) => {
      message.success(res.msg, 2);
      emit('notice');
      emit('update:open', false);
    })
    .finally(() => {
      confirmLoading.value = false;
    });
}
</script>

<style scoped lang="less">
.api-key {
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
  margin-left: 4px;
}

.api-path {
  color: rgba(0, 0, 0, 0.65);
  font-size: 12px;
  margin-left: 8px;
}
</style>
