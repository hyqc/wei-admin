<template>
  <a-spin :spinning="loading">
    <div class="current-account">
      <div class="account-left">
        <a-avatar :size="96" :src="formState.avatar" class="account-avatar">
          {{ displayName[0] }}
        </a-avatar>
        <div class="account-nickname">{{ formState.nickname }}</div>
      </div>
      <div class="account-right">
        <a-form :model="formState" :rules="rules" :label-col="{ span: 4 }">
          <a-form-item label="账号" name="username">
            <a-input v-model:value="formState.username" disabled />
          </a-form-item>
          <a-form-item label="昵称" name="nickname">
            <a-input v-model:value="formState.nickname" placeholder="请输入昵称" allow-clear />
          </a-form-item>
          <a-form-item label="邮箱" name="email">
            <a-input v-model:value="formState.email" placeholder="请输入邮箱" allow-clear />
          </a-form-item>
          <a-form-item label="头像">
            <a-input v-model:value="formState.avatar" placeholder="请输入头像地址" allow-clear />
          </a-form-item>
          <a-form-item :wrapper-col="{ offset: 4 }">
            <a-button type="primary" :loading="saving" @click="onSave">
              保存
            </a-button>
          </a-form-item>
        </a-form>
      </div>
    </div>
  </a-spin>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import { useUserStore } from '@/store/user';
import { currentAdminEdit } from '@/api/admin/account';
import { AdminEmail } from '@/api/pattern';

const store = useUserStore();
const loading = ref(false);
const saving = ref(false);

const formState = reactive({
  adminId: store.userInfo?.adminId,
  username: store.userInfo?.username || '',
  nickname: store.userInfo?.nickname || '',
  email: store.userInfo?.email || '',
  avatar: store.userInfo?.avatar || '',
});

const displayName = formState.nickname || formState.username || '?';

const rules = {
  nickname: [
    { required: true, message: '请输入昵称' },
    { max: 50, message: '昵称长度不能超过50个字符' },
  ],
  email: [
    {
      // 邮箱非必填：未填写时跳过校验，填写后才校验格式
      validator: (_rule: unknown, value: string) => {
        if (!value || AdminEmail.test(value)) {
          return Promise.resolve();
        }
        return Promise.reject(new Error('邮箱格式不正确'));
      },
    },
  ],
};

async function onSave() {
  saving.value = true;
  try {
    const res = await currentAdminEdit({
      adminId: formState.adminId,
      nickname: formState.nickname,
      email: formState.email,
      avatar: formState.avatar,
    });
    // 更新本地用户信息
    if (store.userInfo) {
      store.setCurrentUser({
        ...store.userInfo,
        nickname: formState.nickname,
        email: formState.email,
        avatar: formState.avatar,
      });
    }
    message.success(res.msg, 2);
  } finally {
    saving.value = false;
  }
}
</script>

<style scoped lang="less">
.current-account {
  display: flex;
  padding: 24px;

  .account-left {
    flex: 0 0 160px;
    display: flex;
    flex-direction: column;
    align-items: center;

    .account-avatar {
      background: #1677ff;
      font-size: 36px;
    }

    .account-nickname {
      margin-top: 12px;
      font-size: 16px;
      font-weight: 600;
    }
  }

  .account-right {
    flex: 1;
  }
}
</style>
