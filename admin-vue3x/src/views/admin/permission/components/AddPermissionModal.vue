<template>
  <a-modal
    :open="open"
    title="新建权限"
    :width="900"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-row :gutter="16">
      <a-col :span="10">
        <div class="page-menu-title">选择菜单页面</div>
        <div class="page-menu-body">
          <PageMenus ref="pageMenusRef" :value="formState.menuId" @change="onMenuChange" />
        </div>
      </a-col>
      <a-col :span="14">
        <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 16 }">
          <a-form-item label="权限名称" name="name">
            <a-input v-model:value="formState.name" placeholder="请输入权限名称" allow-clear />
          </a-form-item>
          <a-form-item label="权限类型" name="type">
            <a-radio-group v-model:value="formState.type" @change="onTypeChange">
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
      </a-col>
    </a-row>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import PageMenus from './PageMenus.vue';
import { DEFAULT_PERMISSION_TYPES, handleKey } from './common';
import { addAdminPermission } from '@/api/admin/permission';
import type { MenuTreeItem } from '@/types/admin_menu';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{ open: boolean }>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const pageMenusRef = ref();
const confirmLoading = ref(false);
const menuPath = ref('');

const formState = reactive<{
  menuId?: number;
  name: string;
  type: string;
  key: string;
  describe: string;
  enabled: boolean;
}>({
  menuId: undefined,
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
  key: [{ required: true, pattern: /^([A-Z][a-zA-Z0-9]*)+$/g, message: '请按照驼峰法命名' }],
};

function onMenuChange(node?: MenuTreeItem) {
  if (node) {
    menuPath.value = node.path || '';
    formState.key = handleKey(menuPath.value, formState.type);
  }
}

function onTypeChange() {
  formState.key = handleKey(menuPath.value, formState.type);
}

watch(
  () => formState.type,
  () => {
    if (formState.menuId) {
      formState.key = handleKey(menuPath.value, formState.type);
    }
  },
);

watch(
  () => props,
  (val) => {
    if (val.open) {
      formRef.value?.resetFields();
      formState.menuId = undefined;
      formState.name = '';
      formState.type = 'view';
      formState.key = '';
      formState.describe = '';
      formState.enabled = true;
      menuPath.value = '';
      pageMenusRef.value?.load();
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      addAdminPermission({
        menuId: formState.menuId,
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

function handleCancel() {
  emit('update:open', false);
}
</script>

<style scoped lang="less">
.page-menu-title {
  margin-bottom: 8px;
  font-weight: 600;
}

.page-menu-body {
  max-height: 360px;
  overflow: auto;
  border: 1px solid #f0f0f0;
  border-radius: 4px;
  padding: 8px;
}
</style>
