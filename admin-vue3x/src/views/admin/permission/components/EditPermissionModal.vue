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
        <a-input :value="detailMenuName" disabled />
      </a-form-item>
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
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { DEFAULT_PERMISSION_TYPES, handleKey } from './common';
import { editAdminPermission, getAdminPermissionInfo } from '@/api/admin/permission';
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

// 打开时实时拉取详情回填，避免编辑列表中的过期数据
watch(
  () => props.open,
  async (val) => {
    if (val && props.detailData?.id) {
      const res = await getAdminPermissionInfo({ id: props.detailData.id });
      detailPath.value = res.data.menuPath || '';
      detailMenuName.value = res.data.menuName || '';
      formState.name = res.data.name || '';
      formState.type = res.data.type || 'view';
      formState.key = res.data.key || '';
      formState.describe = res.data.describe || '';
      formState.enabled = res.data.isEnabled ?? true;
    }
  },
);

/** 服务端返回的菜单路径（用于切换类型时重新生成唯一键） */
const detailPath = ref('');
/** 服务端返回的菜单名称 */
const detailMenuName = ref('');

/** 切换权限类型时，唯一键跟随菜单路径 + 类型重新生成 */
function onTypeChange() {
  const menuPath = detailPath.value || props.detailData?.menuPath || '';
  if (menuPath) {
    formState.key = handleKey(menuPath, formState.type);
  }
}

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
