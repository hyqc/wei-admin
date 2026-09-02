<template>
  <div class="tabs-view">
    <a-tabs
      v-model:active-key="activeKey"
      type="editable-card"
      hide-add
      size="small"
      class="tabs"
      @edit="onEdit"
      @change="onChange"
    >
      <a-tab-pane v-for="tab in tabs" :key="tab.path" :closable="tab.closable">
        <template #tab>
          <a-dropdown :trigger="['contextmenu']">
            <span class="tab-title">{{ tab.title }}</span>
            <template #overlay>
              <a-menu @click="handleMenuClick($event, tab.path)">
                <a-menu-item key="current" :disabled="!tab.closable">
                  <template #icon><CloseOutlined /></template>
                  关闭当前
                </a-menu-item>
                <a-menu-item key="other">
                  <template #icon><ColumnWidthOutlined /></template>
                  关闭其他
                </a-menu-item>
                <a-menu-item key="all">
                  <template #icon><MinusOutlined /></template>
                  关闭全部
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </template>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { CloseOutlined, ColumnWidthOutlined, MinusOutlined } from '@ant-design/icons-vue';
import type { MenuProps } from 'ant-design-vue';
import { HomePath } from '@/api/config';
import { useTabsStore } from '@/store/tabs';

const route = useRoute();
const router = useRouter();
const tabsStore = useTabsStore();

const tabs = computed(() => tabsStore.tabs);
const activeKey = ref(route.path);

// 路由变化（点击菜单/切换页签/前进后退）时同步页签，不存在则新增
watch(
  () => route.path,
  (path) => {
    activeKey.value = path;
    tabsStore.addTab({
      path,
      title: (route.meta?.title as string) || String(route.name ?? path),
      closable: path !== HomePath,
    });
  },
  { immediate: true },
);

/** 右键菜单点击：转发为 key + 页签路径 */
function handleMenuClick(info: Parameters<NonNullable<MenuProps['onClick']>>[0], path: string) {
  onMenuClick(String(info.key), path);
}

function onChange(key: string | number) {
  const path = String(key);
  if (path !== route.path) {
    router.push(path);
  }
}

function onEdit(targetKey: string | number, action: string) {
  if (action === 'remove') {
    closeTab(String(targetKey));
  }
}

function closeTab(path: string) {
  const nextPath = tabsStore.getNeighborPath(path, route.path);
  tabsStore.removeTab(path);
  if (nextPath) {
    router.push(nextPath);
  }
}

function onMenuClick(key: string, path: string) {
  if (key === 'current') {
    closeTab(path);
    return;
  }
  if (key === 'other') {
    tabsStore.removeOther(path);
    if (route.path !== path) {
      router.push(path);
    }
    return;
  }
  if (key === 'all') {
    tabsStore.removeAll();
    if (route.path !== HomePath) {
      router.push(HomePath);
    }
  }
}
</script>

<style scoped lang="less">
.tabs-view {
  padding: 6px 8px 0;
  margin-bottom: 8px;
  background: #fff;
  border-radius: 2px;

  .tab-title {
    user-select: none;
  }

  :deep(.ant-tabs-nav) {
    margin-bottom: 0;
  }
}
</style>
