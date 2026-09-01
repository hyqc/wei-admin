<template>
  <a-select
    :value="value"
    :placeholder="placeholder"
    :disabled="disabled"
    :list-height="320"
    show-search
    allow-clear
    :filter-option="filterOption"
    @change="onChange"
  >
    <a-select-option v-for="name in antIconNames" :key="name" :value="name">
      <span class="icon-option">
        <component :is="antIcons[name]" />
        <span class="icon-option-name">{{ name }}</span>
      </span>
    </a-select-option>
  </a-select>
</template>

<script setup lang="ts">
import { antIconNames, antIcons } from '@/utils/icon';

defineProps<{
  value?: string;
  placeholder?: string;
  disabled?: boolean;
}>();
const emit = defineEmits<{
  (e: 'update:value', value?: string): void;
  (e: 'change', value?: string): void;
}>();

/** 按图标名称搜索（忽略大小写） */
const filterOption = (input: string, option: any) =>
  String(option?.value ?? '')
    .toLowerCase()
    .includes(input.toLowerCase());

function onChange(val?: string) {
  emit('update:value', val);
  emit('change', val);
}
</script>

<style scoped lang="less">
.icon-option {
  display: inline-flex;
  align-items: center;
  gap: 8px;
}
</style>
