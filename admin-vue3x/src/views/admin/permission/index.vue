<template>
  <PageContainer :page-info="pageInfo" @page-change="onPageChange" @page-size-change="onPageSizeChange">
    <template #searchArea>
      <a-form layout="inline" :model="searchForm" class="search-form" @finish="onSearch">
        <a-form-item label="菜单" name="menuId">
          <a-tree-select
            v-model:value="searchForm.menuId"
            :tree-data="menuTreeOptions"
            tree-default-expand-all
            placeholder="请选择菜单"
            allow-clear
            style="width: 180px"
          />
        </a-form-item>
        <a-form-item label="唯一键" name="key">
          <a-input v-model:value="searchForm.key" placeholder="请输入唯一键" allow-clear />
        </a-form-item>
        <a-form-item label="名称" name="name">
          <a-input v-model:value="searchForm.name" placeholder="请输入权限名称" allow-clear />
        </a-form-item>
        <a-form-item label="类型" name="type">
          <a-select
            v-model:value="searchForm.type"
            :options="PERMISSION_TYPE_OPTIONS"
            placeholder="全部"
            allow-clear
            style="width: 120px"
          />
        </a-form-item>
        <a-form-item label="接口" name="apiId">
          <a-select
            v-model:value="searchForm.apiId"
            placeholder="全部"
            allow-clear
            show-search
            option-filter-prop="label"
            :options="apiOptions"
            style="width: 240px"
          />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" html-type="submit">
              <template #icon><SearchOutlined /></template>
              查询
            </a-button>
            <a-button @click="onReset">
              <template #icon><ReloadOutlined /></template>
              重置
            </a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </template>
    <template #extra>
      <span class="muted">权限点及其接口绑定由各菜单的“权限配置”统一维护，此页仅用于查询与审计</span>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      row-key="id"
      :scroll="{ x: 1250 }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'menuName'">
          {{ record.menuName || '-' }}
        </template>
        <template v-else-if="column.key === 'type'">
          <a-tag color="blue">{{ record.typeText }}</a-tag>
        </template>
        <template v-else-if="column.key === 'apis'">
          <a-tag v-if="record.apis?.length" color="green">{{ record.apis.length }} 个</a-tag>
          <span v-else class="muted">未绑定</span>
        </template>
        <template v-else-if="column.key === 'enabled'">
          <RowEnabledButton :is-enabled="record.isEnabled" />
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <Authorization permission="AdminPermissionView">
              <a-button type="link" size="small" @click="openDetailModal(record)">
                <template #icon><EyeOutlined /></template>
                详情
              </a-button>
            </Authorization>
            <Authorization permission="AdminMenuEdit">
              <a-button type="link" size="small" @click="goMenuPermission">
                <template #icon><SettingOutlined /></template>
                去配置
              </a-button>
            </Authorization>
          </a-space>
        </template>
      </template>
    </a-table>

    <DetailPermissionDrawer v-model:open="detailModalStatus" :detail-data="detailData" />
  </PageContainer>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { SearchOutlined, ReloadOutlined, EyeOutlined, SettingOutlined } from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import RowEnabledButton from '@/components/RowEnabledButton.vue';
import DetailPermissionDrawer from './components/DetailPermissionDrawer.vue';
import { PERMISSION_TYPE_OPTIONS } from './components/common';
import { getAdminPermissionList } from '@/api/admin/permission';
import { getAdminMenuTree } from '@/api/admin/menu';
import { getAdminApiAll } from '@/api/admin/api';
import type { PermissionListItem } from '@/types/admin_permission';
import type { MenuTreeItem } from '@/types/admin_menu';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const loading = ref(false);
const rows = ref<PermissionListItem[]>([]);
const pageInfo = reactive<PageInfoType>({ ...DEFAULT_PAGE_INFO });
const menuTreeOptions = ref<any[]>([]);
const apiOptions = ref<{ label: string; value: number }[]>([]);
const router = useRouter();

const searchForm = reactive<{
  menuId?: number;
  key?: string;
  name?: string;
  type?: string;
  apiId?: number;
}>({
  menuId: undefined,
  key: undefined,
  name: undefined,
  type: undefined,
  apiId: undefined,
});

const detailModalStatus = ref(false);
const detailData = ref<PermissionListItem>();

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
  { title: '权限名称', dataIndex: 'name', key: 'name', width: 140 },
  { title: '菜单ID', dataIndex: 'menuId', key: 'menuId', width: 80 },
  { title: '所属菜单', dataIndex: 'menuName', key: 'menuName', width: 140 },
  { title: '唯一键', dataIndex: 'key', key: 'key', width: 180 },
  { title: '类型', dataIndex: 'type', key: 'type', width: 100 },
  { title: '接口数', dataIndex: 'apis', key: 'apis', width: 100 },
  { title: '状态', dataIndex: 'enabled', key: 'enabled', width: 100 },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 180 },
];

function loadMenuTree() {
  getAdminMenuTree().then((res) => {
    const build = (list: MenuTreeItem[]): any[] =>
      list.map((item) => ({
        key: item.id,
        value: item.id,
        title: item.name,
        children: item.children?.length ? build(item.children) : undefined,
      }));
    menuTreeOptions.value = build(res.data.list || []);
  });
}

/** 接口下拉：用于反查“绑定了某接口的权限点” */
function loadApiOptions() {
  getAdminApiAll().then((res) => {
    apiOptions.value = (res.data || [])
      .slice()
      .sort((a, b) => (a.path || '').localeCompare(b.path || ''))
      .map((item) => ({
        label: `${item.name} (${item.path})`,
        value: item.id as number,
      }));
  });
}

function getRows() {
  loading.value = true;
  getAdminPermissionList({
    pageNum: pageInfo.pageNum,
    pageSize: pageInfo.pageSize,
    menuId: searchForm.menuId,
    key: searchForm.key,
    name: searchForm.name,
    type: searchForm.type,
    apiId: searchForm.apiId,
  })
    .then((res) => {
      rows.value = res.data.list || [];
      pageInfo.total = res.data.pageInfo?.total || rows.value.length;
    })
    .finally(() => {
      loading.value = false;
    });
}

function onSearch() {
  if (pageInfo.pageNum !== 1) {
    pageInfo.pageNum = 1;
    getRows();
  } else {
    getRows();
  }
}

function onReset() {
  searchForm.menuId = undefined;
  searchForm.key = undefined;
  searchForm.name = undefined;
  searchForm.type = undefined;
  searchForm.apiId = undefined;
  onSearch();
}

function onPageChange(pageNum: number) {
  pageInfo.pageNum = pageNum;
  getRows();
}

function onPageSizeChange(pageSize: number) {
  pageInfo.pageSize = pageSize;
  pageInfo.pageNum = 1;
  getRows();
}

function openDetailModal(record: PermissionListItem) {
  detailData.value = record;
  detailModalStatus.value = true;
}

/** 跳转到菜单管理，在对应菜单的“权限配置”中维护权限点与接口绑定 */
function goMenuPermission() {
  router.push('/admin/menu');
}

onMounted(() => {
  loadMenuTree();
  loadApiOptions();
  getRows();
});
</script>

<style scoped lang="less">
.search-form {
  .ant-form-item {
    margin-bottom: 8px;
  }
}

.muted {
  color: rgba(0, 0, 0, 0.45);
}
</style>
