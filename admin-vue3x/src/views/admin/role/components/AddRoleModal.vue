<template>
  <a-modal
    :open="open"
    title="新建角色"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }">
      <a-form-item label="角色名称" name="name">
        <a-input v-model:value="formState.name" placeholder="请输入角色名称" allow-clear />
      </a-form-item>
      <a-form-item label="描述" name="describe">
        <a-textarea v-model:value="formState.describe" placeholder="请输入角色描述" :rows="4" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { addAdminRole } from '@/api/admin/role';
import { DefaultModalWidth } from '@/api/config';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{ open: boolean }>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);
const formState = reactive<{ name: string; describe: string }>({
  name: '',
  describe: '',
});

const rules = {
  name: [
    { required: true, message: '请输入角色名称' },
    { max: 50, message: '角色名称长度不能超过50个字符' },
  ],
  describe: [{ max: 200, message: '描述长度不能超过200个字符' }],
};

watch(
  () => props.open,
  (val) => {
    if (val) {
      formRef.value?.resetFields();
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      addAdminRole({ ...formState })
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
