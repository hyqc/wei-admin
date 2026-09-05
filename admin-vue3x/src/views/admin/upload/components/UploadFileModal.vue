<template>
  <a-modal
    :open="open"
    title="上传文件"
    :width="560"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="开始上传"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-form :label-col="{ span: 5 }" :wrapper-col="{ span: 18 }">
      <a-form-item label="分组" required>
        <a-input v-model:value="group" placeholder="上传路径前缀，如 /admin/user/" allow-clear />
        <div class="upload-tip">
          分组即上传路径前缀，通常与所属菜单路径保持一致；存储目录结构为「分组/年/月/文件」
        </div>
      </a-form-item>
      <a-form-item label="文件" required>
        <a-upload
          :max-count="1"
          :file-list="fileList"
          :before-upload="beforeUpload"
          @remove="onRemove"
        >
          <a-button>
            <template #icon><UploadOutlined /></template>
            选择文件
          </a-button>
        </a-upload>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { UploadOutlined } from '@ant-design/icons-vue';
import type { UploadProps } from 'ant-design-vue';
import { uploadAdminFile } from '@/api/admin/upload';

const props = defineProps<{ open: boolean }>();
const emit = defineEmits<{
  (e: 'update:open', value: boolean): void;
  (e: 'notice'): void;
}>();

const group = ref('');
const file = ref<File>();
const fileList = ref<UploadProps['fileList']>([]);
const confirmLoading = ref(false);

watch(
  () => props.open,
  (val) => {
    if (val) {
      group.value = '';
      file.value = undefined;
      fileList.value = [];
    }
  },
);

/** 不自动上传：由父级弹窗统一提交，便于先校验分组 */
function beforeUpload(uploadFile: File) {
  file.value = uploadFile;
  fileList.value = [{ uid: '1', name: uploadFile.name, status: 'done' }];
  return false;
}

function onRemove() {
  file.value = undefined;
  fileList.value = [];
}

async function handleOk() {
  const groupVal = group.value.trim();
  if (!groupVal) {
    message.warning('请填写分组（上传路径前缀）');
    return;
  }
  if (!file.value) {
    message.warning('请选择要上传的文件');
    return;
  }
  confirmLoading.value = true;
  try {
    const res = await uploadAdminFile({ file: file.value, uploadGroup: groupVal });
    message.success(res.msg, 2);
    emit('notice');
    emit('update:open', false);
  } finally {
    confirmLoading.value = false;
  }
}
</script>

<style scoped lang="less">
.upload-tip {
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
  line-height: 1.5;
}
</style>
