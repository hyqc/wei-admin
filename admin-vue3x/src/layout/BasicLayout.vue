<template>
  <a-layout class="basic-layout">
    <a-layout-sider
      v-model:collapsed="collapsed"
      :trigger="null"
      collapsible
      :width="220"
      class="layout-sider"
    >
      <SiderMenu :collapsed="collapsed" />
    </a-layout-sider>
    <a-layout class="basic-layout-inner">
      <a-layout-header class="layout-header">
        <div class="header-left">
          <MenuUnfoldOutlined
            v-if="collapsed"
            class="trigger"
            @click="collapsed = !collapsed"
          />
          <MenuFoldOutlined v-else class="trigger" @click="collapsed = !collapsed" />
        </div>
        <div class="header-right">
          <HeaderRight />
        </div>
      </a-layout-header>
      <a-layout-content class="layout-content">
        <router-view />
      </a-layout-content>
      <a-layout-footer class="layout-footer">
        <FooterBar />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons-vue';
import SiderMenu from '@/components/SiderMenu.vue';
import HeaderRight from '@/components/HeaderRight.vue';
import FooterBar from '@/components/FooterBar.vue';

const collapsed = ref(false);
</script>

<style scoped lang="less">
.basic-layout {
  height: 100vh;

  .layout-sider {
    background: #001529;
  }

  .basic-layout-inner {
    height: 100vh;
  }

  .layout-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 48px;
    line-height: 48px;
    padding: 0 16px;
    background: #fff;
    box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

    .trigger {
      font-size: 18px;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }
  }

  .layout-content {
    margin: 16px;
    overflow: auto;
  }

  .layout-footer {
    padding: 12px 16px;
    text-align: center;
  }
}
</style>
