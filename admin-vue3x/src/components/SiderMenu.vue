<template>
  <div class="sider-menu">
    <div class="logo" @click="router.push('/home')">
      <img
        src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg"
        alt="logo"
      />
      <span v-show="!collapsed">Admin Vue3x</span>
    </div>
    <a-menu
      theme="dark"
      mode="inline"
      :selected-keys="selectedKeys"
      :open-keys="openKeys"
      @click="onMenuClick"
      @openChange="onOpenChange"
    >
      <template v-for="item in menuItems" :key="item.key">
        <a-menu-item v-if="!item.children || item.children.length === 0" :key="item.key">
          <router-link :to="item.path">
            <component :is="iconMap[item.icon || '']" v-if="item.icon && iconMap[item.icon]" />
            <span>{{ item.name }}</span>
          </router-link>
        </a-menu-item>
        <a-sub-menu v-else :key="item.key">
          <template #title>
            <component :is="iconMap[item.icon || '']" v-if="item.icon && iconMap[item.icon]" />
            <span>{{ item.name }}</span>
          </template>
          <a-menu-item v-for="child in item.children" :key="child.key">
            <router-link :to="child.path">
              <span>{{ child.name }}</span>
            </router-link>
          </a-menu-item>
        </a-sub-menu>
      </template>
    </a-menu>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
  HomeOutlined,
  SettingOutlined,
  UserOutlined,
  TeamOutlined,
  MenuOutlined,
  ApiOutlined,
} from '@ant-design/icons-vue';
import { useUserStore } from '@/store/user';
import { HandleRemoteMenuIntoLocal, localMenuData, type MenuConfigItem } from '@/router/menu';
import { HomePath } from '@/api/config';

defineProps<{ collapsed: boolean }>();

const route = useRoute();
const router = useRouter();
const store = useUserStore();

const iconMap: Record<string, any> = {
  HomeOutlined,
  SettingOutlined,
  UserOutlined,
  TeamOutlined,
  MenuOutlined,
  ApiOutlined,
};

/** 根据远程菜单过滤本地菜单 */
const menuItems = computed<MenuConfigItem[]>(() => {
  const remoteMenus = store.menus || {};
  let result: MenuConfigItem[] = HandleRemoteMenuIntoLocal([], localMenuData, remoteMenus, 'children');
  // 顶部固定“首页”
  const homeItem: MenuConfigItem = { key: 'Home', path: HomePath, name: '首页', icon: 'HomeOutlined' };
  result = [homeItem, ...result];
  return result;
});

const selectedKeys = computed(() => [String(route.name || '')]);
const openKeys = ref<string[]>([]);

watch(
  () => route.path,
  (path) => {
    // 自动展开当前菜单所在分组
    for (const item of menuItems.value) {
      if (item.children?.some((child) => child.path === path)) {
        if (!openKeys.value.includes(item.key)) {
          openKeys.value = [item.key];
        }
        break;
      }
    }
  },
  { immediate: true },
);

const onMenuClick = ({ key }: { key: string }) => {
  const item = menuItems.value.find((i) => i.key === key);
  if (item) {
    router.push(item.path);
  }
};

const onOpenChange = (keys: string[]) => {
  openKeys.value = keys as string[];
};
</script>

<style scoped lang="less">
.sider-menu {
  height: 100%;

  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 64px;
    color: #fff;
    font-size: 18px;
    font-weight: 600;
    cursor: pointer;

    img {
      width: 32px;
      height: 32px;
      margin-right: 8px;
    }
  }
}
</style>
