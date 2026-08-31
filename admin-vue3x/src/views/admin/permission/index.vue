<template>
  <PageContainer title="权限管理" :page-info="pageInfo" @page-change="onPageChange" @page-size-change="onPageSizeChange">
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
          <a-select v-model:value="searchForm.type" placeholder="全部" allow-clear style="width: 120px">
            <a-select-option
              v-for="item in DEFAULT_PERMISSION_TYPES"
              :key="item.key"
              :value="item.key"
            >
              {{ item.name }}
            </a-select-option>
          </a-select>
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
      <a-button type="primary" @click="openAddModal">
        <template #icon><PlusOutlined /></template>
        新建权限
      </a-button>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      row-key="id"
      :scroll="{ x: 1160 }"
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
          <a-popconfirm
            :title="`确定要${record.enabled ? '禁用' : '启用'}该权限吗？`"
            ok-text="确定"
            cancel-text="取消"
            @confirm="updateEnabled(record)"
          >
            <RowEnabledButton :is-enabled="record.enabled" />
          </a-popconfirm>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <Authorization permission="AdminPermissionView">
              <a-button type="link" size="small" @click="openDetailModal(record)">
                <template #icon><EyeOutlined /></template>
                详情
              </a-button>
            </Authorization>
            <Authorization permission="AdminPermissionEdit">
              <a-button type="link" size="small" @click="openEditModal(record)">
                <template #icon><EditOutlined /></template>
                编辑
              </a-button>
            </Authorization>
            <Authorization permission="AdminPermissionEdit">
              <a-button type="link" size="small" @click="openBindApisModal(record)">
                <template #icon><ApiOutlined /></template>
                绑定接口
              </a-button>
            </Authorization>
            <Authorization permission="AdminPermissionDelete">
              <a-popconfirm
                v-if="!record.enabled"
                title="确定要删除该权限吗？"
                ok-text="确定"
                cancel-text="取消"
                @confirm="onDelete(record)"
              >
                <a-button type="link" size="small" danger>
                  <template #icon><DeleteOutlined /></template>
                  删除
                </a-button>
              </a-popconfirm>
            </Authorization>
          </a-space>
        </template>
      </template>
    </a-table>

    <AddPermissionModal v-model:open="addModalStatus" @notice="onModalNotice" />
    <EditPermissionModal v-model:open="editModalStatus" :detail-data="editData" @notice="onModalNotice" />
    <DetailPermissionDrawer v-model:open="detailModalStatus" :detail-data="detailData" @notice="onModalNotice" />
    <BindApisModal v-model:open="bindApisModalStatus" :detail-data="detailData" @notice="onModalNotice" />
  </PageContainer>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import {
  SearchOutlined,
  ReloadOutlined,
  PlusOutlined,
  EyeOutlined,
  EditOutlined,
  ApiOutlined,
  DeleteOutlined,
} from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import RowEnabledButton from '@/components/RowEnabledButton.vue';
import AddPermissionModal from './components/AddPermissionModal.vue';
import EditPermissionModal from './components/EditPermissionModal.vue';
import DetailPermissionDrawer from './components/DetailPermissionDrawer.vue';
import BindApisModal from './components/BindApisModal.vue';
import { DEFAULT_PERMISSION_TYPES } from './components/common';
import { getAdminPermissionList, deleteAdminPermission, enableAdminPermission } from '@/api/admin/permission';
import { getAdminMenuTree } from '@/api/admin/menu';
import type { PermissionListItem } from '@/types/admin_permission';
import type { MenuTreeItem } from '@/types/admin_menu';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const loading = ref(false);
const rows = ref<PermissionListItem[]>([]);
const pageInfo = reactive<PageInfoType>({ ...DEFAULT_PAGE_INFO });
const menuTreeOptions = ref<any[]>([]);

const searchForm = reactive<{
  menuId?: number;
  key?: string;
  name?: string;
  type?: string;
}>({
  menuId: undefined,
  key: undefined,
  name: undefined,
  type: undefined,
});

const addModalStatus = ref(false);
const editModalStatus = ref(false);
const detailModalStatus = ref(false);
const bindApisModalStatus = ref(false);
const editData = ref<PermissionListItem>();
const detailData = ref<PermissionListItem>();

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
  { title: '权限名称', dataIndex: 'name', key: 'name', width: 140 },
  { title: '所属菜单', dataIndex: 'menuName', key: 'menuName', width: 140 },
  { title: '唯一键', dataIndex: 'key', key: 'key', width: 180 },
  { title: '类型', dataIndex: 'type', key: 'type', width: 100 },
  { title: '接口数', dataIndex: 'apis', key: 'apis', width: 100 },
  { title: '状态', dataIndex: 'enabled', key: 'enabled', width: 100 },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 320 },
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

function getRows() {
  loading.value = true;
  getAdminPermissionList({
    pageNum: pageInfo.pageNum,
    pageSize: pageInfo.pageSize,
    menuId: searchForm.menuId,
    key: searchForm.key,
    name: searchForm.name,
    type: searchForm.type,
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

function updateEnabled(record: PermissionListItem) {
  enableAdminPermission({ id: record.id, enabled: !record.enabled }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function onDelete(record: PermissionListItem) {
  deleteAdminPermission({ id: record.id }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function openAddModal() {
  addModalStatus.value = true;
}

function openEditModal(record: PermissionListItem) {
  editData.value = record;
  editModalStatus.value = true;
}

function openDetailModal(record: PermissionListItem) {
  detailData.value = record;
  detailModalStatus.value = true;
}

function openBindApisModal(record: PermissionListItem) {
  detailData.value = record;
  bindApisModalStatus.value = true;
}

function onModalNotice() {
  getRows();
}

onMounted(() => {
  loadMenuTree();
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
