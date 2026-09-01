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
        <!-- 超管无需绑定角色，自动拥有系统全部权限 -->
        <a-tag v-if="detailData?.isSuperAdmin" color="gold">超级管理员</a-tag>
        <template v-else-if="detailData?.roles?.length">
          <a-tag v-for="role in detailData?.roles" :key="role.roleId" color="blue">
            {{ role.roleName }}
          </a-tag>
        </template>
        <span v-else>-</span>
      </a-descriptions-item>
      <a-descriptions-item label="状态">
        <a-badge :status="detailData?.isEnabled ? 'success' : 'error'" :text="detailData?.isEnabled ? '启用' : '禁用'" />
      </a-descriptions-item>
      <a-descriptions-item label="登录次数">{{ detailData?.loginTotal }}</a-descriptions-item>
      <a-descriptions-item label="上次登录IP">{{ detailData?.lastLoginIp || '-' }}</a-descriptions-item>
      <a-descriptions-item label="本次登录IP">{{ detailData?.currentLoginIp || '-' }}</a-descriptions-item>
      <a-descriptions-item label="本次登录时间">{{ detailData?.currentLoginTime || '-' }}</a-descriptions-item>
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
