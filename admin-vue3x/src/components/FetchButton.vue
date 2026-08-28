<template>
  <a-spin v-if="loading" :size="size" />
</template>

<script setup lang="ts">
defineProps<{ loading?: boolean; size?: 'small' | 'default' | 'large' }>();
</script>
