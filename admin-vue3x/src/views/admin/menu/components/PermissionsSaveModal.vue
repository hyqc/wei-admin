<template>
  <a-modal
    :open="open"
    title="添加菜单操作权限"
    :width="900"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-spin :spinning="loading">
      <a-form :label-col="{ span: 3 }" :wrapper-col="{ span: 20 }">
        <a-form-item label="菜单">
          <a-input :value="detailData?.name" disabled style="width: 260px" />
        </a-form-item>
        <a-form-item label="权限">
          <div v-for="(item, index) in permissions" :key="index" class="permission-row">
            <span class="permission-type">{{ item.typeName }}</span>
            <a-space>
              <a-input
                v-model:value="item.name"
                placeholder="权限名称"
                style="width: 160px"
              />
              <a-input
                v-model:value="item.key"
                placeholder="唯一键"
                style="width: 200px"
                :disabled="item.id !== undefined && item.id > 0"
              />
              <a-switch
                v-model:checked="item.enabled"
                checked-children="启用"
                un-checked-children="禁用"
              />
            </a-space>
          </div>
          <a-empty v-if="permissions.length === 0" description="暂无权限配置" />
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { getAdminMenuPermissions } from '@/api/admin/menu';
import { addAdminMenuPermissions } from '@/api/admin/permission';
import type { MenuTreeItem, MenuPermissionItem } from '@/types/admin_menu';

const props = defineProps<{
  open: boolean;
  detailData?: MenuTreeItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const loading = ref(false);
const confirmLoading = ref(false);
const permissions = ref<MenuPermissionItem[]>([]);

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      permissions.value = [];
      loading.value = true;
      getAdminMenuPermissions({ menuId: props.detailData.id })
        .then((res) => {
          permissions.value = (res.data.permissions || []).map((item) => ({
            ...item,
          }));
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

function handleOk() {
  if (!props.detailData) return;
  confirmLoading.value = true;
  const list = permissions.value.map((item) => ({
    menuId: props.detailData?.id,
    menuName: props.detailData?.name,
    menuPath: props.detailData?.path,
    id: item.id,
    name: item.name,
    key: item.key,
    type: item.type,
    describe: item.describe,
    enabled: item.enabled,
  }));
  addAdminMenuPermissions(list)
    .then((res) => {
      message.success(res.msg, 2);
      emit('update:open', false);
    })
    .finally(() => {
      confirmLoading.value = false;
    });
}

function handleCancel() {
  emit('update:open', false);
}
</script>

<style scoped lang="less">
.permission-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;

  .permission-type {
    flex: 0 0 48px;
    font-weight: 600;
  }
}
</style>
