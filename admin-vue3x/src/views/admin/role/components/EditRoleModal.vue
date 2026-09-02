<template>
  <a-modal
    :open="open"
    title="编辑角色"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="emit('update:open', false)"
  >
    <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }">
      <a-form-item label="角色名称" name="name">
        <a-input
          v-model:value="formState.name"
          placeholder="请输入角色名称"
          allow-clear
          :disabled="isSuperAdmin"
        />
        <div v-if="isSuperAdmin" class="form-tip">超级管理员角色名称不可修改，仅可修改描述</div>
      </a-form-item>
      <a-form-item label="描述" name="describe">
        <a-textarea v-model:value="formState.describe" placeholder="请输入角色描述" :rows="4" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { editAdminRole, getAdminRoleInfo } from '@/api/admin/role';
import { DefaultModalWidth } from '@/api/config';
import type { RoleItem } from '@/types/admin_role';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{
  open: boolean;
  detailData?: RoleItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);
const formState = reactive<{ name: string; describe: string }>({
  name: '',
  describe: '',
});

/** 超管角色仅允许修改描述（以服务端最新数据为准） */
const isSuperAdmin = ref(false);

const rules = {
  name: [
    { required: true, message: '请输入角色名称' },
    { max: 50, message: '角色名称长度不能超过50个字符' },
  ],
  describe: [{ max: 200, message: '描述长度不能超过200个字符' }],
};

// 打开时实时拉取详情回填，避免编辑列表中的过期数据
watch(
  () => props.open,
  async (val) => {
    if (val && props.detailData?.id) {
      const res = await getAdminRoleInfo({ id: props.detailData.id });
      formState.name = res.data.name || '';
      formState.describe = res.data.describe || '';
      isSuperAdmin.value = !!res.data.isSuperAdmin;
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      confirmLoading.value = true;
      editAdminRole({
        id: props.detailData?.id as number,
        name: formState.name,
        describe: formState.describe,
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
</script>

<style scoped lang="less">
.form-tip {
  margin-top: 4px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
}
</style>
