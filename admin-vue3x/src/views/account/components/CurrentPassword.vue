<template>
  <div class="current-password">
    <a-form
      :model="formState"
      :rules="rules"
      :label-col="{ span: 4 }"
      style="max-width: 480px"
    >
      <a-form-item label="原密码" name="oldPassword">
        <a-input-password v-model:value="formState.oldPassword" placeholder="请输入原密码" />
      </a-form-item>
      <a-form-item label="新密码" name="newPassword">
        <a-input-password v-model:value="formState.newPassword" placeholder="请输入新密码" />
      </a-form-item>
      <a-form-item label="确认密码" name="confirmPassword">
        <a-input-password v-model:value="formState.confirmPassword" placeholder="请再次输入新密码" />
      </a-form-item>
      <a-form-item :wrapper-col="{ offset: 4 }">
        <a-button type="primary" :loading="saving" @click="onSave">
          保存
        </a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import { currentAdminEditPassword } from '@/api/admin/account';
import { AdminUserPassword } from '@/api/pattern';

const saving = ref(false);

const formState = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: '',
});

const rules = {
  oldPassword: [{ required: true, message: '请输入原密码' }],
  newPassword: [
    { required: true, message: '请输入新密码' },
    { pattern: AdminUserPassword, message: '密码格式不正确' },
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码' },
    {
      validator: (_rule: unknown, value: string) => {
        if (!value || value === formState.newPassword) {
          return Promise.resolve();
        }
        return Promise.reject(new Error('两次输入的密码不一致'));
      },
    },
  ],
};

async function onSave() {
  saving.value = true;
  try {
    const res = await currentAdminEditPassword({
      oldPassword: formState.oldPassword,
      newPassword: formState.newPassword,
      confirmPassword: formState.confirmPassword,
    });
    message.success(res.msg, 2);
    formState.oldPassword = '';
    formState.newPassword = '';
    formState.confirmPassword = '';
  } finally {
    saving.value = false;
  }
}
</script>
