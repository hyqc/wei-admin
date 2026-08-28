<template>
  <slot v-if="store.hasPermission(permission)" />
</template>

<script setup lang="ts">
import { useUserStore } from '@/store/user';

/** 按钮级权限控制：permission 为空时始终渲染 */
defineProps<{ permission?: string }>();

const store = useUserStore();
</script>
