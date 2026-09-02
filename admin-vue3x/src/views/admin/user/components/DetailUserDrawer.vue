<template>
  <a-drawer
    :open="open"
    title="账号详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-spin :spinning="loading">
      <a-descriptions :column="1" bordered>
        <a-descriptions-item label="账号">{{ detail?.username }}</a-descriptions-item>
      <a-descriptions-item label="昵称">{{ detail?.nickname }}</a-descriptions-item>
      <a-descriptions-item label="邮箱">{{ detail?.email }}</a-descriptions-item>
      <a-descriptions-item label="角色">
        <!-- 超管无需绑定角色，自动拥有系统全部权限 -->
        <a-tag v-if="detail?.isSuperAdmin" color="gold">超级管理员</a-tag>
        <template v-else-if="detail?.roles?.length">
          <a-tag v-for="role in detail?.roles" :key="role.roleId" color="blue">
            {{ role.roleName }}
          </a-tag>
        </template>
        <span v-else>-</span>
      </a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-badge :status="detail?.isEnabled ? 'success' : 'error'" :text="detail?.isEnabled ? '启用' : '禁用'" />
      </a-descriptions-item>
      <a-descriptions-item label="登录次数">{{ detail?.loginTotal }}</a-descriptions-item>
      <a-descriptions-item label="上次登录IP">{{ detail?.lastLoginIp || '-' }}</a-descriptions-item>
      <a-descriptions-item label="本次登录IP">{{ detail?.currentLoginIp || '-' }}</a-descriptions-item>
      <a-descriptions-item label="本次登录时间">{{ detail?.currentLoginTime || '-' }}</a-descriptions-item>
      <a-descriptions-item label="创建时间">{{ detail?.createdAt }}</a-descriptions-item>
      <a-descriptions-item label="更新时间">{{ detail?.updatedAt }}</a-descriptions-item>
      </a-descriptions>
    </a-spin>
  </a-drawer>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { DefaultDrawerWidth } from '@/api/config';
import { getAdminUserInfo } from '@/api/admin/user';
import type { AdminUserListItem } from '@/types/common';

const props = defineProps<{
  open: boolean;
  detailData?: AdminUserListItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const loading = ref(false);
const detail = ref<AdminUserListItem>();

// 打开时实时拉取详情，避免展示列表中的过期数据
watch(
  () => props.open,
  (val) => {
    if (val && props.detailData?.adminId) {
      loading.value = true;
      detail.value = undefined;
      getAdminUserInfo({ adminId: props.detailData.adminId })
        .then((res) => {
          detail.value = res.data;
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);
</script>
