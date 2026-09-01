<template>
  <div class="login-page">
    <a-card class="login-card">
      <div class="login-title">
        <img
          src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg"
          alt="logo"
          class="login-logo"
        />
        <span>Admin Vue3x</span>
      </div>
      <div class="login-subtitle">ant design vue 后台管理系统</div>
      <a-form
        :model="formState"
        :rules="rules"
        name="login"
        size="large"
        @finish="onFinish"
        @finish-failed="onFinishFailed"
      >
        <a-form-item name="username">
          <a-input v-model:value="formState.username" placeholder="用户名：admin" allow-clear>
            <template #prefix><UserOutlined /></template>
          </a-input>
        </a-form-item>
        <a-form-item name="password">
          <a-input-password v-model:value="formState.password" placeholder="密码：123456">
            <template #prefix><LockOutlined /></template>
          </a-input-password>
        </a-form-item>
        <a-form-item name="captchaCode">
          <a-row :gutter="8">
            <a-col :span="15">
              <a-input
                v-model:value="formState.captchaCode"
                placeholder="请输入验证码"
                allow-clear
                :maxlength="8"
              >
                <template #prefix><SafetyOutlined /></template>
              </a-input>
            </a-col>
            <a-col :span="9">
              <img
                v-if="captcha.image"
                :src="captcha.image"
                class="captcha-img"
                alt="验证码"
                title="看不清？点击刷新"
                @click="refreshCaptcha"
              />
              <a-button v-else block @click="refreshCaptcha">获取验证码</a-button>
            </a-col>
          </a-row>
        </a-form-item>
        <a-form-item>
          <a-checkbox v-model:checked="formState.remember">自动登录</a-checkbox>
        </a-form-item>
        <a-form-item style="margin-bottom: 0">
          <a-button type="primary" html-type="submit" block :loading="loading">
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { UserOutlined, LockOutlined, SafetyOutlined } from '@ant-design/icons-vue';
import { useUserStore } from '@/store/user';
import { getCaptcha } from '@/api/admin/account';
import { AdminUsername, AdminUserPassword } from '@/api/pattern';
import { HomePath } from '@/api/config';
import type { ReqLogin } from '@/types/admin_account';

const route = useRoute();
const router = useRouter();
const store = useUserStore();
const loading = ref(false);

const formState = reactive<ReqLogin>({
  username: 'admin',
  password: '123456',
  remember: true,
  captchaCode: '',
});

/** 当前验证码：ID 随登录提交，图片点击可刷新 */
const captcha = reactive<{ captchaId: string; image: string }>({
  captchaId: '',
  image: '',
});

const rules = {
  username: [
    { required: true, message: '请输入用户名' },
    { pattern: AdminUsername, message: '用户名格式不正确' },
  ],
  password: [
    { required: true, message: '请输入密码' },
    { pattern: AdminUserPassword, message: '密码格式不正确' },
  ],
  captchaCode: [{ required: true, message: '请输入验证码' }],
};

async function refreshCaptcha() {
  const res = await getCaptcha();
  captcha.captchaId = res.data.captchaId;
  captcha.image = res.data.image;
  formState.captchaCode = '';
}

onMounted(refreshCaptcha);

async function onFinish() {
  loading.value = true;
  try {
    await store.loginAsync({ ...formState, captchaId: captcha.captchaId });
    const redirect = (route.query.redirect as string) || HomePath;
    router.replace(redirect);
  } catch {
    // 登录失败（账号密码错误或验证码失效）后刷新验证码
    await refreshCaptcha();
  } finally {
    loading.value = false;
  }
}

/** 登录失败（验证码错误/失效）时自动刷新验证码 */
function onFinishFailed() {
  refreshCaptcha();
}
</script>

<style scoped lang="less">
.login-page {
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-card {
  position: relative;
  z-index: 1;
  width: 420px;
  padding: 24px 32px;
}

.login-logo {
  width: 36px;
  height: 36px;
  margin-right: 12px;
}

.login-title {
  font-size: 30px;
}

.login-subtitle {
  margin-bottom: 32px;
}

.captcha-img {
  display: block;
  width: 100%;
  height: 40px;
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
}
</style>
