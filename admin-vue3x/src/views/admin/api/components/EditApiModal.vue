<template>
  <a-modal
    :open="open"
    title="编辑接口"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 14 }">
      <a-form-item label="接口名称" name="name">
        <a-input v-model:value="formState.name" placeholder="请输入接口名称" allow-clear />
      </a-form-item>
      <a-form-item label="唯一键" name="key">
        <a-input v-model:value="formState.key" placeholder="示例：adminUser::list" allow-clear />
      </a-form-item>
      <a-form-item label="接口路径" name="path">
        <a-input v-model:value="formState.path" placeholder="示例：/admin/user/list" allow-clear />
      </a-form-item>
      <a-form-item label="描述" name="describe">
        <a-textarea v-model:value="formState.describe" placeholder="请输入接口描述" :rows="3" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { editAdminApi } from '@/api/admin/api';
import { AdminAPIKey } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { AdminApiItem } from '@/types/common';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{
  open: boolean;
  detailData?: AdminApiItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);

const formState = reactive<{ name: string; key: string; path: string; describe: string }>({
  name: '',
  key: '',
  path: '',
  describe: '',
});

const rules = {
  name: [
    { required: true, message: '请输入接口名称' },
    { max: 50, message: '名称长度不能超过50个字符' },
  ],
  key: [
    { required: true, message: '请输入唯一键' },
    { pattern: AdminAPIKey, message: '唯一键格式不正确（示例：adminUser::list）' },
  ],
  path: [{ required: true, message: '请输入接口路径' }],
  describe: [{ max: 200, message: '描述长度不能超过200个字符' }],
};

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      formState.name = props.detailData.name || '';
      formState.key = props.detailData.key || '';
      formState.path = props.detailData.path || '';
      formState.describe = props.detailData.describe || '';
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      editAdminApi({
        id: props.detailData?.id as number,
        name: formState.name,
        key: formState.key,
        path: formState.path,
        describe: formState.describe,
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
