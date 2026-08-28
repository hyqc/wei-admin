<template>
  <a-modal
    :open="open"
    title="编辑权限"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 14 }">
      <a-form-item label="菜单名称">
        <a-input :value="detailData?.menuName" disabled />
      </a-form-item>
      <a-form-item label="权限名称" name="name">
        <a-input v-model:value="formState.name" placeholder="请输入权限名称" allow-clear />
      </a-form-item>
      <a-form-item label="权限类型" name="type">
        <a-radio-group v-model:value="formState.type">
          <a-radio-button
            v-for="item in DEFAULT_PERMISSION_TYPES"
            :key="item.key"
            :value="item.key"
          >
            {{ item.name }}
          </a-radio-button>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="唯一键" name="key">
        <a-input v-model:value="formState.key" disabled />
      </a-form-item>
      <a-form-item label="权限描述" name="describe">
        <a-textarea v-model:value="formState.describe" placeholder="请输入权限描述" :rows="3" />
      </a-form-item>
      <a-form-item label="是否启用" name="enabled">
        <a-switch v-model:checked="formState.enabled" checked-children="启用" un-checked-children="禁用" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { DEFAULT_PERMISSION_TYPES } from './common';
import { editAdminPermission } from '@/api/admin/permission';
import { AdminPerssionKey } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { PermissionListItem } from '@/types/admin_permission';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{
  open: boolean;
  detailData?: PermissionListItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);

const formState = reactive<{
  name: string;
  type: string;
  key: string;
  describe: string;
  enabled: boolean;
}>({
  name: '',
  type: 'view',
  key: '',
  describe: '',
  enabled: true,
});

const rules = {
  name: [
    { required: true, message: '请添加权限名称' },
    { max: 50, message: '名称长度不能超过50个字符' },
  ],
  key: [{ required: true, pattern: AdminPerssionKey, message: '请按照驼峰法命名' }],
};

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      formState.name = props.detailData.name || '';
      formState.type = props.detailData.type || 'view';
      formState.key = props.detailData.key || '';
      formState.describe = props.detailData.describe || '';
      formState.enabled = props.detailData.enabled ?? true;
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      editAdminPermission({
        id: props.detailData?.id,
        menuId: props.detailData?.menuId,
        name: formState.name,
        type: formState.type,
        key: formState.key,
        describe: formState.describe,
        enabled: formState.enabled,
      })
        .then((res) => {
          message.success(res.msg, 2);
          emit('notice');
          emit('update:open', false);
        })
        .finally(() => {
          confirmLoading.value = false;
        });
    })
    .catch(() => {});
}
</script>
