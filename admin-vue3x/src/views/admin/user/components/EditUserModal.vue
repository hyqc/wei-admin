<template>
  <a-modal
    :open="open"
    title="编辑账号"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }">
      <a-form-item label="账号">
        <a-input :value="detailData?.username" disabled />
      </a-form-item>
      <a-form-item label="昵称" name="nickname">
        <a-input v-model:value="formState.nickname" placeholder="请输入昵称" allow-clear />
      </a-form-item>
      <a-form-item label="邮箱" name="email">
        <a-input v-model:value="formState.email" placeholder="请输入邮箱" allow-clear />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { editAdminUser } from '@/api/admin/user';
import { AdminEmail } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{
  open: boolean;
  detailData?: AdminUserListItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);

const formState = reactive<{ nickname: string; email: string }>({
  nickname: '',
  email: '',
});

const rules = {
  nickname: [
    { required: true, message: '请输入昵称' },
    { max: 50, message: '昵称长度不能超过50个字符' },
  ],
  email: [
    { required: true, message: '请输入邮箱' },
    { pattern: AdminEmail, message: '邮箱格式不正确' },
  ],
};

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      formState.nickname = props.detailData.nickname || '';
      formState.email = props.detailData.email || '';
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      editAdminUser({
        adminId: props.detailData?.adminId as number,
        nickname: formState.nickname,
        email: formState.email,
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
