<template>
  <a-switch
    :checked="isEnabled"
    :disabled="isEnabledButtonDisabled"
    checked-children="启用"
    un-checked-children="禁用"
    @change="onChange"
  />
</template>

<script setup lang="ts">
defineProps<{
  isEnabled?: boolean;
  isEnabledButtonDisabled?: boolean;
}>();

const emit = defineEmits<{
  (e: 'change', checked: boolean): void;
}>();

const onChange = (checked: boolean) => {
  emit('change', checked);
};
</script>
