<template>
  <div class="account-page">
    <a-card>
      <a-tabs v-model:active-key="activeKey">
        <a-tab-pane key="account" tab="个人中心">
          <CurrentAccount />
        </a-tab-pane>
        <a-tab-pane key="password" tab="修改密码">
          <CurrentPassword />
        </a-tab-pane>
      </a-tabs>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import CurrentAccount from './components/CurrentAccount.vue';
import CurrentPassword from './components/CurrentPassword.vue';

const activeKey = ref('account');
</script>

<style scoped lang="less">
.account-page {
  max-width: 800px;
}
</style>
