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
        <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 4 }">
          <a-form-item label="账号" name="username">
            <a-input v-model:value="formState.username" disabled />
          </a-form-item>
          <a-form-item label="昵称" name="nickname">
            <a-input v-model:value="formState.nickname" placeholder="请输入昵称" allow-clear />
          </a-form-item>
          <a-form-item label="邮箱" name="email">
            <a-input v-model:value="formState.email" placeholder="请输入邮箱" allow-clear />
          </a-form-item>
          <a-form-item label="头像" name="avatar">
            <a-input v-model:value="formState.avatar" placeholder="请输入头像图片地址" allow-clear>
              <template #addonAfter>
                <a-upload accept="image/*" :show-upload-list="false" :before-upload="onPickAvatar">
                  <a-button type="link" size="small" :loading="uploading" class="avatar-upload-btn">
                    点击上传
                  </a-button>
                </a-upload>
              </template>
            </a-input>
            <div class="avatar-tip">支持 jpg/jpeg/png/gif/webp/ico，大小不超过 10MB；上传后自动填入地址并预览</div>
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
import { uploadAdminFile } from '@/api/admin/upload';
import { AdminEmail } from '@/api/pattern';

/** 头像上传分组：与个人中心菜单路径前缀保持一致 */
const AVATAR_UPLOAD_GROUP = '/account/';
/** 允许的图片扩展名（与后端 upload.allowed_exts 保持一致） */
const IMAGE_EXTS = ['jpg', 'jpeg', 'png', 'gif', 'webp', 'ico'];
/** 图片大小上限（与后端 upload.max_size 保持一致） */
const IMAGE_MAX_SIZE = 10 * 1024 * 1024;

const store = useUserStore();
const loading = ref(false);
const saving = ref(false);
const uploading = ref(false);
const formRef = ref();

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

/** 选择图片后立即上传（before-upload 同步返回 false，仅拦截选择，不触发 antd 自动上传） */
function onPickAvatar(file: File) {
  const ext = (file.name.split('.').pop() || '').toLowerCase();
  if (!IMAGE_EXTS.includes(ext) || (file.type && !file.type.startsWith('image/'))) {
    message.error(`仅支持上传 ${IMAGE_EXTS.join('/')} 图片`);
    return false;
  }
  if (file.size > IMAGE_MAX_SIZE) {
    message.error('图片大小不能超过 10MB');
    return false;
  }
  void doUploadAvatar(file);
  return false;
}

async function doUploadAvatar(file: File) {
  uploading.value = true;
  try {
    const res = await uploadAdminFile({ file, uploadGroup: AVATAR_UPLOAD_GROUP });
    formState.avatar = res.data?.url || '';
    message.success('头像上传成功');
  } finally {
    uploading.value = false;
  }
}

async function onSave() {
  // 提交前先校验，避免必填项为空直接打到后端
  try {
    await formRef.value?.validate();
  } catch {
    // 校验不通过：错误提示已由表单展示
    return;
  }
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

    .avatar-upload-btn {
      height: 100%;
    }

    .avatar-tip {
      margin-top: 4px;
      color: rgba(0, 0, 0, 0.45);
      font-size: 12px;
      line-height: 1.5;
    }
  }
}
</style>
