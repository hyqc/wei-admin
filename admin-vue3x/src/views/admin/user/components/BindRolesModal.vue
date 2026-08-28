<template>
  <a-modal
    :open="open"
    title="绑定角色"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-form :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }">
        <a-form-item label="账号">
          <a-input :value="detailData?.username" disabled />
        </a-form-item>
        <a-form-item label="角色">
          <a-select v-model:value="roleIds" mode="multiple" placeholder="请选择角色" allow-clear>
            <a-select-option v-for="role in roleOptions" :key="role.id" :value="role.id">
              {{ role.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { bindAdminUserRoles, getAdminUserInfo } from '@/api/admin/user';
import { getAdminRoleAll } from '@/api/admin/role';
import { DefaultModalWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';
import type { RoleItem } from '@/types/admin_role';

const props = defineProps<{
  open: boolean;
  detailData?: AdminUserListItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const loading = ref(false);
const confirmLoading = ref(false);
const roleIds = ref<number[]>([]);
const roleOptions = ref<RoleItem[]>([]);

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      roleIds.value = [];
      loading.value = true;
      getAdminRoleAll()
        .then((res) => {
          roleOptions.value = res.data || [];
        })
        .finally(() => {
          loading.value = false;
        });
      getAdminUserInfo({ adminId: props.detailData.adminId }).then((res) => {
        roleIds.value = res.data.roleIds || [];
      });
    }
  },
);

function handleOk() {
  confirmLoading.value = true;
  bindAdminUserRoles({ adminId: props.detailData?.adminId, roleIds: roleIds.value })
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
