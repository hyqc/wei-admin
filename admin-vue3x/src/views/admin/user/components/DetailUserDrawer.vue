<template>
  <a-drawer
    :open="open"
    title="账号详情"
    :width="DefaultDrawerWidth"
    @close="emit('update:open', false)"
  >
    <a-descriptions :column="1" bordered>
      <a-descriptions-item label="账号">{{ detailData?.username }}</a-descriptions-item>
      <a-descriptions-item label="昵称">{{ detailData?.nickname }}</a-descriptions-item>
      <a-descriptions-item label="邮箱">{{ detailData?.email }}</a-descriptions-item>
      <a-descriptions-item label="角色">
        <a-tag v-for="role in detailData?.roles" :key="role.roleId" color="blue">
          {{ role.roleName }}
        </a-tag>
      </a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-badge :status="detailData?.isEnabled ? 'success' : 'error'" :text="detailData?.isEnabled ? '启用' : '禁用'" />
      </a-descriptions-item>
      <a-descriptions-item label="登录次数">{{ detailData?.loginTotal }}</a-descriptions-item>
      <a-descriptions-item label="上次登录IP">{{ detailData?.lastLoginIp }}</a-descriptions-item>
      <a-descriptions-item label="上次登录时间">{{ detailData?.lastLoginTime }}</a-descriptions-item>
      <a-descriptions-item label="创建时间">{{ detailData?.createdAt }}</a-descriptions-item>
      <a-descriptions-item label="更新时间">{{ detailData?.updatedAt }}</a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script setup lang="ts">
import { DefaultDrawerWidth } from '@/api/config';
import type { AdminUserListItem } from '@/types/common';

defineProps<{
  open: boolean;
  detailData?: AdminUserListItem;
}>();

const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();
</script>
