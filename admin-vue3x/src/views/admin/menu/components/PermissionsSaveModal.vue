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
              <!-- 权限名称与唯一键由系统根据菜单自动生成，不允许手动修改 -->
              <a-input v-model:value="item.name" style="width: 160px" disabled />
              <a-input v-model:value="item.key" style="width: 200px" disabled />
              <a-switch
                v-model:checked="item.enabled"
                checked-children="启用"
                un-checked-children="禁用"
              />
            </a-space>
          </div>
          <div class="permission-tip">权限名称与唯一键由系统根据菜单自动生成，不允许手动修改</div>
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
import { DEFAULT_PERMISSION_TYPES, handleKey } from '@/views/admin/permission/components/common';
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
      // 先取出菜单信息，避免异步回调中 props 变化导致取值异常
      const menuId = props.detailData.id;
      const menuName = props.detailData.name || '';
      const menuPath = props.detailData.path || '';
      getAdminMenuPermissions({ menuId })
        .then((res) => {
          const existing = res.data.permissions || [];
          if (existing.length > 0) {
            // 后端未返回类型名称，按类型补充展示
            permissions.value = existing.map((item) => ({
              ...item,
              typeName: DEFAULT_PERMISSION_TYPES.find((t) => t.key === item.type)?.name || item.type || '',
            }));
            return;
          }
          // 菜单尚未配置权限时，按菜单信息自动生成查看/编辑/删除三类
          permissions.value = DEFAULT_PERMISSION_TYPES.map((type) => ({
            type: type.key,
            typeName: type.name,
            name: `${menuName}${type.name}`,
            key: handleKey(menuPath, type.key),
            enabled: true,
            describe: '',
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

.permission-tip {
  margin: 4px 0 12px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
}
</style>
