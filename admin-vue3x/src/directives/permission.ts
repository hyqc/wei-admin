import type { Directive } from 'vue';
import { useUserStore } from '@/store/user';

/** 按钮级权限指令：v-permission="'AdminUserAdd'" */
export const permission: Directive = {
  mounted(el, binding) {
    const key = binding.value as string | undefined;
    if (!key) return;
    const store = useUserStore();
    if (!store.hasPermission(key)) {
      el.parentNode?.removeChild(el);
    }
  },
  updated(el, binding) {
    const key = binding.value as string | undefined;
    if (!key) return;
    const store = useUserStore();
    if (!store.hasPermission(key)) {
      el.parentNode?.removeChild(el);
    }
  },
};
