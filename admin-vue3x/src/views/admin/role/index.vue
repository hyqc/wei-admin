<template>
  <PageContainer :page-info="pageInfo" @page-change="onPageChange" @page-size-change="onPageSizeChange">
    <template #searchArea>
      <a-form layout="inline" :model="searchForm" class="search-form" @finish="onSearch">
        <a-form-item label="角色名称" name="name">
          <a-input v-model:value="searchForm.name" placeholder="请输入角色名称" allow-clear />
        </a-form-item>
        <a-form-item label="状态" name="enabled">
          <a-select v-model:value="searchForm.enabled" style="width: 120px" placeholder="全部">
            <a-select-option :value="0">全部</a-select-option>
            <a-select-option :value="1">启用</a-select-option>
            <a-select-option :value="2">禁用</a-select-option>
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
      <Authorization permission="AdminRoleAdd">
        <a-button type="primary" @click="openAddModal">
          <template #icon><PlusOutlined /></template>
          新建角色
        </a-button>
      </Authorization>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      row-key="id"
      :scroll="{ x: 1320 }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <a-space>
            <span>{{ record.name }}</span>
            <a-tag v-if="record.isSuperAdmin" color="gold">超级管理员</a-tag>
          </a-space>
        </template>
        <template v-else-if="column.key === 'isEnabled'">
          <Authorization permission="AdminRoleEdit">
            <a-popconfirm
              :title="`确定要${record.isEnabled ? '禁用' : '启用'}该角色吗？`"
              ok-text="确定"
              cancel-text="取消"
              @confirm="updateEnabled(record)"
            >
              <RowEnabledButton :is-enabled="record.isEnabled" />
            </a-popconfirm>
          </Authorization>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <Authorization permission="AdminRoleView">
              <a-button type="link" size="small" @click="openDetailModal(record)">
                <template #icon><EyeOutlined /></template>
                详情
              </a-button>
            </Authorization>
            <Authorization permission="AdminRoleEdit">
              <!-- 超管角色允许编辑，但仅可修改描述 -->
              <a-button type="link" size="small" @click="openEditModal(record)">
                <template #icon><EditOutlined /></template>
                编辑
              </a-button>
            </Authorization>
            <Authorization permission="AdminRoleBindPermissions">
              <!-- 超级管理员角色不允许绑定权限 -->
              <a-button v-if="record.id !== 1" type="link" size="small" @click="openBindModal(record)">
                <template #icon><SafetyCertificateOutlined /></template>
                绑定权限
              </a-button>
            </Authorization>
            <Authorization permission="AdminRoleDelete">
              <a-popconfirm
                v-if="!record.isEnabled"
                title="确定要删除该角色吗？"
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

    <AddRoleModal v-model:open="addModalStatus" @notice="onModalNotice" />
    <EditRoleModal v-model:open="editModalStatus" :detail-data="editData" @notice="onModalNotice" />
    <DetailRoleDrawer v-model:open="detailModalStatus" :detail-data="detailData" />
    <BindPermissionsModal
      v-model:open="bindModalStatus"
      :detail-data="detailData"
      @notice="onModalNotice"
    />
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
  SafetyCertificateOutlined,
  DeleteOutlined,
} from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import RowEnabledButton from '@/components/RowEnabledButton.vue';
import AddRoleModal from './components/AddRoleModal.vue';
import EditRoleModal from './components/EditRoleModal.vue';
import DetailRoleDrawer from './components/DetailRoleDrawer.vue';
import BindPermissionsModal from './components/BindPermissionsModal.vue';
import { getAdminRoleList, deleteAdminRole, enableAdminRole } from '@/api/admin/role';
import type { RoleItem } from '@/types/admin_role';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const loading = ref(false);
const rows = ref<RoleItem[]>([]);
const pageInfo = reactive<PageInfoType>({ ...DEFAULT_PAGE_INFO });

const searchForm = reactive<{ name?: string; enabled?: number }>({
  name: undefined,
  enabled: 0,
});

const addModalStatus = ref(false);
const editModalStatus = ref(false);
const detailModalStatus = ref(false);
const bindModalStatus = ref(false);
const editData = ref<RoleItem>();
const detailData = ref<RoleItem>();

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
  { title: '角色名称', dataIndex: 'name', key: 'name', width: 160 },
  { title: '描述', dataIndex: 'describe', key: 'describe', width: 180 },
  { title: '创建人', dataIndex: 'createAdminName', key: 'createAdminName', width: 120 },
  { title: '状态', dataIndex: 'isEnabled', key: 'isEnabled', width: 100 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 180 },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 320 },
];

function getRows() {
  loading.value = true;
  getAdminRoleList({
    pageNum: pageInfo.pageNum,
    pageSize: pageInfo.pageSize,
    name: searchForm.name,
    enabled: searchForm.enabled,
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
  searchForm.name = undefined;
  searchForm.enabled = 0;
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

function updateEnabled(record: RoleItem) {
  enableAdminRole({ id: record.id, enabled: !record.isEnabled }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function onDelete(record: RoleItem) {
  deleteAdminRole({ id: record.id }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function openAddModal() {
  addModalStatus.value = true;
}

function openEditModal(record: RoleItem) {
  editData.value = record;
  editModalStatus.value = true;
}

function openDetailModal(record: RoleItem) {
  detailData.value = record;
  detailModalStatus.value = true;
}

function openBindModal(record: RoleItem) {
  detailData.value = record;
  bindModalStatus.value = true;
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
