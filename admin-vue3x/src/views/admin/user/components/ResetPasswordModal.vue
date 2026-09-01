<template>
  <a-modal
    :open="open"
    title="重置密码"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="{ span: 6 }"
      :wrapper-col="{ span: 12 }"
      autocomplete="off"
    >
      <a-form-item label="账号">
        <a-input :value="detailData?.username" disabled autocomplete="off" />
      </a-form-item>
      <a-form-item label="新密码" name="password">
        <!-- 禁用浏览器自动填充，避免带入已保存的登录密码 -->
        <a-input-password
          v-model:value="formState.password"
          placeholder="请输入新密码"
          autocomplete="new-password"
        />
      </a-form-item>
      <a-form-item label="确认密码" name="confirmPassword">
        <a-input-password
          v-model:value="formState.confirmPassword"
          placeholder="请再次输入新密码"
          autocomplete="new-password"
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { resetAdminUserPassword } from '@/api/admin/user';
import { AdminUserPassword } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{
  open: boolean;
  detailData?: AdminUserListItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);
const formState = reactive<{ password: string; confirmPassword: string }>({
  password: '',
  confirmPassword: '',
});

const rules = {
  password: [
    { required: true, message: '请输入新密码' },
    { pattern: AdminUserPassword, message: '密码格式不正确' },
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码' },
    {
      validator: (_rule: unknown, value: string) => {
        if (!value || value === formState.password) {
          return Promise.resolve();
        }
        return Promise.reject(new Error('两次输入的密码不一致'));
      },
    },
  ],
};

watch(
  () => props.open,
  (val) => {
    if (val) {
      formState.password = '';
      formState.confirmPassword = '';
      formRef.value?.clearValidate();
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      resetAdminUserPassword({ adminId: props.detailData?.adminId, password: formState.password })
        .then((res) => {
          message.success(res.msg, 2);
          emit('update:open', false);
        })
        .finally(() => {
          confirmLoading.value = false;
        });
    })
    .catch(() => {});
}
</script>
