<template>
  <PageContainer :page-info="pageInfo" @page-change="onPageChange" @page-size-change="onPageSizeChange">
    <template #searchArea>
      <a-form layout="inline" :model="searchForm" class="search-form" @finish="onSearch">
        <a-form-item label="账号" name="username">
          <a-input v-model:value="searchForm.username" placeholder="请输入账号" allow-clear />
        </a-form-item>
        <a-form-item label="昵称" name="nickname">
          <a-input v-model:value="searchForm.nickname" placeholder="请输入昵称" allow-clear />
        </a-form-item>
        <a-form-item label="邮箱" name="email">
          <a-input v-model:value="searchForm.email" placeholder="请输入邮箱" allow-clear />
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
      <a-button type="primary" @click="openAddModal">
        <template #icon><PlusOutlined /></template>
        新建账号
      </a-button>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      row-key="adminId"
      :scroll="{ x: 1940 }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'nickname'">
          <a-tooltip v-if="record.nickname && record.nickname.length > 10" :title="record.nickname">
            <span class="nickname-cell">{{ displayNickname(record.nickname) }}</span>
          </a-tooltip>
          <span v-else>{{ displayNickname(record.nickname) }}</span>
        </template>
        <template v-else-if="column.key === 'lastLoginIp'">
          {{ record.lastLoginIp || '-' }}
        </template>
        <template v-else-if="column.key === 'currentLoginIp'">
          {{ record.currentLoginIp || getLoginIp(record.lastLoginIp, 'current') }}
        </template>
        <template v-else-if="column.key === 'roles'">
          <!-- 超管无需角色，自动拥有系统全部权限 -->
          <a-tag v-if="record.isSuperAdmin" color="gold">超级管理员</a-tag>
          <template v-else>
            <a-tag v-for="role in record.roles" :key="role.roleId" color="blue">
              {{ role.roleName }}
            </a-tag>
          </template>
        </template>
        <template v-else-if="column.key === 'isEnabled'">
          <a-popconfirm
            :title="`确定要${record.isEnabled ? '禁用' : '启用'}该账号吗？`"
            ok-text="确定"
            cancel-text="取消"
            @confirm="updateEnabled(record)"
          >
            <RowEnabledButton :is-enabled="record.isEnabled" :is-enabled-button-disabled="record.adminId === AdminId" />
          </a-popconfirm>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <Authorization permission="AdminUserView">
              <a-button type="link" size="small" @click="openDetailModal(record)">
                <template #icon><EyeOutlined /></template>
                详情
              </a-button>
            </Authorization>
            <Authorization permission="AdminUserEdit">
              <a-button type="link" size="small" @click="openEditModal(record)">
                <template #icon><EditOutlined /></template>
                编辑
              </a-button>
            </Authorization>
            <Authorization permission="AdminUserEdit">
              <a-button type="link" size="small" @click="openResetPwdModal(record)">
                <template #icon><KeyOutlined /></template>
                重置密码
              </a-button>
            </Authorization>
            <Authorization permission="AdminUserEdit">
              <!-- 超管账号自动拥有全部权限，不允许绑定角色 -->
              <a-button v-if="!record.isSuperAdmin" type="link" size="small" @click="openBindRolesModal(record)">
                <template #icon><UserSwitchOutlined /></template>
                绑定角色
              </a-button>
            </Authorization>
            <Authorization permission="AdminUserDelete">
              <!-- 超管不可删除，仅被禁用的账号可删除 -->
              <a-popconfirm
                v-if="record.username !== 'admin' && !record.isEnabled"
                title="确定要删除该账号吗？"
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

    <AddUserModal v-model:open="addModalStatus" @notice="onModalNotice" />
    <EditUserModal
      v-model:open="editModalStatus"
      :detail-data="editData"
      @notice="onModalNotice"
    />
    <DetailUserDrawer v-model:open="detailModalStatus" :detail-data="detailData" />
    <BindRolesModal
      v-model:open="bindRolesModalStatus"
      :detail-data="detailData"
      @notice="onModalNotice"
    />
    <ResetPasswordModal
      v-model:open="resetPwdModalStatus"
      :detail-data="detailData"
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
  KeyOutlined,
  UserSwitchOutlined,
  DeleteOutlined,
} from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import RowEnabledButton from '@/components/RowEnabledButton.vue';
import AddUserModal from './components/AddUserModal.vue';
import EditUserModal from './components/EditUserModal.vue';
import DetailUserDrawer from './components/DetailUserDrawer.vue';
import BindRolesModal from './components/BindRolesModal.vue';
import ResetPasswordModal from './components/ResetPasswordModal.vue';
import {
  getAdminUserList,
  deleteAdminUser,
  enableAdminUser,
} from '@/api/admin/user';
import type { AdminUserListItem } from '@/types/common';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import { AdminId } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const loading = ref(false);
const rows = ref<AdminUserListItem[]>([]);
const pageInfo = reactive<PageInfoType>({ ...DEFAULT_PAGE_INFO });

const searchForm = reactive<{
  username?: string;
  nickname?: string;
  email?: string;
  enabled?: number;
}>({
  username: undefined,
  nickname: undefined,
  email: undefined,
  enabled: 0,
});

// 弹窗状态
const addModalStatus = ref(false);
const editModalStatus = ref(false);
const detailModalStatus = ref(false);
const bindRolesModalStatus = ref(false);
const resetPwdModalStatus = ref(false);
const editData = ref<AdminUserListItem>();
const detailData = ref<AdminUserListItem>();

const columns = [
  { title: 'ID', dataIndex: 'adminId', key: 'adminId', fixed: 'left', width: 80 },
  { title: '账号', dataIndex: 'username', key: 'username', fixed: 'left', width: 120 },
  { title: '昵称', dataIndex: 'nickname', key: 'nickname', width: 160 },
  { title: '邮箱', dataIndex: 'email', key: 'email', width: 180 },
  { title: '角色', dataIndex: 'roles', key: 'roles', width: 180 },
  { title: '状态', dataIndex: 'isEnabled', key: 'isEnabled', width: 100 },
  { title: '登录次数', dataIndex: 'loginTotal', key: 'loginTotal', width: 100 },
  { title: '上次登录IP', dataIndex: 'lastLoginIp', key: 'lastLoginIp', width: 140 },
  { title: '当前登录IP', dataIndex: 'currentLoginIp', key: 'currentLoginIp', width: 140 },
  { title: '上次登录时间', dataIndex: 'lastLoginTime', key: 'lastLoginTime', width: 180 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 420 },
];

/** 昵称超 10 字符截断显示，超出部分用省略号 */
function displayNickname(nickname?: string): string {
  if (!nickname) return '';
  return nickname.length > 10 ? `${nickname.slice(0, 10)}...` : nickname;
}

/** 拆分登录IP：last 取倒数第 2 个（上次登录），current 取最后一个（当前登录） */
function getLoginIp(ipStr?: string, type: 'last' | 'current' = 'last'): string {
  // 兼容 JSON 数组字符串（后端原始存储，如 ["ip1","ip2"]）与逗号分隔字符串（解析后/ mock，如 ip1, ip2）
  const s = ipStr || '';
  let list: string[] = [];
  try {
    const parsed = JSON.parse(s);
    if (Array.isArray(parsed)) list = parsed.map(String);
  } catch {
    list = [];
  }
  if (list.length === 0) {
    list = s
      .split(',')
      .map((x) => x.trim())
      .filter(Boolean);
  }
  const ip = type === 'current' ? list[list.length - 1] : list[list.length - 2];
  return ip || '-';
}

function getRows() {
  loading.value = true;
  getAdminUserList({
    pageNum: pageInfo.pageNum,
    pageSize: pageInfo.pageSize,
    username: searchForm.username,
    nickname: searchForm.nickname,
    email: searchForm.email,
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
  searchForm.username = undefined;
  searchForm.nickname = undefined;
  searchForm.email = undefined;
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

function updateEnabled(record: AdminUserListItem) {
  enableAdminUser({ adminId: record.adminId, enabled: !record.isEnabled }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function onDelete(record: AdminUserListItem) {
  deleteAdminUser({ adminId: record.adminId }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function openAddModal() {
  addModalStatus.value = true;
}

function openEditModal(record: AdminUserListItem) {
  editData.value = record;
  editModalStatus.value = true;
}

function openDetailModal(record: AdminUserListItem) {
  detailData.value = record;
  detailModalStatus.value = true;
}

function openBindRolesModal(record: AdminUserListItem) {
  detailData.value = record;
  bindRolesModalStatus.value = true;
}

function openResetPwdModal(record: AdminUserListItem) {
  detailData.value = record;
  resetPwdModalStatus.value = true;
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

.nickname-cell {
  display: inline-block;
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  vertical-align: middle;
}
</style>
