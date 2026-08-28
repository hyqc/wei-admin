<template>
  <div class="header-right">
    <a-dropdown placement="bottomRight">
      <div class="account">
        <a-avatar :src="store.userInfo?.avatar" :size="28" class="avatar">
          {{ (store.userInfo?.nickname || store.userInfo?.username || '?')[0] }}
        </a-avatar>
        <span class="name">{{ store.userInfo?.nickname || store.userInfo?.username }}</span>
      </div>
      <template #overlay>
        <a-menu @click="onMenuClick">
          <a-menu-item key="account">
            <UserOutlined />
            个人中心
          </a-menu-item>
          <a-menu-divider />
          <a-menu-item key="logout">
            <LogoutOutlined />
            退出登录
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { UserOutlined, LogoutOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import { useUserStore } from '@/store/user';
import { LoginPath } from '@/api/config';

const router = useRouter();
const store = useUserStore();

const onMenuClick = async ({ key }: { key: string }) => {
  if (key === 'account') {
    router.push('/account');
  } else if (key === 'logout') {
    await store.logoutAsync();
    message.success('退出登录成功');
    router.push(LoginPath);
  }
};
</script>

<style scoped lang="less">
.header-right {
  .account {
    display: flex;
    align-items: center;
    cursor: pointer;

    .avatar {
      background: #1677ff;
    }

    .name {
      margin-left: 8px;
      color: rgba(0, 0, 0, 0.85);
    }
  }
}
</style>
