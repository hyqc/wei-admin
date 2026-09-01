<template>
  <a-modal
    :open="open"
    title="新建账号"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="{ span: 6 }"
      :wrapper-col="{ span: 12 }"
      autocomplete="off"
    >
      <a-form-item label="账号" name="username">
        <!-- 禁用浏览器自动填充，避免带入已保存的登录账号密码 -->
        <a-input
          v-model:value="formState.username"
          placeholder="请输入账号"
          allow-clear
          autocomplete="off"
        />
      </a-form-item>
      <a-form-item label="密码" name="password">
        <a-input-password
          v-model:value="formState.password"
          placeholder="请输入密码"
          autocomplete="new-password"
        />
      </a-form-item>
      <a-form-item label="昵称" name="nickname">
        <a-input v-model:value="formState.nickname" placeholder="请输入昵称" allow-clear />
      </a-form-item>
      <a-form-item label="邮箱" name="email">
        <a-input v-model:value="formState.email" placeholder="请输入邮箱" allow-clear />
      </a-form-item>
      <a-form-item label="角色" name="roleIds">
        <a-select v-model:value="formState.roleIds" mode="multiple" placeholder="请选择角色" allow-clear>
          <a-select-option v-for="role in roleOptions" :key="role.id" :value="role.id">
            {{ role.name }}
          </a-select-option>
        </a-select>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { addAdminUser } from '@/api/admin/user';
import { getAdminRoleAll } from '@/api/admin/role';
import { AdminUsername, AdminUserPassword, AdminEmail } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { RoleItem } from '@/types/admin_role';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{ open: boolean }>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);
const roleOptions = ref<RoleItem[]>([]);

const formState = reactive<{
  username: string;
  password: string;
  nickname: string;
  email: string;
  roleIds: number[];
}>({
  username: '',
  password: '',
  nickname: '',
  email: '',
  roleIds: [],
});

const rules = {
  username: [
    { required: true, message: '请输入账号' },
    { pattern: AdminUsername, message: '账号格式不正确' },
  ],
  password: [
    { required: true, message: '请输入密码' },
    { pattern: AdminUserPassword, message: '密码格式不正确' },
  ],
  nickname: [{ max: 50, message: '昵称长度不能超过50个字符' }],
  email: [{ pattern: AdminEmail, message: '邮箱格式不正确' }],
};

watch(
  () => props.open,
  (val) => {
    if (val) {
      formRef.value?.resetFields();
      loadRoles();
    }
  },
);

function loadRoles() {
  getAdminRoleAll().then((res) => {
    roleOptions.value = res.data || [];
  });
}

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      addAdminUser({ ...formState })
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
