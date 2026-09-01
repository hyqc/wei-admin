<template>
  <PageContainer :page-info="pageInfo" @page-change="onPageChange" @page-size-change="onPageSizeChange">
    <template #searchArea>
      <a-form layout="inline" :model="searchForm" class="search-form" @finish="onSearch">
        <a-form-item label="唯一键" name="key">
          <a-input v-model:value="searchForm.key" placeholder="请输入唯一键" allow-clear />
        </a-form-item>
        <a-form-item label="名称" name="name">
          <a-input v-model:value="searchForm.name" placeholder="请输入接口名称" allow-clear />
        </a-form-item>
        <a-form-item label="路径" name="path">
          <a-input v-model:value="searchForm.path" placeholder="请输入接口路径" allow-clear />
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
        新建接口
      </a-button>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      row-key="id"
      :scroll="{ x: 1060 }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'isEnabled'">
          <a-popconfirm
            :title="`确定要${record.isEnabled ? '禁用' : '启用'}该接口吗？`"
            ok-text="确定"
            cancel-text="取消"
            @confirm="updateEnabled(record)"
          >
            <RowEnabledButton :is-enabled="record.isEnabled" />
          </a-popconfirm>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <Authorization permission="AdminApiView">
              <a-button type="link" size="small" @click="openDetailModal(record)">
                <template #icon><EyeOutlined /></template>
                详情
              </a-button>
            </Authorization>
            <Authorization permission="AdminApiEdit">
              <a-button type="link" size="small" @click="openEditModal(record)">
                <template #icon><EditOutlined /></template>
                编辑
              </a-button>
            </Authorization>
            <Authorization permission="AdminApiDelete">
              <a-popconfirm
                v-if="!record.isEnabled"
                title="确定要删除该接口吗？"
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

    <AddApiModal v-model:open="addModalStatus" @notice="onModalNotice" />
    <EditApiModal v-model:open="editModalStatus" :detail-data="editData" @notice="onModalNotice" />
    <DetailApiDrawer v-model:open="detailModalStatus" :detail-data="detailData" />
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
  DeleteOutlined,
} from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import RowEnabledButton from '@/components/RowEnabledButton.vue';
import AddApiModal from './components/AddApiModal.vue';
import EditApiModal from './components/EditApiModal.vue';
import DetailApiDrawer from './components/DetailApiDrawer.vue';
import { getAdminApiList, deleteAdminApi, enableAdminApi } from '@/api/admin/api';
import type { AdminApiItem } from '@/types/common';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const loading = ref(false);
const rows = ref<AdminApiItem[]>([]);
const pageInfo = reactive<PageInfoType>({ ...DEFAULT_PAGE_INFO });

const searchForm = reactive<{ key?: string; name?: string; path?: string }>({
  key: undefined,
  name: undefined,
  path: undefined,
});

const addModalStatus = ref(false);
const editModalStatus = ref(false);
const detailModalStatus = ref(false);
const editData = ref<AdminApiItem>();
const detailData = ref<AdminApiItem>();

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
  { title: '接口名称', dataIndex: 'name', key: 'name', width: 160 },
  { title: '唯一键', dataIndex: 'key', key: 'key', width: 200 },
  { title: '接口路径', dataIndex: 'path', key: 'path', width: 240 },
  { title: '状态', dataIndex: 'isEnabled', key: 'isEnabled', width: 100 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 180 },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 260 },
];

function getRows() {
  loading.value = true;
  getAdminApiList({
    pageNum: pageInfo.pageNum,
    pageSize: pageInfo.pageSize,
    key: searchForm.key,
    name: searchForm.name,
    path: searchForm.path,
  })
    .then((res) => {
      rows.value = res.data.list || [];
      pageInfo.total = res.data.pageInfo?.total || 0;
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
  searchForm.key = undefined;
  searchForm.name = undefined;
  searchForm.path = undefined;
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

function updateEnabled(record: AdminApiItem) {
  enableAdminApi({ id: record.id, enabled: !record.isEnabled }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function onDelete(record: AdminApiItem) {
  deleteAdminApi({ id: record.id }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function openAddModal() {
  addModalStatus.value = true;
}

function openEditModal(record: AdminApiItem) {
  editData.value = record;
  editModalStatus.value = true;
}

function openDetailModal(record: AdminApiItem) {
  detailData.value = record;
  detailModalStatus.value = true;
}

function onModalNotice() {
  getRows();
}

onMounted(() => {
  getRows();
});
</script>

<style scoped lang="less">
.search-form {
  .ant-form-item {
    margin-bottom: 8px;
  }
}
</style>
